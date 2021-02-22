// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

// SetProgramNoContentCode is the HTTP code returned for type SetProgramNoContent
const SetProgramNoContentCode int = 204

/*SetProgramNoContent OK

swagger:response setProgramNoContent
*/
type SetProgramNoContent struct {
}

// NewSetProgramNoContent creates SetProgramNoContent with default headers values
func NewSetProgramNoContent() *SetProgramNoContent {

	return &SetProgramNoContent{}
}

// WriteResponse to the client
func (o *SetProgramNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *SetProgramNoContent) SetProgramResponder() {}

// SetProgramInternalServerErrorCode is the HTTP code returned for type SetProgramInternalServerError
const SetProgramInternalServerErrorCode int = 500

/*SetProgramInternalServerError internal error

swagger:response setProgramInternalServerError
*/
type SetProgramInternalServerError struct {
}

// NewSetProgramInternalServerError creates SetProgramInternalServerError with default headers values
func NewSetProgramInternalServerError() *SetProgramInternalServerError {

	return &SetProgramInternalServerError{}
}

// WriteResponse to the client
func (o *SetProgramInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *SetProgramInternalServerError) SetProgramResponder() {}

type SetProgramNotImplementedResponder struct {
	middleware.Responder
}

func (*SetProgramNotImplementedResponder) SetProgramResponder() {}

func SetProgramNotImplemented() SetProgramResponder {
	return &SetProgramNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.SetProgram has not yet been implemented",
		),
	}
}

type SetProgramResponder interface {
	middleware.Responder
	SetProgramResponder()
}
