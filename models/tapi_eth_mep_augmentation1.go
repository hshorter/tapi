// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiEthMepAugmentation1 tapi eth mep augmentation1
// swagger:model tapi.eth.MepAugmentation1
type TapiEthMepAugmentation1 struct {

	// none
	EthMepSpec *TapiEthEthMepSpec `json:"eth-mep-spec,omitempty"`
}

// Validate validates this tapi eth mep augmentation1
func (m *TapiEthMepAugmentation1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthMepSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiEthMepAugmentation1) validateEthMepSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.EthMepSpec) { // not required
		return nil
	}

	if m.EthMepSpec != nil {
		if err := m.EthMepSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth-mep-spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiEthMepAugmentation1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiEthMepAugmentation1) UnmarshalBinary(b []byte) error {
	var res TapiEthMepAugmentation1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}