// Code generated by go-swagger; DO NOT EDIT.

package tapi_virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/damianoneill/tapi/models"
)

// NewPostOperationsCreateVirtualNetworkServiceParams creates a new PostOperationsCreateVirtualNetworkServiceParams object
// no default values defined in spec.
func NewPostOperationsCreateVirtualNetworkServiceParams() PostOperationsCreateVirtualNetworkServiceParams {

	return PostOperationsCreateVirtualNetworkServiceParams{}
}

// PostOperationsCreateVirtualNetworkServiceParams contains all the bound params for the post operations create virtual network service operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostOperationsCreateVirtualNetworkService
type PostOperationsCreateVirtualNetworkServiceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	BodyParam *models.TapiVirtualNetworkCreateVirtualNetworkServiceInput
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostOperationsCreateVirtualNetworkServiceParams() beforehand.
func (o *PostOperationsCreateVirtualNetworkServiceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.TapiVirtualNetworkCreateVirtualNetworkServiceInput
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("bodyParam", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.BodyParam = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
