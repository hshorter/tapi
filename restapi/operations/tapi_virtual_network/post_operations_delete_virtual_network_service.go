// Code generated by go-swagger; DO NOT EDIT.

package tapi_virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsDeleteVirtualNetworkServiceHandlerFunc turns a function with the right signature into a post operations delete virtual network service handler
type PostOperationsDeleteVirtualNetworkServiceHandlerFunc func(PostOperationsDeleteVirtualNetworkServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsDeleteVirtualNetworkServiceHandlerFunc) Handle(params PostOperationsDeleteVirtualNetworkServiceParams) middleware.Responder {
	return fn(params)
}

// PostOperationsDeleteVirtualNetworkServiceHandler interface for that can handle valid post operations delete virtual network service params
type PostOperationsDeleteVirtualNetworkServiceHandler interface {
	Handle(PostOperationsDeleteVirtualNetworkServiceParams) middleware.Responder
}

// NewPostOperationsDeleteVirtualNetworkService creates a new http.Handler for the post operations delete virtual network service operation
func NewPostOperationsDeleteVirtualNetworkService(ctx *middleware.Context, handler PostOperationsDeleteVirtualNetworkServiceHandler) *PostOperationsDeleteVirtualNetworkService {
	return &PostOperationsDeleteVirtualNetworkService{Context: ctx, Handler: handler}
}

/*PostOperationsDeleteVirtualNetworkService swagger:route POST /operations/delete-virtual-network-service/ tapi-virtual-network postOperationsDeleteVirtualNetworkService

PostOperationsDeleteVirtualNetworkService post operations delete virtual network service API

*/
type PostOperationsDeleteVirtualNetworkService struct {
	Context *middleware.Context
	Handler PostOperationsDeleteVirtualNetworkServiceHandler
}

func (o *PostOperationsDeleteVirtualNetworkService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsDeleteVirtualNetworkServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}