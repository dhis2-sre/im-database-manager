package client

import (
	"context"
	"errors"
	"strconv"

	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/swagger/sdk/client/operations"
	"github.com/dhis2-sre/im-database-manager/swagger/sdk/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Client interface {
	FindById(token string, id uint) (*models.Database, error)
}

type cli struct {
	clientService operations.ClientService
}

func New(host string, basePath string) *cli {
	transport := httptransport.New(host, basePath, nil)
	userService := operations.New(transport, strfmt.Default)
	return &cli{userService}
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
