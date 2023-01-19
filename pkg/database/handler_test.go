package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

func TestHandler_FindById(t *testing.T) {
	groupName := "name"
	repository := &mockRepository{}
	database := &model.Database{
		GroupName: groupName,
	}
	repository.
		On("FindById", uint(1)).
		Return(database, nil)
	service := NewService(config.Config{}, nil, nil, repository)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.AddParam("id", "1")
	user := &models.User{
		Groups: []*models.Group{
			{Name: groupName},
		},
	}
	c.Set("user", user)

	handler.FindById(c)

	assert.Empty(t, c.Errors)
	var actualBody model.Database
	assertResponse(t, w, http.StatusOK, &actualBody, database)
	repository.AssertExpectations(t)
}

func TestHandler_Copy(t *testing.T) {
	token := "token"
	groupName := "name"
	databaseName := "name"
	group := &models.Group{
		Name: groupName,
	}
	userClient := &mockUserClient{}
	userClient.
		On("FindGroupByName", token, groupName).
		Return(group, nil)

	s3Client := &mockS3Client{}
	s3Client.
		On("Copy", mock.AnythingOfType("string"), "path", fmt.Sprintf("%s/%s", group.Name, databaseName)).
		Return(nil)
	repository := &mockRepository{}
	repository.
		On("FindById", uint(1)).
		Return(&model.Database{
			Url: "/path",
		}, nil)
	repository.
		On("Create", mock.Anything).
		Return(nil)
	service := NewService(config.Config{}, nil, s3Client, repository)
	handler := New(userClient, service, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.AddParam("id", "1")

	user := &models.User{
		ID: 1,
		Groups: []*models.Group{
			group,
		},
	}
	c.Set("user", user)

	copyRequest := &CopyDatabaseRequest{
		Name:  databaseName,
		Group: groupName,
	}
	request := newRequest(t, http.MethodPost, "/groups", copyRequest)
	request.Header.Set("Authorization", token)
	c.Request = request

	handler.Copy(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Empty(t, c.Errors)
	userClient.AssertExpectations(t)
	s3Client.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func newRequest(t *testing.T, method string, path string, request any) *http.Request {
	body, err := json.Marshal(request)
	require.NoError(t, err)

	req, err := http.NewRequest(method, path, bytes.NewReader(body))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return req
}

func TestHandler_List(t *testing.T) {
	name := "name"

	groups := []*models.Group{
		{
			Name: name,
		},
	}

	databases := []*model.Database{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "some name",
			GroupName: name,
			Url:       "",
		},
	}

	repository := &mockRepository{}
	repository.
		On("FindByGroupNames", []string{name}).
		Return(databases, nil)
	service := NewService(config.Config{}, nil, nil, repository)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user", &models.User{Groups: groups})

	handler.List(c)

	assert.Empty(t, c.Errors)
	expectedBody := &[]GroupsWithDatabases{
		{
			Name:      name,
			Databases: databases,
		},
	}
	var actualBody []GroupsWithDatabases
	assertResponse(t, w, http.StatusOK, &actualBody, expectedBody)
	repository.AssertExpectations(t)
}

func assertResponse(t *testing.T, rec *httptest.ResponseRecorder, expectedCode int, bodyType any, expectedBody any) {
	assert.Equal(t, expectedCode, rec.Code, "HTTP status code does not match")
	assertJSON(t, rec.Body, bodyType, expectedBody)
}

func assertJSON(t *testing.T, body *bytes.Buffer, v any, expected any) {
	err := json.Unmarshal(body.Bytes(), v)
	require.NoError(t, err)
	assert.Equal(t, expected, v, "HTTP response body does not match")
}

func TestHandler_List_RepositoryError(t *testing.T) {
	groups := []*models.Group{
		{
			Name: "name",
		},
	}

	errorMessage := "some error"

	repository := &mockRepository{}
	repository.
		On("FindByGroupNames", []string{groups[0].Name}).
		Return(nil, errors.New(errorMessage))
	service := NewService(config.Config{}, nil, nil, repository)
	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user", &models.User{Groups: groups})

	handler.List(c)

	assert.Empty(t, w.Body.Bytes())
	assert.Len(t, c.Errors, 1)
	assert.ErrorContains(t, c.Errors[0].Err, errorMessage)
	repository.AssertExpectations(t)
}

type mockRepository struct{ mock.Mock }

func (m *mockRepository) Create(d *model.Database) error {
	return m.Called(d).Error(0)
}

func (m *mockRepository) Save(d *model.Database) error {
	panic("implement me")
}

func (m *mockRepository) FindById(id uint) (*model.Database, error) {
	called := m.Called(id)
	return called.Get(0).(*model.Database), nil
}

func (m *mockRepository) Lock(id, instanceId, userId uint) (*model.Lock, error) {
	panic("implement me")
}

func (m *mockRepository) Unlock(id uint) error {
	panic("implement me")
}

func (m *mockRepository) Delete(id uint) error {
	panic("implement me")
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
	panic("implement me")
}

func (m *mockRepository) CreateExternalDownload(databaseID uint, expiration time.Time) (model.ExternalDownload, error) {
	panic("implement me")
}

func (m *mockRepository) FindExternalDownload(uuid uuid.UUID) (model.ExternalDownload, error) {
	panic("implement me")
}

func (m *mockRepository) PurgeExternalDownload() error {
	panic("implement me")
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
	panic("implement me")
}

func (m *mockS3Client) Download(bucket string, key string, dst io.Writer, cb func(contentLength int64)) error {
	panic("implement me")
}
