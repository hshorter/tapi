// Code generated by go-swagger; DO NOT EDIT.

package tapi_notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostOperationsGetNotificationSubscriptionServiceDetailsHandlerFunc turns a function with the right signature into a post operations get notification subscription service details handler
type PostOperationsGetNotificationSubscriptionServiceDetailsHandlerFunc func(PostOperationsGetNotificationSubscriptionServiceDetailsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOperationsGetNotificationSubscriptionServiceDetailsHandlerFunc) Handle(params PostOperationsGetNotificationSubscriptionServiceDetailsParams) middleware.Responder {
	return fn(params)
}

// PostOperationsGetNotificationSubscriptionServiceDetailsHandler interface for that can handle valid post operations get notification subscription service details params
type PostOperationsGetNotificationSubscriptionServiceDetailsHandler interface {
	Handle(PostOperationsGetNotificationSubscriptionServiceDetailsParams) middleware.Responder
}

// NewPostOperationsGetNotificationSubscriptionServiceDetails creates a new http.Handler for the post operations get notification subscription service details operation
func NewPostOperationsGetNotificationSubscriptionServiceDetails(ctx *middleware.Context, handler PostOperationsGetNotificationSubscriptionServiceDetailsHandler) *PostOperationsGetNotificationSubscriptionServiceDetails {
	return &PostOperationsGetNotificationSubscriptionServiceDetails{Context: ctx, Handler: handler}
}

/*PostOperationsGetNotificationSubscriptionServiceDetails swagger:route POST /operations/get-notification-subscription-service-details/ tapi-notification postOperationsGetNotificationSubscriptionServiceDetails

PostOperationsGetNotificationSubscriptionServiceDetails post operations get notification subscription service details API

*/
type PostOperationsGetNotificationSubscriptionServiceDetails struct {
	Context *middleware.Context
	Handler PostOperationsGetNotificationSubscriptionServiceDetailsHandler
}

func (o *PostOperationsGetNotificationSubscriptionServiceDetails) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOperationsGetNotificationSubscriptionServiceDetailsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
