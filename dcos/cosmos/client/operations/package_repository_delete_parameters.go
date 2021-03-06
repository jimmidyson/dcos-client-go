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

	models "github.com/dcos/client-go/dcos/cosmos/models"
)

// NewPackageRepositoryDeleteParams creates a new PackageRepositoryDeleteParams object
// with the default values initialized.
func NewPackageRepositoryDeleteParams() *PackageRepositoryDeleteParams {
	var ()
	return &PackageRepositoryDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackageRepositoryDeleteParamsWithTimeout creates a new PackageRepositoryDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackageRepositoryDeleteParamsWithTimeout(timeout time.Duration) *PackageRepositoryDeleteParams {
	var ()
	return &PackageRepositoryDeleteParams{

		timeout: timeout,
	}
}

// NewPackageRepositoryDeleteParamsWithContext creates a new PackageRepositoryDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackageRepositoryDeleteParamsWithContext(ctx context.Context) *PackageRepositoryDeleteParams {
	var ()
	return &PackageRepositoryDeleteParams{

		Context: ctx,
	}
}

// NewPackageRepositoryDeleteParamsWithHTTPClient creates a new PackageRepositoryDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackageRepositoryDeleteParamsWithHTTPClient(client *http.Client) *PackageRepositoryDeleteParams {
	var ()
	return &PackageRepositoryDeleteParams{
		HTTPClient: client,
	}
}

/*PackageRepositoryDeleteParams contains all the parameters to send to the API endpoint
for the package repository delete operation typically these are written to a http.Request
*/
type PackageRepositoryDeleteParams struct {

	/*Body*/
	Body *models.PackageDeleteRepoRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the package repository delete params
func (o *PackageRepositoryDeleteParams) WithTimeout(timeout time.Duration) *PackageRepositoryDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package repository delete params
func (o *PackageRepositoryDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package repository delete params
func (o *PackageRepositoryDeleteParams) WithContext(ctx context.Context) *PackageRepositoryDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package repository delete params
func (o *PackageRepositoryDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package repository delete params
func (o *PackageRepositoryDeleteParams) WithHTTPClient(client *http.Client) *PackageRepositoryDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package repository delete params
func (o *PackageRepositoryDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the package repository delete params
func (o *PackageRepositoryDeleteParams) WithBody(body *models.PackageDeleteRepoRequest) *PackageRepositoryDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the package repository delete params
func (o *PackageRepositoryDeleteParams) SetBody(body *models.PackageDeleteRepoRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PackageRepositoryDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
