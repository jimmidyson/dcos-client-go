// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutUsersUIDReader is a Reader for the PutUsersUID structure.
type PutUsersUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUsersUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutUsersUIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewPutUsersUIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutUsersUIDCreated creates a PutUsersUIDCreated with default headers values
func NewPutUsersUIDCreated() *PutUsersUIDCreated {
	return &PutUsersUIDCreated{}
}

/*PutUsersUIDCreated handles this case with default header values.

User created.
*/
type PutUsersUIDCreated struct {
}

func (o *PutUsersUIDCreated) Error() string {
	return fmt.Sprintf("[PUT /users/{uid}][%d] putUsersUidCreated ", 201)
}

func (o *PutUsersUIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUsersUIDConflict creates a PutUsersUIDConflict with default headers values
func NewPutUsersUIDConflict() *PutUsersUIDConflict {
	return &PutUsersUIDConflict{}
}

/*PutUsersUIDConflict handles this case with default header values.

User already exists.
*/
type PutUsersUIDConflict struct {
}

func (o *PutUsersUIDConflict) Error() string {
	return fmt.Sprintf("[PUT /users/{uid}][%d] putUsersUidConflict ", 409)
}

func (o *PutUsersUIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
