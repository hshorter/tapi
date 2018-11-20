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

// TapiOduMappingType tapi odu mapping type
// swagger:model tapi.odu.MappingType
type TapiOduMappingType string

const (

	// TapiOduMappingTypeAMP captures enum value "AMP"
	TapiOduMappingTypeAMP TapiOduMappingType = "AMP"

	// TapiOduMappingTypeBMP captures enum value "BMP"
	TapiOduMappingTypeBMP TapiOduMappingType = "BMP"

	// TapiOduMappingTypeGFPF captures enum value "GFP-F"
	TapiOduMappingTypeGFPF TapiOduMappingType = "GFP-F"

	// TapiOduMappingTypeGMP captures enum value "GMP"
	TapiOduMappingTypeGMP TapiOduMappingType = "GMP"

	// TapiOduMappingTypeTTPGFPBMP captures enum value "TTP_GFP_BMP"
	TapiOduMappingTypeTTPGFPBMP TapiOduMappingType = "TTP_GFP_BMP"

	// TapiOduMappingTypeNULL captures enum value "NULL"
	TapiOduMappingTypeNULL TapiOduMappingType = "NULL"
)

// for schema
var tapiOduMappingTypeEnum []interface{}

func init() {
	var res []TapiOduMappingType
	if err := json.Unmarshal([]byte(`["AMP","BMP","GFP-F","GMP","TTP_GFP_BMP","NULL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiOduMappingTypeEnum = append(tapiOduMappingTypeEnum, v)
	}
}

func (m TapiOduMappingType) validateTapiOduMappingTypeEnum(path, location string, value TapiOduMappingType) error {
	if err := validate.Enum(path, location, value, tapiOduMappingTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi odu mapping type
func (m TapiOduMappingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiOduMappingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}