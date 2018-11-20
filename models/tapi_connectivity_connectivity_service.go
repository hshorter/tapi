// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityConnectivityService tapi connectivity connectivity service
// swagger:model tapi.connectivity.ConnectivityService
type TapiConnectivityConnectivityService struct {
	TapiCommonAdminStatePac

	TapiCommonGlobalClass

	TapiConnectivityConnectivityConstraint

	TapiConnectivityResilienceConstraint

	TapiPathComputationRoutingConstraint

	TapiPathComputationTopologyConstraint

	// none
	Connection []*TapiConnectivityConnectionRef `json:"connection"`

	// none
	EndPoint []*TapiConnectivityConnectivityServiceEndPoint `json:"end-point"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityConnectivityService) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonAdminStatePac
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonAdminStatePac = aO0

	// AO1
	var aO1 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO1

	// AO2
	var aO2 TapiConnectivityConnectivityConstraint
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiConnectivityConnectivityConstraint = aO2

	// AO3
	var aO3 TapiConnectivityResilienceConstraint
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiConnectivityResilienceConstraint = aO3

	// AO4
	var aO4 TapiPathComputationRoutingConstraint
	if err := swag.ReadJSON(raw, &aO4); err != nil {
		return err
	}
	m.TapiPathComputationRoutingConstraint = aO4

	// AO5
	var aO5 TapiPathComputationTopologyConstraint
	if err := swag.ReadJSON(raw, &aO5); err != nil {
		return err
	}
	m.TapiPathComputationTopologyConstraint = aO5

	// AO6
	var dataAO6 struct {
		Connection []*TapiConnectivityConnectionRef `json:"connection"`

		EndPoint []*TapiConnectivityConnectivityServiceEndPoint `json:"end-point"`
	}
	if err := swag.ReadJSON(raw, &dataAO6); err != nil {
		return err
	}

	m.Connection = dataAO6.Connection

	m.EndPoint = dataAO6.EndPoint

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityConnectivityService) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 7)

	aO0, err := swag.WriteJSON(m.TapiCommonAdminStatePac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiConnectivityConnectivityConstraint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiConnectivityResilienceConstraint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	aO4, err := swag.WriteJSON(m.TapiPathComputationRoutingConstraint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO4)

	aO5, err := swag.WriteJSON(m.TapiPathComputationTopologyConstraint)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO5)

	var dataAO6 struct {
		Connection []*TapiConnectivityConnectionRef `json:"connection"`

		EndPoint []*TapiConnectivityConnectivityServiceEndPoint `json:"end-point"`
	}

	dataAO6.Connection = m.Connection

	dataAO6.EndPoint = m.EndPoint

	jsonDataAO6, errAO6 := swag.WriteJSON(dataAO6)
	if errAO6 != nil {
		return nil, errAO6
	}
	_parts = append(_parts, jsonDataAO6)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity connectivity service
func (m *TapiConnectivityConnectivityService) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonAdminStatePac
	if err := m.TapiCommonAdminStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiConnectivityConnectivityConstraint
	if err := m.TapiConnectivityConnectivityConstraint.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiConnectivityResilienceConstraint
	if err := m.TapiConnectivityResilienceConstraint.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiPathComputationRoutingConstraint
	if err := m.TapiPathComputationRoutingConstraint.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiPathComputationTopologyConstraint
	if err := m.TapiPathComputationTopologyConstraint.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityConnectivityService) validateConnection(formats strfmt.Registry) error {

	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	for i := 0; i < len(m.Connection); i++ {
		if swag.IsZero(m.Connection[i]) { // not required
			continue
		}

		if m.Connection[i] != nil {
			if err := m.Connection[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connection" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityConnectivityService) validateEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPoint); i++ {
		if swag.IsZero(m.EndPoint[i]) { // not required
			continue
		}

		if m.EndPoint[i] != nil {
			if err := m.EndPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("end-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityConnectivityService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityConnectivityService) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityConnectivityService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
