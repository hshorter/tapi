// Code generated by go-swagger; DO NOT EDIT.

package tapi_notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsDeleteNotificationSubscriptionServiceOKCode is the HTTP code returned for type PostOperationsDeleteNotificationSubscriptionServiceOK
const PostOperationsDeleteNotificationSubscriptionServiceOKCode int = 200

/*PostOperationsDeleteNotificationSubscriptionServiceOK Correct response

swagger:response postOperationsDeleteNotificationSubscriptionServiceOK
*/
type PostOperationsDeleteNotificationSubscriptionServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiNotificationDeleteNotificationSubscriptionServiceOutput `json:"body,omitempty"`
}

// NewPostOperationsDeleteNotificationSubscriptionServiceOK creates PostOperationsDeleteNotificationSubscriptionServiceOK with default headers values
func NewPostOperationsDeleteNotificationSubscriptionServiceOK() *PostOperationsDeleteNotificationSubscriptionServiceOK {

	return &PostOperationsDeleteNotificationSubscriptionServiceOK{}
}

// WithPayload adds the payload to the post operations delete notification subscription service o k response
func (o *PostOperationsDeleteNotificationSubscriptionServiceOK) WithPayload(payload *models.TapiNotificationDeleteNotificationSubscriptionServiceOutput) *PostOperationsDeleteNotificationSubscriptionServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations delete notification subscription service o k response
func (o *PostOperationsDeleteNotificationSubscriptionServiceOK) SetPayload(payload *models.TapiNotificationDeleteNotificationSubscriptionServiceOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsDeleteNotificationSubscriptionServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsDeleteNotificationSubscriptionServiceCreatedCode is the HTTP code returned for type PostOperationsDeleteNotificationSubscriptionServiceCreated
const PostOperationsDeleteNotificationSubscriptionServiceCreatedCode int = 201

/*PostOperationsDeleteNotificationSubscriptionServiceCreated No response

swagger:response postOperationsDeleteNotificationSubscriptionServiceCreated
*/
type PostOperationsDeleteNotificationSubscriptionServiceCreated struct {
}

// NewPostOperationsDeleteNotificationSubscriptionServiceCreated creates PostOperationsDeleteNotificationSubscriptionServiceCreated with default headers values
func NewPostOperationsDeleteNotificationSubscriptionServiceCreated() *PostOperationsDeleteNotificationSubscriptionServiceCreated {

	return &PostOperationsDeleteNotificationSubscriptionServiceCreated{}
}

// WriteResponse to the client
func (o *PostOperationsDeleteNotificationSubscriptionServiceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsDeleteNotificationSubscriptionServiceBadRequestCode is the HTTP code returned for type PostOperationsDeleteNotificationSubscriptionServiceBadRequest
const PostOperationsDeleteNotificationSubscriptionServiceBadRequestCode int = 400

/*PostOperationsDeleteNotificationSubscriptionServiceBadRequest Internal error

swagger:response postOperationsDeleteNotificationSubscriptionServiceBadRequest
*/
type PostOperationsDeleteNotificationSubscriptionServiceBadRequest struct {
}

// NewPostOperationsDeleteNotificationSubscriptionServiceBadRequest creates PostOperationsDeleteNotificationSubscriptionServiceBadRequest with default headers values
func NewPostOperationsDeleteNotificationSubscriptionServiceBadRequest() *PostOperationsDeleteNotificationSubscriptionServiceBadRequest {

	return &PostOperationsDeleteNotificationSubscriptionServiceBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsDeleteNotificationSubscriptionServiceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
