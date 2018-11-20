// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthOnDemand1wayMeasurementJob tapi eth eth on demand1way measurement job
// swagger:model tapi.eth.EthOnDemand1wayMeasurementJob
type TapiEthEthOnDemand1wayMeasurementJob struct {

	// none
	OnDemandControl1waySink *TapiEthEthOnDemandMeasurementJobControlSink `json:"on-demand-control-1way-sink,omitempty"`

	// none
	OnDemandControl1waySource *TapiEthEthOnDemandMeasurementJobControlSource `json:"on-demand-control-1way-source,omitempty"`
}

// Validate validates this tapi eth eth on demand1way measurement job
func (m *TapiEthEthOnDemand1wayMeasurementJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOnDemandControl1waySink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnDemandControl1waySource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthOnDemand1wayMeasurementJob) validateOnDemandControl1waySink(formats strfmt.Registry) error {

	if swag.IsZero(m.OnDemandControl1waySink) { // not required
		return nil
	}

	if m.OnDemandControl1waySink != nil {
		if err := m.OnDemandControl1waySink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on-demand-control-1way-sink")
			}
			return err
		}
	}

	return nil
}

func (m *TapiEthEthOnDemand1wayMeasurementJob) validateOnDemandControl1waySource(formats strfmt.Registry) error {

	if swag.IsZero(m.OnDemandControl1waySource) { // not required
		return nil
	}

	if m.OnDemandControl1waySource != nil {
		if err := m.OnDemandControl1waySource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on-demand-control-1way-source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthOnDemand1wayMeasurementJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthOnDemand1wayMeasurementJob) UnmarshalBinary(b []byte) error {
	var res TapiEthEthOnDemand1wayMeasurementJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}