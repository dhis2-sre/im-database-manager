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
)

// NewFindDatabaseURLByIDParams creates a new FindDatabaseURLByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindDatabaseURLByIDParams() *FindDatabaseURLByIDParams {
	return &FindDatabaseURLByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindDatabaseURLByIDParamsWithTimeout creates a new FindDatabaseURLByIDParams object
// with the ability to set a timeout on a request.
func NewFindDatabaseURLByIDParamsWithTimeout(timeout time.Duration) *FindDatabaseURLByIDParams {
	return &FindDatabaseURLByIDParams{
		timeout: timeout,
	}
}

// NewFindDatabaseURLByIDParamsWithContext creates a new FindDatabaseURLByIDParams object
// with the ability to set a context for a request.
func NewFindDatabaseURLByIDParamsWithContext(ctx context.Context) *FindDatabaseURLByIDParams {
	return &FindDatabaseURLByIDParams{
		Context: ctx,
	}
}

// NewFindDatabaseURLByIDParamsWithHTTPClient creates a new FindDatabaseURLByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindDatabaseURLByIDParamsWithHTTPClient(client *http.Client) *FindDatabaseURLByIDParams {
	return &FindDatabaseURLByIDParams{
		HTTPClient: client,
	}
}

/* FindDatabaseURLByIDParams contains all the parameters to send to the API endpoint
   for the find database Url by Id operation.

   Typically these are written to a http.Request.
*/
type FindDatabaseURLByIDParams struct {

	// ID.
	//
	// Format: uint64
	ID uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find database Url by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindDatabaseURLByIDParams) WithDefaults() *FindDatabaseURLByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find database Url by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindDatabaseURLByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find database Url by Id params
func (o *FindDatabaseURLByIDParams) WithTimeout(timeout time.Duration) *FindDatabaseURLByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find database Url by Id params
func (o *FindDatabaseURLByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find database Url by Id params
func (o *FindDatabaseURLByIDParams) WithContext(ctx context.Context) *FindDatabaseURLByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find database Url by Id params
func (o *FindDatabaseURLByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find database Url by Id params
func (o *FindDatabaseURLByIDParams) WithHTTPClient(client *http.Client) *FindDatabaseURLByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find database Url by Id params
func (o *FindDatabaseURLByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find database Url by Id params
func (o *FindDatabaseURLByIDParams) WithID(id uint64) *FindDatabaseURLByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find database Url by Id params
func (o *FindDatabaseURLByIDParams) SetID(id uint64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindDatabaseURLByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatUint64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
