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

	"github.com/DiaElectronics/lea-central-wash/storageapi"
	"github.com/DiaElectronics/lea-central-wash/storageapi/model"
)

// UpdateUserPasswordHandlerFunc turns a function with the right signature into a update user password handler
type UpdateUserPasswordHandlerFunc func(UpdateUserPasswordParams, *storageapi.Profile) UpdateUserPasswordResponder

// Handle executing the request and returning a response
func (fn UpdateUserPasswordHandlerFunc) Handle(params UpdateUserPasswordParams, principal *storageapi.Profile) UpdateUserPasswordResponder {
	return fn(params, principal)
}

// UpdateUserPasswordHandler interface for that can handle valid update user password params
type UpdateUserPasswordHandler interface {
	Handle(UpdateUserPasswordParams, *storageapi.Profile) UpdateUserPasswordResponder
}

// NewUpdateUserPassword creates a new http.Handler for the update user password operation
func NewUpdateUserPassword(ctx *middleware.Context, handler UpdateUserPasswordHandler) *UpdateUserPassword {
	return &UpdateUserPassword{Context: ctx, Handler: handler}
}

/*UpdateUserPassword swagger:route POST /user-password updateUserPassword

UpdateUserPassword update user password API

*/
type UpdateUserPassword struct {
	Context *middleware.Context
	Handler UpdateUserPasswordHandler
}

func (o *UpdateUserPassword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateUserPasswordParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *storageapi.Profile
	if uprinc != nil {
		principal = uprinc.(*storageapi.Profile) // this is really a storageapi.Profile, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateUserPasswordBody update user password body
// swagger:model UpdateUserPasswordBody
type UpdateUserPasswordBody struct {

	// login
	// Required: true
	Login model.Login `json:"login"`

	// new password
	// Required: true
	NewPassword model.Password `json:"newPassword"`

	// old password
	// Required: true
	OldPassword model.Password `json:"oldPassword"`
}

// Validate validates this update user password body
func (o *UpdateUserPasswordBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNewPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOldPassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUserPasswordBody) validateLogin(formats strfmt.Registry) error {

	if err := o.Login.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "login")
		}
		return err
	}

	return nil
}

func (o *UpdateUserPasswordBody) validateNewPassword(formats strfmt.Registry) error {

	if err := o.NewPassword.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "newPassword")
		}
		return err
	}

	return nil
}

func (o *UpdateUserPasswordBody) validateOldPassword(formats strfmt.Registry) error {

	if err := o.OldPassword.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "oldPassword")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserPasswordBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserPasswordBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserPasswordBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateUserPasswordOKBody update user password o k body
// swagger:model UpdateUserPasswordOKBody
type UpdateUserPasswordOKBody struct {

	// id
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this update user password o k body
func (o *UpdateUserPasswordOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUserPasswordOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("updateUserPasswordOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserPasswordOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserPasswordOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserPasswordOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
