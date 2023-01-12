package database

import (
	"github.com/dhis2-sre/im-database-manager/pkg/config"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_service_Create(t *testing.T) {
	c := config.Config{}

	database := &model.Database{}
	repository := &mockRepository{}
	repository.
		On("Create", database).
		Return(nil)
	service := NewService(c, nil, nil, repository)

	err := service.Create(database)

	require.NoError(t, err)

	repository.AssertExpectations(t)
}

type mockRepository struct{ mock.Mock }

func (m *mockRepository) Create(d *model.Database) error {
	called := m.Called(d)
	return called.Error(0)
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
	panic("implement me")
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
