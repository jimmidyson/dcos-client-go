// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dcos/client-go/dcos/iam/models"
)

// PostAuthLoginReader is a Reader for the PostAuthLogin structure.
type PostAuthLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAuthLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostAuthLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAuthLoginOK creates a PostAuthLoginOK with default headers values
func NewPostAuthLoginOK() *PostAuthLoginOK {
	return &PostAuthLoginOK{}
}

/*PostAuthLoginOK handles this case with default header values.

Login successful. The response body contains a JSON object providing the authentication token.
*/
type PostAuthLoginOK struct {
	/*A cookie containing the auth token (implementation detail for browser support, should not be of interest to general API consumers).
	 */
	SetCookie string

	Payload *models.AuthToken
}

func (o *PostAuthLoginOK) Error() string {
	return fmt.Sprintf("[POST /auth/login][%d] postAuthLoginOK  %+v", 200, o.Payload)
}

func (o *PostAuthLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Set-Cookie
	o.SetCookie = response.GetHeader("Set-Cookie")

	o.Payload = new(models.AuthToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthLoginUnauthorized creates a PostAuthLoginUnauthorized with default headers values
func NewPostAuthLoginUnauthorized() *PostAuthLoginUnauthorized {
	return &PostAuthLoginUnauthorized{}
}

/*PostAuthLoginUnauthorized handles this case with default header values.

Login failed.
*/
type PostAuthLoginUnauthorized struct {
}

func (o *PostAuthLoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /auth/login][%d] postAuthLoginUnauthorized ", 401)
}

func (o *PostAuthLoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
