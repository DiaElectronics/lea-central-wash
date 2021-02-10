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

// NewCardReaderConfigParams creates a new CardReaderConfigParams object
// with the default values initialized.
func NewCardReaderConfigParams() *CardReaderConfigParams {
	var ()
	return &CardReaderConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCardReaderConfigParamsWithTimeout creates a new CardReaderConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCardReaderConfigParamsWithTimeout(timeout time.Duration) *CardReaderConfigParams {
	var ()
	return &CardReaderConfigParams{

		timeout: timeout,
	}
}

// NewCardReaderConfigParamsWithContext creates a new CardReaderConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewCardReaderConfigParamsWithContext(ctx context.Context) *CardReaderConfigParams {
	var ()
	return &CardReaderConfigParams{

		Context: ctx,
	}
}

// NewCardReaderConfigParamsWithHTTPClient creates a new CardReaderConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCardReaderConfigParamsWithHTTPClient(client *http.Client) *CardReaderConfigParams {
	var ()
	return &CardReaderConfigParams{
		HTTPClient: client,
	}
}

/*CardReaderConfigParams contains all the parameters to send to the API endpoint
for the card reader config operation typically these are written to a http.Request
*/
type CardReaderConfigParams struct {

	/*Args*/
	Args CardReaderConfigBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the card reader config params
func (o *CardReaderConfigParams) WithTimeout(timeout time.Duration) *CardReaderConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the card reader config params
func (o *CardReaderConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the card reader config params
func (o *CardReaderConfigParams) WithContext(ctx context.Context) *CardReaderConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the card reader config params
func (o *CardReaderConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the card reader config params
func (o *CardReaderConfigParams) WithHTTPClient(client *http.Client) *CardReaderConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the card reader config params
func (o *CardReaderConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the card reader config params
func (o *CardReaderConfigParams) WithArgs(args CardReaderConfigBody) *CardReaderConfigParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the card reader config params
func (o *CardReaderConfigParams) SetArgs(args CardReaderConfigBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *CardReaderConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
