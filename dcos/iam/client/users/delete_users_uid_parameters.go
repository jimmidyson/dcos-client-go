// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewDeleteUsersUIDParams creates a new DeleteUsersUIDParams object
// with the default values initialized.
func NewDeleteUsersUIDParams() *DeleteUsersUIDParams {
	var ()
	return &DeleteUsersUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsersUIDParamsWithTimeout creates a new DeleteUsersUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUsersUIDParamsWithTimeout(timeout time.Duration) *DeleteUsersUIDParams {
	var ()
	return &DeleteUsersUIDParams{

		timeout: timeout,
	}
}

// NewDeleteUsersUIDParamsWithContext creates a new DeleteUsersUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUsersUIDParamsWithContext(ctx context.Context) *DeleteUsersUIDParams {
	var ()
	return &DeleteUsersUIDParams{

		Context: ctx,
	}
}

// NewDeleteUsersUIDParamsWithHTTPClient creates a new DeleteUsersUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUsersUIDParamsWithHTTPClient(client *http.Client) *DeleteUsersUIDParams {
	var ()
	return &DeleteUsersUIDParams{
		HTTPClient: client,
	}
}

/*DeleteUsersUIDParams contains all the parameters to send to the API endpoint
for the delete users UID operation typically these are written to a http.Request
*/
type DeleteUsersUIDParams struct {

	/*UID
	  The ID of the user account to delete.

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete users UID params
func (o *DeleteUsersUIDParams) WithTimeout(timeout time.Duration) *DeleteUsersUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete users UID params
func (o *DeleteUsersUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete users UID params
func (o *DeleteUsersUIDParams) WithContext(ctx context.Context) *DeleteUsersUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete users UID params
func (o *DeleteUsersUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete users UID params
func (o *DeleteUsersUIDParams) WithHTTPClient(client *http.Client) *DeleteUsersUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete users UID params
func (o *DeleteUsersUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the delete users UID params
func (o *DeleteUsersUIDParams) WithUID(uid string) *DeleteUsersUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the delete users UID params
func (o *DeleteUsersUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsersUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
