package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/runner-mei/gofish/models"
)

// RemoveAccountReader is a Reader for the RemoveAccount structure.
type RemoveAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRemoveAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRemoveAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRemoveAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveAccountNoContent creates a RemoveAccountNoContent with default headers values
func NewRemoveAccountNoContent() *RemoveAccountNoContent {
	return &RemoveAccountNoContent{}
}

/*RemoveAccountNoContent handles this case with default header values.

Success
*/
type RemoveAccountNoContent struct {
}

func (o *RemoveAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /AccountService/Accounts/{name}][%d] removeAccountNoContent ", 204)
}

func (o *RemoveAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccountUnauthorized creates a RemoveAccountUnauthorized with default headers values
func NewRemoveAccountUnauthorized() *RemoveAccountUnauthorized {
	return &RemoveAccountUnauthorized{}
}

/*RemoveAccountUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type RemoveAccountUnauthorized struct {
}

func (o *RemoveAccountUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /AccountService/Accounts/{name}][%d] removeAccountUnauthorized ", 401)
}

func (o *RemoveAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccountForbidden creates a RemoveAccountForbidden with default headers values
func NewRemoveAccountForbidden() *RemoveAccountForbidden {
	return &RemoveAccountForbidden{}
}

/*RemoveAccountForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type RemoveAccountForbidden struct {
}

func (o *RemoveAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /AccountService/Accounts/{name}][%d] removeAccountForbidden ", 403)
}

func (o *RemoveAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAccountInternalServerError creates a RemoveAccountInternalServerError with default headers values
func NewRemoveAccountInternalServerError() *RemoveAccountInternalServerError {
	return &RemoveAccountInternalServerError{}
}

/*RemoveAccountInternalServerError handles this case with default header values.

Error
*/
type RemoveAccountInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *RemoveAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /AccountService/Accounts/{name}][%d] removeAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
