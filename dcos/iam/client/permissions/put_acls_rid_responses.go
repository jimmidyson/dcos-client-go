// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutAclsRidReader is a Reader for the PutAclsRid structure.
type PutAclsRidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAclsRidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutAclsRidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewPutAclsRidConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAclsRidCreated creates a PutAclsRidCreated with default headers values
func NewPutAclsRidCreated() *PutAclsRidCreated {
	return &PutAclsRidCreated{}
}

/*PutAclsRidCreated handles this case with default header values.

ACL created.
*/
type PutAclsRidCreated struct {
}

func (o *PutAclsRidCreated) Error() string {
	return fmt.Sprintf("[PUT /acls/{rid}][%d] putAclsRidCreated ", 201)
}

func (o *PutAclsRidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAclsRidConflict creates a PutAclsRidConflict with default headers values
func NewPutAclsRidConflict() *PutAclsRidConflict {
	return &PutAclsRidConflict{}
}

/*PutAclsRidConflict handles this case with default header values.

Already exists (this resource already has an ACL set).
*/
type PutAclsRidConflict struct {
}

func (o *PutAclsRidConflict) Error() string {
	return fmt.Sprintf("[PUT /acls/{rid}][%d] putAclsRidConflict ", 409)
}

func (o *PutAclsRidConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}