// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dhis2-sre/im-database-manager/swagger/sdk/models"
)

// FindDatabaseByIDReader is a Reader for the FindDatabaseByID structure.
type FindDatabaseByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindDatabaseByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindDatabaseByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFindDatabaseByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFindDatabaseByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindDatabaseByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindDatabaseByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewFindDatabaseByIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindDatabaseByIDOK creates a FindDatabaseByIDOK with default headers values
func NewFindDatabaseByIDOK() *FindDatabaseByIDOK {
	return &FindDatabaseByIDOK{}
}

/* FindDatabaseByIDOK describes a response with status code 200, with default header values.

Database
*/
type FindDatabaseByIDOK struct {
	Payload *models.Database
}

func (o *FindDatabaseByIDOK) Error() string {
	return fmt.Sprintf("[GET /databases/{id}][%d] findDatabaseByIdOK  %+v", 200, o.Payload)
}
func (o *FindDatabaseByIDOK) GetPayload() *models.Database {
	return o.Payload
}

func (o *FindDatabaseByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Database)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindDatabaseByIDBadRequest creates a FindDatabaseByIDBadRequest with default headers values
func NewFindDatabaseByIDBadRequest() *FindDatabaseByIDBadRequest {
	return &FindDatabaseByIDBadRequest{}
}

/* FindDatabaseByIDBadRequest describes a response with status code 400, with default header values.

FindDatabaseByIDBadRequest find database by Id bad request
*/
type FindDatabaseByIDBadRequest struct {
}

func (o *FindDatabaseByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /databases/{id}][%d] findDatabaseByIdBadRequest ", 400)
}

func (o *FindDatabaseByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindDatabaseByIDUnauthorized creates a FindDatabaseByIDUnauthorized with default headers values
func NewFindDatabaseByIDUnauthorized() *FindDatabaseByIDUnauthorized {
	return &FindDatabaseByIDUnauthorized{}
}

/* FindDatabaseByIDUnauthorized describes a response with status code 401, with default header values.

FindDatabaseByIDUnauthorized find database by Id unauthorized
*/
type FindDatabaseByIDUnauthorized struct {
}

func (o *FindDatabaseByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /databases/{id}][%d] findDatabaseByIdUnauthorized ", 401)
}

func (o *FindDatabaseByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindDatabaseByIDForbidden creates a FindDatabaseByIDForbidden with default headers values
func NewFindDatabaseByIDForbidden() *FindDatabaseByIDForbidden {
	return &FindDatabaseByIDForbidden{}
}

/* FindDatabaseByIDForbidden describes a response with status code 403, with default header values.

FindDatabaseByIDForbidden find database by Id forbidden
*/
type FindDatabaseByIDForbidden struct {
}

func (o *FindDatabaseByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /databases/{id}][%d] findDatabaseByIdForbidden ", 403)
}

func (o *FindDatabaseByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindDatabaseByIDNotFound creates a FindDatabaseByIDNotFound with default headers values
func NewFindDatabaseByIDNotFound() *FindDatabaseByIDNotFound {
	return &FindDatabaseByIDNotFound{}
}

/* FindDatabaseByIDNotFound describes a response with status code 404, with default header values.

FindDatabaseByIDNotFound find database by Id not found
*/
type FindDatabaseByIDNotFound struct {
}

func (o *FindDatabaseByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /databases/{id}][%d] findDatabaseByIdNotFound ", 404)
}

func (o *FindDatabaseByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindDatabaseByIDUnsupportedMediaType creates a FindDatabaseByIDUnsupportedMediaType with default headers values
func NewFindDatabaseByIDUnsupportedMediaType() *FindDatabaseByIDUnsupportedMediaType {
	return &FindDatabaseByIDUnsupportedMediaType{}
}

/* FindDatabaseByIDUnsupportedMediaType describes a response with status code 415, with default header values.

FindDatabaseByIDUnsupportedMediaType find database by Id unsupported media type
*/
type FindDatabaseByIDUnsupportedMediaType struct {
}

func (o *FindDatabaseByIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /databases/{id}][%d] findDatabaseByIdUnsupportedMediaType ", 415)
}

func (o *FindDatabaseByIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
