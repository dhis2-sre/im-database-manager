// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExternalDownloadDatabaseReader is a Reader for the ExternalDownloadDatabase structure.
type ExternalDownloadDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExternalDownloadDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExternalDownloadDatabaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExternalDownloadDatabaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExternalDownloadDatabaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExternalDownloadDatabaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewExternalDownloadDatabaseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExternalDownloadDatabaseOK creates a ExternalDownloadDatabaseOK with default headers values
func NewExternalDownloadDatabaseOK() *ExternalDownloadDatabaseOK {
	return &ExternalDownloadDatabaseOK{}
}

/*
ExternalDownloadDatabaseOK describes a response with status code 200, with default header values.

ExternalDownloadDatabaseOK external download database o k
*/
type ExternalDownloadDatabaseOK struct {
	Payload []uint8
}

// IsSuccess returns true when this external download database o k response has a 2xx status code
func (o *ExternalDownloadDatabaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this external download database o k response has a 3xx status code
func (o *ExternalDownloadDatabaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external download database o k response has a 4xx status code
func (o *ExternalDownloadDatabaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this external download database o k response has a 5xx status code
func (o *ExternalDownloadDatabaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this external download database o k response a status code equal to that given
func (o *ExternalDownloadDatabaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the external download database o k response
func (o *ExternalDownloadDatabaseOK) Code() int {
	return 200
}

func (o *ExternalDownloadDatabaseOK) Error() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseOK  %+v", 200, o.Payload)
}

func (o *ExternalDownloadDatabaseOK) String() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseOK  %+v", 200, o.Payload)
}

func (o *ExternalDownloadDatabaseOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *ExternalDownloadDatabaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExternalDownloadDatabaseUnauthorized creates a ExternalDownloadDatabaseUnauthorized with default headers values
func NewExternalDownloadDatabaseUnauthorized() *ExternalDownloadDatabaseUnauthorized {
	return &ExternalDownloadDatabaseUnauthorized{}
}

/*
ExternalDownloadDatabaseUnauthorized describes a response with status code 401, with default header values.

ExternalDownloadDatabaseUnauthorized external download database unauthorized
*/
type ExternalDownloadDatabaseUnauthorized struct {
}

// IsSuccess returns true when this external download database unauthorized response has a 2xx status code
func (o *ExternalDownloadDatabaseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this external download database unauthorized response has a 3xx status code
func (o *ExternalDownloadDatabaseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external download database unauthorized response has a 4xx status code
func (o *ExternalDownloadDatabaseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this external download database unauthorized response has a 5xx status code
func (o *ExternalDownloadDatabaseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this external download database unauthorized response a status code equal to that given
func (o *ExternalDownloadDatabaseUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the external download database unauthorized response
func (o *ExternalDownloadDatabaseUnauthorized) Code() int {
	return 401
}

func (o *ExternalDownloadDatabaseUnauthorized) Error() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseUnauthorized ", 401)
}

func (o *ExternalDownloadDatabaseUnauthorized) String() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseUnauthorized ", 401)
}

func (o *ExternalDownloadDatabaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExternalDownloadDatabaseForbidden creates a ExternalDownloadDatabaseForbidden with default headers values
func NewExternalDownloadDatabaseForbidden() *ExternalDownloadDatabaseForbidden {
	return &ExternalDownloadDatabaseForbidden{}
}

/*
ExternalDownloadDatabaseForbidden describes a response with status code 403, with default header values.

ExternalDownloadDatabaseForbidden external download database forbidden
*/
type ExternalDownloadDatabaseForbidden struct {
}

// IsSuccess returns true when this external download database forbidden response has a 2xx status code
func (o *ExternalDownloadDatabaseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this external download database forbidden response has a 3xx status code
func (o *ExternalDownloadDatabaseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external download database forbidden response has a 4xx status code
func (o *ExternalDownloadDatabaseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this external download database forbidden response has a 5xx status code
func (o *ExternalDownloadDatabaseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this external download database forbidden response a status code equal to that given
func (o *ExternalDownloadDatabaseForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the external download database forbidden response
func (o *ExternalDownloadDatabaseForbidden) Code() int {
	return 403
}

func (o *ExternalDownloadDatabaseForbidden) Error() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseForbidden ", 403)
}

func (o *ExternalDownloadDatabaseForbidden) String() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseForbidden ", 403)
}

func (o *ExternalDownloadDatabaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExternalDownloadDatabaseNotFound creates a ExternalDownloadDatabaseNotFound with default headers values
func NewExternalDownloadDatabaseNotFound() *ExternalDownloadDatabaseNotFound {
	return &ExternalDownloadDatabaseNotFound{}
}

/*
ExternalDownloadDatabaseNotFound describes a response with status code 404, with default header values.

ExternalDownloadDatabaseNotFound external download database not found
*/
type ExternalDownloadDatabaseNotFound struct {
}

// IsSuccess returns true when this external download database not found response has a 2xx status code
func (o *ExternalDownloadDatabaseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this external download database not found response has a 3xx status code
func (o *ExternalDownloadDatabaseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external download database not found response has a 4xx status code
func (o *ExternalDownloadDatabaseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this external download database not found response has a 5xx status code
func (o *ExternalDownloadDatabaseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this external download database not found response a status code equal to that given
func (o *ExternalDownloadDatabaseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the external download database not found response
func (o *ExternalDownloadDatabaseNotFound) Code() int {
	return 404
}

func (o *ExternalDownloadDatabaseNotFound) Error() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseNotFound ", 404)
}

func (o *ExternalDownloadDatabaseNotFound) String() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseNotFound ", 404)
}

func (o *ExternalDownloadDatabaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExternalDownloadDatabaseUnsupportedMediaType creates a ExternalDownloadDatabaseUnsupportedMediaType with default headers values
func NewExternalDownloadDatabaseUnsupportedMediaType() *ExternalDownloadDatabaseUnsupportedMediaType {
	return &ExternalDownloadDatabaseUnsupportedMediaType{}
}

/*
ExternalDownloadDatabaseUnsupportedMediaType describes a response with status code 415, with default header values.

ExternalDownloadDatabaseUnsupportedMediaType external download database unsupported media type
*/
type ExternalDownloadDatabaseUnsupportedMediaType struct {
}

// IsSuccess returns true when this external download database unsupported media type response has a 2xx status code
func (o *ExternalDownloadDatabaseUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this external download database unsupported media type response has a 3xx status code
func (o *ExternalDownloadDatabaseUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external download database unsupported media type response has a 4xx status code
func (o *ExternalDownloadDatabaseUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this external download database unsupported media type response has a 5xx status code
func (o *ExternalDownloadDatabaseUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this external download database unsupported media type response a status code equal to that given
func (o *ExternalDownloadDatabaseUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

// Code gets the status code for the external download database unsupported media type response
func (o *ExternalDownloadDatabaseUnsupportedMediaType) Code() int {
	return 415
}

func (o *ExternalDownloadDatabaseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseUnsupportedMediaType ", 415)
}

func (o *ExternalDownloadDatabaseUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /databases/external/{uuid}][%d] externalDownloadDatabaseUnsupportedMediaType ", 415)
}

func (o *ExternalDownloadDatabaseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
