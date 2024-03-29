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

// UploadDatabaseReader is a Reader for the UploadDatabase structure.
type UploadDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUploadDatabaseCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUploadDatabaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadDatabaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadDatabaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewUploadDatabaseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadDatabaseCreated creates a UploadDatabaseCreated with default headers values
func NewUploadDatabaseCreated() *UploadDatabaseCreated {
	return &UploadDatabaseCreated{}
}

/*
UploadDatabaseCreated describes a response with status code 201, with default header values.

UploadDatabaseCreated upload database created
*/
type UploadDatabaseCreated struct {
	Payload *models.Database
}

// IsSuccess returns true when this upload database created response has a 2xx status code
func (o *UploadDatabaseCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload database created response has a 3xx status code
func (o *UploadDatabaseCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload database created response has a 4xx status code
func (o *UploadDatabaseCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload database created response has a 5xx status code
func (o *UploadDatabaseCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this upload database created response a status code equal to that given
func (o *UploadDatabaseCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the upload database created response
func (o *UploadDatabaseCreated) Code() int {
	return 201
}

func (o *UploadDatabaseCreated) Error() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseCreated  %+v", 201, o.Payload)
}

func (o *UploadDatabaseCreated) String() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseCreated  %+v", 201, o.Payload)
}

func (o *UploadDatabaseCreated) GetPayload() *models.Database {
	return o.Payload
}

func (o *UploadDatabaseCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Database)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadDatabaseUnauthorized creates a UploadDatabaseUnauthorized with default headers values
func NewUploadDatabaseUnauthorized() *UploadDatabaseUnauthorized {
	return &UploadDatabaseUnauthorized{}
}

/*
UploadDatabaseUnauthorized describes a response with status code 401, with default header values.

UploadDatabaseUnauthorized upload database unauthorized
*/
type UploadDatabaseUnauthorized struct {
}

// IsSuccess returns true when this upload database unauthorized response has a 2xx status code
func (o *UploadDatabaseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload database unauthorized response has a 3xx status code
func (o *UploadDatabaseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload database unauthorized response has a 4xx status code
func (o *UploadDatabaseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload database unauthorized response has a 5xx status code
func (o *UploadDatabaseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this upload database unauthorized response a status code equal to that given
func (o *UploadDatabaseUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the upload database unauthorized response
func (o *UploadDatabaseUnauthorized) Code() int {
	return 401
}

func (o *UploadDatabaseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseUnauthorized ", 401)
}

func (o *UploadDatabaseUnauthorized) String() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseUnauthorized ", 401)
}

func (o *UploadDatabaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadDatabaseForbidden creates a UploadDatabaseForbidden with default headers values
func NewUploadDatabaseForbidden() *UploadDatabaseForbidden {
	return &UploadDatabaseForbidden{}
}

/*
UploadDatabaseForbidden describes a response with status code 403, with default header values.

UploadDatabaseForbidden upload database forbidden
*/
type UploadDatabaseForbidden struct {
}

// IsSuccess returns true when this upload database forbidden response has a 2xx status code
func (o *UploadDatabaseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload database forbidden response has a 3xx status code
func (o *UploadDatabaseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload database forbidden response has a 4xx status code
func (o *UploadDatabaseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload database forbidden response has a 5xx status code
func (o *UploadDatabaseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upload database forbidden response a status code equal to that given
func (o *UploadDatabaseForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upload database forbidden response
func (o *UploadDatabaseForbidden) Code() int {
	return 403
}

func (o *UploadDatabaseForbidden) Error() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseForbidden ", 403)
}

func (o *UploadDatabaseForbidden) String() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseForbidden ", 403)
}

func (o *UploadDatabaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadDatabaseNotFound creates a UploadDatabaseNotFound with default headers values
func NewUploadDatabaseNotFound() *UploadDatabaseNotFound {
	return &UploadDatabaseNotFound{}
}

/*
UploadDatabaseNotFound describes a response with status code 404, with default header values.

UploadDatabaseNotFound upload database not found
*/
type UploadDatabaseNotFound struct {
}

// IsSuccess returns true when this upload database not found response has a 2xx status code
func (o *UploadDatabaseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload database not found response has a 3xx status code
func (o *UploadDatabaseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload database not found response has a 4xx status code
func (o *UploadDatabaseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload database not found response has a 5xx status code
func (o *UploadDatabaseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this upload database not found response a status code equal to that given
func (o *UploadDatabaseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the upload database not found response
func (o *UploadDatabaseNotFound) Code() int {
	return 404
}

func (o *UploadDatabaseNotFound) Error() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseNotFound ", 404)
}

func (o *UploadDatabaseNotFound) String() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseNotFound ", 404)
}

func (o *UploadDatabaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadDatabaseUnsupportedMediaType creates a UploadDatabaseUnsupportedMediaType with default headers values
func NewUploadDatabaseUnsupportedMediaType() *UploadDatabaseUnsupportedMediaType {
	return &UploadDatabaseUnsupportedMediaType{}
}

/*
UploadDatabaseUnsupportedMediaType describes a response with status code 415, with default header values.

UploadDatabaseUnsupportedMediaType upload database unsupported media type
*/
type UploadDatabaseUnsupportedMediaType struct {
}

// IsSuccess returns true when this upload database unsupported media type response has a 2xx status code
func (o *UploadDatabaseUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload database unsupported media type response has a 3xx status code
func (o *UploadDatabaseUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload database unsupported media type response has a 4xx status code
func (o *UploadDatabaseUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload database unsupported media type response has a 5xx status code
func (o *UploadDatabaseUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this upload database unsupported media type response a status code equal to that given
func (o *UploadDatabaseUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

// Code gets the status code for the upload database unsupported media type response
func (o *UploadDatabaseUnsupportedMediaType) Code() int {
	return 415
}

func (o *UploadDatabaseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseUnsupportedMediaType ", 415)
}

func (o *UploadDatabaseUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /databases][%d] uploadDatabaseUnsupportedMediaType ", 415)
}

func (o *UploadDatabaseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
