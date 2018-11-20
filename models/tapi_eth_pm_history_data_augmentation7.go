// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmHistoryDataAugmentation7 tapi eth pm history data augmentation7
// swagger:model tapi.eth.PmHistoryDataAugmentation7
type TapiEthPmHistoryDataAugmentation7 struct {

	// none
	EthProActive1DmPerformanceData *TapiEthEthProActive1DmPerformanceData `json:"eth-pro-active-1-dm-performance-data,omitempty"`
}

// Validate validates this tapi eth pm history data augmentation7
func (m *TapiEthPmHistoryDataAugmentation7) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthProActive1DmPerformanceData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmHistoryDataAugmentation7) validateEthProActive1DmPerformanceData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthProActive1DmPerformanceData) { // not required
		return nil
	}

	if m.EthProActive1DmPerformanceData != nil {
		if err := m.EthProActive1DmPerformanceData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-pro-active-1-dm-performance-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthPmHistoryDataAugmentation7) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmHistoryDataAugmentation7) UnmarshalBinary(b []byte) error {
	var res TapiEthPmHistoryDataAugmentation7
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
