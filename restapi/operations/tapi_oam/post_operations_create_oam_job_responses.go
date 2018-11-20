// Code generated by go-swagger; DO NOT EDIT.

package tapi_oam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsCreateOamJobOKCode is the HTTP code returned for type PostOperationsCreateOamJobOK
const PostOperationsCreateOamJobOKCode int = 200

/*PostOperationsCreateOamJobOK Correct response

swagger:response postOperationsCreateOamJobOK
*/
type PostOperationsCreateOamJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiOamCreateOamJobOutput `json:"body,omitempty"`
}

// NewPostOperationsCreateOamJobOK creates PostOperationsCreateOamJobOK with default headers values
func NewPostOperationsCreateOamJobOK() *PostOperationsCreateOamJobOK {

	return &PostOperationsCreateOamJobOK{}
}

// WithPayload adds the payload to the post operations create oam job o k response
func (o *PostOperationsCreateOamJobOK) WithPayload(payload *models.TapiOamCreateOamJobOutput) *PostOperationsCreateOamJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations create oam job o k response
func (o *PostOperationsCreateOamJobOK) SetPayload(payload *models.TapiOamCreateOamJobOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsCreateOamJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsCreateOamJobCreatedCode is the HTTP code returned for type PostOperationsCreateOamJobCreated
const PostOperationsCreateOamJobCreatedCode int = 201

/*PostOperationsCreateOamJobCreated No response

swagger:response postOperationsCreateOamJobCreated
*/
type PostOperationsCreateOamJobCreated struct {
}

// NewPostOperationsCreateOamJobCreated creates PostOperationsCreateOamJobCreated with default headers values
func NewPostOperationsCreateOamJobCreated() *PostOperationsCreateOamJobCreated {

	return &PostOperationsCreateOamJobCreated{}
}

// WriteResponse to the client
func (o *PostOperationsCreateOamJobCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsCreateOamJobBadRequestCode is the HTTP code returned for type PostOperationsCreateOamJobBadRequest
const PostOperationsCreateOamJobBadRequestCode int = 400

/*PostOperationsCreateOamJobBadRequest Internal error

swagger:response postOperationsCreateOamJobBadRequest
*/
type PostOperationsCreateOamJobBadRequest struct {
}

// NewPostOperationsCreateOamJobBadRequest creates PostOperationsCreateOamJobBadRequest with default headers values
func NewPostOperationsCreateOamJobBadRequest() *PostOperationsCreateOamJobBadRequest {

	return &PostOperationsCreateOamJobBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsCreateOamJobBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}