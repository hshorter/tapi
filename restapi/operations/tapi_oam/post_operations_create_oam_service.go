// Code generated by go-swagger; DO NOT EDIT.

package tapi_oam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsCreateOamServiceHandlerFunc turns a function with the right signature into a post operations create oam service handler
type PostOperationsCreateOamServiceHandlerFunc func(PostOperationsCreateOamServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsCreateOamServiceHandlerFunc) Handle(params PostOperationsCreateOamServiceParams) middleware.Responder {
	return fn(params)
}

// PostOperationsCreateOamServiceHandler interface for that can handle valid post operations create oam service params
type PostOperationsCreateOamServiceHandler interface {
	Handle(PostOperationsCreateOamServiceParams) middleware.Responder
}

// NewPostOperationsCreateOamService creates a new http.Handler for the post operations create oam service operation
func NewPostOperationsCreateOamService(ctx *middleware.Context, handler PostOperationsCreateOamServiceHandler) *PostOperationsCreateOamService {
	return &PostOperationsCreateOamService{Context: ctx, Handler: handler}
}

/*PostOperationsCreateOamService swagger:route POST /operations/create-oam-service/ tapi-oam postOperationsCreateOamService

PostOperationsCreateOamService post operations create oam service API

*/
type PostOperationsCreateOamService struct {
	Context *middleware.Context
	Handler PostOperationsCreateOamServiceHandler
}

func (o *PostOperationsCreateOamService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsCreateOamServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
