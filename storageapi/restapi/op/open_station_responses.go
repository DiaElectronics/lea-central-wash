// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

// OpenStationNoContentCode is the HTTP code returned for type OpenStationNoContent
const OpenStationNoContentCode int = 204

/*OpenStationNoContent OK

swagger:response openStationNoContent
*/
type OpenStationNoContent struct {
}

// NewOpenStationNoContent creates OpenStationNoContent with default headers values
func NewOpenStationNoContent() *OpenStationNoContent {

	return &OpenStationNoContent{}
}

// WriteResponse to the client
func (o *OpenStationNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *OpenStationNoContent) OpenStationResponder() {}

// OpenStationNotFoundCode is the HTTP code returned for type OpenStationNotFound
const OpenStationNotFoundCode int = 404

/*OpenStationNotFound not found

swagger:response openStationNotFound
*/
type OpenStationNotFound struct {
}

// NewOpenStationNotFound creates OpenStationNotFound with default headers values
func NewOpenStationNotFound() *OpenStationNotFound {

	return &OpenStationNotFound{}
}

// WriteResponse to the client
func (o *OpenStationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *OpenStationNotFound) OpenStationResponder() {}

// OpenStationInternalServerErrorCode is the HTTP code returned for type OpenStationInternalServerError
const OpenStationInternalServerErrorCode int = 500

/*OpenStationInternalServerError internal error

swagger:response openStationInternalServerError
*/
type OpenStationInternalServerError struct {
}

// NewOpenStationInternalServerError creates OpenStationInternalServerError with default headers values
func NewOpenStationInternalServerError() *OpenStationInternalServerError {

	return &OpenStationInternalServerError{}
}

// WriteResponse to the client
func (o *OpenStationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *OpenStationInternalServerError) OpenStationResponder() {}

type OpenStationNotImplementedResponder struct {
	middleware.Responder
}

func (*OpenStationNotImplementedResponder) OpenStationResponder() {}

func OpenStationNotImplemented() OpenStationResponder {
	return &OpenStationNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.OpenStation has not yet been implemented",
		),
	}
}

type OpenStationResponder interface {
	middleware.Responder
	OpenStationResponder()
}