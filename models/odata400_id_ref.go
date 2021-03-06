package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Odata400IDRef odata 4 0 0 id ref
// swagger:model odata.4.0.0_idRef
type Odata400IDRef struct {

	// at odata id
	// Read Only: true
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`
}

// Validate validates this odata 4 0 0 id ref
func (m *Odata400IDRef) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
