// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

// LoadRelayOKCode is the HTTP code returned for type LoadRelayOK
const LoadRelayOKCode int = 200

/*LoadRelayOK OK

swagger:response loadRelayOK
*/
type LoadRelayOK struct {

	/*
	  In: Body
	*/
	Payload *model.RelayReport `json:"body,omitempty"`
}

// NewLoadRelayOK creates LoadRelayOK with default headers values
func NewLoadRelayOK() *LoadRelayOK {

	return &LoadRelayOK{}
}

// WithPayload adds the payload to the load relay o k response
func (o *LoadRelayOK) WithPayload(payload *model.RelayReport) *LoadRelayOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the load relay o k response
func (o *LoadRelayOK) SetPayload(payload *model.RelayReport) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoadRelayOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *LoadRelayOK) LoadRelayResponder() {}

// LoadRelayNotFoundCode is the HTTP code returned for type LoadRelayNotFound
const LoadRelayNotFoundCode int = 404

/*LoadRelayNotFound not found

swagger:response loadRelayNotFound
*/
type LoadRelayNotFound struct {
}

// NewLoadRelayNotFound creates LoadRelayNotFound with default headers values
func NewLoadRelayNotFound() *LoadRelayNotFound {

	return &LoadRelayNotFound{}
}

// WriteResponse to the client
func (o *LoadRelayNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *LoadRelayNotFound) LoadRelayResponder() {}

// LoadRelayInternalServerErrorCode is the HTTP code returned for type LoadRelayInternalServerError
const LoadRelayInternalServerErrorCode int = 500

/*LoadRelayInternalServerError internal error

swagger:response loadRelayInternalServerError
*/
type LoadRelayInternalServerError struct {
}

// NewLoadRelayInternalServerError creates LoadRelayInternalServerError with default headers values
func NewLoadRelayInternalServerError() *LoadRelayInternalServerError {

	return &LoadRelayInternalServerError{}
}

// WriteResponse to the client
func (o *LoadRelayInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *LoadRelayInternalServerError) LoadRelayResponder() {}

type LoadRelayNotImplementedResponder struct {
	middleware.Responder
}

func (*LoadRelayNotImplementedResponder) LoadRelayResponder() {}

func LoadRelayNotImplemented() LoadRelayResponder {
	return &LoadRelayNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.LoadRelay has not yet been implemented",
		),
	}
}

type LoadRelayResponder interface {
	middleware.Responder
	LoadRelayResponder()
}
