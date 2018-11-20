// Code generated by go-swagger; DO NOT EDIT.

package tapi_notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsGetNotificationListHandlerFunc turns a function with the right signature into a post operations get notification list handler
type PostOperationsGetNotificationListHandlerFunc func(PostOperationsGetNotificationListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsGetNotificationListHandlerFunc) Handle(params PostOperationsGetNotificationListParams) middleware.Responder {
	return fn(params)
}

// PostOperationsGetNotificationListHandler interface for that can handle valid post operations get notification list params
type PostOperationsGetNotificationListHandler interface {
	Handle(PostOperationsGetNotificationListParams) middleware.Responder
}

// NewPostOperationsGetNotificationList creates a new http.Handler for the post operations get notification list operation
func NewPostOperationsGetNotificationList(ctx *middleware.Context, handler PostOperationsGetNotificationListHandler) *PostOperationsGetNotificationList {
	return &PostOperationsGetNotificationList{Context: ctx, Handler: handler}
}

/*PostOperationsGetNotificationList swagger:route POST /operations/get-notification-list/ tapi-notification postOperationsGetNotificationList

PostOperationsGetNotificationList post operations get notification list API

*/
type PostOperationsGetNotificationList struct {
	Context *middleware.Context
	Handler PostOperationsGetNotificationListHandler
}

func (o *PostOperationsGetNotificationList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsGetNotificationListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
