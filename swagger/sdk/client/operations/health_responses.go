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

// HealthReader is a Reader for the Health structure.
type HealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHealthOK creates a HealthOK with default headers values
func NewHealthOK() *HealthOK {
	return &HealthOK{}
}

/* HealthOK describes a response with status code 200, with default header values.

HealthOK health o k
*/
type HealthOK struct {
	Payload *models.Response
}

func (o *HealthOK) Error() string {
	return fmt.Sprintf("[GET /health][%d] healthOK  %+v", 200, o.Payload)
}
func (o *HealthOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *HealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
