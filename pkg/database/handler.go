package database

import (
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/dhis2-sre/im-database-manager/pkg/model"

	"github.com/dhis2-sre/im-database-manager/internal/apperror"
	"github.com/dhis2-sre/im-database-manager/internal/handler"
	"github.com/dhis2-sre/im-user/swagger/sdk/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userClient      userClientHandler
	databaseService Service
}

func New(userClient userClientHandler, databaseService Service) Handler {
	return Handler{
		userClient,
		databaseService,
	}
}

type userClientHandler interface {
	FindGroupByName(token string, name string) (*models.Group, error)
	FindUserById(token string, id uint) (*models.User, error)
}

type UploadDatabaseRequest struct {
	Group    string                `form:"group" binding:"required"`
	Database *multipart.FileHeader `form:"database" binding:"required"`
}

// Upload database
// swagger:route POST /databases uploadDatabase
//
// Upload database
//
// Security:
//   oauth2:
//
// responses:
//   201: Database
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Upload(c *gin.Context) {
	var request UploadDatabaseRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	if request.Database == nil {
		badRequest := apperror.NewBadRequest("file not found")
		_ = c.Error(badRequest)
		return
	}

	d := &model.Database{
		Name:      request.Database.Filename,
		GroupName: request.Group,
	}

	err := h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	token, err := handler.GetTokenFromHttpAuthHeader(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	group, err := h.userClient.FindGroupByName(token, d.GroupName)
	if err != nil {
		_ = c.Error(err)
		return
	}

	file, err := request.Database.Open()
	if err != nil {
		_ = c.Error(err)
		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			_ = c.Error(err)
			return
		}
	}(file)

	save, err := h.databaseService.Upload(d, group, file)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, save)
}

type CopyDatabaseRequest struct {
	Name  string `json:"name" binding:"required"`
	Group string `json:"group" binding:"required"`
}

// Copy database
// swagger:route POST /databases/{id}/copy copyDatabase
//
// Copy database
//
// Security:
//   oauth2:
//
// responses:
//   202: Database
//   401: Error
//   403: Error
//   415: Error
func (h Handler) Copy(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	var request CopyDatabaseRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	d := &model.Database{
		Name:      request.Name,
		GroupName: request.Group,
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	token, err := handler.GetTokenFromHttpAuthHeader(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	group, err := h.userClient.FindGroupByName(token, d.GroupName)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if err := h.databaseService.Copy(uint(id), d, group); err != nil {
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
//   oauth2:
//
// responses:
//   200: Database
//   400: Error
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, d)
}

// FindUrlById database
// swagger:route GET /databases/{id}/url findDatabaseUrlById
//
// Find database URL by id
//
// responses:
//   200: DatabaseUrl
//   400: Error
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) FindUrlById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, d.Url)
}

type LockDatabaseRequest struct {
	InstanceId uint `json:"instanceId" binding:"required"`
}

// Lock database
// swagger:route POST /databases/{id}/lock lockDatabaseById
//
// Lock database by id
//
// Security:
//   oauth2:
//
// responses:
//   200: Database
//   401: Error
//   403: Error
//   404: Error
//   409: Error
//   415: Error
func (h Handler) Lock(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
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

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	d, err = h.databaseService.Lock(uint(id), request.InstanceId)
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
//   oauth2:
//
// responses:
//   202:
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Unlock(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	userWithGroups, err := h.getUserWithGroups(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	// This is a bit hacky. All other handlers are using the h.canAccess method only group admins can unlock (admins can't)
	isGroupAdministrator := handler.IsGroupAdministrator(d.GroupName, userWithGroups.AdminGroups)
	if !isGroupAdministrator {
		forbidden := apperror.NewForbidden("access denied")
		_ = c.Error(forbidden)
		return
	}

	err = h.databaseService.Unlock(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusAccepted)
}

/*
type SaveDatabaseRequest struct {
	InstanceId uint `json:"instanceId" binding:"required"`
}

type SaveDatabaseResponse struct {
	RunId string `json:"runId"`
}

// Save database
// swagger:route POST /databases/{id}/save saveDatabaseById
//
// Save database by id
//
// Security:
//   oauth2:
//
// responses:
//   202:
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Save(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	var request SaveDatabaseRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	token, err := handler.GetTokenFromHttpAuthHeader(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if d.InstanceID != request.InstanceId && d.InstanceID != 0 {
		message := fmt.Sprintf("database locked by instance with id %d", d.InstanceID)
		conflict := apperror.NewConflict(message)
		_ = c.Error(conflict)
		return
	}

	runId, err := h.databaseService.Save(token, uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, SaveDatabaseResponse{runId})
}
*/

// Download database
// swagger:route GET /databases/{id}/download downloadDatabase
//
// Download database
//
// Security:
//   oauth2:
//
// responses:
//   200: DownloadDatabaseResponse
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Download(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	_, file := path.Split(d.Url)
	c.Header("Content-Disposition", "attachment; filename="+file)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Type", "application/octet-stream")

	err = h.databaseService.Download(uint(id), c.Writer, func(contentLength int64) {
		c.Header("Content-Length", strconv.FormatInt(contentLength, 10))
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Delete database
// swagger:route DELETE /databases/{id} deleteDatabaseById
//
// Delete database by id
//
// Security:
//   oauth2:
//
// responses:
//   202:
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.databaseService.Delete(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusAccepted)
}

// swagger:model GroupsWithDatabases
type GroupsWithDatabases struct {
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
//   oauth2:
//
// responses:
//   200: []GroupsWithDatabases
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

	c.JSON(http.StatusOK, h.groupsWithDatabases(groups, d))
}

func (h Handler) groupsWithDatabases(groups []*models.Group, databases []*model.Database) []GroupsWithDatabases {
	groupsWithDatabases := make([]GroupsWithDatabases, len(groups))
	for i, group := range groups {
		groupsWithDatabases[i].Name = group.Name
		groupsWithDatabases[i].Hostname = group.Hostname
		groupsWithDatabases[i].Databases = h.filterByGroupId(databases, func(instance *model.Database) bool {
			return instance.GroupName == group.Name
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

// TODO: Make all? properties optional... ensure we only update if a new value is present
type UpdateDatabaseRequest struct {
	Name      string `json:"name" binding:"required"`
	GroupName string `json:"groupName" binding:"required"`
}

// Update database
// swagger:route PUT /databases/{id} updateDatabaseById
//
// Update database by id
//
// TODO: Race condition? If two clients request at the same time... Do we need a transaction between find and update
//
// Security:
//   oauth2:
//
// TODO: document return type
// responses:
//   200:
//   401: Error
//   403: Error
//   415: Error
func (h Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	var request UpdateDatabaseRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	d.Name = request.Name
	d.GroupName = request.GroupName

	err = h.databaseService.Update(d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, d)
}

func (h Handler) getUserWithGroups(c *gin.Context) (*models.User, error) {
	user, err := handler.GetUserFromContext(c)
	if err != nil {
		return nil, err
	}

	token, err := handler.GetTokenFromHttpAuthHeader(c)
	if err != nil {
		return nil, err
	}

	userWithGroups, err := h.userClient.FindUserById(token, user.ID)
	if err != nil {
		return nil, err
	}

	return userWithGroups, nil
}

func (h Handler) canAccess(c *gin.Context, d *model.Database) error {
	userWithGroups, err := h.getUserWithGroups(c)
	if err != nil {
		return err
	}

	can := handler.CanAccess(userWithGroups, d)
	if !can {
		return apperror.NewForbidden("access denied")
	}

	return nil
}

type CreateExternalDatabaseRequest struct {
	Expiration time.Time `json:"expiration" binding:"required"`
}

// CreateExternalDownload
// swagger:route POST /databases/{id}/external createExternalDownloadDatabase
//
// Create external database download
//
// Security:
//   oauth2:
//
// responses:
//   200: CreateExternalDownloadResponse
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) CreateExternalDownload(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		badRequest := apperror.NewBadRequest("error parsing id")
		_ = c.Error(badRequest)
		return
	}

	var request CreateExternalDatabaseRequest
	if err := handler.DataBinder(c, &request); err != nil {
		_ = c.Error(err)
		return
	}

	d, err := h.databaseService.FindById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.canAccess(c, d)
	if err != nil {
		_ = c.Error(err)
		return
	}

	externalDownload, err := h.databaseService.CreateExternalDownload(d.ID, request.Expiration)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, externalDownload)
}

// ExternalDownload database
// swagger:route GET /databases/external/{uuid} externalDownloadDatabase
//
// Download database
//
// Security:
//   oauth2:
//
// responses:
//   200: DownloadDatabaseResponse
//   401: Error
//   403: Error
//   404: Error
//   415: Error
func (h Handler) ExternalDownload(c *gin.Context) {
	uuidParam := c.Param("uuid")
	if uuidParam == "" {
		badRequest := apperror.NewBadRequest("error missing uuid")
		_ = c.Error(badRequest)
		return
	}

	id, err := uuid.Parse(uuidParam)
	if err != nil {
		_ = c.Error(err)
		return
	}

	download, err := h.databaseService.FindExternalDownload(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	d, err := h.databaseService.FindById(download.DatabaseID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	_, file := path.Split(d.Url)
	c.Header("Content-Disposition", "attachment; filename="+file)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Type", "application/octet-stream")

	err = h.databaseService.Download(d.ID, c.Writer, func(contentLength int64) {
		c.Header("Content-Length", strconv.FormatInt(contentLength, 10))
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
}
