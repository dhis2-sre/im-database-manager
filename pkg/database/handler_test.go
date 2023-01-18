package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
	panic("implement me")
}

func (m *mockRepository) Save(d *model.Database) error {
	panic("implement me")
}

func (m *mockRepository) FindById(id uint) (*model.Database, error) {
	panic("implement me")
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
