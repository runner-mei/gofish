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

// NewCreateAccountParams creates a new CreateAccountParams object
// with the default values initialized.
func NewCreateAccountParams() *CreateAccountParams {
	var ()
	return &CreateAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccountParamsWithTimeout creates a new CreateAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccountParamsWithTimeout(timeout time.Duration) *CreateAccountParams {
	var ()
	return &CreateAccountParams{

		timeout: timeout,
	}
}

// NewCreateAccountParamsWithContext creates a new CreateAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAccountParamsWithContext(ctx context.Context) *CreateAccountParams {
	var ()
	return &CreateAccountParams{

		Context: ctx,
	}
}

/*CreateAccountParams contains all the parameters to send to the API endpoint
for the create account operation typically these are written to a http.Request
*/
type CreateAccountParams struct {

	/*Payload*/
	Payload *models.ManagerAccount100ManagerAccountPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create account params
func (o *CreateAccountParams) WithTimeout(timeout time.Duration) *CreateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create account params
func (o *CreateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create account params
func (o *CreateAccountParams) WithContext(ctx context.Context) *CreateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create account params
func (o *CreateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithPayload adds the payload to the create account params
func (o *CreateAccountParams) WithPayload(payload *models.ManagerAccount100ManagerAccountPost) *CreateAccountParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the create account params
func (o *CreateAccountParams) SetPayload(payload *models.ManagerAccount100ManagerAccountPost) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Payload == nil {
		o.Payload = new(models.ManagerAccount100ManagerAccountPost)
	}

	if err := r.SetBodyParam(o.Payload); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}