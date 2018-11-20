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

// TapiPathComputationDiversityPolicy tapi path computation diversity policy
// swagger:model tapi.path.computation.DiversityPolicy
type TapiPathComputationDiversityPolicy string

const (

	// TapiPathComputationDiversityPolicySRLG captures enum value "SRLG"
	TapiPathComputationDiversityPolicySRLG TapiPathComputationDiversityPolicy = "SRLG"

	// TapiPathComputationDiversityPolicySRNG captures enum value "SRNG"
	TapiPathComputationDiversityPolicySRNG TapiPathComputationDiversityPolicy = "SRNG"

	// TapiPathComputationDiversityPolicySNG captures enum value "SNG"
	TapiPathComputationDiversityPolicySNG TapiPathComputationDiversityPolicy = "SNG"

	// TapiPathComputationDiversityPolicyNODE captures enum value "NODE"
	TapiPathComputationDiversityPolicyNODE TapiPathComputationDiversityPolicy = "NODE"

	// TapiPathComputationDiversityPolicyLINK captures enum value "LINK"
	TapiPathComputationDiversityPolicyLINK TapiPathComputationDiversityPolicy = "LINK"
)

// for schema
var tapiPathComputationDiversityPolicyEnum []interface{}

func init() {
	var res []TapiPathComputationDiversityPolicy
	if err := json.Unmarshal([]byte(`["SRLG","SRNG","SNG","NODE","LINK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiPathComputationDiversityPolicyEnum = append(tapiPathComputationDiversityPolicyEnum, v)
	}
}

func (m TapiPathComputationDiversityPolicy) validateTapiPathComputationDiversityPolicyEnum(path, location string, value TapiPathComputationDiversityPolicy) error {
	if err := validate.Enum(path, location, value, tapiPathComputationDiversityPolicyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi path computation diversity policy
func (m TapiPathComputationDiversityPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiPathComputationDiversityPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
