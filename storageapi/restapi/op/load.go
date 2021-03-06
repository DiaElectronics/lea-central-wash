// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/DiaElectronics/lea-central-wash/storageapi/model"
	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// LoadHandlerFunc turns a function with the right signature into a load handler
type LoadHandlerFunc func(LoadParams) LoadResponder

// Handle executing the request and returning a response
func (fn LoadHandlerFunc) Handle(params LoadParams) LoadResponder {
	return fn(params)
}

// LoadHandler interface for that can handle valid load params
type LoadHandler interface {
	Handle(LoadParams) LoadResponder
}

// NewLoad creates a new http.Handler for the load operation
func NewLoad(ctx *middleware.Context, handler LoadHandler) *Load {
	return &Load{Context: ctx, Handler: handler}
}

/*Load swagger:route POST /load load

Load load API

*/
type Load struct {
	Context *middleware.Context
	Handler LoadHandler
}

func (o *Load) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// LoadBody load body
// swagger:model LoadBody
type LoadBody struct {

	// hash
	// Required: true
	Hash model.Hash `json:"hash"`

	// key
	// Required: true
	// Min Length: 1
	Key *string `json:"key"`
}

// Validate validates this load body
func (o *LoadBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoadBody) validateHash(formats strfmt.Registry) error {

	if err := o.Hash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "hash")
		}
		return err
	}

	return nil
}

func (o *LoadBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	if err := validate.MinLength("args"+"."+"key", "body", string(*o.Key), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoadBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoadBody) UnmarshalBinary(b []byte) error {
	var res LoadBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
