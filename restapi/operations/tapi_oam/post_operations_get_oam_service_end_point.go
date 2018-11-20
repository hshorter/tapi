// Code generated by go-swagger; DO NOT EDIT.

package tapi_oam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsGetOamServiceEndPointHandlerFunc turns a function with the right signature into a post operations get oam service end point handler
type PostOperationsGetOamServiceEndPointHandlerFunc func(PostOperationsGetOamServiceEndPointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsGetOamServiceEndPointHandlerFunc) Handle(params PostOperationsGetOamServiceEndPointParams) middleware.Responder {
	return fn(params)
}

// PostOperationsGetOamServiceEndPointHandler interface for that can handle valid post operations get oam service end point params
type PostOperationsGetOamServiceEndPointHandler interface {
	Handle(PostOperationsGetOamServiceEndPointParams) middleware.Responder
}

// NewPostOperationsGetOamServiceEndPoint creates a new http.Handler for the post operations get oam service end point operation
func NewPostOperationsGetOamServiceEndPoint(ctx *middleware.Context, handler PostOperationsGetOamServiceEndPointHandler) *PostOperationsGetOamServiceEndPoint {
	return &PostOperationsGetOamServiceEndPoint{Context: ctx, Handler: handler}
}

/*PostOperationsGetOamServiceEndPoint swagger:route POST /operations/get-oam-service-end-point/ tapi-oam postOperationsGetOamServiceEndPoint

PostOperationsGetOamServiceEndPoint post operations get oam service end point API

*/
type PostOperationsGetOamServiceEndPoint struct {
	Context *middleware.Context
	Handler PostOperationsGetOamServiceEndPointHandler
}

func (o *PostOperationsGetOamServiceEndPoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsGetOamServiceEndPointParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
