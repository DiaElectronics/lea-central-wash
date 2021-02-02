// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// StatusCollectionHandlerFunc turns a function with the right signature into a status collection handler
type StatusCollectionHandlerFunc func(StatusCollectionParams, interface{}) StatusCollectionResponder

// Handle executing the request and returning a response
func (fn StatusCollectionHandlerFunc) Handle(params StatusCollectionParams, principal interface{}) StatusCollectionResponder {
	return fn(params, principal)
}

// StatusCollectionHandler interface for that can handle valid status collection params
type StatusCollectionHandler interface {
	Handle(StatusCollectionParams, interface{}) StatusCollectionResponder
}

// NewStatusCollection creates a new http.Handler for the status collection operation
func NewStatusCollection(ctx *middleware.Context, handler StatusCollectionHandler) *StatusCollection {
	return &StatusCollection{Context: ctx, Handler: handler}
}

/*StatusCollection swagger:route GET /status-collection statusCollection

StatusCollection status collection API

*/
type StatusCollection struct {
	Context *middleware.Context
	Handler StatusCollectionHandler
}

func (o *StatusCollection) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewStatusCollectionParams()

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
