// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewKubernetesAuthDataParams creates a new KubernetesAuthDataParams object
// with the default values initialized.
func NewKubernetesAuthDataParams() *KubernetesAuthDataParams {
	var ()
	return &KubernetesAuthDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesAuthDataParamsWithTimeout creates a new KubernetesAuthDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKubernetesAuthDataParamsWithTimeout(timeout time.Duration) *KubernetesAuthDataParams {
	var ()
	return &KubernetesAuthDataParams{

		timeout: timeout,
	}
}

// NewKubernetesAuthDataParamsWithContext creates a new KubernetesAuthDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewKubernetesAuthDataParamsWithContext(ctx context.Context) *KubernetesAuthDataParams {
	var ()
	return &KubernetesAuthDataParams{

		Context: ctx,
	}
}

// NewKubernetesAuthDataParamsWithHTTPClient creates a new KubernetesAuthDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKubernetesAuthDataParamsWithHTTPClient(client *http.Client) *KubernetesAuthDataParams {
	var ()
	return &KubernetesAuthDataParams{
		HTTPClient: client,
	}
}

/*KubernetesAuthDataParams contains all the parameters to send to the API endpoint
for the kubernetes auth data operation typically these are written to a http.Request
*/
type KubernetesAuthDataParams struct {

	/*AppID*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the kubernetes auth data params
func (o *KubernetesAuthDataParams) WithTimeout(timeout time.Duration) *KubernetesAuthDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes auth data params
func (o *KubernetesAuthDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes auth data params
func (o *KubernetesAuthDataParams) WithContext(ctx context.Context) *KubernetesAuthDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes auth data params
func (o *KubernetesAuthDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes auth data params
func (o *KubernetesAuthDataParams) WithHTTPClient(client *http.Client) *KubernetesAuthDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes auth data params
func (o *KubernetesAuthDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the kubernetes auth data params
func (o *KubernetesAuthDataParams) WithAppID(appID string) *KubernetesAuthDataParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the kubernetes auth data params
func (o *KubernetesAuthDataParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesAuthDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
