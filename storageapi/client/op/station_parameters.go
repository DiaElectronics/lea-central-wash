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

// NewStationParams creates a new StationParams object
// with the default values initialized.
func NewStationParams() *StationParams {
	var ()
	return &StationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStationParamsWithTimeout creates a new StationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStationParamsWithTimeout(timeout time.Duration) *StationParams {
	var ()
	return &StationParams{

		timeout: timeout,
	}
}

// NewStationParamsWithContext creates a new StationParams object
// with the default values initialized, and the ability to set a context for a request
func NewStationParamsWithContext(ctx context.Context) *StationParams {
	var ()
	return &StationParams{

		Context: ctx,
	}
}

// NewStationParamsWithHTTPClient creates a new StationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStationParamsWithHTTPClient(client *http.Client) *StationParams {
	var ()
	return &StationParams{
		HTTPClient: client,
	}
}

/*StationParams contains all the parameters to send to the API endpoint
for the station operation typically these are written to a http.Request
*/
type StationParams struct {

	/*Args*/
	Args StationBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the station params
func (o *StationParams) WithTimeout(timeout time.Duration) *StationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the station params
func (o *StationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the station params
func (o *StationParams) WithContext(ctx context.Context) *StationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the station params
func (o *StationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the station params
func (o *StationParams) WithHTTPClient(client *http.Client) *StationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the station params
func (o *StationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the station params
func (o *StationParams) WithArgs(args StationBody) *StationParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the station params
func (o *StationParams) SetArgs(args StationBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *StationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
