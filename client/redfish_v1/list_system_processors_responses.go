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

// ListSystemProcessorsReader is a Reader for the ListSystemProcessors structure.
type ListSystemProcessorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSystemProcessorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSystemProcessorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSystemProcessorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListSystemProcessorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListSystemProcessorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListSystemProcessorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListSystemProcessorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSystemProcessorsOK creates a ListSystemProcessorsOK with default headers values
func NewListSystemProcessorsOK() *ListSystemProcessorsOK {
	return &ListSystemProcessorsOK{}
}

/*ListSystemProcessorsOK handles this case with default header values.

Success
*/
type ListSystemProcessorsOK struct {
	Payload *models.ProcessorCollectionProcessorCollection
}

func (o *ListSystemProcessorsOK) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors][%d] listSystemProcessorsOK  %+v", 200, o.Payload)
}

func (o *ListSystemProcessorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorCollectionProcessorCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSystemProcessorsBadRequest creates a ListSystemProcessorsBadRequest with default headers values
func NewListSystemProcessorsBadRequest() *ListSystemProcessorsBadRequest {
	return &ListSystemProcessorsBadRequest{}
}

/*ListSystemProcessorsBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type ListSystemProcessorsBadRequest struct {
}

func (o *ListSystemProcessorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors][%d] listSystemProcessorsBadRequest ", 400)
}

func (o *ListSystemProcessorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemProcessorsUnauthorized creates a ListSystemProcessorsUnauthorized with default headers values
func NewListSystemProcessorsUnauthorized() *ListSystemProcessorsUnauthorized {
	return &ListSystemProcessorsUnauthorized{}
}

/*ListSystemProcessorsUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type ListSystemProcessorsUnauthorized struct {
}

func (o *ListSystemProcessorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors][%d] listSystemProcessorsUnauthorized ", 401)
}

func (o *ListSystemProcessorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemProcessorsForbidden creates a ListSystemProcessorsForbidden with default headers values
func NewListSystemProcessorsForbidden() *ListSystemProcessorsForbidden {
	return &ListSystemProcessorsForbidden{}
}

/*ListSystemProcessorsForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type ListSystemProcessorsForbidden struct {
}

func (o *ListSystemProcessorsForbidden) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors][%d] listSystemProcessorsForbidden ", 403)
}

func (o *ListSystemProcessorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemProcessorsNotFound creates a ListSystemProcessorsNotFound with default headers values
func NewListSystemProcessorsNotFound() *ListSystemProcessorsNotFound {
	return &ListSystemProcessorsNotFound{}
}

/*ListSystemProcessorsNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type ListSystemProcessorsNotFound struct {
}

func (o *ListSystemProcessorsNotFound) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors][%d] listSystemProcessorsNotFound ", 404)
}

func (o *ListSystemProcessorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSystemProcessorsInternalServerError creates a ListSystemProcessorsInternalServerError with default headers values
func NewListSystemProcessorsInternalServerError() *ListSystemProcessorsInternalServerError {
	return &ListSystemProcessorsInternalServerError{}
}

/*ListSystemProcessorsInternalServerError handles this case with default header values.

Error
*/
type ListSystemProcessorsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ListSystemProcessorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Systems/{identifier}/Processors][%d] listSystemProcessorsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSystemProcessorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
