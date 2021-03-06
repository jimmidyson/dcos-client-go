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

// NewPackageDescribeParams creates a new PackageDescribeParams object
// with the default values initialized.
func NewPackageDescribeParams() *PackageDescribeParams {
	var ()
	return &PackageDescribeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackageDescribeParamsWithTimeout creates a new PackageDescribeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackageDescribeParamsWithTimeout(timeout time.Duration) *PackageDescribeParams {
	var ()
	return &PackageDescribeParams{

		timeout: timeout,
	}
}

// NewPackageDescribeParamsWithContext creates a new PackageDescribeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackageDescribeParamsWithContext(ctx context.Context) *PackageDescribeParams {
	var ()
	return &PackageDescribeParams{

		Context: ctx,
	}
}

// NewPackageDescribeParamsWithHTTPClient creates a new PackageDescribeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackageDescribeParamsWithHTTPClient(client *http.Client) *PackageDescribeParams {
	var ()
	return &PackageDescribeParams{
		HTTPClient: client,
	}
}

/*PackageDescribeParams contains all the parameters to send to the API endpoint
for the package describe operation typically these are written to a http.Request
*/
type PackageDescribeParams struct {

	/*Body*/
	Body *models.PackageDescribeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the package describe params
func (o *PackageDescribeParams) WithTimeout(timeout time.Duration) *PackageDescribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package describe params
func (o *PackageDescribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package describe params
func (o *PackageDescribeParams) WithContext(ctx context.Context) *PackageDescribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package describe params
func (o *PackageDescribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package describe params
func (o *PackageDescribeParams) WithHTTPClient(client *http.Client) *PackageDescribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package describe params
func (o *PackageDescribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the package describe params
func (o *PackageDescribeParams) WithBody(body *models.PackageDescribeRequest) *PackageDescribeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the package describe params
func (o *PackageDescribeParams) SetBody(body *models.PackageDescribeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PackageDescribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
