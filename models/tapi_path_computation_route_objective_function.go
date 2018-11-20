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

// TapiPathComputationRouteObjectiveFunction tapi path computation route objective function
// swagger:model tapi.path.computation.RouteObjectiveFunction
type TapiPathComputationRouteObjectiveFunction string

const (

	// TapiPathComputationRouteObjectiveFunctionMINWORKROUTEHOP captures enum value "MIN_WORK_ROUTE_HOP"
	TapiPathComputationRouteObjectiveFunctionMINWORKROUTEHOP TapiPathComputationRouteObjectiveFunction = "MIN_WORK_ROUTE_HOP"

	// TapiPathComputationRouteObjectiveFunctionMINWORKROUTECOST captures enum value "MIN_WORK_ROUTE_COST"
	TapiPathComputationRouteObjectiveFunctionMINWORKROUTECOST TapiPathComputationRouteObjectiveFunction = "MIN_WORK_ROUTE_COST"

	// TapiPathComputationRouteObjectiveFunctionMINWORKROUTELATENCY captures enum value "MIN_WORK_ROUTE_LATENCY"
	TapiPathComputationRouteObjectiveFunctionMINWORKROUTELATENCY TapiPathComputationRouteObjectiveFunction = "MIN_WORK_ROUTE_LATENCY"

	// TapiPathComputationRouteObjectiveFunctionMINSUMOFWORKANDPROTECTIONROUTEHOP captures enum value "MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_HOP"
	TapiPathComputationRouteObjectiveFunctionMINSUMOFWORKANDPROTECTIONROUTEHOP TapiPathComputationRouteObjectiveFunction = "MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_HOP"

	// TapiPathComputationRouteObjectiveFunctionMINSUMOFWORKANDPROTECTIONROUTECOST captures enum value "MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_COST"
	TapiPathComputationRouteObjectiveFunctionMINSUMOFWORKANDPROTECTIONROUTECOST TapiPathComputationRouteObjectiveFunction = "MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_COST"

	// TapiPathComputationRouteObjectiveFunctionMINSUMOFWORKANDPROTECTIONROUTELATENCY captures enum value "MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_LATENCY"
	TapiPathComputationRouteObjectiveFunctionMINSUMOFWORKANDPROTECTIONROUTELATENCY TapiPathComputationRouteObjectiveFunction = "MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_LATENCY"

	// TapiPathComputationRouteObjectiveFunctionLOADBALANCEMAXUNUSEDCAPACITY captures enum value "LOAD_BALANCE_MAX_UNUSED_CAPACITY"
	TapiPathComputationRouteObjectiveFunctionLOADBALANCEMAXUNUSEDCAPACITY TapiPathComputationRouteObjectiveFunction = "LOAD_BALANCE_MAX_UNUSED_CAPACITY"
)

// for schema
var tapiPathComputationRouteObjectiveFunctionEnum []interface{}

func init() {
	var res []TapiPathComputationRouteObjectiveFunction
	if err := json.Unmarshal([]byte(`["MIN_WORK_ROUTE_HOP","MIN_WORK_ROUTE_COST","MIN_WORK_ROUTE_LATENCY","MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_HOP","MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_COST","MIN_SUM_OF_WORK_AND_PROTECTION_ROUTE_LATENCY","LOAD_BALANCE_MAX_UNUSED_CAPACITY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tapiPathComputationRouteObjectiveFunctionEnum = append(tapiPathComputationRouteObjectiveFunctionEnum, v)
	}
}

func (m TapiPathComputationRouteObjectiveFunction) validateTapiPathComputationRouteObjectiveFunctionEnum(path, location string, value TapiPathComputationRouteObjectiveFunction) error {
	if err := validate.Enum(path, location, value, tapiPathComputationRouteObjectiveFunctionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tapi path computation route objective function
func (m TapiPathComputationRouteObjectiveFunction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTapiPathComputationRouteObjectiveFunctionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
