// Code generated by go-swagger; DO NOT EDIT.

package tapi_oam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsGetOamServiceOKCode is the HTTP code returned for type PostOperationsGetOamServiceOK
const PostOperationsGetOamServiceOKCode int = 200

/*PostOperationsGetOamServiceOK Correct response

swagger:response postOperationsGetOamServiceOK
*/
type PostOperationsGetOamServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiOamGetOamServiceOutput `json:"body,omitempty"`
}

// NewPostOperationsGetOamServiceOK creates PostOperationsGetOamServiceOK with default headers values
func NewPostOperationsGetOamServiceOK() *PostOperationsGetOamServiceOK {

	return &PostOperationsGetOamServiceOK{}
}

// WithPayload adds the payload to the post operations get oam service o k response
func (o *PostOperationsGetOamServiceOK) WithPayload(payload *models.TapiOamGetOamServiceOutput) *PostOperationsGetOamServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get oam service o k response
func (o *PostOperationsGetOamServiceOK) SetPayload(payload *models.TapiOamGetOamServiceOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetOamServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetOamServiceCreatedCode is the HTTP code returned for type PostOperationsGetOamServiceCreated
const PostOperationsGetOamServiceCreatedCode int = 201

/*PostOperationsGetOamServiceCreated No response

swagger:response postOperationsGetOamServiceCreated
*/
type PostOperationsGetOamServiceCreated struct {
}

// NewPostOperationsGetOamServiceCreated creates PostOperationsGetOamServiceCreated with default headers values
func NewPostOperationsGetOamServiceCreated() *PostOperationsGetOamServiceCreated {

	return &PostOperationsGetOamServiceCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetOamServiceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetOamServiceBadRequestCode is the HTTP code returned for type PostOperationsGetOamServiceBadRequest
const PostOperationsGetOamServiceBadRequestCode int = 400

/*PostOperationsGetOamServiceBadRequest Internal error

swagger:response postOperationsGetOamServiceBadRequest
*/
type PostOperationsGetOamServiceBadRequest struct {
}

// NewPostOperationsGetOamServiceBadRequest creates PostOperationsGetOamServiceBadRequest with default headers values
func NewPostOperationsGetOamServiceBadRequest() *PostOperationsGetOamServiceBadRequest {

	return &PostOperationsGetOamServiceBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetOamServiceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
