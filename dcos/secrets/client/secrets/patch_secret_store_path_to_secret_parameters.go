// Code generated by go-swagger; DO NOT EDIT.

package secrets

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

	models "github.com/dcos/client-go/dcos/secrets/models"
)

// NewPatchSecretStorePathToSecretParams creates a new PatchSecretStorePathToSecretParams object
// with the default values initialized.
func NewPatchSecretStorePathToSecretParams() *PatchSecretStorePathToSecretParams {
	var ()
	return &PatchSecretStorePathToSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSecretStorePathToSecretParamsWithTimeout creates a new PatchSecretStorePathToSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSecretStorePathToSecretParamsWithTimeout(timeout time.Duration) *PatchSecretStorePathToSecretParams {
	var ()
	return &PatchSecretStorePathToSecretParams{

		timeout: timeout,
	}
}

// NewPatchSecretStorePathToSecretParamsWithContext creates a new PatchSecretStorePathToSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSecretStorePathToSecretParamsWithContext(ctx context.Context) *PatchSecretStorePathToSecretParams {
	var ()
	return &PatchSecretStorePathToSecretParams{

		Context: ctx,
	}
}

// NewPatchSecretStorePathToSecretParamsWithHTTPClient creates a new PatchSecretStorePathToSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSecretStorePathToSecretParamsWithHTTPClient(client *http.Client) *PatchSecretStorePathToSecretParams {
	var ()
	return &PatchSecretStorePathToSecretParams{
		HTTPClient: client,
	}
}

/*PatchSecretStorePathToSecretParams contains all the parameters to send to the API endpoint
for the patch secret store path to secret operation typically these are written to a http.Request
*/
type PatchSecretStorePathToSecretParams struct {

	/*Body
	  Secret value.

	*/
	Body *models.Secret
	/*PathToSecret
	  The path for the secret to update.

	*/
	PathToSecret string
	/*Store
	  The backend to store the secret in.

	*/
	Store string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) WithTimeout(timeout time.Duration) *PatchSecretStorePathToSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) WithContext(ctx context.Context) *PatchSecretStorePathToSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) WithHTTPClient(client *http.Client) *PatchSecretStorePathToSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) WithBody(body *models.Secret) *PatchSecretStorePathToSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) SetBody(body *models.Secret) {
	o.Body = body
}

// WithPathToSecret adds the pathToSecret to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) WithPathToSecret(pathToSecret string) *PatchSecretStorePathToSecretParams {
	o.SetPathToSecret(pathToSecret)
	return o
}

// SetPathToSecret adds the pathToSecret to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) SetPathToSecret(pathToSecret string) {
	o.PathToSecret = pathToSecret
}

// WithStore adds the store to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) WithStore(store string) *PatchSecretStorePathToSecretParams {
	o.SetStore(store)
	return o
}

// SetStore adds the store to the patch secret store path to secret params
func (o *PatchSecretStorePathToSecretParams) SetStore(store string) {
	o.Store = store
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSecretStorePathToSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param path-to-secret
	if err := r.SetPathParam("path-to-secret", o.PathToSecret); err != nil {
		return err
	}

	// path param store
	if err := r.SetPathParam("store", o.Store); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
