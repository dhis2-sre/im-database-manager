package database

import (
	"encoding/json"
	"errors"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	instanceModels "github.com/dhis2-sre/im-manager/swagger/sdk/models"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_List(t *testing.T) {
	groups := []*models.Group{
		{
			Name:     "name",
			Hostname: "hostname",
		},
	}

	databases := []*model.Database{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "some name",
			GroupName: groups[0].Name,
			Url:       "",
		},
	}

	service := &mockDatabaseService{}
	service.
		On("List", groups).
		Return(databases, nil)

	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user", &models.User{Groups: groups})

	handler.List(c)

	assert.Empty(t, c.Errors)
	assert.Equal(t, http.StatusOK, w.Code)
	var body []GroupsWithDatabases
	err := json.Unmarshal(w.Body.Bytes(), &body)
	require.NoError(t, err)
	assert.Equal(t, groupsWithDatabases(groups, databases), body)

	service.AssertExpectations(t)
}

func TestHandler_List_ServiceError(t *testing.T) {
	groups := []*models.Group{
		{
			Name:     "name",
			Hostname: "hostname",
		},
	}

	service := &mockDatabaseService{}
	errorMessage := "some error"
	service.
		On("List", groups).
		Return(nil, errors.New(errorMessage))

	handler := New(nil, service, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user", &models.User{Groups: groups})

	handler.List(c)

	assert.Empty(t, w.Body.Bytes())
	assert.NotEmpty(t, c.Errors)

	service.AssertExpectations(t)
}

type mockDatabaseService struct{ mock.Mock }

func (m *mockDatabaseService) Create(d *model.Database) error {
	panic("implement me")
}

func (m *mockDatabaseService) Copy(id uint, d *model.Database, group *models.Group) error {
	panic("implement me")
}

func (m *mockDatabaseService) FindById(id uint) (*model.Database, error) {
	panic("implement me")
}

func (m *mockDatabaseService) Lock(id uint, instanceId uint, userId uint) (*model.Lock, error) {
	panic("implement me")
}

func (m *mockDatabaseService) Unlock(id uint) error {
	panic("implement me")
}

func (m *mockDatabaseService) Upload(d *model.Database, group *models.Group, file io.Reader) (*model.Database, error) {
	panic("implement me")
}

func (m *mockDatabaseService) Download(id uint, dst io.Writer, headers func(contentLength int64)) error {
	panic("implement me")
}

func (m *mockDatabaseService) Delete(id uint) error {
	panic("implement me")
}

func (m *mockDatabaseService) List(groups []*models.Group) ([]*model.Database, error) {
	called := m.Called(groups)
	databases, ok := called.Get(0).([]*model.Database)
	if ok {
		return databases, nil
	} else {
		return nil, called.Error(1)
	}
}

func (m *mockDatabaseService) Update(d *model.Database) error {
	panic("implement me")
}

func (m *mockDatabaseService) CreateExternalDownload(databaseID uint, expiration time.Time) (model.ExternalDownload, error) {
	panic("implement me")
}

func (m *mockDatabaseService) FindExternalDownload(uuid uuid.UUID) (model.ExternalDownload, error) {
	panic("implement me")
}

func (m *mockDatabaseService) SaveAs(token string, database *model.Database, instance *instanceModels.Instance, stack *instanceModels.Stack, newName string, format string) (*model.Database, error) {
	panic("implement me")
}
