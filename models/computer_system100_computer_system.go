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

// ComputerSystem100ComputerSystem This schema defines a computer system and its respective properties.  A computer system represents a machine (physical or virtual) and the local resources such as memory, cpu and other devices that can be accessed from that machine.
// swagger:model ComputerSystem.1.0.0_ComputerSystem
type ComputerSystem100ComputerSystem struct {

	// at odata context
	// Read Only: true
	AtOdataContext strfmt.URI `json:"@odata.context,omitempty"`

	// at odata id
	// Read Only: true
	AtOdataID strfmt.URI `json:"@odata.id,omitempty"`

	// at odata type
	// Read Only: true
	AtOdataType string `json:"@odata.type,omitempty"`

	// actions
	Actions *ComputerSystem100ComputerSystemActions `json:"Actions,omitempty"`

	// The user definable tag that can be used to track this computer system for inventory or other client purposes
	AssetTag string `json:"AssetTag,omitempty"`

	// The version of the system BIOS or primary system firmware.
	BiosVersion string `json:"BiosVersion,omitempty"`

	// Information about the boot settings for this system
	Boot *ComputerSystem100Boot `json:"Boot,omitempty"`

	// Provides a description of this resource and is used for commonality  in the schema definitions.
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// A reference to the collection of Ethernet interfaces associated with this system
	// Read Only: true
	EthernetInterfaces *EthernetInterfaceCollectionEthernetInterfaceCollection `json:"EthernetInterfaces,omitempty"`

	// The DNS Host Name, without any domain information
	HostName string `json:"HostName,omitempty"`

	// Uniquely identifies the resource within the collection of like resources.
	// Read Only: true
	ID string `json:"Id,omitempty"`

	// The state of the indicator LED, used to identify the system
	IndicatorLED string `json:"IndicatorLED,omitempty"`

	// links
	Links *ComputerSystem100ComputerSystemLinks `json:"Links,omitempty"`

	// A reference to the collection of Log Services associated with this system
	// Read Only: true
	LogServices *LogServiceCollectionLogServiceCollection `json:"LogServices,omitempty"`

	// The manufacturer or OEM of this system.
	// Read Only: true
	Manufacturer string `json:"Manufacturer,omitempty"`

	// This object describes the central memory of the system in general detail.
	MemorySummary *ComputerSystem100MemorySummary `json:"MemorySummary,omitempty"`

	// The model number for this system
	// Read Only: true
	Model string `json:"Model,omitempty"`

	// The name of the resource or array element.
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	Oem ResourceOem `json:"Oem,omitempty"`

	// The part number for this system
	// Read Only: true
	PartNumber string `json:"PartNumber,omitempty"`

	// This is the current power state of the system
	// Read Only: true
	PowerState string `json:"PowerState,omitempty"`

	// This object describes the central processors of the system in general detail.
	ProcessorSummary *ComputerSystem100ProcessorSummary `json:"ProcessorSummary,omitempty"`

	// A reference to the collection of Processors associated with this system
	// Read Only: true
	Processors *ProcessorCollectionProcessorCollection `json:"Processors,omitempty"`

	// The manufacturer SKU for this system
	// Read Only: true
	SKU string `json:"SKU,omitempty"`

	// The serial number for this system
	// Read Only: true
	SerialNumber string `json:"SerialNumber,omitempty"`

	// A reference to the collection of storage devices associated with this system
	// Read Only: true
	SimpleStorage *SimpleStorageCollectionSimpleStorageCollection `json:"SimpleStorage,omitempty"`

	// status
	Status *ResourceStatus `json:"Status,omitempty"`

	// The type of computer system represented by this resource.
	// Read Only: true
	SystemType string `json:"SystemType,omitempty"`

	// The universal unique identifier (UUID) for this system
	// Read Only: true
	// Pattern: ([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})
	UUID string `json:"UUID,omitempty"`
}

// Validate validates this computer system 1 0 0 computer system
func (m *ComputerSystem100ComputerSystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBoot(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEthernetInterfaces(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIndicatorLED(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLogServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemorySummary(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePowerState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcessorSummary(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcessors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSimpleStorage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSystemType(formats); err != nil {
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

func (m *ComputerSystem100ComputerSystem) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {

		if err := m.Actions.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateBoot(formats strfmt.Registry) error {

	if swag.IsZero(m.Boot) { // not required
		return nil
	}

	if m.Boot != nil {

		if err := m.Boot.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateEthernetInterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.EthernetInterfaces) { // not required
		return nil
	}

	if m.EthernetInterfaces != nil {

		if err := m.EthernetInterfaces.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var computerSystem100ComputerSystemTypeIndicatorLEDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Lit","Blinking","Off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerSystem100ComputerSystemTypeIndicatorLEDPropEnum = append(computerSystem100ComputerSystemTypeIndicatorLEDPropEnum, v)
	}
}

const (
	// ComputerSystem100ComputerSystemIndicatorLEDUnknown captures enum value "Unknown"
	ComputerSystem100ComputerSystemIndicatorLEDUnknown string = "Unknown"
	// ComputerSystem100ComputerSystemIndicatorLEDLit captures enum value "Lit"
	ComputerSystem100ComputerSystemIndicatorLEDLit string = "Lit"
	// ComputerSystem100ComputerSystemIndicatorLEDBlinking captures enum value "Blinking"
	ComputerSystem100ComputerSystemIndicatorLEDBlinking string = "Blinking"
	// ComputerSystem100ComputerSystemIndicatorLEDOff captures enum value "Off"
	ComputerSystem100ComputerSystemIndicatorLEDOff string = "Off"
)

// prop value enum
func (m *ComputerSystem100ComputerSystem) validateIndicatorLEDEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, computerSystem100ComputerSystemTypeIndicatorLEDPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComputerSystem100ComputerSystem) validateIndicatorLED(formats strfmt.Registry) error {

	if swag.IsZero(m.IndicatorLED) { // not required
		return nil
	}

	// value enum
	if err := m.validateIndicatorLEDEnum("IndicatorLED", "body", m.IndicatorLED); err != nil {
		return err
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateLogServices(formats strfmt.Registry) error {

	if swag.IsZero(m.LogServices) { // not required
		return nil
	}

	if m.LogServices != nil {

		if err := m.LogServices.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateMemorySummary(formats strfmt.Registry) error {

	if swag.IsZero(m.MemorySummary) { // not required
		return nil
	}

	if m.MemorySummary != nil {

		if err := m.MemorySummary.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var computerSystem100ComputerSystemTypePowerStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["On","Off","Unknown","Reset"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerSystem100ComputerSystemTypePowerStatePropEnum = append(computerSystem100ComputerSystemTypePowerStatePropEnum, v)
	}
}

const (
	// ComputerSystem100ComputerSystemPowerStateOn captures enum value "On"
	ComputerSystem100ComputerSystemPowerStateOn string = "On"
	// ComputerSystem100ComputerSystemPowerStateOff captures enum value "Off"
	ComputerSystem100ComputerSystemPowerStateOff string = "Off"
	// ComputerSystem100ComputerSystemPowerStateUnknown captures enum value "Unknown"
	ComputerSystem100ComputerSystemPowerStateUnknown string = "Unknown"
	// ComputerSystem100ComputerSystemPowerStateReset captures enum value "Reset"
	ComputerSystem100ComputerSystemPowerStateReset string = "Reset"
)

// prop value enum
func (m *ComputerSystem100ComputerSystem) validatePowerStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, computerSystem100ComputerSystemTypePowerStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComputerSystem100ComputerSystem) validatePowerState(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerState) { // not required
		return nil
	}

	// value enum
	if err := m.validatePowerStateEnum("PowerState", "body", m.PowerState); err != nil {
		return err
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateProcessorSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessorSummary) { // not required
		return nil
	}

	if m.ProcessorSummary != nil {

		if err := m.ProcessorSummary.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateProcessors(formats strfmt.Registry) error {

	if swag.IsZero(m.Processors) { // not required
		return nil
	}

	if m.Processors != nil {

		if err := m.Processors.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateSimpleStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.SimpleStorage) { // not required
		return nil
	}

	if m.SimpleStorage != nil {

		if err := m.SimpleStorage.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateStatus(formats strfmt.Registry) error {

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

var computerSystem100ComputerSystemTypeSystemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Physical","Virtual","OS","PhysicallyPartitioned","VirtuallyPartitioned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerSystem100ComputerSystemTypeSystemTypePropEnum = append(computerSystem100ComputerSystemTypeSystemTypePropEnum, v)
	}
}

const (
	// ComputerSystem100ComputerSystemSystemTypePhysical captures enum value "Physical"
	ComputerSystem100ComputerSystemSystemTypePhysical string = "Physical"
	// ComputerSystem100ComputerSystemSystemTypeVirtual captures enum value "Virtual"
	ComputerSystem100ComputerSystemSystemTypeVirtual string = "Virtual"
	// ComputerSystem100ComputerSystemSystemTypeOS captures enum value "OS"
	ComputerSystem100ComputerSystemSystemTypeOS string = "OS"
	// ComputerSystem100ComputerSystemSystemTypePhysicallyPartitioned captures enum value "PhysicallyPartitioned"
	ComputerSystem100ComputerSystemSystemTypePhysicallyPartitioned string = "PhysicallyPartitioned"
	// ComputerSystem100ComputerSystemSystemTypeVirtuallyPartitioned captures enum value "VirtuallyPartitioned"
	ComputerSystem100ComputerSystemSystemTypeVirtuallyPartitioned string = "VirtuallyPartitioned"
)

// prop value enum
func (m *ComputerSystem100ComputerSystem) validateSystemTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, computerSystem100ComputerSystemTypeSystemTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComputerSystem100ComputerSystem) validateSystemType(formats strfmt.Registry) error {

	if swag.IsZero(m.SystemType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSystemTypeEnum("SystemType", "body", m.SystemType); err != nil {
		return err
	}

	return nil
}

func (m *ComputerSystem100ComputerSystem) validateUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.Pattern("UUID", "body", string(m.UUID), `([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})`); err != nil {
		return err
	}

	return nil
}

// ComputerSystem100ComputerSystemActions The Actions object contains the available custom actions on this resource.
// swagger:model ComputerSystem100ComputerSystemActions
type ComputerSystem100ComputerSystemActions struct {

	// computer system reset
	NrComputerSystemReset *ComputerSystem100Reset `json:"#ComputerSystem.Reset,omitempty"`

	// oem
	Oem interface{} `json:"Oem,omitempty"`
}

// Validate validates this computer system100 computer system actions
func (m *ComputerSystem100ComputerSystemActions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNrComputerSystemReset(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerSystem100ComputerSystemActions) validateNrComputerSystemReset(formats strfmt.Registry) error {

	if swag.IsZero(m.NrComputerSystemReset) { // not required
		return nil
	}

	if m.NrComputerSystemReset != nil {

		if err := m.NrComputerSystemReset.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// ComputerSystem100ComputerSystemLinks Contains links to other resources that are related to this resource.
// swagger:model ComputerSystem100ComputerSystemLinks
type ComputerSystem100ComputerSystemLinks struct {

	// An array of references to the chassis in which this system is contained
	// Read Only: true
	Chassis []*Odata400IDRef `json:"Chassis"`

	// chassis at odata count
	// Read Only: true
	ChassisAtOdataCount float64 `json:"Chassis@odata.count,omitempty"`

	// chassis at odata navigation link
	ChassisAtOdataNavigationLink *Odata400IDRef `json:"Chassis@odata.navigationLink,omitempty"`

	// An array of ID[s] of resources that cool this computer system. Normally the ID will be a chassis or a specific set of fans.
	// Read Only: true
	CooledBy []*Odata400IDRef `json:"CooledBy"`

	// cooled by at odata count
	// Read Only: true
	CooledByAtOdataCount float64 `json:"CooledBy@odata.count,omitempty"`

	// cooled by at odata navigation link
	CooledByAtOdataNavigationLink *Odata400IDRef `json:"CooledBy@odata.navigationLink,omitempty"`

	// An array of references to the Managers responsible for this system
	// Read Only: true
	ManagedBy []*Odata400IDRef `json:"ManagedBy"`

	// managed by at odata count
	// Read Only: true
	ManagedByAtOdataCount float64 `json:"ManagedBy@odata.count,omitempty"`

	// managed by at odata navigation link
	ManagedByAtOdataNavigationLink *Odata400IDRef `json:"ManagedBy@odata.navigationLink,omitempty"`

	// Oem extension object.
	Oem ResourceOem `json:"Oem,omitempty"`

	// An array of ID[s] of resources that power this computer system. Normally the ID will be a chassis or a specific set of powerSupplies
	// Read Only: true
	PoweredBy []*Odata400IDRef `json:"PoweredBy"`

	// powered by at odata count
	// Read Only: true
	PoweredByAtOdataCount float64 `json:"PoweredBy@odata.count,omitempty"`

	// powered by at odata navigation link
	PoweredByAtOdataNavigationLink *Odata400IDRef `json:"PoweredBy@odata.navigationLink,omitempty"`
}

// Validate validates this computer system100 computer system links
func (m *ComputerSystem100ComputerSystemLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChassis(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChassisAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCooledBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCooledByAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagedByAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePoweredBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePoweredByAtOdataNavigationLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerSystem100ComputerSystemLinks) validateChassis(formats strfmt.Registry) error {

	if swag.IsZero(m.Chassis) { // not required
		return nil
	}

	for i := 0; i < len(m.Chassis); i++ {

		if swag.IsZero(m.Chassis[i]) { // not required
			continue
		}

		if m.Chassis[i] != nil {

			if err := m.Chassis[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ComputerSystem100ComputerSystemLinks) validateChassisAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.ChassisAtOdataNavigationLink) { // not required
		return nil
	}

	if m.ChassisAtOdataNavigationLink != nil {

		if err := m.ChassisAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystemLinks) validateCooledBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CooledBy) { // not required
		return nil
	}

	for i := 0; i < len(m.CooledBy); i++ {

		if swag.IsZero(m.CooledBy[i]) { // not required
			continue
		}

		if m.CooledBy[i] != nil {

			if err := m.CooledBy[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ComputerSystem100ComputerSystemLinks) validateCooledByAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.CooledByAtOdataNavigationLink) { // not required
		return nil
	}

	if m.CooledByAtOdataNavigationLink != nil {

		if err := m.CooledByAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystemLinks) validateManagedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedBy) { // not required
		return nil
	}

	for i := 0; i < len(m.ManagedBy); i++ {

		if swag.IsZero(m.ManagedBy[i]) { // not required
			continue
		}

		if m.ManagedBy[i] != nil {

			if err := m.ManagedBy[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ComputerSystem100ComputerSystemLinks) validateManagedByAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedByAtOdataNavigationLink) { // not required
		return nil
	}

	if m.ManagedByAtOdataNavigationLink != nil {

		if err := m.ManagedByAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ComputerSystem100ComputerSystemLinks) validatePoweredBy(formats strfmt.Registry) error {

	if swag.IsZero(m.PoweredBy) { // not required
		return nil
	}

	for i := 0; i < len(m.PoweredBy); i++ {

		if swag.IsZero(m.PoweredBy[i]) { // not required
			continue
		}

		if m.PoweredBy[i] != nil {

			if err := m.PoweredBy[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ComputerSystem100ComputerSystemLinks) validatePoweredByAtOdataNavigationLink(formats strfmt.Registry) error {

	if swag.IsZero(m.PoweredByAtOdataNavigationLink) { // not required
		return nil
	}

	if m.PoweredByAtOdataNavigationLink != nil {

		if err := m.PoweredByAtOdataNavigationLink.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
