// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

// ProgramsOKCode is the HTTP code returned for type ProgramsOK
const ProgramsOKCode int = 200

/*ProgramsOK OK

swagger:response programsOK
*/
type ProgramsOK struct {

	/*
	  In: Body
	*/
	Payload []*model.ProgramInfo `json:"body,omitempty"`
}

// NewProgramsOK creates ProgramsOK with default headers values
func NewProgramsOK() *ProgramsOK {

	return &ProgramsOK{}
}

// WithPayload adds the payload to the programs o k response
func (o *ProgramsOK) WithPayload(payload []*model.ProgramInfo) *ProgramsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the programs o k response
func (o *ProgramsOK) SetPayload(payload []*model.ProgramInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProgramsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*model.ProgramInfo, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *ProgramsOK) ProgramsResponder() {}

// ProgramsInternalServerErrorCode is the HTTP code returned for type ProgramsInternalServerError
const ProgramsInternalServerErrorCode int = 500

/*ProgramsInternalServerError internal error

swagger:response programsInternalServerError
*/
type ProgramsInternalServerError struct {
}

// NewProgramsInternalServerError creates ProgramsInternalServerError with default headers values
func NewProgramsInternalServerError() *ProgramsInternalServerError {

	return &ProgramsInternalServerError{}
}

// WriteResponse to the client
func (o *ProgramsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *ProgramsInternalServerError) ProgramsResponder() {}

type ProgramsNotImplementedResponder struct {
	middleware.Responder
}

func (*ProgramsNotImplementedResponder) ProgramsResponder() {}

func ProgramsNotImplemented() ProgramsResponder {
	return &ProgramsNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Programs has not yet been implemented",
		),
	}
}

type ProgramsResponder interface {
	middleware.Responder
	ProgramsResponder()
}
