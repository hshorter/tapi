// Code generated by go-swagger; DO NOT EDIT.

package tapi_connectivity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsGetConnectionDetailsOKCode is the HTTP code returned for type PostOperationsGetConnectionDetailsOK
const PostOperationsGetConnectionDetailsOKCode int = 200

/*PostOperationsGetConnectionDetailsOK Correct response

swagger:response postOperationsGetConnectionDetailsOK
*/
type PostOperationsGetConnectionDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiConnectivityGetConnectionDetailsOutput `json:"body,omitempty"`
}

// NewPostOperationsGetConnectionDetailsOK creates PostOperationsGetConnectionDetailsOK with default headers values
func NewPostOperationsGetConnectionDetailsOK() *PostOperationsGetConnectionDetailsOK {

	return &PostOperationsGetConnectionDetailsOK{}
}

// WithPayload adds the payload to the post operations get connection details o k response
func (o *PostOperationsGetConnectionDetailsOK) WithPayload(payload *models.TapiConnectivityGetConnectionDetailsOutput) *PostOperationsGetConnectionDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get connection details o k response
func (o *PostOperationsGetConnectionDetailsOK) SetPayload(payload *models.TapiConnectivityGetConnectionDetailsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetConnectionDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetConnectionDetailsCreatedCode is the HTTP code returned for type PostOperationsGetConnectionDetailsCreated
const PostOperationsGetConnectionDetailsCreatedCode int = 201

/*PostOperationsGetConnectionDetailsCreated No response

swagger:response postOperationsGetConnectionDetailsCreated
*/
type PostOperationsGetConnectionDetailsCreated struct {
}

// NewPostOperationsGetConnectionDetailsCreated creates PostOperationsGetConnectionDetailsCreated with default headers values
func NewPostOperationsGetConnectionDetailsCreated() *PostOperationsGetConnectionDetailsCreated {

	return &PostOperationsGetConnectionDetailsCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetConnectionDetailsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetConnectionDetailsBadRequestCode is the HTTP code returned for type PostOperationsGetConnectionDetailsBadRequest
const PostOperationsGetConnectionDetailsBadRequestCode int = 400

/*PostOperationsGetConnectionDetailsBadRequest Internal error

swagger:response postOperationsGetConnectionDetailsBadRequest
*/
type PostOperationsGetConnectionDetailsBadRequest struct {
}

// NewPostOperationsGetConnectionDetailsBadRequest creates PostOperationsGetConnectionDetailsBadRequest with default headers values
func NewPostOperationsGetConnectionDetailsBadRequest() *PostOperationsGetConnectionDetailsBadRequest {

	return &PostOperationsGetConnectionDetailsBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetConnectionDetailsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
