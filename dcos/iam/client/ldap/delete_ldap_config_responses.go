// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLdapConfigReader is a Reader for the DeleteLdapConfig structure.
type DeleteLdapConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLdapConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteLdapConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteLdapConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLdapConfigNoContent creates a DeleteLdapConfigNoContent with default headers values
func NewDeleteLdapConfigNoContent() *DeleteLdapConfigNoContent {
	return &DeleteLdapConfigNoContent{}
}

/*DeleteLdapConfigNoContent handles this case with default header values.

Configuration deleted.
*/
type DeleteLdapConfigNoContent struct {
}

func (o *DeleteLdapConfigNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ldap/config][%d] deleteLdapConfigNoContent ", 204)
}

func (o *DeleteLdapConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLdapConfigBadRequest creates a DeleteLdapConfigBadRequest with default headers values
func NewDeleteLdapConfigBadRequest() *DeleteLdapConfigBadRequest {
	return &DeleteLdapConfigBadRequest{}
}

/*DeleteLdapConfigBadRequest handles this case with default header values.

Various errors. If no config has yet been stored, the custom error code `ERR_LDAP_CONFIG_NOT_AVAILABLE` is set in the response.
*/
type DeleteLdapConfigBadRequest struct {
}

func (o *DeleteLdapConfigBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ldap/config][%d] deleteLdapConfigBadRequest ", 400)
}

func (o *DeleteLdapConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
