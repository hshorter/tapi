// Code generated by go-swagger; DO NOT EDIT.

package tapi_oam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsUpdateOamServiceEndPointHandlerFunc turns a function with the right signature into a post operations update oam service end point handler
type PostOperationsUpdateOamServiceEndPointHandlerFunc func(PostOperationsUpdateOamServiceEndPointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsUpdateOamServiceEndPointHandlerFunc) Handle(params PostOperationsUpdateOamServiceEndPointParams) middleware.Responder {
	return fn(params)
}

// PostOperationsUpdateOamServiceEndPointHandler interface for that can handle valid post operations update oam service end point params
type PostOperationsUpdateOamServiceEndPointHandler interface {
	Handle(PostOperationsUpdateOamServiceEndPointParams) middleware.Responder
}

// NewPostOperationsUpdateOamServiceEndPoint creates a new http.Handler for the post operations update oam service end point operation
func NewPostOperationsUpdateOamServiceEndPoint(ctx *middleware.Context, handler PostOperationsUpdateOamServiceEndPointHandler) *PostOperationsUpdateOamServiceEndPoint {
	return &PostOperationsUpdateOamServiceEndPoint{Context: ctx, Handler: handler}
}

/*PostOperationsUpdateOamServiceEndPoint swagger:route POST /operations/update-oam-service-end-point/ tapi-oam postOperationsUpdateOamServiceEndPoint

PostOperationsUpdateOamServiceEndPoint post operations update oam service end point API

*/
type PostOperationsUpdateOamServiceEndPoint struct {
	Context *middleware.Context
	Handler PostOperationsUpdateOamServiceEndPointHandler
}

func (o *PostOperationsUpdateOamServiceEndPoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsUpdateOamServiceEndPointParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
