// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostnameResponse hostname response
// swagger:model HostnameResponse
type HostnameResponse struct {

	// hostname
	// Required: true
	Hostname *string `json:"hostname"`
}

// Validate validates this hostname response
func (m *HostnameResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostnameResponse) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostnameResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostnameResponse) UnmarshalBinary(b []byte) error {
	var res HostnameResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
