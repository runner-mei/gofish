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

// GetTaskReader is a Reader for the GetTask structure.
type GetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTaskOK creates a GetTaskOK with default headers values
func NewGetTaskOK() *GetTaskOK {
	return &GetTaskOK{}
}

/*GetTaskOK handles this case with default header values.

Success
*/
type GetTaskOK struct {
	Payload *models.Task100Task
}

func (o *GetTaskOK) Error() string {
	return fmt.Sprintf("[GET /TaskService/Tasks/{identifier}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task100Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskBadRequest creates a GetTaskBadRequest with default headers values
func NewGetTaskBadRequest() *GetTaskBadRequest {
	return &GetTaskBadRequest{}
}

/*GetTaskBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type GetTaskBadRequest struct {
}

func (o *GetTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /TaskService/Tasks/{identifier}][%d] getTaskBadRequest ", 400)
}

func (o *GetTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskUnauthorized creates a GetTaskUnauthorized with default headers values
func NewGetTaskUnauthorized() *GetTaskUnauthorized {
	return &GetTaskUnauthorized{}
}

/*GetTaskUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetTaskUnauthorized struct {
}

func (o *GetTaskUnauthorized) Error() string {
	return fmt.Sprintf("[GET /TaskService/Tasks/{identifier}][%d] getTaskUnauthorized ", 401)
}

func (o *GetTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskForbidden creates a GetTaskForbidden with default headers values
func NewGetTaskForbidden() *GetTaskForbidden {
	return &GetTaskForbidden{}
}

/*GetTaskForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetTaskForbidden struct {
}

func (o *GetTaskForbidden) Error() string {
	return fmt.Sprintf("[GET /TaskService/Tasks/{identifier}][%d] getTaskForbidden ", 403)
}

func (o *GetTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskNotFound creates a GetTaskNotFound with default headers values
func NewGetTaskNotFound() *GetTaskNotFound {
	return &GetTaskNotFound{}
}

/*GetTaskNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetTaskNotFound struct {
}

func (o *GetTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /TaskService/Tasks/{identifier}][%d] getTaskNotFound ", 404)
}

func (o *GetTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskInternalServerError creates a GetTaskInternalServerError with default headers values
func NewGetTaskInternalServerError() *GetTaskInternalServerError {
	return &GetTaskInternalServerError{}
}

/*GetTaskInternalServerError handles this case with default header values.

Error
*/
type GetTaskInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /TaskService/Tasks/{identifier}][%d] getTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
