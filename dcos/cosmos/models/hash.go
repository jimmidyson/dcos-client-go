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

// Hash hash
// swagger:model hash
type Hash struct {

	// algo
	// Required: true
	Algo Algo `json:"algo"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this hash
func (m *Hash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hash) validateAlgo(formats strfmt.Registry) error {

	if err := m.Algo.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("algo")
		}
		return err
	}

	return nil
}

func (m *Hash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Hash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hash) UnmarshalBinary(b []byte) error {
	var res Hash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
