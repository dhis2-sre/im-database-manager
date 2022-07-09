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
	CopyDatabase(params *CopyDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CopyDatabaseAccepted, error)

	CreateExternalDownloadDatabase(params *CreateExternalDownloadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExternalDownloadDatabaseOK, error)

	DeleteDatabaseByID(params *DeleteDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDatabaseByIDAccepted, error)

	DownloadDatabase(params *DownloadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DownloadDatabaseOK, error)

	ExternalDownloadDatabase(params *ExternalDownloadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalDownloadDatabaseOK, error)

	FindDatabaseByID(params *FindDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindDatabaseByIDOK, error)

	FindDatabaseURLByID(params *FindDatabaseURLByIDParams, opts ...ClientOption) (*FindDatabaseURLByIDOK, error)

	Health(params *HealthParams, opts ...ClientOption) (*HealthOK, error)

	ListDatabases(params *ListDatabasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDatabasesOK, error)

	LockDatabaseByID(params *LockDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LockDatabaseByIDOK, error)

	UnlockDatabaseByID(params *UnlockDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnlockDatabaseByIDAccepted, error)

	UpdateDatabaseByID(params *UpdateDatabaseByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDatabaseByIDOK, error)

	UploadDatabase(params *UploadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadDatabaseCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CopyDatabase Copy database
*/
func (a *Client) CopyDatabase(params *CopyDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CopyDatabaseAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyDatabaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "copyDatabase",
		Method:             "POST",
		PathPattern:        "/databases/{id}/copy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CopyDatabaseReader{formats: a.formats},
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
	success, ok := result.(*CopyDatabaseAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for copyDatabase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateExternalDownloadDatabase Create external database download
*/
func (a *Client) CreateExternalDownloadDatabase(params *CreateExternalDownloadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExternalDownloadDatabaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExternalDownloadDatabaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createExternalDownloadDatabase",
		Method:             "POST",
		PathPattern:        "/databases/{id}/external",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateExternalDownloadDatabaseReader{formats: a.formats},
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
	success, ok := result.(*CreateExternalDownloadDatabaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createExternalDownloadDatabase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
  DownloadDatabase Download database
*/
func (a *Client) DownloadDatabase(params *DownloadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DownloadDatabaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadDatabaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadDatabase",
		Method:             "GET",
		PathPattern:        "/databases/{id}/download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadDatabaseReader{formats: a.formats},
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
	success, ok := result.(*DownloadDatabaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadDatabase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalDownloadDatabase Download database
*/
func (a *Client) ExternalDownloadDatabase(params *ExternalDownloadDatabaseParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalDownloadDatabaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalDownloadDatabaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "externalDownloadDatabase",
		Method:             "GET",
		PathPattern:        "/databases/external/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExternalDownloadDatabaseReader{formats: a.formats},
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
	success, ok := result.(*ExternalDownloadDatabaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for externalDownloadDatabase: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
  FindDatabaseURLByID Find database URL by id
*/
func (a *Client) FindDatabaseURLByID(params *FindDatabaseURLByIDParams, opts ...ClientOption) (*FindDatabaseURLByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindDatabaseURLByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findDatabaseUrlById",
		Method:             "GET",
		PathPattern:        "/databases/{id}/url",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindDatabaseURLByIDReader{formats: a.formats},
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
	success, ok := result.(*FindDatabaseURLByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findDatabaseUrlById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Health Service health status
*/
func (a *Client) Health(params *HealthParams, opts ...ClientOption) (*HealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "health",
		Method:             "GET",
		PathPattern:        "/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthReader{formats: a.formats},
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
	success, ok := result.(*HealthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for health: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
  UpdateDatabaseByID updates database by id

  TODO: Race condition? If two clients request at the same time... Do we need a transaction between find and update
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
		PathPattern:        "/databases",
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
