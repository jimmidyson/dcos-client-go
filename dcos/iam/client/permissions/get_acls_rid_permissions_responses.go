// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/iam/models"
)

// GetAclsRidPermissionsReader is a Reader for the GetAclsRidPermissions structure.
type GetAclsRidPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAclsRidPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAclsRidPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAclsRidPermissionsOK creates a GetAclsRidPermissionsOK with default headers values
func NewGetAclsRidPermissionsOK() *GetAclsRidPermissionsOK {
	return &GetAclsRidPermissionsOK{}
}

/*GetAclsRidPermissionsOK handles this case with default header values.

Success.
*/
type GetAclsRidPermissionsOK struct {
	Payload *models.ACLPermissions
}

func (o *GetAclsRidPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /acls/{rid}/permissions][%d] getAclsRidPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetAclsRidPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ACLPermissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
