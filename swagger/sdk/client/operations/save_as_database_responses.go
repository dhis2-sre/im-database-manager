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

// SaveAsDatabaseReader is a Reader for the SaveAsDatabase structure.
type SaveAsDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveAsDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSaveAsDatabaseCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSaveAsDatabaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSaveAsDatabaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSaveAsDatabaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewSaveAsDatabaseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSaveAsDatabaseCreated creates a SaveAsDatabaseCreated with default headers values
func NewSaveAsDatabaseCreated() *SaveAsDatabaseCreated {
	return &SaveAsDatabaseCreated{}
}

/* SaveAsDatabaseCreated describes a response with status code 201, with default header values.

Database
*/
type SaveAsDatabaseCreated struct {
	Payload *models.Database
}

func (o *SaveAsDatabaseCreated) Error() string {
	return fmt.Sprintf("[POST /databases/save-as/{instanceId}][%d] saveAsDatabaseCreated  %+v", 201, o.Payload)
}
func (o *SaveAsDatabaseCreated) GetPayload() *models.Database {
	return o.Payload
}

func (o *SaveAsDatabaseCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Database)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveAsDatabaseUnauthorized creates a SaveAsDatabaseUnauthorized with default headers values
func NewSaveAsDatabaseUnauthorized() *SaveAsDatabaseUnauthorized {
	return &SaveAsDatabaseUnauthorized{}
}

/* SaveAsDatabaseUnauthorized describes a response with status code 401, with default header values.

SaveAsDatabaseUnauthorized save as database unauthorized
*/
type SaveAsDatabaseUnauthorized struct {
}

func (o *SaveAsDatabaseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /databases/save-as/{instanceId}][%d] saveAsDatabaseUnauthorized ", 401)
}

func (o *SaveAsDatabaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveAsDatabaseForbidden creates a SaveAsDatabaseForbidden with default headers values
func NewSaveAsDatabaseForbidden() *SaveAsDatabaseForbidden {
	return &SaveAsDatabaseForbidden{}
}

/* SaveAsDatabaseForbidden describes a response with status code 403, with default header values.

SaveAsDatabaseForbidden save as database forbidden
*/
type SaveAsDatabaseForbidden struct {
}

func (o *SaveAsDatabaseForbidden) Error() string {
	return fmt.Sprintf("[POST /databases/save-as/{instanceId}][%d] saveAsDatabaseForbidden ", 403)
}

func (o *SaveAsDatabaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveAsDatabaseNotFound creates a SaveAsDatabaseNotFound with default headers values
func NewSaveAsDatabaseNotFound() *SaveAsDatabaseNotFound {
	return &SaveAsDatabaseNotFound{}
}

/* SaveAsDatabaseNotFound describes a response with status code 404, with default header values.

SaveAsDatabaseNotFound save as database not found
*/
type SaveAsDatabaseNotFound struct {
}

func (o *SaveAsDatabaseNotFound) Error() string {
	return fmt.Sprintf("[POST /databases/save-as/{instanceId}][%d] saveAsDatabaseNotFound ", 404)
}

func (o *SaveAsDatabaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveAsDatabaseUnsupportedMediaType creates a SaveAsDatabaseUnsupportedMediaType with default headers values
func NewSaveAsDatabaseUnsupportedMediaType() *SaveAsDatabaseUnsupportedMediaType {
	return &SaveAsDatabaseUnsupportedMediaType{}
}

/* SaveAsDatabaseUnsupportedMediaType describes a response with status code 415, with default header values.

SaveAsDatabaseUnsupportedMediaType save as database unsupported media type
*/
type SaveAsDatabaseUnsupportedMediaType struct {
}

func (o *SaveAsDatabaseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /databases/save-as/{instanceId}][%d] saveAsDatabaseUnsupportedMediaType ", 415)
}

func (o *SaveAsDatabaseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}