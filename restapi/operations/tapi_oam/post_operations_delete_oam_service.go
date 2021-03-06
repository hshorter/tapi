// Code generated by go-swagger; DO NOT EDIT.

package tapi_oam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsDeleteOamServiceHandlerFunc turns a function with the right signature into a post operations delete oam service handler
type PostOperationsDeleteOamServiceHandlerFunc func(PostOperationsDeleteOamServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsDeleteOamServiceHandlerFunc) Handle(params PostOperationsDeleteOamServiceParams) middleware.Responder {
	return fn(params)
}

// PostOperationsDeleteOamServiceHandler interface for that can handle valid post operations delete oam service params
type PostOperationsDeleteOamServiceHandler interface {
	Handle(PostOperationsDeleteOamServiceParams) middleware.Responder
}

// NewPostOperationsDeleteOamService creates a new http.Handler for the post operations delete oam service operation
func NewPostOperationsDeleteOamService(ctx *middleware.Context, handler PostOperationsDeleteOamServiceHandler) *PostOperationsDeleteOamService {
	return &PostOperationsDeleteOamService{Context: ctx, Handler: handler}
}

/*PostOperationsDeleteOamService swagger:route POST /operations/delete-oam-service/ tapi-oam postOperationsDeleteOamService

PostOperationsDeleteOamService post operations delete oam service API

*/
type PostOperationsDeleteOamService struct {
	Context *middleware.Context
	Handler PostOperationsDeleteOamServiceHandler
}

func (o *PostOperationsDeleteOamService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsDeleteOamServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
