package database

import (
	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	userClient "github.com/dhis2-sre/im-user/pkg/client"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ProvideHandler(userClient userClient.Client, databaseService Service) Handler {
	return Handler{
		userClient,
		databaseService,
	}
}

type Handler struct {
	userClient      userClient.Client
	databaseService Service
}

type CreateDatabaseRequest struct {
	Name    string `json:"name" binding:"required"`
	GroupId uint   `json:"groupId" binding:"required"`
}

// Create database
// swagger:route POST /databases createDatabase
//
// Create database
//
// Security:
//  oauth2:
//
// responses:
//   202: Database
//   401: Error
//   403: Error
//   415: Error
func (h Handler) Create(c *gin.Context) {
	var request CreateDatabaseRequest

	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	// TODO: Authorize
	/*
		user, err := handler.GetUserFromContext(c)
		if err != nil {
			_ = c.Error(err)
			return
		}

		token, err := handler.GetTokenFromHttpAuthHeader(c)
		if err != nil {
			_ = c.Error(err)
			return
		}

		userWithGroups, err := h.userClient.FindUserById(token, user.ID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		canWrite := handler.CanWriteInstance(userWithGroups, instance)
		if !canWrite {
			unauthorized := apperror.NewUnauthorized("Write access denied")
			_ = c.Error(unauthorized)
			return
		}
	*/

	d := &model.Database{
		Name:    request.Name,
		GroupID: request.GroupId,
	}

	if err := h.databaseService.Create(d); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, d)
}

// FindById database
// swagger:route GET /databases/{id} findDatabaseById
//
// Find database by id
//
// Security:
//  oauth2:
//
// responses:
//   200: Database
//   401: Error
//   403: Error
//   415: Error
func (h Handler) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	// TODO: Authorize

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, d)
}

// LockDatabaseRequest ...
type LockDatabaseRequest struct {
	InstanceId uint `json:"instanceId" binding:"required"`
}

// Lock database
// swagger:route POST /databases/{id}/lock lockDatabaseById
//
// Lock database by id
//
// Security:
//  oauth2:
//
// responses:
//   200: Database
//   401: Error
//   403: Error
//   409: Error
//   415: Error
func (h Handler) Lock(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	var request LockDatabaseRequest

	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	// TODO: Authorize

	d, err := h.databaseService.Lock(uint(id), request.InstanceId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, d)
}

// Unlock database
// swagger:route DELETE /databases/{id}/lock unlockDatabaseById
//
// Unlock database by id
//
// Security:
//  oauth2:
//
// responses:
//   202:
//   401: Error
//   403: Error
//   415: Error
func (h Handler) Unlock(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	// TODO: Authorize

	err = h.databaseService.Unlock(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusAccepted)
}
