// Code generated by go-swagger; DO NOT EDIT.

package tapi_virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsGetVirtualNetworkServiceListHandlerFunc turns a function with the right signature into a post operations get virtual network service list handler
type PostOperationsGetVirtualNetworkServiceListHandlerFunc func(PostOperationsGetVirtualNetworkServiceListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsGetVirtualNetworkServiceListHandlerFunc) Handle(params PostOperationsGetVirtualNetworkServiceListParams) middleware.Responder {
	return fn(params)
}

// PostOperationsGetVirtualNetworkServiceListHandler interface for that can handle valid post operations get virtual network service list params
type PostOperationsGetVirtualNetworkServiceListHandler interface {
	Handle(PostOperationsGetVirtualNetworkServiceListParams) middleware.Responder
}

// NewPostOperationsGetVirtualNetworkServiceList creates a new http.Handler for the post operations get virtual network service list operation
func NewPostOperationsGetVirtualNetworkServiceList(ctx *middleware.Context, handler PostOperationsGetVirtualNetworkServiceListHandler) *PostOperationsGetVirtualNetworkServiceList {
	return &PostOperationsGetVirtualNetworkServiceList{Context: ctx, Handler: handler}
}

/*PostOperationsGetVirtualNetworkServiceList swagger:route POST /operations/get-virtual-network-service-list/ tapi-virtual-network postOperationsGetVirtualNetworkServiceList

PostOperationsGetVirtualNetworkServiceList post operations get virtual network service list API

*/
type PostOperationsGetVirtualNetworkServiceList struct {
	Context *middleware.Context
	Handler PostOperationsGetVirtualNetworkServiceListHandler
}

func (o *PostOperationsGetVirtualNetworkServiceList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsGetVirtualNetworkServiceListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
