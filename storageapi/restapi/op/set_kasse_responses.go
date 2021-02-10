// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"srp-be/media/mediaapi/model"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

// SetKasseNoContentCode is the HTTP code returned for type SetKasseNoContent
const SetKasseNoContentCode int = 204

/*SetKasseNoContent OK

swagger:response setKasseNoContent
*/
type SetKasseNoContent struct {
}

// NewSetKasseNoContent creates SetKasseNoContent with default headers values
func NewSetKasseNoContent() *SetKasseNoContent {

	return &SetKasseNoContent{}
}

// WriteResponse to the client
func (o *SetKasseNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

func (o *SetKasseNoContent) SetKasseResponder() {}

/*SetKasseDefault - 409.1600: email is not available
- 409.1601: account is not available
- 422.1602: password is too weak
- 409.1604: code is not available


swagger:response setKasseDefault
*/
type SetKasseDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewSetKasseDefault creates SetKasseDefault with default headers values
func NewSetKasseDefault(code int) *SetKasseDefault {
	if code <= 0 {
		code = 500
	}

	return &SetKasseDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the set kasse default response
func (o *SetKasseDefault) WithStatusCode(code int) *SetKasseDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the set kasse default response
func (o *SetKasseDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the set kasse default response
func (o *SetKasseDefault) WithPayload(payload *model.Error) *SetKasseDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set kasse default response
func (o *SetKasseDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetKasseDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *SetKasseDefault) SetKasseResponder() {}

type SetKasseNotImplementedResponder struct {
	middleware.Responder
}

func (*SetKasseNotImplementedResponder) SetKasseResponder() {}

func SetKasseNotImplemented() SetKasseResponder {
	return &SetKasseNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.SetKasse has not yet been implemented",
		),
	}
}

type SetKasseResponder interface {
	middleware.Responder
	SetKasseResponder()
}
