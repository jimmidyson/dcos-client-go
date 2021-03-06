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

// KubernetesAuthDataResponse kubernetes auth data response
// swagger:model kubernetesAuthDataResponse
type KubernetesAuthDataResponse struct {

	// ca crt
	// Required: true
	CaCrt *string `json:"caCrt"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this kubernetes auth data response
func (m *KubernetesAuthDataResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaCrt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesAuthDataResponse) validateCaCrt(formats strfmt.Registry) error {

	if err := validate.Required("caCrt", "body", m.CaCrt); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesAuthDataResponse) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesAuthDataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesAuthDataResponse) UnmarshalBinary(b []byte) error {
	var res KubernetesAuthDataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
