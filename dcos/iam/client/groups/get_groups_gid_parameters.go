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

// NewGetGroupsGidParams creates a new GetGroupsGidParams object
// with the default values initialized.
func NewGetGroupsGidParams() *GetGroupsGidParams {
	var ()
	return &GetGroupsGidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupsGidParamsWithTimeout creates a new GetGroupsGidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupsGidParamsWithTimeout(timeout time.Duration) *GetGroupsGidParams {
	var ()
	return &GetGroupsGidParams{

		timeout: timeout,
	}
}

// NewGetGroupsGidParamsWithContext creates a new GetGroupsGidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupsGidParamsWithContext(ctx context.Context) *GetGroupsGidParams {
	var ()
	return &GetGroupsGidParams{

		Context: ctx,
	}
}

// NewGetGroupsGidParamsWithHTTPClient creates a new GetGroupsGidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupsGidParamsWithHTTPClient(client *http.Client) *GetGroupsGidParams {
	var ()
	return &GetGroupsGidParams{
		HTTPClient: client,
	}
}

/*GetGroupsGidParams contains all the parameters to send to the API endpoint
for the get groups gid operation typically these are written to a http.Request
*/
type GetGroupsGidParams struct {

	/*Gid
	  The ID of the group to retrieve.

	*/
	Gid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get groups gid params
func (o *GetGroupsGidParams) WithTimeout(timeout time.Duration) *GetGroupsGidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get groups gid params
func (o *GetGroupsGidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get groups gid params
func (o *GetGroupsGidParams) WithContext(ctx context.Context) *GetGroupsGidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get groups gid params
func (o *GetGroupsGidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get groups gid params
func (o *GetGroupsGidParams) WithHTTPClient(client *http.Client) *GetGroupsGidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get groups gid params
func (o *GetGroupsGidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGid adds the gid to the get groups gid params
func (o *GetGroupsGidParams) WithGid(gid string) *GetGroupsGidParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the get groups gid params
func (o *GetGroupsGidParams) SetGid(gid string) {
	o.Gid = gid
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupsGidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
