// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiOduMepAugmentation2 tapi odu mep augmentation2
// swagger:model tapi.odu.MepAugmentation2
type TapiOduMepAugmentation2 struct {

	// none
	OduConnectionEndPointSpec *TapiOduOduConnectionEndPointSpec `json:"odu-connection-end-point-spec,omitempty"`
}

// Validate validates this tapi odu mep augmentation2
func (m *TapiOduMepAugmentation2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOduConnectionEndPointSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOduMepAugmentation2) validateOduConnectionEndPointSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.OduConnectionEndPointSpec) { // not required
		return nil
	}

	if m.OduConnectionEndPointSpec != nil {
		if err := m.OduConnectionEndPointSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("odu-connection-end-point-spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOduMepAugmentation2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOduMepAugmentation2) UnmarshalBinary(b []byte) error {
	var res TapiOduMepAugmentation2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
