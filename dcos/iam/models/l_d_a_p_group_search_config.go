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

// LDAPGroupSearchConfig l d a p group search config
// swagger:model LDAPGroupSearchConfig
type LDAPGroupSearchConfig struct {

	// member attributes
	MemberAttributes []string `json:"member-attributes"`

	// search base
	// Required: true
	SearchBase *string `json:"search-base"`

	// search filter template
	// Required: true
	SearchFilterTemplate *string `json:"search-filter-template"`
}

// Validate validates this l d a p group search config
func (m *LDAPGroupSearchConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearchBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchFilterTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LDAPGroupSearchConfig) validateSearchBase(formats strfmt.Registry) error {

	if err := validate.Required("search-base", "body", m.SearchBase); err != nil {
		return err
	}

	return nil
}

func (m *LDAPGroupSearchConfig) validateSearchFilterTemplate(formats strfmt.Registry) error {

	if err := validate.Required("search-filter-template", "body", m.SearchFilterTemplate); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LDAPGroupSearchConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LDAPGroupSearchConfig) UnmarshalBinary(b []byte) error {
	var res LDAPGroupSearchConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}