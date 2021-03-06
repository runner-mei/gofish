package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CIMCResetActionResetAction This is the base type for the reset action.
// swagger:model CIMC.ResetAction_ResetAction
type CIMCResetActionResetAction struct {

	// reset type
	// Required: true
	ResetType *string `json:"reset_type"`
}

// Validate validates this c i m c reset action reset action
func (m *CIMCResetActionResetAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResetType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cIMCResetActionResetActionTypeResetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["On","ForceOff","GracefulShutdown","GracefulRestart","ForceRestart","Nmi","ForceOn","PushPowerButton"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cIMCResetActionResetActionTypeResetTypePropEnum = append(cIMCResetActionResetActionTypeResetTypePropEnum, v)
	}
}

const (
	// CIMCResetActionResetActionResetTypeOn captures enum value "On"
	CIMCResetActionResetActionResetTypeOn string = "On"
	// CIMCResetActionResetActionResetTypeForceOff captures enum value "ForceOff"
	CIMCResetActionResetActionResetTypeForceOff string = "ForceOff"
	// CIMCResetActionResetActionResetTypeGracefulShutdown captures enum value "GracefulShutdown"
	CIMCResetActionResetActionResetTypeGracefulShutdown string = "GracefulShutdown"
	// CIMCResetActionResetActionResetTypeGracefulRestart captures enum value "GracefulRestart"
	CIMCResetActionResetActionResetTypeGracefulRestart string = "GracefulRestart"
	// CIMCResetActionResetActionResetTypeForceRestart captures enum value "ForceRestart"
	CIMCResetActionResetActionResetTypeForceRestart string = "ForceRestart"
	// CIMCResetActionResetActionResetTypeNmi captures enum value "Nmi"
	CIMCResetActionResetActionResetTypeNmi string = "Nmi"
	// CIMCResetActionResetActionResetTypeForceOn captures enum value "ForceOn"
	CIMCResetActionResetActionResetTypeForceOn string = "ForceOn"
	// CIMCResetActionResetActionResetTypePushPowerButton captures enum value "PushPowerButton"
	CIMCResetActionResetActionResetTypePushPowerButton string = "PushPowerButton"
)

// prop value enum
func (m *CIMCResetActionResetAction) validateResetTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cIMCResetActionResetActionTypeResetTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CIMCResetActionResetAction) validateResetType(formats strfmt.Registry) error {

	if err := validate.Required("reset_type", "body", m.ResetType); err != nil {
		return err
	}

	// value enum
	if err := m.validateResetTypeEnum("reset_type", "body", *m.ResetType); err != nil {
		return err
	}

	return nil
}
