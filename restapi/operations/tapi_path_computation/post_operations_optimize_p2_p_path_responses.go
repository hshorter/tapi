// Code generated by go-swagger; DO NOT EDIT.

package tapi_path_computation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/tapi/models"
)

// PostOperationsOptimizeP2PPathOKCode is the HTTP code returned for type PostOperationsOptimizeP2PPathOK
const PostOperationsOptimizeP2PPathOKCode int = 200

/*PostOperationsOptimizeP2PPathOK Correct response

swagger:response postOperationsOptimizeP2PPathOK
*/
type PostOperationsOptimizeP2PPathOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiPathComputationOptimizeP2PPathOutput `json:"body,omitempty"`
}

// NewPostOperationsOptimizeP2PPathOK creates PostOperationsOptimizeP2PPathOK with default headers values
func NewPostOperationsOptimizeP2PPathOK() *PostOperationsOptimizeP2PPathOK {

	return &PostOperationsOptimizeP2PPathOK{}
}

// WithPayload adds the payload to the post operations optimize p2 p path o k response
func (o *PostOperationsOptimizeP2PPathOK) WithPayload(payload *models.TapiPathComputationOptimizeP2PPathOutput) *PostOperationsOptimizeP2PPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations optimize p2 p path o k response
func (o *PostOperationsOptimizeP2PPathOK) SetPayload(payload *models.TapiPathComputationOptimizeP2PPathOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsOptimizeP2PPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsOptimizeP2PPathCreatedCode is the HTTP code returned for type PostOperationsOptimizeP2PPathCreated
const PostOperationsOptimizeP2PPathCreatedCode int = 201

/*PostOperationsOptimizeP2PPathCreated No response

swagger:response postOperationsOptimizeP2PPathCreated
*/
type PostOperationsOptimizeP2PPathCreated struct {
}

// NewPostOperationsOptimizeP2PPathCreated creates PostOperationsOptimizeP2PPathCreated with default headers values
func NewPostOperationsOptimizeP2PPathCreated() *PostOperationsOptimizeP2PPathCreated {

	return &PostOperationsOptimizeP2PPathCreated{}
}

// WriteResponse to the client
func (o *PostOperationsOptimizeP2PPathCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsOptimizeP2PPathBadRequestCode is the HTTP code returned for type PostOperationsOptimizeP2PPathBadRequest
const PostOperationsOptimizeP2PPathBadRequestCode int = 400

/*PostOperationsOptimizeP2PPathBadRequest Internal error

swagger:response postOperationsOptimizeP2PPathBadRequest
*/
type PostOperationsOptimizeP2PPathBadRequest struct {
}

// NewPostOperationsOptimizeP2PPathBadRequest creates PostOperationsOptimizeP2PPathBadRequest with default headers values
func NewPostOperationsOptimizeP2PPathBadRequest() *PostOperationsOptimizeP2PPathBadRequest {

	return &PostOperationsOptimizeP2PPathBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsOptimizeP2PPathBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
