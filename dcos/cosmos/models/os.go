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

// Os os
// swagger:model os
type Os struct {

	// x86 64
	// Required: true
	X8664 *CliInfo `json:"x86-64"`
}

// Validate validates this os
func (m *Os) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateX8664(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Os) validateX8664(formats strfmt.Registry) error {

	if err := validate.Required("x86-64", "body", m.X8664); err != nil {
		return err
	}

	if m.X8664 != nil {
		if err := m.X8664.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("x86-64")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Os) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Os) UnmarshalBinary(b []byte) error {
	var res Os
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}