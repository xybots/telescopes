// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetServicesPathParams GetServicesPathParams GetServicesPathParams is a placeholder for the services related route path parameters
// swagger:model GetServicesPathParams
type GetServicesPathParams struct {

	// in:path
	Provider string `json:"provider,omitempty"`

	// in:path
	Service string `json:"service,omitempty"`
}

// Validate validates this get services path params
func (m *GetServicesPathParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetServicesPathParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetServicesPathParams) UnmarshalBinary(b []byte) error {
	var res GetServicesPathParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
