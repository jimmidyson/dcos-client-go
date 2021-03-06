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

	models "github.com/dcos/client-go/dcos/iam/models"
)

// NewPostLdapConfigTestParams creates a new PostLdapConfigTestParams object
// with the default values initialized.
func NewPostLdapConfigTestParams() *PostLdapConfigTestParams {
	var ()
	return &PostLdapConfigTestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLdapConfigTestParamsWithTimeout creates a new PostLdapConfigTestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLdapConfigTestParamsWithTimeout(timeout time.Duration) *PostLdapConfigTestParams {
	var ()
	return &PostLdapConfigTestParams{

		timeout: timeout,
	}
}

// NewPostLdapConfigTestParamsWithContext creates a new PostLdapConfigTestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLdapConfigTestParamsWithContext(ctx context.Context) *PostLdapConfigTestParams {
	var ()
	return &PostLdapConfigTestParams{

		Context: ctx,
	}
}

// NewPostLdapConfigTestParamsWithHTTPClient creates a new PostLdapConfigTestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLdapConfigTestParamsWithHTTPClient(client *http.Client) *PostLdapConfigTestParams {
	var ()
	return &PostLdapConfigTestParams{
		HTTPClient: client,
	}
}

/*PostLdapConfigTestParams contains all the parameters to send to the API endpoint
for the post ldap config test operation typically these are written to a http.Request
*/
type PostLdapConfigTestParams struct {

	/*TestUserCredentials
	  JSON object containing `uid` and password of an LDAP user. For the most expressive test result, choose credentials different from the lookup credentials. The `uid` is the string the user is supposed to log in with after successful LDAP back-end configuration.

	*/
	TestUserCredentials *models.LDAPTestCredentials

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ldap config test params
func (o *PostLdapConfigTestParams) WithTimeout(timeout time.Duration) *PostLdapConfigTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ldap config test params
func (o *PostLdapConfigTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ldap config test params
func (o *PostLdapConfigTestParams) WithContext(ctx context.Context) *PostLdapConfigTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ldap config test params
func (o *PostLdapConfigTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ldap config test params
func (o *PostLdapConfigTestParams) WithHTTPClient(client *http.Client) *PostLdapConfigTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ldap config test params
func (o *PostLdapConfigTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTestUserCredentials adds the testUserCredentials to the post ldap config test params
func (o *PostLdapConfigTestParams) WithTestUserCredentials(testUserCredentials *models.LDAPTestCredentials) *PostLdapConfigTestParams {
	o.SetTestUserCredentials(testUserCredentials)
	return o
}

// SetTestUserCredentials adds the testUserCredentials to the post ldap config test params
func (o *PostLdapConfigTestParams) SetTestUserCredentials(testUserCredentials *models.LDAPTestCredentials) {
	o.TestUserCredentials = testUserCredentials
}

// WriteToRequest writes these params to a swagger request
func (o *PostLdapConfigTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TestUserCredentials != nil {
		if err := r.SetBodyParam(o.TestUserCredentials); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
