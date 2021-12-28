package database

import (
	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	userClient "github.com/dhis2-sre/im-user/pkg/client"
	"github.com/gin-gonic/gin"
	"net/http"
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
