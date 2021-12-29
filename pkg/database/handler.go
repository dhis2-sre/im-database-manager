package database

import (
	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-database-manager/pkg/model"
	userClient "github.com/dhis2-sre/im-user/pkg/client"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/gin-gonic/gin"
	"mime/multipart"
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

type createDatabaseRequest struct {
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
	var request createDatabaseRequest
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

type lockDatabaseRequest struct {
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

	var request lockDatabaseRequest
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

type uploadDatabaseRequest struct {
	Database *multipart.FileHeader `form:"database" binding:"required"`
}

// Upload database
// swagger:route POST /databases/{id}/upload uploadDatabase
//
// Upload database
//
// Security:
//  oauth2:
//
// responses:
//   201: Database
//   401: Error
//   403: Error
//   415: Error
func (h Handler) Upload(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	var request uploadDatabaseRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	// TODO: Authorize

	if request.Database == nil {
		badRequest := apperror.NewBadRequest("file not found")
		_ = c.Error(badRequest)
		return
	}

	save, err := h.databaseService.Upload(d, request.Database)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, save)
}

// Delete database
// swagger:route DELETE /databases/{id} deleteDatabaseById
//
// Delete database by id
//
// Security:
//  oauth2:
//
// responses:
//   202:
//   401: Error
//   403: Error
//   415: Error
func (h Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	// TODO: Authorize

	err = h.databaseService.Delete(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusAccepted)
}

type groupWithDatabases struct {
	ID        uint
	Name      string
	Hostname  string
	Databases []*model.Database
}

// List databases
// swagger:route GET /databases listDatabases
//
// List databases
//
// Security:
//  oauth2:
//
// responses:
//   200: []Database
//   401: Error
//   403: Error
//   415: Error
func (h Handler) List(c *gin.Context) {
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

	groups := userWithGroups.Groups
	d, err := h.databaseService.List(groups)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, h.groupsWithInstances(groups, d))
}

func (h Handler) groupsWithInstances(groups []*models.Group, databases []*model.Database) []groupWithDatabases {
	groupsWithDatabases := make([]groupWithDatabases, len(groups))
	for i, group := range groups {
		groupsWithDatabases[i].ID = uint(group.ID)
		groupsWithDatabases[i].Name = group.Name
		groupsWithDatabases[i].Hostname = group.Hostname
		groupsWithDatabases[i].Databases = h.filterByGroupId(databases, func(instance *model.Database) bool {
			return instance.GroupID == uint(group.ID)
		})
	}
	return groupsWithDatabases
}

func (h Handler) filterByGroupId(databases []*model.Database, test func(instance *model.Database) bool) (ret []*model.Database) {
	for _, database := range databases {
		if test(database) {
			ret = append(ret, database)
		}
	}
	return
}

type updateDatabaseRequest struct {
	Name    string `json:"name" binding:"required"`
	GroupId uint   `json:"groupId" binding:"required"`
}

// Update database
// swagger:route PUT /databases/{id} updateDatabaseById
//
// Update database by id
//
// Security:
//  oauth2:
//
// responses:
//   200:
//   401: Error
//   403: Error
//   415: Error
func (h Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	// TODO: Authorize

	var request createDatabaseRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	d.Name = request.Name
	d.GroupID = request.GroupId

	err = h.databaseService.Update(d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, d)
}
