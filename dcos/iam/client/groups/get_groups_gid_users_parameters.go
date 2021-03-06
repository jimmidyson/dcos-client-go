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

// NewGetGroupsGidUsersParams creates a new GetGroupsGidUsersParams object
// with the default values initialized.
func NewGetGroupsGidUsersParams() *GetGroupsGidUsersParams {
	var ()
	return &GetGroupsGidUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupsGidUsersParamsWithTimeout creates a new GetGroupsGidUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupsGidUsersParamsWithTimeout(timeout time.Duration) *GetGroupsGidUsersParams {
	var ()
	return &GetGroupsGidUsersParams{

		timeout: timeout,
	}
}

// NewGetGroupsGidUsersParamsWithContext creates a new GetGroupsGidUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupsGidUsersParamsWithContext(ctx context.Context) *GetGroupsGidUsersParams {
	var ()
	return &GetGroupsGidUsersParams{

		Context: ctx,
	}
}

// NewGetGroupsGidUsersParamsWithHTTPClient creates a new GetGroupsGidUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupsGidUsersParamsWithHTTPClient(client *http.Client) *GetGroupsGidUsersParams {
	var ()
	return &GetGroupsGidUsersParams{
		HTTPClient: client,
	}
}

/*GetGroupsGidUsersParams contains all the parameters to send to the API endpoint
for the get groups gid users operation typically these are written to a http.Request
*/
type GetGroupsGidUsersParams struct {

	/*Gid
	  The group ID.

	*/
	Gid string
	/*Type
	  If set to `service`, list only service accounts. If unset, default to only listing user accounts members of a group.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get groups gid users params
func (o *GetGroupsGidUsersParams) WithTimeout(timeout time.Duration) *GetGroupsGidUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get groups gid users params
func (o *GetGroupsGidUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get groups gid users params
func (o *GetGroupsGidUsersParams) WithContext(ctx context.Context) *GetGroupsGidUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get groups gid users params
func (o *GetGroupsGidUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get groups gid users params
func (o *GetGroupsGidUsersParams) WithHTTPClient(client *http.Client) *GetGroupsGidUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get groups gid users params
func (o *GetGroupsGidUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGid adds the gid to the get groups gid users params
func (o *GetGroupsGidUsersParams) WithGid(gid string) *GetGroupsGidUsersParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the get groups gid users params
func (o *GetGroupsGidUsersParams) SetGid(gid string) {
	o.Gid = gid
}

// WithType adds the typeVar to the get groups gid users params
func (o *GetGroupsGidUsersParams) WithType(typeVar *string) *GetGroupsGidUsersParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get groups gid users params
func (o *GetGroupsGidUsersParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupsGidUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gid
	if err := r.SetPathParam("gid", o.Gid); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
