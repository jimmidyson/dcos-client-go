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

// PackageDescribeRequest package describe request
// swagger:model packageDescribeRequest
type PackageDescribeRequest struct {

	// package name
	// Required: true
	PackageName *string `json:"packageName"`

	// package version
	PackageVersion string `json:"packageVersion,omitempty"`
}

// Validate validates this package describe request
func (m *PackageDescribeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackageName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageDescribeRequest) validatePackageName(formats strfmt.Registry) error {

	if err := validate.Required("packageName", "body", m.PackageName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageDescribeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageDescribeRequest) UnmarshalBinary(b []byte) error {
	var res PackageDescribeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}