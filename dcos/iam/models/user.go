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

// User user
// swagger:model User
type User struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// is remote
	// Required: true
	IsRemote *bool `json:"is_remote"`

	// is service
	IsService bool `json:"is_service,omitempty"`

	// provider id
	// Required: true
	ProviderID *string `json:"provider_id"`

	// provider type
	// Required: true
	ProviderType *string `json:"provider_type"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// uid
	// Required: true
	UID *string `json:"uid"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *User) validateIsRemote(formats strfmt.Registry) error {

	if err := validate.Required("is_remote", "body", m.IsRemote); err != nil {
		return err
	}

	return nil
}

func (m *User) validateProviderID(formats strfmt.Registry) error {

	if err := validate.Required("provider_id", "body", m.ProviderID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateProviderType(formats strfmt.Registry) error {

	if err := validate.Required("provider_type", "body", m.ProviderType); err != nil {
		return err
	}

	return nil
}

func (m *User) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
