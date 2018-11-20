// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsUpdateServiceInterfacePointHandlerFunc turns a function with the right signature into a post operations update service interface point handler
type PostOperationsUpdateServiceInterfacePointHandlerFunc func(PostOperationsUpdateServiceInterfacePointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsUpdateServiceInterfacePointHandlerFunc) Handle(params PostOperationsUpdateServiceInterfacePointParams) middleware.Responder {
	return fn(params)
}

// PostOperationsUpdateServiceInterfacePointHandler interface for that can handle valid post operations update service interface point params
type PostOperationsUpdateServiceInterfacePointHandler interface {
	Handle(PostOperationsUpdateServiceInterfacePointParams) middleware.Responder
}

// NewPostOperationsUpdateServiceInterfacePoint creates a new http.Handler for the post operations update service interface point operation
func NewPostOperationsUpdateServiceInterfacePoint(ctx *middleware.Context, handler PostOperationsUpdateServiceInterfacePointHandler) *PostOperationsUpdateServiceInterfacePoint {
	return &PostOperationsUpdateServiceInterfacePoint{Context: ctx, Handler: handler}
}

/*PostOperationsUpdateServiceInterfacePoint swagger:route POST /operations/update-service-interface-point/ tapi-common postOperationsUpdateServiceInterfacePoint

PostOperationsUpdateServiceInterfacePoint post operations update service interface point API

*/
type PostOperationsUpdateServiceInterfacePoint struct {
	Context *middleware.Context
	Handler PostOperationsUpdateServiceInterfacePointHandler
}

func (o *PostOperationsUpdateServiceInterfacePoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsUpdateServiceInterfacePointParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
