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
)

// NewGetSelLogServiceParams creates a new GetSelLogServiceParams object
// with the default values initialized.
func NewGetSelLogServiceParams() *GetSelLogServiceParams {
	var ()
	return &GetSelLogServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSelLogServiceParamsWithTimeout creates a new GetSelLogServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSelLogServiceParamsWithTimeout(timeout time.Duration) *GetSelLogServiceParams {
	var ()
	return &GetSelLogServiceParams{

		timeout: timeout,
	}
}

// NewGetSelLogServiceParamsWithContext creates a new GetSelLogServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSelLogServiceParamsWithContext(ctx context.Context) *GetSelLogServiceParams {
	var ()
	return &GetSelLogServiceParams{

		Context: ctx,
	}
}

/*GetSelLogServiceParams contains all the parameters to send to the API endpoint
for the get sel log service operation typically these are written to a http.Request
*/
type GetSelLogServiceParams struct {

	/*Identifier*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sel log service params
func (o *GetSelLogServiceParams) WithTimeout(timeout time.Duration) *GetSelLogServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sel log service params
func (o *GetSelLogServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sel log service params
func (o *GetSelLogServiceParams) WithContext(ctx context.Context) *GetSelLogServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sel log service params
func (o *GetSelLogServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithIdentifier adds the identifier to the get sel log service params
func (o *GetSelLogServiceParams) WithIdentifier(identifier string) *GetSelLogServiceParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the get sel log service params
func (o *GetSelLogServiceParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetSelLogServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}