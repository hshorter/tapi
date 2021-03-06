// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsGetServiceInterfacePointDetailsHandlerFunc turns a function with the right signature into a post operations get service interface point details handler
type PostOperationsGetServiceInterfacePointDetailsHandlerFunc func(PostOperationsGetServiceInterfacePointDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsGetServiceInterfacePointDetailsHandlerFunc) Handle(params PostOperationsGetServiceInterfacePointDetailsParams) middleware.Responder {
	return fn(params)
}

// PostOperationsGetServiceInterfacePointDetailsHandler interface for that can handle valid post operations get service interface point details params
type PostOperationsGetServiceInterfacePointDetailsHandler interface {
	Handle(PostOperationsGetServiceInterfacePointDetailsParams) middleware.Responder
}

// NewPostOperationsGetServiceInterfacePointDetails creates a new http.Handler for the post operations get service interface point details operation
func NewPostOperationsGetServiceInterfacePointDetails(ctx *middleware.Context, handler PostOperationsGetServiceInterfacePointDetailsHandler) *PostOperationsGetServiceInterfacePointDetails {
	return &PostOperationsGetServiceInterfacePointDetails{Context: ctx, Handler: handler}
}

/*PostOperationsGetServiceInterfacePointDetails swagger:route POST /operations/get-service-interface-point-details/ tapi-common postOperationsGetServiceInterfacePointDetails

PostOperationsGetServiceInterfacePointDetails post operations get service interface point details API

*/
type PostOperationsGetServiceInterfacePointDetails struct {
	Context *middleware.Context
	Handler PostOperationsGetServiceInterfacePointDetailsHandler
}

func (o *PostOperationsGetServiceInterfacePointDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsGetServiceInterfacePointDetailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
