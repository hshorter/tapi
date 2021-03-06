// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthPmCurrentDataAugmentation10 tapi eth pm current data augmentation10
// swagger:model tapi.eth.PmCurrentDataAugmentation10
type TapiEthPmCurrentDataAugmentation10 struct {

	// none
	EthTestResultData *TapiEthEthTestResultData `json:"eth-test-result-data,omitempty"`
}

// Validate validates this tapi eth pm current data augmentation10
func (m *TapiEthPmCurrentDataAugmentation10) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthTestResultData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthPmCurrentDataAugmentation10) validateEthTestResultData(formats strfmt.Registry) error {

	if swag.IsZero(m.EthTestResultData) { // not required
		return nil
	}

	if m.EthTestResultData != nil {
		if err := m.EthTestResultData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-test-result-data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation10) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthPmCurrentDataAugmentation10) UnmarshalBinary(b []byte) error {
	var res TapiEthPmCurrentDataAugmentation10
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
