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
	validate "github.com/go-openapi/validate"
)

// SaveCollectionHandlerFunc turns a function with the right signature into a save collection handler
type SaveCollectionHandlerFunc func(SaveCollectionParams, interface{}) SaveCollectionResponder

// Handle executing the request and returning a response
func (fn SaveCollectionHandlerFunc) Handle(params SaveCollectionParams, principal interface{}) SaveCollectionResponder {
	return fn(params, principal)
}

// SaveCollectionHandler interface for that can handle valid save collection params
type SaveCollectionHandler interface {
	Handle(SaveCollectionParams, interface{}) SaveCollectionResponder
}

// NewSaveCollection creates a new http.Handler for the save collection operation
func NewSaveCollection(ctx *middleware.Context, handler SaveCollectionHandler) *SaveCollection {
	return &SaveCollection{Context: ctx, Handler: handler}
}

/*SaveCollection swagger:route POST /save-collection saveCollection

SaveCollection save collection API

*/
type SaveCollection struct {
	Context *middleware.Context
	Handler SaveCollectionHandler
}

func (o *SaveCollection) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSaveCollectionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SaveCollectionBody save collection body
// swagger:model SaveCollectionBody
type SaveCollectionBody struct {

	// id
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this save collection body
func (o *SaveCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SaveCollectionBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("args"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SaveCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SaveCollectionBody) UnmarshalBinary(b []byte) error {
	var res SaveCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
