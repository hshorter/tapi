// Code generated by go-swagger; DO NOT EDIT.

package tapi_topology

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsGetTopologyDetailsHandlerFunc turns a function with the right signature into a post operations get topology details handler
type PostOperationsGetTopologyDetailsHandlerFunc func(PostOperationsGetTopologyDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsGetTopologyDetailsHandlerFunc) Handle(params PostOperationsGetTopologyDetailsParams) middleware.Responder {
	return fn(params)
}

// PostOperationsGetTopologyDetailsHandler interface for that can handle valid post operations get topology details params
type PostOperationsGetTopologyDetailsHandler interface {
	Handle(PostOperationsGetTopologyDetailsParams) middleware.Responder
}

// NewPostOperationsGetTopologyDetails creates a new http.Handler for the post operations get topology details operation
func NewPostOperationsGetTopologyDetails(ctx *middleware.Context, handler PostOperationsGetTopologyDetailsHandler) *PostOperationsGetTopologyDetails {
	return &PostOperationsGetTopologyDetails{Context: ctx, Handler: handler}
}

/*PostOperationsGetTopologyDetails swagger:route POST /operations/get-topology-details/ tapi-topology postOperationsGetTopologyDetails

PostOperationsGetTopologyDetails post operations get topology details API

*/
type PostOperationsGetTopologyDetails struct {
	Context *middleware.Context
	Handler PostOperationsGetTopologyDetailsHandler
}

func (o *PostOperationsGetTopologyDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsGetTopologyDetailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
