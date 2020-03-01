// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	model "github.com/DiaElectronics/lea-central-wash/storageapi/model"
)

// LoadRelayHandlerFunc turns a function with the right signature into a load relay handler
type LoadRelayHandlerFunc func(LoadRelayParams) LoadRelayResponder

// Handle executing the request and returning a response
func (fn LoadRelayHandlerFunc) Handle(params LoadRelayParams) LoadRelayResponder {
	return fn(params)
}

// LoadRelayHandler interface for that can handle valid load relay params
type LoadRelayHandler interface {
	Handle(LoadRelayParams) LoadRelayResponder
}

// NewLoadRelay creates a new http.Handler for the load relay operation
func NewLoadRelay(ctx *middleware.Context, handler LoadRelayHandler) *LoadRelay {
	return &LoadRelay{Context: ctx, Handler: handler}
}

/*LoadRelay swagger:route POST /load-relay loadRelay

LoadRelay load relay API

*/
type LoadRelay struct {
	Context *middleware.Context
	Handler LoadRelayHandler
}

func (o *LoadRelay) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoadRelayParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// LoadRelayBody load relay body
// swagger:model LoadRelayBody
type LoadRelayBody struct {

	// hash
	// Required: true
	Hash model.Hash `json:"hash"`
}

// Validate validates this load relay body
func (o *LoadRelayBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoadRelayBody) validateHash(formats strfmt.Registry) error {

	if err := o.Hash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "hash")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoadRelayBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoadRelayBody) UnmarshalBinary(b []byte) error {
	var res LoadRelayBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}