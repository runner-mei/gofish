package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ResourceStatus resource status
// swagger:model Resource_Status
type ResourceStatus struct {

	// This represents the health state of this resource in the absence of its dependent resources.
	// Read Only: true
	Health string `json:"Health,omitempty"`

	// This represents the overall health state from the view of this resource.
	// Read Only: true
	HealthRollup string `json:"HealthRollup,omitempty"`

	// oem
	Oem ResourceOem `json:"Oem,omitempty"`

	// This indicates the known state of the resource, such as if it is enabled.
	// Read Only: true
	State string `json:"State,omitempty"`
}

// Validate validates this resource status
func (m *ResourceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHealthRollup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var resourceStatusTypeHealthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","Warning","Critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceStatusTypeHealthPropEnum = append(resourceStatusTypeHealthPropEnum, v)
	}
}

const (
	// ResourceStatusHealthOK captures enum value "OK"
	ResourceStatusHealthOK string = "OK"
	// ResourceStatusHealthWarning captures enum value "Warning"
	ResourceStatusHealthWarning string = "Warning"
	// ResourceStatusHealthCritical captures enum value "Critical"
	ResourceStatusHealthCritical string = "Critical"
)

// prop value enum
func (m *ResourceStatus) validateHealthEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceStatusTypeHealthPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceStatus) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	// value enum
	if err := m.validateHealthEnum("Health", "body", m.Health); err != nil {
		return err
	}

	return nil
}

var resourceStatusTypeHealthRollupPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","Warning","Critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceStatusTypeHealthRollupPropEnum = append(resourceStatusTypeHealthRollupPropEnum, v)
	}
}

const (
	// ResourceStatusHealthRollupOK captures enum value "OK"
	ResourceStatusHealthRollupOK string = "OK"
	// ResourceStatusHealthRollupWarning captures enum value "Warning"
	ResourceStatusHealthRollupWarning string = "Warning"
	// ResourceStatusHealthRollupCritical captures enum value "Critical"
	ResourceStatusHealthRollupCritical string = "Critical"
)

// prop value enum
func (m *ResourceStatus) validateHealthRollupEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceStatusTypeHealthRollupPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceStatus) validateHealthRollup(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthRollup) { // not required
		return nil
	}

	// value enum
	if err := m.validateHealthRollupEnum("HealthRollup", "body", m.HealthRollup); err != nil {
		return err
	}

	return nil
}

var resourceStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Enabled","Disabled","StandbyOffline","StandbySpare","InTest","Starting","Absent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceStatusTypeStatePropEnum = append(resourceStatusTypeStatePropEnum, v)
	}
}

const (
	// ResourceStatusStateEnabled captures enum value "Enabled"
	ResourceStatusStateEnabled string = "Enabled"
	// ResourceStatusStateDisabled captures enum value "Disabled"
	ResourceStatusStateDisabled string = "Disabled"
	// ResourceStatusStateStandbyOffline captures enum value "StandbyOffline"
	ResourceStatusStateStandbyOffline string = "StandbyOffline"
	// ResourceStatusStateStandbySpare captures enum value "StandbySpare"
	ResourceStatusStateStandbySpare string = "StandbySpare"
	// ResourceStatusStateInTest captures enum value "InTest"
	ResourceStatusStateInTest string = "InTest"
	// ResourceStatusStateStarting captures enum value "Starting"
	ResourceStatusStateStarting string = "Starting"
	// ResourceStatusStateAbsent captures enum value "Absent"
	ResourceStatusStateAbsent string = "Absent"
)

// prop value enum
func (m *ResourceStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("State", "body", m.State); err != nil {
		return err
	}

	return nil
}