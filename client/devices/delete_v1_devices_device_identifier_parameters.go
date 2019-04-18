// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// NewDeleteV1DevicesDeviceIdentifierParams creates a new DeleteV1DevicesDeviceIdentifierParams object
// with the default values initialized.
func NewDeleteV1DevicesDeviceIdentifierParams() *DeleteV1DevicesDeviceIdentifierParams {
	var ()
	return &DeleteV1DevicesDeviceIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1DevicesDeviceIdentifierParamsWithTimeout creates a new DeleteV1DevicesDeviceIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1DevicesDeviceIdentifierParamsWithTimeout(timeout time.Duration) *DeleteV1DevicesDeviceIdentifierParams {
	var ()
	return &DeleteV1DevicesDeviceIdentifierParams{

		timeout: timeout,
	}
}

// NewDeleteV1DevicesDeviceIdentifierParamsWithContext creates a new DeleteV1DevicesDeviceIdentifierParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1DevicesDeviceIdentifierParamsWithContext(ctx context.Context) *DeleteV1DevicesDeviceIdentifierParams {
	var ()
	return &DeleteV1DevicesDeviceIdentifierParams{

		Context: ctx,
	}
}

// NewDeleteV1DevicesDeviceIdentifierParamsWithHTTPClient creates a new DeleteV1DevicesDeviceIdentifierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV1DevicesDeviceIdentifierParamsWithHTTPClient(client *http.Client) *DeleteV1DevicesDeviceIdentifierParams {
	var ()
	return &DeleteV1DevicesDeviceIdentifierParams{
		HTTPClient: client,
	}
}

/*DeleteV1DevicesDeviceIdentifierParams contains all the parameters to send to the API endpoint
for the delete v1 devices device identifier operation typically these are written to a http.Request
*/
type DeleteV1DevicesDeviceIdentifierParams struct {

	/*DeviceIdentifier*/
	DeviceIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v1 devices device identifier params
func (o *DeleteV1DevicesDeviceIdentifierParams) WithTimeout(timeout time.Duration) *DeleteV1DevicesDeviceIdentifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 devices device identifier params
func (o *DeleteV1DevicesDeviceIdentifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 devices device identifier params
func (o *DeleteV1DevicesDeviceIdentifierParams) WithContext(ctx context.Context) *DeleteV1DevicesDeviceIdentifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 devices device identifier params
func (o *DeleteV1DevicesDeviceIdentifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 devices device identifier params
func (o *DeleteV1DevicesDeviceIdentifierParams) WithHTTPClient(client *http.Client) *DeleteV1DevicesDeviceIdentifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 devices device identifier params
func (o *DeleteV1DevicesDeviceIdentifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceIdentifier adds the deviceIdentifier to the delete v1 devices device identifier params
func (o *DeleteV1DevicesDeviceIdentifierParams) WithDeviceIdentifier(deviceIdentifier string) *DeleteV1DevicesDeviceIdentifierParams {
	o.SetDeviceIdentifier(deviceIdentifier)
	return o
}

// SetDeviceIdentifier adds the deviceIdentifier to the delete v1 devices device identifier params
func (o *DeleteV1DevicesDeviceIdentifierParams) SetDeviceIdentifier(deviceIdentifier string) {
	o.DeviceIdentifier = deviceIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1DevicesDeviceIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_identifier
	if err := r.SetPathParam("device_identifier", o.DeviceIdentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}