package docs

// swagger:response
type Error struct {
	// The error message
	//in: body
	Message string
}

// swagger:parameters findDatabaseById lockDatabaseById
type IdParam struct {
	// in: path
	// required: true
	ID uint `json:"id"`
}
