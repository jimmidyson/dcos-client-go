// Code generated by go-swagger; DO NOT EDIT.

package saml

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

// NewPostAuthSamlProvidersProviderIDAcsCallbackParams creates a new PostAuthSamlProvidersProviderIDAcsCallbackParams object
// with the default values initialized.
func NewPostAuthSamlProvidersProviderIDAcsCallbackParams() *PostAuthSamlProvidersProviderIDAcsCallbackParams {
	var ()
	return &PostAuthSamlProvidersProviderIDAcsCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthSamlProvidersProviderIDAcsCallbackParamsWithTimeout creates a new PostAuthSamlProvidersProviderIDAcsCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAuthSamlProvidersProviderIDAcsCallbackParamsWithTimeout(timeout time.Duration) *PostAuthSamlProvidersProviderIDAcsCallbackParams {
	var ()
	return &PostAuthSamlProvidersProviderIDAcsCallbackParams{

		timeout: timeout,
	}
}

// NewPostAuthSamlProvidersProviderIDAcsCallbackParamsWithContext creates a new PostAuthSamlProvidersProviderIDAcsCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAuthSamlProvidersProviderIDAcsCallbackParamsWithContext(ctx context.Context) *PostAuthSamlProvidersProviderIDAcsCallbackParams {
	var ()
	return &PostAuthSamlProvidersProviderIDAcsCallbackParams{

		Context: ctx,
	}
}

// NewPostAuthSamlProvidersProviderIDAcsCallbackParamsWithHTTPClient creates a new PostAuthSamlProvidersProviderIDAcsCallbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAuthSamlProvidersProviderIDAcsCallbackParamsWithHTTPClient(client *http.Client) *PostAuthSamlProvidersProviderIDAcsCallbackParams {
	var ()
	return &PostAuthSamlProvidersProviderIDAcsCallbackParams{
		HTTPClient: client,
	}
}

/*PostAuthSamlProvidersProviderIDAcsCallbackParams contains all the parameters to send to the API endpoint
for the post auth saml providers provider ID acs callback operation typically these are written to a http.Request
*/
type PostAuthSamlProvidersProviderIDAcsCallbackParams struct {

	/*ProviderID
	  The ID of the provider the authentication response is meant for.

	*/
	ProviderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post auth saml providers provider ID acs callback params
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) WithTimeout(timeout time.Duration) *PostAuthSamlProvidersProviderIDAcsCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auth saml providers provider ID acs callback params
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auth saml providers provider ID acs callback params
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) WithContext(ctx context.Context) *PostAuthSamlProvidersProviderIDAcsCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auth saml providers provider ID acs callback params
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auth saml providers provider ID acs callback params
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) WithHTTPClient(client *http.Client) *PostAuthSamlProvidersProviderIDAcsCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auth saml providers provider ID acs callback params
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProviderID adds the providerID to the post auth saml providers provider ID acs callback params
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) WithProviderID(providerID string) *PostAuthSamlProvidersProviderIDAcsCallbackParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the post auth saml providers provider ID acs callback params
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) SetProviderID(providerID string) {
	o.ProviderID = providerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthSamlProvidersProviderIDAcsCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param provider-id
	if err := r.SetPathParam("provider-id", o.ProviderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
