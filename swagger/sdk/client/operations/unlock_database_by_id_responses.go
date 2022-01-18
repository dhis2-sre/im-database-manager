// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UnlockDatabaseByIDReader is a Reader for the UnlockDatabaseByID structure.
type UnlockDatabaseByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnlockDatabaseByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUnlockDatabaseByIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUnlockDatabaseByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUnlockDatabaseByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnlockDatabaseByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewUnlockDatabaseByIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnlockDatabaseByIDAccepted creates a UnlockDatabaseByIDAccepted with default headers values
func NewUnlockDatabaseByIDAccepted() *UnlockDatabaseByIDAccepted {
	return &UnlockDatabaseByIDAccepted{}
}

/* UnlockDatabaseByIDAccepted describes a response with status code 202, with default header values.

UnlockDatabaseByIDAccepted unlock database by Id accepted
*/
type UnlockDatabaseByIDAccepted struct {
}

func (o *UnlockDatabaseByIDAccepted) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}/lock][%d] unlockDatabaseByIdAccepted ", 202)
}

func (o *UnlockDatabaseByIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlockDatabaseByIDUnauthorized creates a UnlockDatabaseByIDUnauthorized with default headers values
func NewUnlockDatabaseByIDUnauthorized() *UnlockDatabaseByIDUnauthorized {
	return &UnlockDatabaseByIDUnauthorized{}
}

/* UnlockDatabaseByIDUnauthorized describes a response with status code 401, with default header values.

UnlockDatabaseByIDUnauthorized unlock database by Id unauthorized
*/
type UnlockDatabaseByIDUnauthorized struct {
}

func (o *UnlockDatabaseByIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}/lock][%d] unlockDatabaseByIdUnauthorized ", 401)
}

func (o *UnlockDatabaseByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlockDatabaseByIDForbidden creates a UnlockDatabaseByIDForbidden with default headers values
func NewUnlockDatabaseByIDForbidden() *UnlockDatabaseByIDForbidden {
	return &UnlockDatabaseByIDForbidden{}
}

/* UnlockDatabaseByIDForbidden describes a response with status code 403, with default header values.

UnlockDatabaseByIDForbidden unlock database by Id forbidden
*/
type UnlockDatabaseByIDForbidden struct {
}

func (o *UnlockDatabaseByIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}/lock][%d] unlockDatabaseByIdForbidden ", 403)
}

func (o *UnlockDatabaseByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlockDatabaseByIDNotFound creates a UnlockDatabaseByIDNotFound with default headers values
func NewUnlockDatabaseByIDNotFound() *UnlockDatabaseByIDNotFound {
	return &UnlockDatabaseByIDNotFound{}
}

/* UnlockDatabaseByIDNotFound describes a response with status code 404, with default header values.

UnlockDatabaseByIDNotFound unlock database by Id not found
*/
type UnlockDatabaseByIDNotFound struct {
}

func (o *UnlockDatabaseByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}/lock][%d] unlockDatabaseByIdNotFound ", 404)
}

func (o *UnlockDatabaseByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlockDatabaseByIDUnsupportedMediaType creates a UnlockDatabaseByIDUnsupportedMediaType with default headers values
func NewUnlockDatabaseByIDUnsupportedMediaType() *UnlockDatabaseByIDUnsupportedMediaType {
	return &UnlockDatabaseByIDUnsupportedMediaType{}
}

/* UnlockDatabaseByIDUnsupportedMediaType describes a response with status code 415, with default header values.

UnlockDatabaseByIDUnsupportedMediaType unlock database by Id unsupported media type
*/
type UnlockDatabaseByIDUnsupportedMediaType struct {
}

func (o *UnlockDatabaseByIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}/lock][%d] unlockDatabaseByIdUnsupportedMediaType ", 415)
}

func (o *UnlockDatabaseByIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}