// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateDatabaseRequest update database request
//
// swagger:model UpdateDatabaseRequest
type UpdateDatabaseRequest struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this update database request
func (m *UpdateDatabaseRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update database request based on context it is used
func (m *UpdateDatabaseRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDatabaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDatabaseRequest) UnmarshalBinary(b []byte) error {
	var res UpdateDatabaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
