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

// NewCardReaderConfigByHashParams creates a new CardReaderConfigByHashParams object
// with the default values initialized.
func NewCardReaderConfigByHashParams() *CardReaderConfigByHashParams {
	var ()
	return &CardReaderConfigByHashParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCardReaderConfigByHashParamsWithTimeout creates a new CardReaderConfigByHashParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCardReaderConfigByHashParamsWithTimeout(timeout time.Duration) *CardReaderConfigByHashParams {
	var ()
	return &CardReaderConfigByHashParams{

		timeout: timeout,
	}
}

// NewCardReaderConfigByHashParamsWithContext creates a new CardReaderConfigByHashParams object
// with the default values initialized, and the ability to set a context for a request
func NewCardReaderConfigByHashParamsWithContext(ctx context.Context) *CardReaderConfigByHashParams {
	var ()
	return &CardReaderConfigByHashParams{

		Context: ctx,
	}
}

// NewCardReaderConfigByHashParamsWithHTTPClient creates a new CardReaderConfigByHashParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCardReaderConfigByHashParamsWithHTTPClient(client *http.Client) *CardReaderConfigByHashParams {
	var ()
	return &CardReaderConfigByHashParams{
		HTTPClient: client,
	}
}

/*CardReaderConfigByHashParams contains all the parameters to send to the API endpoint
for the card reader config by hash operation typically these are written to a http.Request
*/
type CardReaderConfigByHashParams struct {

	/*Args*/
	Args CardReaderConfigByHashBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the card reader config by hash params
func (o *CardReaderConfigByHashParams) WithTimeout(timeout time.Duration) *CardReaderConfigByHashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the card reader config by hash params
func (o *CardReaderConfigByHashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the card reader config by hash params
func (o *CardReaderConfigByHashParams) WithContext(ctx context.Context) *CardReaderConfigByHashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the card reader config by hash params
func (o *CardReaderConfigByHashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the card reader config by hash params
func (o *CardReaderConfigByHashParams) WithHTTPClient(client *http.Client) *CardReaderConfigByHashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the card reader config by hash params
func (o *CardReaderConfigByHashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the card reader config by hash params
func (o *CardReaderConfigByHashParams) WithArgs(args CardReaderConfigByHashBody) *CardReaderConfigByHashParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the card reader config by hash params
func (o *CardReaderConfigByHashParams) SetArgs(args CardReaderConfigByHashBody) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *CardReaderConfigByHashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
