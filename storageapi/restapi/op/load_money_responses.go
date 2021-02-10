// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

// LoadMoneyOKCode is the HTTP code returned for type LoadMoneyOK
const LoadMoneyOKCode int = 200

/*LoadMoneyOK OK

swagger:response loadMoneyOK
*/
type LoadMoneyOK struct {

	/*
	  In: Body
	*/
	Payload *model.MoneyReport `json:"body,omitempty"`
}

// NewLoadMoneyOK creates LoadMoneyOK with default headers values
func NewLoadMoneyOK() *LoadMoneyOK {

	return &LoadMoneyOK{}
}

// WithPayload adds the payload to the load money o k response
func (o *LoadMoneyOK) WithPayload(payload *model.MoneyReport) *LoadMoneyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the load money o k response
func (o *LoadMoneyOK) SetPayload(payload *model.MoneyReport) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoadMoneyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *LoadMoneyOK) LoadMoneyResponder() {}

// LoadMoneyNotFoundCode is the HTTP code returned for type LoadMoneyNotFound
const LoadMoneyNotFoundCode int = 404

/*LoadMoneyNotFound not found

swagger:response loadMoneyNotFound
*/
type LoadMoneyNotFound struct {
}

// NewLoadMoneyNotFound creates LoadMoneyNotFound with default headers values
func NewLoadMoneyNotFound() *LoadMoneyNotFound {

	return &LoadMoneyNotFound{}
}

// WriteResponse to the client
func (o *LoadMoneyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

func (o *LoadMoneyNotFound) LoadMoneyResponder() {}

// LoadMoneyInternalServerErrorCode is the HTTP code returned for type LoadMoneyInternalServerError
const LoadMoneyInternalServerErrorCode int = 500

/*LoadMoneyInternalServerError internal error

swagger:response loadMoneyInternalServerError
*/
type LoadMoneyInternalServerError struct {
}

// NewLoadMoneyInternalServerError creates LoadMoneyInternalServerError with default headers values
func NewLoadMoneyInternalServerError() *LoadMoneyInternalServerError {

	return &LoadMoneyInternalServerError{}
}

// WriteResponse to the client
func (o *LoadMoneyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func (o *LoadMoneyInternalServerError) LoadMoneyResponder() {}

type LoadMoneyNotImplementedResponder struct {
	middleware.Responder
}

func (*LoadMoneyNotImplementedResponder) LoadMoneyResponder() {}

func LoadMoneyNotImplemented() LoadMoneyResponder {
	return &LoadMoneyNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.LoadMoney has not yet been implemented",
		),
	}
}

type LoadMoneyResponder interface {
	middleware.Responder
	LoadMoneyResponder()
}
