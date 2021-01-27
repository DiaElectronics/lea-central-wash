// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSetProgramRelaysParams creates a new SetProgramRelaysParams object
// with the default values initialized.
func NewSetProgramRelaysParams() *SetProgramRelaysParams {
	var ()
	return &SetProgramRelaysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetProgramRelaysParamsWithTimeout creates a new SetProgramRelaysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetProgramRelaysParamsWithTimeout(timeout time.Duration) *SetProgramRelaysParams {
	var ()
	return &SetProgramRelaysParams{

		timeout: timeout,
	}
}

// NewSetProgramRelaysParamsWithContext creates a new SetProgramRelaysParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetProgramRelaysParamsWithContext(ctx context.Context) *SetProgramRelaysParams {
	var ()
	return &SetProgramRelaysParams{

		Context: ctx,
	}
}

// NewSetProgramRelaysParamsWithHTTPClient creates a new SetProgramRelaysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetProgramRelaysParamsWithHTTPClient(client *http.Client) *SetProgramRelaysParams {
	var ()
	return &SetProgramRelaysParams{
		HTTPClient: client,
	}
}

/*SetProgramRelaysParams contains all the parameters to send to the API endpoint
for the set program relays operation typically these are written to a http.Request
*/
type SetProgramRelaysParams struct {

	/*Args*/
	Args SetProgramRelaysBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set program relays params
func (o *SetProgramRelaysParams) WithTimeout(timeout time.Duration) *SetProgramRelaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set program relays params
func (o *SetProgramRelaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set program relays params
func (o *SetProgramRelaysParams) WithContext(ctx context.Context) *SetProgramRelaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set program relays params
func (o *SetProgramRelaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set program relays params
func (o *SetProgramRelaysParams) WithHTTPClient(client *http.Client) *SetProgramRelaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set program relays params
func (o *SetProgramRelaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the set program relays params
func (o *SetProgramRelaysParams) WithArgs(args SetProgramRelaysBody) *SetProgramRelaysParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the set program relays params
func (o *SetProgramRelaysParams) SetArgs(args SetProgramRelaysBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *SetProgramRelaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Args); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}