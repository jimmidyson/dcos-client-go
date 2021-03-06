// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Binaries binaries
// swagger:model binaries
type Binaries struct {

	// darwin
	Darwin *Os `json:"darwin,omitempty"`

	// linux
	Linux *Os `json:"linux,omitempty"`

	// windows
	Windows *Os `json:"windows,omitempty"`
}

// Validate validates this binaries
func (m *Binaries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDarwin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinux(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWindows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Binaries) validateDarwin(formats strfmt.Registry) error {

	if swag.IsZero(m.Darwin) { // not required
		return nil
	}

	if m.Darwin != nil {
		if err := m.Darwin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("darwin")
			}
			return err
		}
	}

	return nil
}

func (m *Binaries) validateLinux(formats strfmt.Registry) error {

	if swag.IsZero(m.Linux) { // not required
		return nil
	}

	if m.Linux != nil {
		if err := m.Linux.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linux")
			}
			return err
		}
	}

	return nil
}

func (m *Binaries) validateWindows(formats strfmt.Registry) error {

	if swag.IsZero(m.Windows) { // not required
		return nil
	}

	if m.Windows != nil {
		if err := m.Windows.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("windows")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Binaries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Binaries) UnmarshalBinary(b []byte) error {
	var res Binaries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
