package database

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	pg "github.com/habx/pg-commands"

	h "github.com/dhis2-sre/im-database-manager/internal/handler"

	instanceModels "github.com/dhis2-sre/im-manager/swagger/sdk/models"
	userModels "github.com/dhis2-sre/im-user/swagger/sdk/models"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/dhis2-sre/im-database-manager/pkg/storage"

	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestHandler_SaveAs(t *testing.T) {
	err := h.RegisterValidation()
	require.NoError(t, err)
	userClient := &mockUserClient{}
	userClient.
		On("FindGroupByName", "token", "group-name").
		Return(&userModels.Group{
			Name:                 "group-name",
			ClusterConfiguration: &models.ClusterConfiguration{KubernetesConfiguration: nil},
		}, nil)
	database := &model.Database{Name: "new-name", GroupName: "group-name"}
	repository := &mockRepository{}
	repository.
		On("FindById", uint(1)).
		Return(database, nil)
	repository.
		On("Save", database).
		Return(nil)
	newPgDump := &mockNewPgDump{}
	postgres := &pg.Postgres{Host: "%!(EXTRA string=, string=group-name)", Port: 5432, DB: "database-name", Username: "database-username", Password: "database-password"}
	pgDump := &pg.Dump{Postgres: postgres}
	newPgDump.
		On("NewDump", postgres).
		Return(pgDump, nil)
	service := NewService(config.Config{}, userClient, nil, repository, newPgDump)
	instanceClient := &mockInstanceClient{}
	instanceClient.
		On("FindByIdDecrypted", "token", uint(1)).
		Return(&instanceModels.Instance{
			ID:        1,
			UserID:    1,
			GroupName: "group-name",
			StackName: "stack-name",
			RequiredParameters: []*instanceModels.InstanceRequiredParameter{
				{StackRequiredParameterID: "DATABASE_ID", Value: "1"},
				{StackRequiredParameterID: "DATABASE_NAME", Value: "database-name"},
				{StackRequiredParameterID: "DATABASE_USERNAME", Value: "database-username"},
				{StackRequiredParameterID: "DATABASE_PASSWORD", Value: "database-password"},
			},
		}, nil)
	instanceClient.
		On("FindStack", "token", "stack-name").
		Return(&instanceModels.Stack{Name: "stack-name"}, nil)
	handler := New(nil, service, instanceClient)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	c.AddParam("instanceId", "1")
	c.Request = newPost(t, "", &saveAsRequest{Name: "new-name", Format: "custom"})

	handler.SaveAs(c)

	assert.Empty(t, c.Errors)
	assertResponse(t, w, http.StatusCreated, database)
	repository.AssertExpectations(t)
	newPgDump.AssertExpectations(t)
	userClient.AssertExpectations(t)
	instanceClient.AssertExpectations(t)
}

type mockNewPgDump struct{ mock.Mock }

func (m *mockNewPgDump) NewDump(opts *pg.Postgres) (*pg.Dump, error) {
	called := m.Called(opts)
	return called.Get(0).(*pg.Dump), nil
}

type mockInstanceClient struct{ mock.Mock }

func (m *mockInstanceClient) FindByIdDecrypted(token string, id uint) (*instanceModels.Instance, error) {
	called := m.Called(token, id)
	return called.Get(0).(*instanceModels.Instance), nil
}

func (m *mockInstanceClient) FindStack(token string, name string) (*instanceModels.Stack, error) {
	called := m.Called(token, name)
	return called.Get(0).(*instanceModels.Stack), nil
}

func TestHandler_CreateExternalDownload(t *testing.T) {
	repository := &mockRepository{}
	repository.
		On("FindById", uint(1)).
		Return(&model.Database{
			Model:     gorm.Model{ID: 1},
			GroupName: "group-name",
		}, nil)
	repository.
		On("PurgeExternalDownload").
		Return(nil)
	expiration := time.Now().Add(time.Duration(1) * time.Hour).Round(time.Duration(1)).UTC()
	externalDownload := model.ExternalDownload{
		UUID:       uuid.UUID{},
		Expiration: expiration,
		DatabaseID: 1,
	}
	repository.
		On("CreateExternalDownload", uint(1), expiration).
		Return(externalDownload, nil)
	service := NewService(config.Config{}, nil, nil, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	c.AddParam("id", "1")
	createExternalDatabaseRequest := &CreateExternalDatabaseRequest{Expiration: expiration}
	c.Request = newPost(t, "/databases/1/external", createExternalDatabaseRequest)

	handler.CreateExternalDownload(c)

	require.Empty(t, c.Errors)
	assertResponse(t, w, http.StatusCreated, &externalDownload)
	repository.AssertExpectations(t)
}

func TestHandler_Download(t *testing.T) {
	awsS3Client := &mockAWSS3Client{}
	awsS3Client.
		On("GetObject", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*s3.GetObjectInput"), mock.AnythingOfType("[]func(*s3.Options)")).
		Return(&s3.GetObjectOutput{
			Body:          io.NopCloser(strings.NewReader("Hello, World!")),
			ContentLength: 13,
		}, nil)
	s3Client, err := storage.NewS3Client(awsS3Client, nil)
	require.NoError(t, err)
	repository := &mockRepository{}
	repository.
		On("FindById", uint(1)).
		Return(&model.Database{
			GroupName: "group-name",
			Url:       "s3://whatever",
		}, nil)
	service := NewService(config.Config{}, nil, s3Client, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.AddParam("id", "1")
	user := &models.User{
		Groups: []*models.Group{
			{Name: "group-name"},
		},
	}
	c.Set("user", user)

	handler.Download(c)

	assert.Empty(t, c.Errors)
	headers := w.Header()
	assert.Equal(t, "attachment; filename=whatever", headers.Get("Content-Disposition"))
	assert.Equal(t, "File Transfer", headers.Get("Content-Description"))
	assert.Equal(t, "binary", headers.Get("Content-Transfer-Encoding"))
	assert.Equal(t, "application/octet-stream", headers.Get("Content-Type"))
	assert.Equal(t, "13", headers.Get("Content-Length"))
	assert.Equal(t, "Hello, World!", w.Body.String())
	repository.AssertExpectations(t)
	awsS3Client.AssertExpectations(t)
}

type mockAWSS3Client struct{ mock.Mock }

func (m *mockAWSS3Client) CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error) {
	panic("implement me")
}

func (m *mockAWSS3Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	panic("implement me")
}

func (m *mockAWSS3Client) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	called := m.Called(ctx, params, optFns)
	return called.Get(0).(*s3.GetObjectOutput), nil
}

func (m *mockAWSS3Client) PutObject(ctx context.Context, input *s3.PutObjectInput, f ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	panic("implement me")
}

func (m *mockAWSS3Client) UploadPart(ctx context.Context, input *s3.UploadPartInput, f ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	panic("implement me")
}

func (m *mockAWSS3Client) CreateMultipartUpload(ctx context.Context, input *s3.CreateMultipartUploadInput, f ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	panic("implement me")
}

func (m *mockAWSS3Client) CompleteMultipartUpload(ctx context.Context, input *s3.CompleteMultipartUploadInput, f ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	panic("implement me")
}

func (m *mockAWSS3Client) AbortMultipartUpload(ctx context.Context, input *s3.AbortMultipartUploadInput, f ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
	panic("implement me")
}

func TestHandler_Update(t *testing.T) {
	database := &model.Database{GroupName: "group-name"}
	repository := &mockRepository{}
	repository.
		On("FindById", uint(1)).
		Return(database, nil)
	repository.
		On("Update", database).
		Return(nil)
	service := NewService(config.Config{}, nil, nil, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	c.AddParam("id", "1")
	updateDatabaseRequest := &UpdateDatabaseRequest{Name: "database-name"}
	c.Request = newPost(t, "/databases/1", updateDatabaseRequest)

	handler.Update(c)

	assert.Empty(t, c.Errors)
	assertResponse(t, w, http.StatusOK, database)
	repository.AssertExpectations(t)
}

func TestHandler_Unlock(t *testing.T) {
	repository := &mockRepository{}
	repository.
		On("FindById", uint(1)).
		Return(&model.Database{
			GroupName: "group-name",
			Lock: &model.Lock{
				DatabaseID: 1,
				InstanceID: 1,
				UserID:     1,
			},
		}, nil)
	repository.
		On("Unlock", uint(1)).
		Return(nil)
	service := NewService(config.Config{}, nil, nil, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	c.AddParam("id", "1")

	handler.Unlock(c)

	assert.Empty(t, c.Errors)
	assert.Empty(t, w.Body)
	c.Writer.Flush()
	assert.Equal(t, http.StatusAccepted, w.Code)
	repository.AssertExpectations(t)
}

func TestHandler_Lock(t *testing.T) {
	repository := &mockRepository{}
	database := &model.Database{GroupName: "group-name"}
	repository.
		On("FindById", uint(1)).
		Return(database, nil)
	lock := &model.Lock{
		DatabaseID: 1,
		InstanceID: 1,
		UserID:     1,
	}
	repository.
		On("Lock", uint(1), uint(1), uint(1)).
		Return(lock, nil)
	service := NewService(config.Config{}, nil, nil, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	c.AddParam("id", "1")
	lockDatabaseRequest := &LockDatabaseRequest{InstanceId: 1}
	c.Request = newPost(t, "/whatever", lockDatabaseRequest)

	handler.Lock(c)

	assert.Empty(t, c.Errors)
	assertResponse(t, w, http.StatusCreated, lock)
	repository.AssertExpectations(t)
}

func TestHandler_Delete(t *testing.T) {
	s3Client := &mockS3Client{}
	s3Client.
		On("Delete", "", "path").
		Return(nil)
	database := &model.Database{
		GroupName: "group-name",
		Url:       "/path",
	}
	repository := &mockRepository{}
	repository.
		On("FindById", uint(1)).
		Return(database, nil)
	repository.
		On("Delete", uint(1)).
		Return(nil)
	service := NewService(config.Config{}, nil, s3Client, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	c.AddParam("id", "1")

	handler.Delete(c)

	assert.Empty(t, c.Errors)
	assert.Empty(t, w.Body)
	c.Writer.Flush()
	assert.Equal(t, http.StatusAccepted, w.Code)
	repository.AssertExpectations(t)
	s3Client.AssertExpectations(t)
}

func newContext(w *httptest.ResponseRecorder, group string) *gin.Context {
	user := &models.User{
		ID: uint64(1),
		Groups: []*models.Group{
			{Name: group},
		},
	}
	c, _ := gin.CreateTestContext(w)
	c.Set("user", user)
	return c
}

func TestHandler_FindById(t *testing.T) {
	repository := &mockRepository{}
	database := &model.Database{
		GroupName: "group-name",
	}
	repository.
		On("FindById", uint(1)).
		Return(database, nil)
	service := NewService(config.Config{}, nil, nil, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	c.AddParam("id", "1")

	handler.FindById(c)

	assert.Empty(t, c.Errors)
	assertResponse(t, w, http.StatusOK, database)
	repository.AssertExpectations(t)
}

func TestHandler_Copy(t *testing.T) {
	userClient := &mockUserClient{}
	userClient.
		On("FindGroupByName", "token", "group-name").
		Return(&models.Group{
			Name: "group-name",
		}, nil)
	s3Client := &mockS3Client{}
	s3Client.
		On("Copy", mock.AnythingOfType("string"), "path", "group-name/database-name").
		Return(nil)
	repository := &mockRepository{}
	repository.
		On("FindById", uint(1)).
		Return(&model.Database{
			Url: "/path",
		}, nil)
	repository.
		On("Create", mock.AnythingOfType("*model.Database")).
		Return(nil)
	service := NewService(config.Config{}, nil, s3Client, repository, nil)
	handler := New(userClient, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")
	c.AddParam("id", "1")
	copyRequest := &CopyDatabaseRequest{
		Name:  "database-name",
		Group: "group-name",
	}
	c.Request = newPost(t, "/groups", copyRequest)

	handler.Copy(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Empty(t, c.Errors)
	userClient.AssertExpectations(t)
	s3Client.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func newPost(t *testing.T, path string, jsonBody any) *http.Request {
	body, err := json.Marshal(jsonBody)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "token")

	return req
}

func TestHandler_List(t *testing.T) {
	databases := []*model.Database{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "some name",
			GroupName: "group-name",
			Url:       "",
		},
	}
	repository := &mockRepository{}
	repository.
		On("FindByGroupNames", []string{"group-name"}).
		Return(databases, nil)
	service := NewService(config.Config{}, nil, nil, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")

	handler.List(c)

	assert.Empty(t, c.Errors)
	expectedBody := &[]GroupsWithDatabases{
		{
			Name:      "group-name",
			Databases: databases,
		},
	}
	assertResponse(t, w, http.StatusOK, expectedBody)
	repository.AssertExpectations(t)
}

func assertResponse[V any](t *testing.T, rec *httptest.ResponseRecorder, expectedCode int, expectedBody V) {
	require.Equal(t, expectedCode, rec.Code, "HTTP status code does not match")
	assertJSON(t, rec.Body, expectedBody)
}

func assertJSON[V any](t *testing.T, body *bytes.Buffer, expected V) {
	actualBody := new(V)
	err := json.Unmarshal(body.Bytes(), &actualBody)
	require.NoError(t, err)
	require.Equal(t, expected, *actualBody, "HTTP response body does not match")
}
func TestHandler_List_RepositoryError(t *testing.T) {
	repository := &mockRepository{}
	repository.
		On("FindByGroupNames", []string{"group-name"}).
		Return(nil, errors.New("some error"))
	service := NewService(config.Config{}, nil, nil, repository, nil)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c := newContext(w, "group-name")

	handler.List(c)

	assert.Empty(t, w.Body.Bytes())
	assert.Len(t, c.Errors, 1)
	assert.ErrorContains(t, c.Errors[0].Err, "some error")
	repository.AssertExpectations(t)
}

type mockRepository struct{ mock.Mock }

func (m *mockRepository) Create(d *model.Database) error {
	return m.Called(d).Error(0)
}

func (m *mockRepository) Save(d *model.Database) error {
	called := m.Called(d)
	return called.Error(0)
}

func (m *mockRepository) FindById(id uint) (*model.Database, error) {
	called := m.Called(id)
	return called.Get(0).(*model.Database), nil
}

func (m *mockRepository) Lock(id, instanceId, userId uint) (*model.Lock, error) {
	called := m.Called(id, instanceId, userId)
	return called.Get(0).(*model.Lock), nil
}

func (m *mockRepository) Unlock(id uint) error {
	called := m.Called(id)
	return called.Error(0)
}

func (m *mockRepository) Delete(id uint) error {
	return m.Called(id).Error(0)
}

func (m *mockRepository) FindByGroupNames(names []string) ([]*model.Database, error) {
	called := m.Called(names)
	databases, ok := called.Get(0).([]*model.Database)
	if ok {
		return databases, nil
	}
	return nil, called.Error(1)
}

func (m *mockRepository) Update(d *model.Database) error {
	called := m.Called(d)
	return called.Error(0)
}

func (m *mockRepository) CreateExternalDownload(databaseID uint, expiration time.Time) (model.ExternalDownload, error) {
	called := m.Called(databaseID, expiration)
	return called.Get(0).(model.ExternalDownload), nil
}

func (m *mockRepository) FindExternalDownload(uuid uuid.UUID) (model.ExternalDownload, error) {
	panic("implement me")
}

func (m *mockRepository) PurgeExternalDownload() error {
	return m.Called().Error(0)
}

type mockUserClient struct{ mock.Mock }

func (m *mockUserClient) FindGroupByName(token string, name string) (*models.Group, error) {
	called := m.Called(token, name)
	return called.Get(0).(*models.Group), nil
}

func (m *mockUserClient) FindUserById(token string, id uint) (*models.User, error) {
	panic("implement me")
}

type mockS3Client struct{ mock.Mock }

func (m *mockS3Client) Copy(bucket string, source string, destination string) error {
	called := m.Called(bucket, source, destination)
	return called.Error(0)
}

func (m *mockS3Client) Upload(bucket string, key string, body *bytes.Buffer) error {
	panic("implement me")
}

func (m *mockS3Client) Delete(bucket string, key string) error {
	return m.Called(bucket, key).Error(0)
}

func (m *mockS3Client) Download(bucket string, key string, dst io.Writer, cb func(contentLength int64)) error {
	panic("implement me")
}
