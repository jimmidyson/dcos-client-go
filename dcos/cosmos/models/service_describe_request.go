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

// ServiceDescribeRequest service describe request
// swagger:model serviceDescribeRequest
type ServiceDescribeRequest struct {

	// app Id
	// Required: true
	AppID *string `json:"appId"`

	// manager Id
	// Required: true
	ManagerID *string `json:"managerId"`
}

// Validate validates this service describe request
func (m *ServiceDescribeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagerID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDescribeRequest) validateAppID(formats strfmt.Registry) error {

	if err := validate.Required("appId", "body", m.AppID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDescribeRequest) validateManagerID(formats strfmt.Registry) error {

	if err := validate.Required("managerId", "body", m.ManagerID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDescribeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDescribeRequest) UnmarshalBinary(b []byte) error {
	var res ServiceDescribeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}