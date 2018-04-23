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

// ListManagerEthernetInterfacesReader is a Reader for the ListManagerEthernetInterfaces structure.
type ListManagerEthernetInterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListManagerEthernetInterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListManagerEthernetInterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListManagerEthernetInterfacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListManagerEthernetInterfacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListManagerEthernetInterfacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListManagerEthernetInterfacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListManagerEthernetInterfacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListManagerEthernetInterfacesOK creates a ListManagerEthernetInterfacesOK with default headers values
func NewListManagerEthernetInterfacesOK() *ListManagerEthernetInterfacesOK {
	return &ListManagerEthernetInterfacesOK{}
}

/*ListManagerEthernetInterfacesOK handles this case with default header values.

Success
*/
type ListManagerEthernetInterfacesOK struct {
	Payload *models.EthernetInterfaceCollectionEthernetInterfaceCollection
}

func (o *ListManagerEthernetInterfacesOK) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/EthernetInterfaces][%d] listManagerEthernetInterfacesOK  %+v", 200, o.Payload)
}

func (o *ListManagerEthernetInterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EthernetInterfaceCollectionEthernetInterfaceCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManagerEthernetInterfacesBadRequest creates a ListManagerEthernetInterfacesBadRequest with default headers values
func NewListManagerEthernetInterfacesBadRequest() *ListManagerEthernetInterfacesBadRequest {
	return &ListManagerEthernetInterfacesBadRequest{}
}

/*ListManagerEthernetInterfacesBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type ListManagerEthernetInterfacesBadRequest struct {
}

func (o *ListManagerEthernetInterfacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/EthernetInterfaces][%d] listManagerEthernetInterfacesBadRequest ", 400)
}

func (o *ListManagerEthernetInterfacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManagerEthernetInterfacesUnauthorized creates a ListManagerEthernetInterfacesUnauthorized with default headers values
func NewListManagerEthernetInterfacesUnauthorized() *ListManagerEthernetInterfacesUnauthorized {
	return &ListManagerEthernetInterfacesUnauthorized{}
}

/*ListManagerEthernetInterfacesUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type ListManagerEthernetInterfacesUnauthorized struct {
}

func (o *ListManagerEthernetInterfacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/EthernetInterfaces][%d] listManagerEthernetInterfacesUnauthorized ", 401)
}

func (o *ListManagerEthernetInterfacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManagerEthernetInterfacesForbidden creates a ListManagerEthernetInterfacesForbidden with default headers values
func NewListManagerEthernetInterfacesForbidden() *ListManagerEthernetInterfacesForbidden {
	return &ListManagerEthernetInterfacesForbidden{}
}

/*ListManagerEthernetInterfacesForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type ListManagerEthernetInterfacesForbidden struct {
}

func (o *ListManagerEthernetInterfacesForbidden) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/EthernetInterfaces][%d] listManagerEthernetInterfacesForbidden ", 403)
}

func (o *ListManagerEthernetInterfacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManagerEthernetInterfacesNotFound creates a ListManagerEthernetInterfacesNotFound with default headers values
func NewListManagerEthernetInterfacesNotFound() *ListManagerEthernetInterfacesNotFound {
	return &ListManagerEthernetInterfacesNotFound{}
}

/*ListManagerEthernetInterfacesNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type ListManagerEthernetInterfacesNotFound struct {
}

func (o *ListManagerEthernetInterfacesNotFound) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/EthernetInterfaces][%d] listManagerEthernetInterfacesNotFound ", 404)
}

func (o *ListManagerEthernetInterfacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListManagerEthernetInterfacesInternalServerError creates a ListManagerEthernetInterfacesInternalServerError with default headers values
func NewListManagerEthernetInterfacesInternalServerError() *ListManagerEthernetInterfacesInternalServerError {
	return &ListManagerEthernetInterfacesInternalServerError{}
}

/*ListManagerEthernetInterfacesInternalServerError handles this case with default header values.

Error
*/
type ListManagerEthernetInterfacesInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ListManagerEthernetInterfacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/EthernetInterfaces][%d] listManagerEthernetInterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListManagerEthernetInterfacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
