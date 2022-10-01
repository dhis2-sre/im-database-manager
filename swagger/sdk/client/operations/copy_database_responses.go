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

// CopyDatabaseReader is a Reader for the CopyDatabase structure.
type CopyDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCopyDatabaseAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCopyDatabaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCopyDatabaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCopyDatabaseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCopyDatabaseAccepted creates a CopyDatabaseAccepted with default headers values
func NewCopyDatabaseAccepted() *CopyDatabaseAccepted {
	return &CopyDatabaseAccepted{}
}

/* CopyDatabaseAccepted describes a response with status code 202, with default header values.

Database
*/
type CopyDatabaseAccepted struct {
	Payload *models.Database
}

// IsSuccess returns true when this copy database accepted response has a 2xx status code
func (o *CopyDatabaseAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this copy database accepted response has a 3xx status code
func (o *CopyDatabaseAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy database accepted response has a 4xx status code
func (o *CopyDatabaseAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this copy database accepted response has a 5xx status code
func (o *CopyDatabaseAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this copy database accepted response a status code equal to that given
func (o *CopyDatabaseAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CopyDatabaseAccepted) Error() string {
	return fmt.Sprintf("[POST /databases/{id}/copy][%d] copyDatabaseAccepted  %+v", 202, o.Payload)
}

func (o *CopyDatabaseAccepted) String() string {
	return fmt.Sprintf("[POST /databases/{id}/copy][%d] copyDatabaseAccepted  %+v", 202, o.Payload)
}

func (o *CopyDatabaseAccepted) GetPayload() *models.Database {
	return o.Payload
}

func (o *CopyDatabaseAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Database)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyDatabaseUnauthorized creates a CopyDatabaseUnauthorized with default headers values
func NewCopyDatabaseUnauthorized() *CopyDatabaseUnauthorized {
	return &CopyDatabaseUnauthorized{}
}

/* CopyDatabaseUnauthorized describes a response with status code 401, with default header values.

CopyDatabaseUnauthorized copy database unauthorized
*/
type CopyDatabaseUnauthorized struct {
}

// IsSuccess returns true when this copy database unauthorized response has a 2xx status code
func (o *CopyDatabaseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy database unauthorized response has a 3xx status code
func (o *CopyDatabaseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy database unauthorized response has a 4xx status code
func (o *CopyDatabaseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy database unauthorized response has a 5xx status code
func (o *CopyDatabaseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this copy database unauthorized response a status code equal to that given
func (o *CopyDatabaseUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CopyDatabaseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /databases/{id}/copy][%d] copyDatabaseUnauthorized ", 401)
}

func (o *CopyDatabaseUnauthorized) String() string {
	return fmt.Sprintf("[POST /databases/{id}/copy][%d] copyDatabaseUnauthorized ", 401)
}

func (o *CopyDatabaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyDatabaseForbidden creates a CopyDatabaseForbidden with default headers values
func NewCopyDatabaseForbidden() *CopyDatabaseForbidden {
	return &CopyDatabaseForbidden{}
}

/* CopyDatabaseForbidden describes a response with status code 403, with default header values.

CopyDatabaseForbidden copy database forbidden
*/
type CopyDatabaseForbidden struct {
}

// IsSuccess returns true when this copy database forbidden response has a 2xx status code
func (o *CopyDatabaseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy database forbidden response has a 3xx status code
func (o *CopyDatabaseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy database forbidden response has a 4xx status code
func (o *CopyDatabaseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy database forbidden response has a 5xx status code
func (o *CopyDatabaseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this copy database forbidden response a status code equal to that given
func (o *CopyDatabaseForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CopyDatabaseForbidden) Error() string {
	return fmt.Sprintf("[POST /databases/{id}/copy][%d] copyDatabaseForbidden ", 403)
}

func (o *CopyDatabaseForbidden) String() string {
	return fmt.Sprintf("[POST /databases/{id}/copy][%d] copyDatabaseForbidden ", 403)
}

func (o *CopyDatabaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyDatabaseUnsupportedMediaType creates a CopyDatabaseUnsupportedMediaType with default headers values
func NewCopyDatabaseUnsupportedMediaType() *CopyDatabaseUnsupportedMediaType {
	return &CopyDatabaseUnsupportedMediaType{}
}

/* CopyDatabaseUnsupportedMediaType describes a response with status code 415, with default header values.

CopyDatabaseUnsupportedMediaType copy database unsupported media type
*/
type CopyDatabaseUnsupportedMediaType struct {
}

// IsSuccess returns true when this copy database unsupported media type response has a 2xx status code
func (o *CopyDatabaseUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy database unsupported media type response has a 3xx status code
func (o *CopyDatabaseUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy database unsupported media type response has a 4xx status code
func (o *CopyDatabaseUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy database unsupported media type response has a 5xx status code
func (o *CopyDatabaseUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this copy database unsupported media type response a status code equal to that given
func (o *CopyDatabaseUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CopyDatabaseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /databases/{id}/copy][%d] copyDatabaseUnsupportedMediaType ", 415)
}

func (o *CopyDatabaseUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /databases/{id}/copy][%d] copyDatabaseUnsupportedMediaType ", 415)
}

func (o *CopyDatabaseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
