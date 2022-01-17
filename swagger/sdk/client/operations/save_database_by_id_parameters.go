// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/dhis2-sre/im-database-manager/swagger/sdk/models"
)

// NewSaveDatabaseByIDParams creates a new SaveDatabaseByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSaveDatabaseByIDParams() *SaveDatabaseByIDParams {
	return &SaveDatabaseByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSaveDatabaseByIDParamsWithTimeout creates a new SaveDatabaseByIDParams object
// with the ability to set a timeout on a request.
func NewSaveDatabaseByIDParamsWithTimeout(timeout time.Duration) *SaveDatabaseByIDParams {
	return &SaveDatabaseByIDParams{
		timeout: timeout,
	}
}

// NewSaveDatabaseByIDParamsWithContext creates a new SaveDatabaseByIDParams object
// with the ability to set a context for a request.
func NewSaveDatabaseByIDParamsWithContext(ctx context.Context) *SaveDatabaseByIDParams {
	return &SaveDatabaseByIDParams{
		Context: ctx,
	}
}

// NewSaveDatabaseByIDParamsWithHTTPClient creates a new SaveDatabaseByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewSaveDatabaseByIDParamsWithHTTPClient(client *http.Client) *SaveDatabaseByIDParams {
	return &SaveDatabaseByIDParams{
		HTTPClient: client,
	}
}

/* SaveDatabaseByIDParams contains all the parameters to send to the API endpoint
   for the save database by Id operation.

   Typically these are written to a http.Request.
*/
type SaveDatabaseByIDParams struct {

	/* Body.

	   Save database request body parameter
	*/
	Body *models.SaveDatabaseRequest

	// ID.
	//
	// Format: uint64
	ID uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the save database by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveDatabaseByIDParams) WithDefaults() *SaveDatabaseByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the save database by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveDatabaseByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the save database by Id params
func (o *SaveDatabaseByIDParams) WithTimeout(timeout time.Duration) *SaveDatabaseByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save database by Id params
func (o *SaveDatabaseByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save database by Id params
func (o *SaveDatabaseByIDParams) WithContext(ctx context.Context) *SaveDatabaseByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save database by Id params
func (o *SaveDatabaseByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save database by Id params
func (o *SaveDatabaseByIDParams) WithHTTPClient(client *http.Client) *SaveDatabaseByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save database by Id params
func (o *SaveDatabaseByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the save database by Id params
func (o *SaveDatabaseByIDParams) WithBody(body *models.SaveDatabaseRequest) *SaveDatabaseByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the save database by Id params
func (o *SaveDatabaseByIDParams) SetBody(body *models.SaveDatabaseRequest) {
	o.Body = body
}

// WithID adds the id to the save database by Id params
func (o *SaveDatabaseByIDParams) WithID(id uint64) *SaveDatabaseByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the save database by Id params
func (o *SaveDatabaseByIDParams) SetID(id uint64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SaveDatabaseByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatUint64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
