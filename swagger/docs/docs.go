package docs

import "github.com/dhis2-sre/im-database-manager/pkg/database"

// swagger:response
type Error struct {
	// The error message
	//in: body
	Message string
}

//swagger:parameters findDatabaseById lockDatabaseById unlockDatabaseById uploadDatabase downloadDatabase deleteDatabaseById updateDatabaseById saveDatabaseById findDatabaseUrlById
type IdParam struct {
	// in: path
	// required: true
	ID uint `json:"id"`
}

// swagger:parameters createDatabase
type _ struct {
	// Create database request body parameter
	// in: body
	// required: true
	Body database.CreateDatabaseRequest
}

// swagger:parameters lockDatabaseById unlockDatabaseById
type _ struct {
	// Lock/unlock database request body parameter
	// in: body
	// required: true
	Body database.LockDatabaseRequest
}

// swagger:parameters uploadDatabase
type _ struct {
	// Upload database request body parameter
	// in: formData
	// required: true
	// swagger:file
	File database.UploadDatabaseRequest
}

// swagger:parameters updateDatabaseById
type _ struct {
	// Update database request body parameter
	// in: body
	// required: true
	Body database.UpdateDatabaseRequest
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
