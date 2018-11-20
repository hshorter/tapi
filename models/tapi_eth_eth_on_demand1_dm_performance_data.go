// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthEthOnDemand1DmPerformanceData tapi eth eth on demand1 dm performance data
// swagger:model tapi.eth.EthOnDemand1DmPerformanceData
type TapiEthEthOnDemand1DmPerformanceData struct {

	// This attribute contains the results of an on-demand frame delay measurement job in the ingress direction.
	OnDemandNearEnd1DmParameters *TapiEthOnDemandDmPerformanceParameters `json:"on-demand-near-end-1-dm-parameters,omitempty"`
}

// Validate validates this tapi eth eth on demand1 dm performance data
func (m *TapiEthEthOnDemand1DmPerformanceData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOnDemandNearEnd1DmParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthEthOnDemand1DmPerformanceData) validateOnDemandNearEnd1DmParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.OnDemandNearEnd1DmParameters) { // not required
		return nil
	}

	if m.OnDemandNearEnd1DmParameters != nil {
		if err := m.OnDemandNearEnd1DmParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on-demand-near-end-1-dm-parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthEthOnDemand1DmPerformanceData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthEthOnDemand1DmPerformanceData) UnmarshalBinary(b []byte) error {
	var res TapiEthEthOnDemand1DmPerformanceData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}