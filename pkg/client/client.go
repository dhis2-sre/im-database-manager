package client

import (
	"context"
	"errors"
	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/swagger/sdk/client/operations"
	"github.com/dhis2-sre/im-database-manager/swagger/sdk/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"strconv"
)

type Client interface {
	Create(token string, createDatabaseRequest *models.CreateDatabaseRequest) (*models.Database, error)
	FindById(token string, id uint) (*models.Database, error)
	Save(token string, databaseId uint, instanceId uint) (string, error)
}

func ProvideClient(host string, basePath string) Client {
	transport := httptransport.New(host, basePath, nil)
	userService := operations.New(transport, strfmt.Default)
	return &cli{userService}
}

type cli struct {
	clientService operations.ClientService
}

func (c cli) Create(token string, createDatabaseRequest *models.CreateDatabaseRequest) (*models.Database, error) {
	params := &operations.CreateDatabaseParams{
		Body:    createDatabaseRequest,
		Context: context.Background(),
	}
	authInfo := httptransport.BearerToken(token)
	db, err := c.clientService.CreateDatabase(params, authInfo)
	if err != nil {
		return nil, err
	}
	return db.GetPayload(), nil
}

func (c cli) FindById(token string, id uint) (*models.Database, error) {
	params := &operations.FindDatabaseByIDParams{ID: uint64(id), Context: context.Background()}
	authInfo := httptransport.BearerToken(token)
	db, err := c.clientService.FindDatabaseByID(params, authInfo)
	if err != nil {
		var e *operations.FindDatabaseByIDNotFound
		if errors.As(err, &e) {
			return nil, apperror.NewNotFound("database", strconv.FormatUint(uint64(id), 10))
		}
		return nil, err
	}
	return db.GetPayload(), nil
}

func (c cli) Save(token string, databaseId uint, instanceId uint) (string, error) {
	body := &models.SaveDatabaseRequest{InstanceID: uint64(instanceId)}
	params := &operations.SaveDatabaseByIDParams{
		Body:    body,
		ID:      uint64(databaseId),
		Context: context.Background(),
	}
	authInfo := httptransport.BearerToken(token)

	response, err := c.clientService.SaveDatabaseByID(params, authInfo)
	if err != nil {
		return "", err
	}

	return response.Payload.RunID, nil
}
