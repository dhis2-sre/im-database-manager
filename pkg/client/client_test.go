package client

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dhis2-sre/im-database-manager/pkg/storage"

	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/internal/middleware"
	"github.com/dhis2-sre/im-database-manager/pkg/database"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	instanceModels "github.com/dhis2-sre/im-manager/swagger/sdk/models"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestFindDatabaseById(t *testing.T) {
	databaseService := &mockDatabaseService{}
	databaseService.
		On("FindById", uint(1)).
		Return(&model.Database{
			Model: gorm.Model{ID: 1},
		}, nil)
	r := gin.Default()
	user := &models.User{
		Groups: []*models.Group{
			{Name: handler.AdministratorGroupName},
		},
	}
	r.Use(middleware.ErrorHandler(), mockTokenAuthentication(user))
	h := database.New(nil, databaseService, nil)
	r.GET("/databases/:id", h.FindById)
	srv := httptest.NewServer(r)
	defer srv.Close()
	host := strings.Replace(srv.URL, "http://", "", 1)
	cli := New(host, "")

	db, err := cli.FindById("token", 1)

	require.NoError(t, err)
	assert.Equal(t, uint64(1), db.ID)
	databaseService.AssertExpectations(t)
}

type mockDatabaseService struct{ mock.Mock }

func (m *mockDatabaseService) Create(d *model.Database) error {
	panic("implement me")
}

func (m *mockDatabaseService) Copy(id uint, d *model.Database, group *models.Group) error {
	panic("implement me")
}

func (m *mockDatabaseService) FindById(id uint) (*model.Database, error) {
	called := m.Called(id)
	d, ok := called.Get(0).(*model.Database)
	if ok {
		return d, nil
	} else {
		return nil, called.Error(1)
	}
}

func (m *mockDatabaseService) Lock(id uint, instanceId uint, userId uint) (*model.Lock, error) {
	panic("implement me")
}

func (m *mockDatabaseService) Unlock(id uint) error {
	panic("implement me")
}

func (m *mockDatabaseService) Upload(d *model.Database, group *models.Group, file storage.ReadAtSeeker, length int64) (*model.Database, error) {
	panic("implement me")
}

func (m *mockDatabaseService) Download(id uint, dst io.Writer, headers func(contentLength int64)) error {
	panic("implement me")
}

func (m *mockDatabaseService) Delete(id uint) error {
	panic("implement me")
}

func (m *mockDatabaseService) List(groups []*models.Group) ([]*model.Database, error) {
	panic("implement me")
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

func mockTokenAuthentication(user *models.User) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("user", user)
	}
}
