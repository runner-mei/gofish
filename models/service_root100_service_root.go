package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ServiceRoot100ServiceRoot This object represents the root Redfish service.
// swagger:model ServiceRoot.1.0.0_ServiceRoot
type ServiceRoot100ServiceRoot struct {

	// at odata context
	// Read Only: true
	AtOdataContext strfmt.URI `json:"@odata.context,omitempty"`

	// at odata id
	// Read Only: true
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`

	// at odata type
	// Read Only: true
	AtOdataType string `json:"@odata.type,omitempty"`

	// This is a link to the Account Service.
	// Read Only: true
	AccountService *Odata400IDRef `json:"AccountService,omitempty"`

	// This is a link to a collection of Chassis.
	// Read Only: true
	Chassis *ChassisCollectionChassisCollection `json:"Chassis,omitempty"`

	// Provides a description of this resource and is used for commonality  in the schema definitions.
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// This is a link to the EventService.
	// Read Only: true
	EventService *Odata400IDRef `json:"EventService,omitempty"`

	// Uniquely identifies the resource within the collection of like resources.
	// Read Only: true
	ID string `json:"Id,omitempty"`

	// This is a link to a collection of Json-Schema files.
	// Read Only: true
	JSONSchemas *JSONSchemaFileCollectionJSONSchemaFileCollection `json:"JsonSchemas,omitempty"`

	// links
	// Required: true
	Links *ServiceRoot100ServiceRootLinks `json:"Links"`

	// This is a link to a collection of Managers.
	// Read Only: true
	Managers *ManagerCollectionManagerCollection `json:"Managers,omitempty"`

	// The name of the resource or array element.
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	Oem ResourceOem `json:"Oem,omitempty"`

	// The version of the Redfish service
	// Read Only: true
	// Pattern: ^\d+\.\d+\.\d+$
	RedfishVersion string `json:"RedfishVersion,omitempty"`

	// This is a link to a collection of Registries.
	// Read Only: true
	Registries *MessageRegistryFileCollectionMessageRegistryFileCollection `json:"Registries,omitempty"`

	// This is a link to the Sessions Service.
	// Read Only: true
	SessionService *Odata400IDRef `json:"SessionService,omitempty"`

	// This is a link to a collection of Systems.
	// Read Only: true
	Systems *ComputerSystemCollectionComputerSystemCollection `json:"Systems,omitempty"`

	// This is a link to the Task Service.
	// Read Only: true
	Tasks *Odata400IDRef `json:"Tasks,omitempty"`

	// Unique identifier for a service instance. When SSDP is used, this value should be an exact match of the UUID value returned in a 200OK from an SSDP M-SEARCH request during discovery.
	// Read Only: true
	// Pattern: ([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})
	UUID string `json:"UUID,omitempty"`
}

// Validate validates this service root 1 0 0 service root
func (m *ServiceRoot100ServiceRoot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountService(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChassis(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventService(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJSONSchemas(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRedfishVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegistries(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSessionService(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSystems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceRoot100ServiceRoot) validateAccountService(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountService) { // not required
		return nil
	}

	if m.AccountService != nil {

		if err := m.AccountService.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateChassis(formats strfmt.Registry) error {

	if swag.IsZero(m.Chassis) { // not required
		return nil
	}

	if m.Chassis != nil {

		if err := m.Chassis.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateEventService(formats strfmt.Registry) error {

	if swag.IsZero(m.EventService) { // not required
		return nil
	}

	if m.EventService != nil {

		if err := m.EventService.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateJSONSchemas(formats strfmt.Registry) error {

	if swag.IsZero(m.JSONSchemas) { // not required
		return nil
	}

	if m.JSONSchemas != nil {

		if err := m.JSONSchemas.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("Links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateManagers(formats strfmt.Registry) error {

	if swag.IsZero(m.Managers) { // not required
		return nil
	}

	if m.Managers != nil {

		if err := m.Managers.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateRedfishVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.RedfishVersion) { // not required
		return nil
	}

	if err := validate.Pattern("RedfishVersion", "body", string(m.RedfishVersion), `^\d+\.\d+\.\d+$`); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateRegistries(formats strfmt.Registry) error {

	if swag.IsZero(m.Registries) { // not required
		return nil
	}

	if m.Registries != nil {

		if err := m.Registries.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateSessionService(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionService) { // not required
		return nil
	}

	if m.SessionService != nil {

		if err := m.SessionService.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateSystems(formats strfmt.Registry) error {

	if swag.IsZero(m.Systems) { // not required
		return nil
	}

	if m.Systems != nil {

		if err := m.Systems.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	if m.Tasks != nil {

		if err := m.Tasks.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ServiceRoot100ServiceRoot) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.Pattern("UUID", "body", string(m.UUID), `([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})`); err != nil {
		return err
	}

	return nil
}

// ServiceRoot100ServiceRootLinks The links object contains the links to other resources that are related to this resource.
// swagger:model ServiceRoot100ServiceRootLinks
type ServiceRoot100ServiceRootLinks struct {

	// Oem extension object.
	Oem ResourceOem `json:"Oem,omitempty"`

	// Link to a collection of Sessions
	// Read Only: true
	Sessions *SessionCollectionSessionCollection `json:"Sessions,omitempty"`
}

// Validate validates this service root100 service root links
func (m *ServiceRoot100ServiceRootLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSessions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceRoot100ServiceRootLinks) validateSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.Sessions) { // not required
		return nil
	}

	if m.Sessions != nil {

		if err := m.Sessions.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
