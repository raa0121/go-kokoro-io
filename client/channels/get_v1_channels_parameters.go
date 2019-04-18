// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1ChannelsParams creates a new GetV1ChannelsParams object
// with the default values initialized.
func NewGetV1ChannelsParams() *GetV1ChannelsParams {
	var ()
	return &GetV1ChannelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ChannelsParamsWithTimeout creates a new GetV1ChannelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1ChannelsParamsWithTimeout(timeout time.Duration) *GetV1ChannelsParams {
	var ()
	return &GetV1ChannelsParams{

		timeout: timeout,
	}
}

// NewGetV1ChannelsParamsWithContext creates a new GetV1ChannelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1ChannelsParamsWithContext(ctx context.Context) *GetV1ChannelsParams {
	var ()
	return &GetV1ChannelsParams{

		Context: ctx,
	}
}

// NewGetV1ChannelsParamsWithHTTPClient creates a new GetV1ChannelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1ChannelsParamsWithHTTPClient(client *http.Client) *GetV1ChannelsParams {
	var ()
	return &GetV1ChannelsParams{
		HTTPClient: client,
	}
}

/*GetV1ChannelsParams contains all the parameters to send to the API endpoint
for the get v1 channels operation typically these are written to a http.Request
*/
type GetV1ChannelsParams struct {

	/*Archived*/
	Archived *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 channels params
func (o *GetV1ChannelsParams) WithTimeout(timeout time.Duration) *GetV1ChannelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 channels params
func (o *GetV1ChannelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 channels params
func (o *GetV1ChannelsParams) WithContext(ctx context.Context) *GetV1ChannelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 channels params
func (o *GetV1ChannelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 channels params
func (o *GetV1ChannelsParams) WithHTTPClient(client *http.Client) *GetV1ChannelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 channels params
func (o *GetV1ChannelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchived adds the archived to the get v1 channels params
func (o *GetV1ChannelsParams) WithArchived(archived *bool) *GetV1ChannelsParams {
	o.SetArchived(archived)
	return o
}

// SetArchived adds the archived to the get v1 channels params
func (o *GetV1ChannelsParams) SetArchived(archived *bool) {
	o.Archived = archived
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ChannelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Archived != nil {

		// query param archived
		var qrArchived bool
		if o.Archived != nil {
			qrArchived = *o.Archived
		}
		qArchived := swag.FormatBool(qrArchived)
		if qArchived != "" {
			if err := r.SetQueryParam("archived", qArchived); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}