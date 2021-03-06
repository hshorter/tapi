// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TapiCommonOperationalState tapi common operational state
// swagger:model tapi.common.OperationalState
type TapiCommonOperationalState string

const (

	// TapiCommonOperationalStateDISABLED captures enum value "DISABLED"
	TapiCommonOperationalStateDISABLED TapiCommonOperationalState = "DISABLED"

	// TapiCommonOperationalStateENABLED captures enum value "ENABLED"
	TapiCommonOperationalStateENABLED TapiCommonOperationalState = "ENABLED"
)

// for schema
var tapiCommonOperationalStateEnum []interface{}

func init() {
	var res []TapiCommonOperationalState
	if err := json.Unmarshal([]byte(`["DISABLED","ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiCommonOperationalStateEnum = append(tapiCommonOperationalStateEnum, v)
	}
}

func (m TapiCommonOperationalState) validateTapiCommonOperationalStateEnum(path, location string, value TapiCommonOperationalState) error {
	if err := validate.Enum(path, location, value, tapiCommonOperationalStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi common operational state
func (m TapiCommonOperationalState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiCommonOperationalStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
