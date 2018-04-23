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

// GetServiceRootReader is a Reader for the GetServiceRoot structure.
type GetServiceRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetServiceRootBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetServiceRootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetServiceRootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetServiceRootInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceRootOK creates a GetServiceRootOK with default headers values
func NewGetServiceRootOK() *GetServiceRootOK {
	return &GetServiceRootOK{}
}

/*GetServiceRootOK handles this case with default header values.

Success
*/
type GetServiceRootOK struct {
	Payload *models.ServiceRoot100ServiceRoot
}

func (o *GetServiceRootOK) Error() string {
	return fmt.Sprintf("[GET /][%d] getServiceRootOK  %+v", 200, o.Payload)
}

func (o *GetServiceRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceRoot100ServiceRoot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceRootBadRequest creates a GetServiceRootBadRequest with default headers values
func NewGetServiceRootBadRequest() *GetServiceRootBadRequest {
	return &GetServiceRootBadRequest{}
}

/*GetServiceRootBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type GetServiceRootBadRequest struct {
}

func (o *GetServiceRootBadRequest) Error() string {
	return fmt.Sprintf("[GET /][%d] getServiceRootBadRequest ", 400)
}

func (o *GetServiceRootBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceRootUnauthorized creates a GetServiceRootUnauthorized with default headers values
func NewGetServiceRootUnauthorized() *GetServiceRootUnauthorized {
	return &GetServiceRootUnauthorized{}
}

/*GetServiceRootUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetServiceRootUnauthorized struct {
}

func (o *GetServiceRootUnauthorized) Error() string {
	return fmt.Sprintf("[GET /][%d] getServiceRootUnauthorized ", 401)
}

func (o *GetServiceRootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceRootForbidden creates a GetServiceRootForbidden with default headers values
func NewGetServiceRootForbidden() *GetServiceRootForbidden {
	return &GetServiceRootForbidden{}
}

/*GetServiceRootForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetServiceRootForbidden struct {
}

func (o *GetServiceRootForbidden) Error() string {
	return fmt.Sprintf("[GET /][%d] getServiceRootForbidden ", 403)
}

func (o *GetServiceRootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceRootInternalServerError creates a GetServiceRootInternalServerError with default headers values
func NewGetServiceRootInternalServerError() *GetServiceRootInternalServerError {
	return &GetServiceRootInternalServerError{}
}

/*GetServiceRootInternalServerError handles this case with default header values.

Error
*/
type GetServiceRootInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetServiceRootInternalServerError) Error() string {
	return fmt.Sprintf("[GET /][%d] getServiceRootInternalServerError  %+v", 500, o.Payload)
}

func (o *GetServiceRootInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
