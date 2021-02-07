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

	model "github.com/DiaElectronics/lea-central-wash/storageapi/model"
)

// NewSetCardReaderConfigParams creates a new SetCardReaderConfigParams object
// with the default values initialized.
func NewSetCardReaderConfigParams() *SetCardReaderConfigParams {
	var ()
	return &SetCardReaderConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetCardReaderConfigParamsWithTimeout creates a new SetCardReaderConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetCardReaderConfigParamsWithTimeout(timeout time.Duration) *SetCardReaderConfigParams {
	var ()
	return &SetCardReaderConfigParams{

		timeout: timeout,
	}
}

// NewSetCardReaderConfigParamsWithContext creates a new SetCardReaderConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetCardReaderConfigParamsWithContext(ctx context.Context) *SetCardReaderConfigParams {
	var ()
	return &SetCardReaderConfigParams{

		Context: ctx,
	}
}

// NewSetCardReaderConfigParamsWithHTTPClient creates a new SetCardReaderConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetCardReaderConfigParamsWithHTTPClient(client *http.Client) *SetCardReaderConfigParams {
	var ()
	return &SetCardReaderConfigParams{
		HTTPClient: client,
	}
}

/*SetCardReaderConfigParams contains all the parameters to send to the API endpoint
for the set card reader config operation typically these are written to a http.Request
*/
type SetCardReaderConfigParams struct {

	/*Args*/
	Args *model.CardReaderConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set card reader config params
func (o *SetCardReaderConfigParams) WithTimeout(timeout time.Duration) *SetCardReaderConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set card reader config params
func (o *SetCardReaderConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set card reader config params
func (o *SetCardReaderConfigParams) WithContext(ctx context.Context) *SetCardReaderConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set card reader config params
func (o *SetCardReaderConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set card reader config params
func (o *SetCardReaderConfigParams) WithHTTPClient(client *http.Client) *SetCardReaderConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set card reader config params
func (o *SetCardReaderConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the set card reader config params
func (o *SetCardReaderConfigParams) WithArgs(args *model.CardReaderConfig) *SetCardReaderConfigParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the set card reader config params
func (o *SetCardReaderConfigParams) SetArgs(args *model.CardReaderConfig) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *SetCardReaderConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Args != nil {
		if err := r.SetBodyParam(o.Args); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}