// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/iam/models"
)

// GetUsersUIDReader is a Reader for the GetUsersUID structure.
type GetUsersUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersUIDOK creates a GetUsersUIDOK with default headers values
func NewGetUsersUIDOK() *GetUsersUIDOK {
	return &GetUsersUIDOK{}
}

/*GetUsersUIDOK handles this case with default header values.

Success.
*/
type GetUsersUIDOK struct {
	Payload *models.User
}

func (o *GetUsersUIDOK) Error() string {
	return fmt.Sprintf("[GET /users/{uid}][%d] getUsersUidOK  %+v", 200, o.Payload)
}

func (o *GetUsersUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
