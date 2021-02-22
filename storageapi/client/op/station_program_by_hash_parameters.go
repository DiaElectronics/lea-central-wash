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

// NewStationProgramByHashParams creates a new StationProgramByHashParams object
// with the default values initialized.
func NewStationProgramByHashParams() *StationProgramByHashParams {
	var ()
	return &StationProgramByHashParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStationProgramByHashParamsWithTimeout creates a new StationProgramByHashParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStationProgramByHashParamsWithTimeout(timeout time.Duration) *StationProgramByHashParams {
	var ()
	return &StationProgramByHashParams{

		timeout: timeout,
	}
}

// NewStationProgramByHashParamsWithContext creates a new StationProgramByHashParams object
// with the default values initialized, and the ability to set a context for a request
func NewStationProgramByHashParamsWithContext(ctx context.Context) *StationProgramByHashParams {
	var ()
	return &StationProgramByHashParams{

		Context: ctx,
	}
}

// NewStationProgramByHashParamsWithHTTPClient creates a new StationProgramByHashParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStationProgramByHashParamsWithHTTPClient(client *http.Client) *StationProgramByHashParams {
	var ()
	return &StationProgramByHashParams{
		HTTPClient: client,
	}
}

/*StationProgramByHashParams contains all the parameters to send to the API endpoint
for the station program by hash operation typically these are written to a http.Request
*/
type StationProgramByHashParams struct {

	/*Args*/
	Args StationProgramByHashBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the station program by hash params
func (o *StationProgramByHashParams) WithTimeout(timeout time.Duration) *StationProgramByHashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the station program by hash params
func (o *StationProgramByHashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the station program by hash params
func (o *StationProgramByHashParams) WithContext(ctx context.Context) *StationProgramByHashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the station program by hash params
func (o *StationProgramByHashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the station program by hash params
func (o *StationProgramByHashParams) WithHTTPClient(client *http.Client) *StationProgramByHashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the station program by hash params
func (o *StationProgramByHashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the station program by hash params
func (o *StationProgramByHashParams) WithArgs(args StationProgramByHashBody) *StationProgramByHashParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the station program by hash params
func (o *StationProgramByHashParams) SetArgs(args StationProgramByHashBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *StationProgramByHashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
