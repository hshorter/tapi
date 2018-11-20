// Code generated by go-swagger; DO NOT EDIT.

package tapi_notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsGetNotificationSubscriptionServiceListHandlerFunc turns a function with the right signature into a post operations get notification subscription service list handler
type PostOperationsGetNotificationSubscriptionServiceListHandlerFunc func(PostOperationsGetNotificationSubscriptionServiceListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsGetNotificationSubscriptionServiceListHandlerFunc) Handle(params PostOperationsGetNotificationSubscriptionServiceListParams) middleware.Responder {
	return fn(params)
}

// PostOperationsGetNotificationSubscriptionServiceListHandler interface for that can handle valid post operations get notification subscription service list params
type PostOperationsGetNotificationSubscriptionServiceListHandler interface {
	Handle(PostOperationsGetNotificationSubscriptionServiceListParams) middleware.Responder
}

// NewPostOperationsGetNotificationSubscriptionServiceList creates a new http.Handler for the post operations get notification subscription service list operation
func NewPostOperationsGetNotificationSubscriptionServiceList(ctx *middleware.Context, handler PostOperationsGetNotificationSubscriptionServiceListHandler) *PostOperationsGetNotificationSubscriptionServiceList {
	return &PostOperationsGetNotificationSubscriptionServiceList{Context: ctx, Handler: handler}
}

/*PostOperationsGetNotificationSubscriptionServiceList swagger:route POST /operations/get-notification-subscription-service-list/ tapi-notification postOperationsGetNotificationSubscriptionServiceList

PostOperationsGetNotificationSubscriptionServiceList post operations get notification subscription service list API

*/
type PostOperationsGetNotificationSubscriptionServiceList struct {
	Context *middleware.Context
	Handler PostOperationsGetNotificationSubscriptionServiceListHandler
}

func (o *PostOperationsGetNotificationSubscriptionServiceList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsGetNotificationSubscriptionServiceListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
