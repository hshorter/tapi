// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiNetworkElementNetworkElement tapi network element network element
// swagger:model tapi.network.element.NetworkElement
type TapiNetworkElementNetworkElement struct {
	TapiNetworkElementAttributes

	TapiNetworkElementConnection

	TapiNetworkElementSite

	// The Network Element id.
	NetworkElementID string `json:"network-element-id,omitempty"`

	// The last time the Network Element was successfully discovered.
	NetworkElementLastSuccessfulDiscovery string `json:"network-element-lastSuccessfulDiscovery,omitempty"`

	// The Network Element model.
	NetworkElementModel string `json:"network-element-model,omitempty"`

	// The Network Element name.
	NetworkElementName string `json:"network-element-name,omitempty"`

	// The Network Element type.
	NetworkElementType string `json:"network-element-type,omitempty"`

	// The Network Element vendor.
	NetworkElementVendor string `json:"network-element-vendor,omitempty"`

	// The Network Element version.
	NetworkElementVersion string `json:"network-element-version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiNetworkElementNetworkElement) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiNetworkElementAttributes
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiNetworkElementAttributes = aO0

	// AO1
	var aO1 TapiNetworkElementConnection
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiNetworkElementConnection = aO1

	// AO2
	var aO2 TapiNetworkElementSite
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiNetworkElementSite = aO2

	// AO3
	var dataAO3 struct {
		NetworkElementID string `json:"network-element-id,omitempty"`

		NetworkElementLastSuccessfulDiscovery string `json:"network-element-lastSuccessfulDiscovery,omitempty"`

		NetworkElementModel string `json:"network-element-model,omitempty"`

		NetworkElementName string `json:"network-element-name,omitempty"`

		NetworkElementType string `json:"network-element-type,omitempty"`

		NetworkElementVendor string `json:"network-element-vendor,omitempty"`

		NetworkElementVersion string `json:"network-element-version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO3); err != nil {
		return err
	}

	m.NetworkElementID = dataAO3.NetworkElementID

	m.NetworkElementLastSuccessfulDiscovery = dataAO3.NetworkElementLastSuccessfulDiscovery

	m.NetworkElementModel = dataAO3.NetworkElementModel

	m.NetworkElementName = dataAO3.NetworkElementName

	m.NetworkElementType = dataAO3.NetworkElementType

	m.NetworkElementVendor = dataAO3.NetworkElementVendor

	m.NetworkElementVersion = dataAO3.NetworkElementVersion

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiNetworkElementNetworkElement) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.TapiNetworkElementAttributes)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiNetworkElementConnection)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiNetworkElementSite)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	var dataAO3 struct {
		NetworkElementID string `json:"network-element-id,omitempty"`

		NetworkElementLastSuccessfulDiscovery string `json:"network-element-lastSuccessfulDiscovery,omitempty"`

		NetworkElementModel string `json:"network-element-model,omitempty"`

		NetworkElementName string `json:"network-element-name,omitempty"`

		NetworkElementType string `json:"network-element-type,omitempty"`

		NetworkElementVendor string `json:"network-element-vendor,omitempty"`

		NetworkElementVersion string `json:"network-element-version,omitempty"`
	}

	dataAO3.NetworkElementID = m.NetworkElementID

	dataAO3.NetworkElementLastSuccessfulDiscovery = m.NetworkElementLastSuccessfulDiscovery

	dataAO3.NetworkElementModel = m.NetworkElementModel

	dataAO3.NetworkElementName = m.NetworkElementName

	dataAO3.NetworkElementType = m.NetworkElementType

	dataAO3.NetworkElementVendor = m.NetworkElementVendor

	dataAO3.NetworkElementVersion = m.NetworkElementVersion

	jsonDataAO3, errAO3 := swag.WriteJSON(dataAO3)
	if errAO3 != nil {
		return nil, errAO3
	}
	_parts = append(_parts, jsonDataAO3)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi network element network element
func (m *TapiNetworkElementNetworkElement) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiNetworkElementAttributes
	if err := m.TapiNetworkElementAttributes.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiNetworkElementConnection
	if err := m.TapiNetworkElementConnection.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiNetworkElementSite
	if err := m.TapiNetworkElementSite.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiNetworkElementNetworkElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiNetworkElementNetworkElement) UnmarshalBinary(b []byte) error {
	var res TapiNetworkElementNetworkElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
