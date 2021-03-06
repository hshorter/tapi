// Code generated by go-swagger; DO NOT EDIT.

package tapi_topology

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsGetLinkDetailsHandlerFunc turns a function with the right signature into a post operations get link details handler
type PostOperationsGetLinkDetailsHandlerFunc func(PostOperationsGetLinkDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsGetLinkDetailsHandlerFunc) Handle(params PostOperationsGetLinkDetailsParams) middleware.Responder {
	return fn(params)
}

// PostOperationsGetLinkDetailsHandler interface for that can handle valid post operations get link details params
type PostOperationsGetLinkDetailsHandler interface {
	Handle(PostOperationsGetLinkDetailsParams) middleware.Responder
}

// NewPostOperationsGetLinkDetails creates a new http.Handler for the post operations get link details operation
func NewPostOperationsGetLinkDetails(ctx *middleware.Context, handler PostOperationsGetLinkDetailsHandler) *PostOperationsGetLinkDetails {
	return &PostOperationsGetLinkDetails{Context: ctx, Handler: handler}
}

/*PostOperationsGetLinkDetails swagger:route POST /operations/get-link-details/ tapi-topology postOperationsGetLinkDetails

PostOperationsGetLinkDetails post operations get link details API

*/
type PostOperationsGetLinkDetails struct {
	Context *middleware.Context
	Handler PostOperationsGetLinkDetailsHandler
}

func (o *PostOperationsGetLinkDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsGetLinkDetailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
