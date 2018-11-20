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

// TapiTopologyRuleType tapi topology rule type
// swagger:model tapi.topology.RuleType
type TapiTopologyRuleType string

const (

	// TapiTopologyRuleTypeFORWARDING captures enum value "FORWARDING"
	TapiTopologyRuleTypeFORWARDING TapiTopologyRuleType = "FORWARDING"

	// TapiTopologyRuleTypeCAPACITY captures enum value "CAPACITY"
	TapiTopologyRuleTypeCAPACITY TapiTopologyRuleType = "CAPACITY"

	// TapiTopologyRuleTypeCOST captures enum value "COST"
	TapiTopologyRuleTypeCOST TapiTopologyRuleType = "COST"

	// TapiTopologyRuleTypeTIMING captures enum value "TIMING"
	TapiTopologyRuleTypeTIMING TapiTopologyRuleType = "TIMING"

	// TapiTopologyRuleTypeRISK captures enum value "RISK"
	TapiTopologyRuleTypeRISK TapiTopologyRuleType = "RISK"

	// TapiTopologyRuleTypeGROUPING captures enum value "GROUPING"
	TapiTopologyRuleTypeGROUPING TapiTopologyRuleType = "GROUPING"
)

// for schema
var tapiTopologyRuleTypeEnum []interface{}

func init() {
	var res []TapiTopologyRuleType
	if err := json.Unmarshal([]byte(`["FORWARDING","CAPACITY","COST","TIMING","RISK","GROUPING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiTopologyRuleTypeEnum = append(tapiTopologyRuleTypeEnum, v)
	}
}

func (m TapiTopologyRuleType) validateTapiTopologyRuleTypeEnum(path, location string, value TapiTopologyRuleType) error {
	if err := validate.Enum(path, location, value, tapiTopologyRuleTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi topology rule type
func (m TapiTopologyRuleType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiTopologyRuleTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
