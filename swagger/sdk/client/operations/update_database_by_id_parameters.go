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

// NewUpdateDatabaseByIDParams creates a new UpdateDatabaseByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDatabaseByIDParams() *UpdateDatabaseByIDParams {
	return &UpdateDatabaseByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDatabaseByIDParamsWithTimeout creates a new UpdateDatabaseByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDatabaseByIDParamsWithTimeout(timeout time.Duration) *UpdateDatabaseByIDParams {
	return &UpdateDatabaseByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDatabaseByIDParamsWithContext creates a new UpdateDatabaseByIDParams object
// with the ability to set a context for a request.
func NewUpdateDatabaseByIDParamsWithContext(ctx context.Context) *UpdateDatabaseByIDParams {
	return &UpdateDatabaseByIDParams{
		Context: ctx,
	}
}

// NewUpdateDatabaseByIDParamsWithHTTPClient creates a new UpdateDatabaseByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDatabaseByIDParamsWithHTTPClient(client *http.Client) *UpdateDatabaseByIDParams {
	return &UpdateDatabaseByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateDatabaseByIDParams contains all the parameters to send to the API endpoint

	for the update database by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateDatabaseByIDParams struct {

	/* Body.

	   Update database request body parameter
	*/
	Body *models.UpdateDatabaseRequest

	// ID.
	//
	// Format: uint64
	ID uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update database by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDatabaseByIDParams) WithDefaults() *UpdateDatabaseByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update database by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDatabaseByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update database by Id params
func (o *UpdateDatabaseByIDParams) WithTimeout(timeout time.Duration) *UpdateDatabaseByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update database by Id params
func (o *UpdateDatabaseByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update database by Id params
func (o *UpdateDatabaseByIDParams) WithContext(ctx context.Context) *UpdateDatabaseByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update database by Id params
func (o *UpdateDatabaseByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update database by Id params
func (o *UpdateDatabaseByIDParams) WithHTTPClient(client *http.Client) *UpdateDatabaseByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update database by Id params
func (o *UpdateDatabaseByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update database by Id params
func (o *UpdateDatabaseByIDParams) WithBody(body *models.UpdateDatabaseRequest) *UpdateDatabaseByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update database by Id params
func (o *UpdateDatabaseByIDParams) SetBody(body *models.UpdateDatabaseRequest) {
	o.Body = body
}

// WithID adds the id to the update database by Id params
func (o *UpdateDatabaseByIDParams) WithID(id uint64) *UpdateDatabaseByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update database by Id params
func (o *UpdateDatabaseByIDParams) SetID(id uint64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDatabaseByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
