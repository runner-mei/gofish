package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// SimpleStorage100SimpleStorage This is the schema definition for the Simple Storage resource.  It represents the properties of a storage controller and its directly-attached devices.
// swagger:model SimpleStorage.1.0.0_SimpleStorage
type SimpleStorage100SimpleStorage struct {

	// at odata context
	// Read Only: true
	AtOdataContext strfmt.URI `json:"@odata.context,omitempty"`

	// at odata id
	// Read Only: true
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`

	// at odata type
	// Read Only: true
	AtOdataType string `json:"@odata.type,omitempty"`

	// Provides a description of this resource and is used for commonality  in the schema definitions.
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// The storage devices associated with this resource
	// Read Only: true
	Devices []*SimpleStorage100Device `json:"Devices"`

	// Uniquely identifies the resource within the collection of like resources.
	// Read Only: true
	ID string `json:"Id,omitempty"`

	// The name of the resource or array element.
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	Oem ResourceOem `json:"Oem,omitempty"`

	// status
	Status *ResourceStatus `json:"Status,omitempty"`

	// The UEFI device path used to access this storage controller.
	// Read Only: true
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`
}

// Validate validates this simple storage 1 0 0 simple storage
func (m *SimpleStorage100SimpleStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleStorage100SimpleStorage) validateDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	for i := 0; i < len(m.Devices); i++ {

		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {

			if err := m.Devices[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SimpleStorage100SimpleStorage) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
