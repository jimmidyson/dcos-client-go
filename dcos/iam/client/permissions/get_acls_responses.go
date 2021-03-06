// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/iam/models"
)

// GetAclsReader is a Reader for the GetAcls structure.
type GetAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAclsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAclsOK creates a GetAclsOK with default headers values
func NewGetAclsOK() *GetAclsOK {
	return &GetAclsOK{}
}

/*GetAclsOK handles this case with default header values.

Success
*/
type GetAclsOK struct {
	Payload *GetAclsOKBody
}

func (o *GetAclsOK) Error() string {
	return fmt.Sprintf("[GET /acls][%d] getAclsOK  %+v", 200, o.Payload)
}

func (o *GetAclsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAclsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAclsOKBody get acls o k body
swagger:model GetAclsOKBody
*/
type GetAclsOKBody struct {

	// array
	Array []*models.ACL `json:"array"`
}

// Validate validates this get acls o k body
func (o *GetAclsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateArray(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAclsOKBody) validateArray(formats strfmt.Registry) error {

	if swag.IsZero(o.Array) { // not required
		return nil
	}

	for i := 0; i < len(o.Array); i++ {
		if swag.IsZero(o.Array[i]) { // not required
			continue
		}

		if o.Array[i] != nil {
			if err := o.Array[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAclsOK" + "." + "array" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAclsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAclsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAclsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
