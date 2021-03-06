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

// TapiOamOamcontextOamProfile tapi oam oamcontext oam profile
// swagger:model tapi.oam.oamcontext.OamProfile
type TapiOamOamcontextOamProfile struct {
	TapiCommonGlobalClass

	// none
	PmBinData []*TapiOamPmBinData `json:"pm-bin-data"`

	// none
	PmThresholdData []*TapiOamOamprofilePmThresholdData `json:"pm-threshold-data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiOamOamcontextOamProfile) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO0

	// AO1
	var dataAO1 struct {
		PmBinData []*TapiOamPmBinData `json:"pm-bin-data"`

		PmThresholdData []*TapiOamOamprofilePmThresholdData `json:"pm-threshold-data"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.PmBinData = dataAO1.PmBinData

	m.PmThresholdData = dataAO1.PmThresholdData

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiOamOamcontextOamProfile) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		PmBinData []*TapiOamPmBinData `json:"pm-bin-data"`

		PmThresholdData []*TapiOamOamprofilePmThresholdData `json:"pm-threshold-data"`
	}

	dataAO1.PmBinData = m.PmBinData

	dataAO1.PmThresholdData = m.PmThresholdData

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi oam oamcontext oam profile
func (m *TapiOamOamcontextOamProfile) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmBinData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmThresholdData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiOamOamcontextOamProfile) validatePmBinData(formats strfmt.Registry) error {

	if swag.IsZero(m.PmBinData) { // not required
		return nil
	}

	for i := 0; i < len(m.PmBinData); i++ {
		if swag.IsZero(m.PmBinData[i]) { // not required
			continue
		}

		if m.PmBinData[i] != nil {
			if err := m.PmBinData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pm-bin-data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiOamOamcontextOamProfile) validatePmThresholdData(formats strfmt.Registry) error {

	if swag.IsZero(m.PmThresholdData) { // not required
		return nil
	}

	for i := 0; i < len(m.PmThresholdData); i++ {
		if swag.IsZero(m.PmThresholdData[i]) { // not required
			continue
		}

		if m.PmThresholdData[i] != nil {
			if err := m.PmThresholdData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pm-threshold-data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiOamOamcontextOamProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiOamOamcontextOamProfile) UnmarshalBinary(b []byte) error {
	var res TapiOamOamcontextOamProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
