package database

import "github.com/dhis2-sre/im-database-manager/pkg/model"

// swagger:response
type Error struct {
	// The error message
	//in: body
	Message string
}

//swagger:parameters findDatabaseById lockDatabaseById unlockDatabaseById uploadDatabase downloadDatabase deleteDatabaseById updateDatabaseById saveDatabaseById findDatabaseUrlById createExternalDownloadDatabase
type _ struct {
	// in: path
	// required: true
	ID uint `json:"id"`
}

// swagger:parameters createDatabase
type _ struct {
	// Create database request body parameter
	// in: body
	// required: true
	Body CreateDatabaseRequest
}

// swagger:parameters lockDatabaseById unlockDatabaseById
type _ struct {
	// Lock/unlock database request body parameter
	// in: body
	// required: true
	Body LockDatabaseRequest
}

// swagger:parameters uploadDatabase
type _ struct {
	// Upload database request body parameter
	// in: formData
	// required: true
	// swagger:file
	File UploadDatabaseRequest
}

// swagger:parameters updateDatabaseById
type _ struct {
	// Update database request body parameter
	// in: body
	// required: true
	Body UpdateDatabaseRequest
}

/*
// swagger:parameters saveDatabaseById
type _ struct {
	// Save database request body parameter
	// in: body
	// required: true
	Body database.SaveDatabaseRequest
}

// swagger:response
type SaveDatabaseResponse struct {
	//in: body
	Body database.SaveDatabaseResponse
}
*/

// swagger:response DatabaseUrl
type _ struct {
	//in: body
	_ string
}

//swagger:parameters externalDownloadDatabase
type _ struct {
	// in: path
	// required: true
	UUID uint `json:"uuid"`
}

// swagger:response DownloadDatabaseResponse
type _ struct {
	//in: body
	_ string
}

// swagger:response CreateExternalDownloadResponse
type _ struct {
	//in: body
	_ model.ExternalDownload
}
