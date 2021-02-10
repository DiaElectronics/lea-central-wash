// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

// StatusOKCode is the HTTP code returned for type StatusOK
const StatusOKCode int = 200

/*StatusOK OK

swagger:response statusOK
*/
type StatusOK struct {

	/*
	  In: Body
	*/
	Payload *model.StatusReport `json:"body,omitempty"`
}

// NewStatusOK creates StatusOK with default headers values
func NewStatusOK() *StatusOK {

	return &StatusOK{}
}

// WithPayload adds the payload to the status o k response
func (o *StatusOK) WithPayload(payload *model.StatusReport) *StatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the status o k response
func (o *StatusOK) SetPayload(payload *model.StatusReport) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *StatusOK) StatusResponder() {}

// StatusInternalServerErrorCode is the HTTP code returned for type StatusInternalServerError
const StatusInternalServerErrorCode int = 500

/*StatusInternalServerError internal error

swagger:response statusInternalServerError
*/
type StatusInternalServerError struct {
}

// NewStatusInternalServerError creates StatusInternalServerError with default headers values
func NewStatusInternalServerError() *StatusInternalServerError {

	return &StatusInternalServerError{}
}

// WriteResponse to the client
func (o *StatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *StatusInternalServerError) StatusResponder() {}

type StatusNotImplementedResponder struct {
	middleware.Responder
}

func (*StatusNotImplementedResponder) StatusResponder() {}

func StatusNotImplemented() StatusResponder {
	return &StatusNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Status has not yet been implemented",
		),
	}
}

type StatusResponder interface {
	middleware.Responder
	StatusResponder()
}
