// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLdapConfigParams creates a new GetLdapConfigParams object
// with the default values initialized.
func NewGetLdapConfigParams() *GetLdapConfigParams {

	return &GetLdapConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapConfigParamsWithTimeout creates a new GetLdapConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLdapConfigParamsWithTimeout(timeout time.Duration) *GetLdapConfigParams {

	return &GetLdapConfigParams{

		timeout: timeout,
	}
}

// NewGetLdapConfigParamsWithContext creates a new GetLdapConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLdapConfigParamsWithContext(ctx context.Context) *GetLdapConfigParams {

	return &GetLdapConfigParams{

		Context: ctx,
	}
}

// NewGetLdapConfigParamsWithHTTPClient creates a new GetLdapConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLdapConfigParamsWithHTTPClient(client *http.Client) *GetLdapConfigParams {

	return &GetLdapConfigParams{
		HTTPClient: client,
	}
}

/*GetLdapConfigParams contains all the parameters to send to the API endpoint
for the get ldap config operation typically these are written to a http.Request
*/
type GetLdapConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ldap config params
func (o *GetLdapConfigParams) WithTimeout(timeout time.Duration) *GetLdapConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap config params
func (o *GetLdapConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap config params
func (o *GetLdapConfigParams) WithContext(ctx context.Context) *GetLdapConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap config params
func (o *GetLdapConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap config params
func (o *GetLdapConfigParams) WithHTTPClient(client *http.Client) *GetLdapConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap config params
func (o *GetLdapConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
