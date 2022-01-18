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

// NewLockDatabaseByIDParams creates a new LockDatabaseByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLockDatabaseByIDParams() *LockDatabaseByIDParams {
	return &LockDatabaseByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLockDatabaseByIDParamsWithTimeout creates a new LockDatabaseByIDParams object
// with the ability to set a timeout on a request.
func NewLockDatabaseByIDParamsWithTimeout(timeout time.Duration) *LockDatabaseByIDParams {
	return &LockDatabaseByIDParams{
		timeout: timeout,
	}
}

// NewLockDatabaseByIDParamsWithContext creates a new LockDatabaseByIDParams object
// with the ability to set a context for a request.
func NewLockDatabaseByIDParamsWithContext(ctx context.Context) *LockDatabaseByIDParams {
	return &LockDatabaseByIDParams{
		Context: ctx,
	}
}

// NewLockDatabaseByIDParamsWithHTTPClient creates a new LockDatabaseByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewLockDatabaseByIDParamsWithHTTPClient(client *http.Client) *LockDatabaseByIDParams {
	return &LockDatabaseByIDParams{
		HTTPClient: client,
	}
}

/* LockDatabaseByIDParams contains all the parameters to send to the API endpoint
   for the lock database by Id operation.

   Typically these are written to a http.Request.
*/
type LockDatabaseByIDParams struct {

	/* Body.

	   Lock/unlock database request body parameter
	*/
	Body *models.LockDatabaseRequest

	// ID.
	//
	// Format: uint64
	ID uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lock database by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockDatabaseByIDParams) WithDefaults() *LockDatabaseByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lock database by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockDatabaseByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lock database by Id params
func (o *LockDatabaseByIDParams) WithTimeout(timeout time.Duration) *LockDatabaseByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lock database by Id params
func (o *LockDatabaseByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lock database by Id params
func (o *LockDatabaseByIDParams) WithContext(ctx context.Context) *LockDatabaseByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lock database by Id params
func (o *LockDatabaseByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lock database by Id params
func (o *LockDatabaseByIDParams) WithHTTPClient(client *http.Client) *LockDatabaseByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lock database by Id params
func (o *LockDatabaseByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lock database by Id params
func (o *LockDatabaseByIDParams) WithBody(body *models.LockDatabaseRequest) *LockDatabaseByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lock database by Id params
func (o *LockDatabaseByIDParams) SetBody(body *models.LockDatabaseRequest) {
	o.Body = body
}

// WithID adds the id to the lock database by Id params
func (o *LockDatabaseByIDParams) WithID(id uint64) *LockDatabaseByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the lock database by Id params
func (o *LockDatabaseByIDParams) SetID(id uint64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LockDatabaseByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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