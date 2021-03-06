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

// GetThermalReader is a Reader for the GetThermal structure.
type GetThermalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetThermalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetThermalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetThermalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetThermalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetThermalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetThermalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetThermalOK creates a GetThermalOK with default headers values
func NewGetThermalOK() *GetThermalOK {
	return &GetThermalOK{}
}

/*GetThermalOK handles this case with default header values.

Success
*/
type GetThermalOK struct {
	Payload *models.Thermal100Thermal
}

func (o *GetThermalOK) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Thermal][%d] getThermalOK  %+v", 200, o.Payload)
}

func (o *GetThermalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thermal100Thermal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetThermalBadRequest creates a GetThermalBadRequest with default headers values
func NewGetThermalBadRequest() *GetThermalBadRequest {
	return &GetThermalBadRequest{}
}

/*GetThermalBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information (such as validation error on an input field, a missing required value, and so on). An extended error shall be returned in the response body, as defined in section Extended Error Handling.

*/
type GetThermalBadRequest struct {
}

func (o *GetThermalBadRequest) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Thermal][%d] getThermalBadRequest ", 400)
}

func (o *GetThermalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetThermalUnauthorized creates a GetThermalUnauthorized with default headers values
func NewGetThermalUnauthorized() *GetThermalUnauthorized {
	return &GetThermalUnauthorized{}
}

/*GetThermalUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetThermalUnauthorized struct {
}

func (o *GetThermalUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Thermal][%d] getThermalUnauthorized ", 401)
}

func (o *GetThermalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetThermalForbidden creates a GetThermalForbidden with default headers values
func NewGetThermalForbidden() *GetThermalForbidden {
	return &GetThermalForbidden{}
}

/*GetThermalForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not possess authorization to perform this request.

*/
type GetThermalForbidden struct {
}

func (o *GetThermalForbidden) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Thermal][%d] getThermalForbidden ", 403)
}

func (o *GetThermalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetThermalInternalServerError creates a GetThermalInternalServerError with default headers values
func NewGetThermalInternalServerError() *GetThermalInternalServerError {
	return &GetThermalInternalServerError{}
}

/*GetThermalInternalServerError handles this case with default header values.

Error
*/
type GetThermalInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetThermalInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Chassis/{identifier}/Thermal][%d] getThermalInternalServerError  %+v", 500, o.Payload)
}

func (o *GetThermalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
