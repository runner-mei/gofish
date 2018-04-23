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

// ListManagersReader is a Reader for the ListManagers structure.
type ListManagersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListManagersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListManagersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListManagersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListManagersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListManagersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListManagersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListManagersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListManagersOK creates a ListManagersOK with default headers values
func NewListManagersOK() *ListManagersOK {
	return &ListManagersOK{}
}

/*ListManagersOK handles this case with default header values.

Success
*/
type ListManagersOK struct {
	Payload *models.ManagerCollectionManagerCollection
}

func (o *ListManagersOK) Error() string {
	return fmt.Sprintf("[GET /Managers][%d] listManagersOK  %+v", 200, o.Payload)
}

func (o *ListManagersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagerCollectionManagerCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManagersBadRequest creates a ListManagersBadRequest with default headers values
func NewListManagersBadRequest() *ListManagersBadRequest {
	return &ListManagersBadRequest{}
}

/*ListManagersBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type ListManagersBadRequest struct {
}

func (o *ListManagersBadRequest) Error() string {
	return fmt.Sprintf("[GET /Managers][%d] listManagersBadRequest ", 400)
}

func (o *ListManagersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManagersUnauthorized creates a ListManagersUnauthorized with default headers values
func NewListManagersUnauthorized() *ListManagersUnauthorized {
	return &ListManagersUnauthorized{}
}

/*ListManagersUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type ListManagersUnauthorized struct {
}

func (o *ListManagersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Managers][%d] listManagersUnauthorized ", 401)
}

func (o *ListManagersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManagersForbidden creates a ListManagersForbidden with default headers values
func NewListManagersForbidden() *ListManagersForbidden {
	return &ListManagersForbidden{}
}

/*ListManagersForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type ListManagersForbidden struct {
}

func (o *ListManagersForbidden) Error() string {
	return fmt.Sprintf("[GET /Managers][%d] listManagersForbidden ", 403)
}

func (o *ListManagersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManagersNotFound creates a ListManagersNotFound with default headers values
func NewListManagersNotFound() *ListManagersNotFound {
	return &ListManagersNotFound{}
}

/*ListManagersNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type ListManagersNotFound struct {
}

func (o *ListManagersNotFound) Error() string {
	return fmt.Sprintf("[GET /Managers][%d] listManagersNotFound ", 404)
}

func (o *ListManagersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManagersInternalServerError creates a ListManagersInternalServerError with default headers values
func NewListManagersInternalServerError() *ListManagersInternalServerError {
	return &ListManagersInternalServerError{}
}

/*ListManagersInternalServerError handles this case with default header values.

Error
*/
type ListManagersInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ListManagersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Managers][%d] listManagersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListManagersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
