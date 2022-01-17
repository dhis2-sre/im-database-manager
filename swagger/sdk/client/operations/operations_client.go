// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDatabase(params *CreateDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDatabaseAccepted, error)

	DeleteDatabaseByID(params *DeleteDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDatabaseByIDAccepted, error)

	FindDatabaseByID(params *FindDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDatabaseByIDOK, error)

	ListDatabases(params *ListDatabasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDatabasesOK, error)

	LockDatabaseByID(params *LockDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LockDatabaseByIDOK, error)

	SaveDatabaseByID(params *SaveDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SaveDatabaseByIDAccepted, error)

	UnlockDatabaseByID(params *UnlockDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnlockDatabaseByIDAccepted, error)

	UpdateDatabaseByID(params *UpdateDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDatabaseByIDOK, error)

	UploadDatabase(params *UploadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadDatabaseCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDatabase Create database
*/
func (a *Client) CreateDatabase(params *CreateDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDatabaseAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDatabaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDatabase",
		Method:             "POST",
		PathPattern:        "/databases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDatabaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDatabaseAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDatabase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDatabaseByID Delete database by id
*/
func (a *Client) DeleteDatabaseByID(params *DeleteDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDatabaseByIDAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDatabaseByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDatabaseById",
		Method:             "DELETE",
		PathPattern:        "/databases/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDatabaseByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDatabaseByIDAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDatabaseById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindDatabaseByID Find database by id
*/
func (a *Client) FindDatabaseByID(params *FindDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDatabaseByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDatabaseByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDatabaseById",
		Method:             "GET",
		PathPattern:        "/databases/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindDatabaseByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindDatabaseByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDatabaseById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDatabases List databases
*/
func (a *Client) ListDatabases(params *ListDatabasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDatabasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDatabasesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listDatabases",
		Method:             "GET",
		PathPattern:        "/databases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListDatabasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDatabasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listDatabases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LockDatabaseByID Lock database by id
*/
func (a *Client) LockDatabaseByID(params *LockDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LockDatabaseByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLockDatabaseByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "lockDatabaseById",
		Method:             "POST",
		PathPattern:        "/databases/{id}/lock",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LockDatabaseByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LockDatabaseByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for lockDatabaseById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SaveDatabaseByID Save database by id
*/
func (a *Client) SaveDatabaseByID(params *SaveDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SaveDatabaseByIDAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveDatabaseByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "saveDatabaseById",
		Method:             "POST",
		PathPattern:        "/databases/{id}/save",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveDatabaseByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SaveDatabaseByIDAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for saveDatabaseById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnlockDatabaseByID Unlock database by id
*/
func (a *Client) UnlockDatabaseByID(params *UnlockDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnlockDatabaseByIDAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnlockDatabaseByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unlockDatabaseById",
		Method:             "DELETE",
		PathPattern:        "/databases/{id}/lock",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UnlockDatabaseByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnlockDatabaseByIDAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unlockDatabaseById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDatabaseByID Update database by id
*/
func (a *Client) UpdateDatabaseByID(params *UpdateDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDatabaseByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDatabaseByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDatabaseById",
		Method:             "PUT",
		PathPattern:        "/databases/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDatabaseByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDatabaseByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDatabaseById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadDatabase Upload database
*/
func (a *Client) UploadDatabase(params *UploadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadDatabaseCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadDatabaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadDatabase",
		Method:             "POST",
		PathPattern:        "/databases/{id}/upload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadDatabaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadDatabaseCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadDatabase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
