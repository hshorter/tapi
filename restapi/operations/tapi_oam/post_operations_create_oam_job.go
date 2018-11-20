// Code generated by go-swagger; DO NOT EDIT.

package tapi_oam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsCreateOamJobHandlerFunc turns a function with the right signature into a post operations create oam job handler
type PostOperationsCreateOamJobHandlerFunc func(PostOperationsCreateOamJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsCreateOamJobHandlerFunc) Handle(params PostOperationsCreateOamJobParams) middleware.Responder {
	return fn(params)
}

// PostOperationsCreateOamJobHandler interface for that can handle valid post operations create oam job params
type PostOperationsCreateOamJobHandler interface {
	Handle(PostOperationsCreateOamJobParams) middleware.Responder
}

// NewPostOperationsCreateOamJob creates a new http.Handler for the post operations create oam job operation
func NewPostOperationsCreateOamJob(ctx *middleware.Context, handler PostOperationsCreateOamJobHandler) *PostOperationsCreateOamJob {
	return &PostOperationsCreateOamJob{Context: ctx, Handler: handler}
}

/*PostOperationsCreateOamJob swagger:route POST /operations/create-oam-job/ tapi-oam postOperationsCreateOamJob

PostOperationsCreateOamJob post operations create oam job API

*/
type PostOperationsCreateOamJob struct {
	Context *middleware.Context
	Handler PostOperationsCreateOamJobHandler
}

func (o *PostOperationsCreateOamJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsCreateOamJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}