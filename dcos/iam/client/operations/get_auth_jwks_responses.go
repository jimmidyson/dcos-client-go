// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAuthJwksReader is a Reader for the GetAuthJwks structure.
type GetAuthJwksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthJwksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuthJwksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuthJwksOK creates a GetAuthJwksOK with default headers values
func NewGetAuthJwksOK() *GetAuthJwksOK {
	return &GetAuthJwksOK{}
}

/*GetAuthJwksOK handles this case with default header values.

The response body contains a JSON Web Key Set document.
*/
type GetAuthJwksOK struct {
}

func (o *GetAuthJwksOK) Error() string {
	return fmt.Sprintf("[GET /auth/jwks][%d] getAuthJwksOK ", 200)
}

func (o *GetAuthJwksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
