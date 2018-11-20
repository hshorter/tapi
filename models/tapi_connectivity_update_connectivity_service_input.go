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

// TapiConnectivityUpdateConnectivityServiceInput tapi connectivity update connectivity service input
// swagger:model tapi.connectivity.UpdateConnectivityServiceInput
type TapiConnectivityUpdateConnectivityServiceInput struct {

	// input
	Input *TapiConnectivityUpdateConnectivityServiceInputInput `json:"input,omitempty"`
}

// Validate validates this tapi connectivity update connectivity service input
func (m *TapiConnectivityUpdateConnectivityServiceInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityUpdateConnectivityServiceInput) validateInput(formats strfmt.Registry) error {

	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityUpdateConnectivityServiceInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityUpdateConnectivityServiceInput) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityUpdateConnectivityServiceInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiConnectivityUpdateConnectivityServiceInputInput tapi connectivity update connectivity service input input
// swagger:model TapiConnectivityUpdateConnectivityServiceInputInput
type TapiConnectivityUpdateConnectivityServiceInputInput struct {

	// none
	ConnectivityConstraint *TapiConnectivityConnectivityConstraint `json:"connectivity-constraint,omitempty"`

	// none
	EndPoint *TapiConnectivityConnectivityServiceEndPoint `json:"end-point,omitempty"`

	// none
	ResilienceConstraint []*TapiConnectivityResilienceConstraint `json:"resilience-constraint"`

	// none
	RoutingConstraint *TapiPathComputationRoutingConstraint `json:"routing-constraint,omitempty"`

	// none
	ServiceIDOrName string `json:"service-id-or-name,omitempty"`

	// none
	State string `json:"state,omitempty"`

	// none
	TopologyConstraint *TapiPathComputationTopologyConstraint `json:"topology-constraint,omitempty"`
}

// Validate validates this tapi connectivity update connectivity service input input
func (m *TapiConnectivityUpdateConnectivityServiceInputInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectivityConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResilienceConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopologyConstraint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityUpdateConnectivityServiceInputInput) validateConnectivityConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectivityConstraint) { // not required
		return nil
	}

	if m.ConnectivityConstraint != nil {
		if err := m.ConnectivityConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "connectivity-constraint")
			}
			return err
		}
	}

	return nil
}

func (m *TapiConnectivityUpdateConnectivityServiceInputInput) validateEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.EndPoint) { // not required
		return nil
	}

	if m.EndPoint != nil {
		if err := m.EndPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "end-point")
			}
			return err
		}
	}

	return nil
}

func (m *TapiConnectivityUpdateConnectivityServiceInputInput) validateResilienceConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.ResilienceConstraint) { // not required
		return nil
	}

	for i := 0; i < len(m.ResilienceConstraint); i++ {
		if swag.IsZero(m.ResilienceConstraint[i]) { // not required
			continue
		}

		if m.ResilienceConstraint[i] != nil {
			if err := m.ResilienceConstraint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input" + "." + "resilience-constraint" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityUpdateConnectivityServiceInputInput) validateRoutingConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingConstraint) { // not required
		return nil
	}

	if m.RoutingConstraint != nil {
		if err := m.RoutingConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "routing-constraint")
			}
			return err
		}
	}

	return nil
}

func (m *TapiConnectivityUpdateConnectivityServiceInputInput) validateTopologyConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.TopologyConstraint) { // not required
		return nil
	}

	if m.TopologyConstraint != nil {
		if err := m.TopologyConstraint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input" + "." + "topology-constraint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityUpdateConnectivityServiceInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityUpdateConnectivityServiceInputInput) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityUpdateConnectivityServiceInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
