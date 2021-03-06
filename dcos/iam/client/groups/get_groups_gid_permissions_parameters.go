// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewGetGroupsGidPermissionsParams creates a new GetGroupsGidPermissionsParams object
// with the default values initialized.
func NewGetGroupsGidPermissionsParams() *GetGroupsGidPermissionsParams {
	var ()
	return &GetGroupsGidPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupsGidPermissionsParamsWithTimeout creates a new GetGroupsGidPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupsGidPermissionsParamsWithTimeout(timeout time.Duration) *GetGroupsGidPermissionsParams {
	var ()
	return &GetGroupsGidPermissionsParams{

		timeout: timeout,
	}
}

// NewGetGroupsGidPermissionsParamsWithContext creates a new GetGroupsGidPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupsGidPermissionsParamsWithContext(ctx context.Context) *GetGroupsGidPermissionsParams {
	var ()
	return &GetGroupsGidPermissionsParams{

		Context: ctx,
	}
}

// NewGetGroupsGidPermissionsParamsWithHTTPClient creates a new GetGroupsGidPermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupsGidPermissionsParamsWithHTTPClient(client *http.Client) *GetGroupsGidPermissionsParams {
	var ()
	return &GetGroupsGidPermissionsParams{
		HTTPClient: client,
	}
}

/*GetGroupsGidPermissionsParams contains all the parameters to send to the API endpoint
for the get groups gid permissions operation typically these are written to a http.Request
*/
type GetGroupsGidPermissionsParams struct {

	/*Gid
	  The group ID.

	*/
	Gid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get groups gid permissions params
func (o *GetGroupsGidPermissionsParams) WithTimeout(timeout time.Duration) *GetGroupsGidPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get groups gid permissions params
func (o *GetGroupsGidPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get groups gid permissions params
func (o *GetGroupsGidPermissionsParams) WithContext(ctx context.Context) *GetGroupsGidPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get groups gid permissions params
func (o *GetGroupsGidPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get groups gid permissions params
func (o *GetGroupsGidPermissionsParams) WithHTTPClient(client *http.Client) *GetGroupsGidPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get groups gid permissions params
func (o *GetGroupsGidPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGid adds the gid to the get groups gid permissions params
func (o *GetGroupsGidPermissionsParams) WithGid(gid string) *GetGroupsGidPermissionsParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the get groups gid permissions params
func (o *GetGroupsGidPermissionsParams) SetGid(gid string) {
	o.Gid = gid
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupsGidPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gid
	if err := r.SetPathParam("gid", o.Gid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
