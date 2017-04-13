package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/danehans/gofish/models"
)

// NewPatchManagerParams creates a new PatchManagerParams object
// with the default values initialized.
func NewPatchManagerParams() *PatchManagerParams {
	var ()
	return &PatchManagerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchManagerParamsWithTimeout creates a new PatchManagerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchManagerParamsWithTimeout(timeout time.Duration) *PatchManagerParams {
	var ()
	return &PatchManagerParams{

		timeout: timeout,
	}
}

// NewPatchManagerParamsWithContext creates a new PatchManagerParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchManagerParamsWithContext(ctx context.Context) *PatchManagerParams {
	var ()
	return &PatchManagerParams{

		Context: ctx,
	}
}

/*PatchManagerParams contains all the parameters to send to the API endpoint
for the patch manager operation typically these are written to a http.Request
*/
type PatchManagerParams struct {

	/*Identifier*/
	Identifier string
	/*Payload*/
	Payload *models.Manager100Manager

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch manager params
func (o *PatchManagerParams) WithTimeout(timeout time.Duration) *PatchManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch manager params
func (o *PatchManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch manager params
func (o *PatchManagerParams) WithContext(ctx context.Context) *PatchManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch manager params
func (o *PatchManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIdentifier adds the identifier to the patch manager params
func (o *PatchManagerParams) WithIdentifier(identifier string) *PatchManagerParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the patch manager params
func (o *PatchManagerParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WithPayload adds the payload to the patch manager params
func (o *PatchManagerParams) WithPayload(payload *models.Manager100Manager) *PatchManagerParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the patch manager params
func (o *PatchManagerParams) SetPayload(payload *models.Manager100Manager) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *PatchManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if o.Payload == nil {
		o.Payload = new(models.Manager100Manager)
	}

	if err := r.SetBodyParam(o.Payload); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}