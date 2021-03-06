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

// OIDCProviderConfig o ID c provider config
// swagger:model OIDCProviderConfig
type OIDCProviderConfig struct {

	// base url
	// Required: true
	BaseURL *string `json:"base_url"`

	// ca certs
	CaCerts string `json:"ca_certs,omitempty"`

	// client id
	// Required: true
	ClientID *string `json:"client_id"`

	// client secret
	// Required: true
	ClientSecret *string `json:"client_secret"`

	// description
	// Required: true
	Description *string `json:"description"`

	// issuer
	// Required: true
	Issuer *string `json:"issuer"`

	// verify server certificate
	VerifyServerCertificate bool `json:"verify_server_certificate,omitempty"`
}

// Validate validates this o ID c provider config
func (m *OIDCProviderConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OIDCProviderConfig) validateBaseURL(formats strfmt.Registry) error {

	if err := validate.Required("base_url", "body", m.BaseURL); err != nil {
		return err
	}

	return nil
}

func (m *OIDCProviderConfig) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *OIDCProviderConfig) validateClientSecret(formats strfmt.Registry) error {

	if err := validate.Required("client_secret", "body", m.ClientSecret); err != nil {
		return err
	}

	return nil
}

func (m *OIDCProviderConfig) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *OIDCProviderConfig) validateIssuer(formats strfmt.Registry) error {

	if err := validate.Required("issuer", "body", m.Issuer); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OIDCProviderConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OIDCProviderConfig) UnmarshalBinary(b []byte) error {
	var res OIDCProviderConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
