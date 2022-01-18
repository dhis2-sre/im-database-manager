// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateDatabaseRequest create database request
//
// swagger:model CreateDatabaseRequest
type CreateDatabaseRequest struct {

	// group Id
	GroupID uint64 `json:"groupId,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create database request
func (m *CreateDatabaseRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create database request based on context it is used
func (m *CreateDatabaseRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateDatabaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDatabaseRequest) UnmarshalBinary(b []byte) error {
	var res CreateDatabaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}