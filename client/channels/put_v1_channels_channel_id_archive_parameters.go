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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutV1ChannelsChannelIDArchiveParams creates a new PutV1ChannelsChannelIDArchiveParams object
// with the default values initialized.
func NewPutV1ChannelsChannelIDArchiveParams() *PutV1ChannelsChannelIDArchiveParams {
	var ()
	return &PutV1ChannelsChannelIDArchiveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1ChannelsChannelIDArchiveParamsWithTimeout creates a new PutV1ChannelsChannelIDArchiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutV1ChannelsChannelIDArchiveParamsWithTimeout(timeout time.Duration) *PutV1ChannelsChannelIDArchiveParams {
	var ()
	return &PutV1ChannelsChannelIDArchiveParams{

		timeout: timeout,
	}
}

// NewPutV1ChannelsChannelIDArchiveParamsWithContext creates a new PutV1ChannelsChannelIDArchiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutV1ChannelsChannelIDArchiveParamsWithContext(ctx context.Context) *PutV1ChannelsChannelIDArchiveParams {
	var ()
	return &PutV1ChannelsChannelIDArchiveParams{

		Context: ctx,
	}
}

// NewPutV1ChannelsChannelIDArchiveParamsWithHTTPClient creates a new PutV1ChannelsChannelIDArchiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutV1ChannelsChannelIDArchiveParamsWithHTTPClient(client *http.Client) *PutV1ChannelsChannelIDArchiveParams {
	var ()
	return &PutV1ChannelsChannelIDArchiveParams{
		HTTPClient: client,
	}
}

/*PutV1ChannelsChannelIDArchiveParams contains all the parameters to send to the API endpoint
for the put v1 channels channel Id archive operation typically these are written to a http.Request
*/
type PutV1ChannelsChannelIDArchiveParams struct {

	/*ChannelID*/
	ChannelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put v1 channels channel Id archive params
func (o *PutV1ChannelsChannelIDArchiveParams) WithTimeout(timeout time.Duration) *PutV1ChannelsChannelIDArchiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 channels channel Id archive params
func (o *PutV1ChannelsChannelIDArchiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 channels channel Id archive params
func (o *PutV1ChannelsChannelIDArchiveParams) WithContext(ctx context.Context) *PutV1ChannelsChannelIDArchiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 channels channel Id archive params
func (o *PutV1ChannelsChannelIDArchiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 channels channel Id archive params
func (o *PutV1ChannelsChannelIDArchiveParams) WithHTTPClient(client *http.Client) *PutV1ChannelsChannelIDArchiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 channels channel Id archive params
func (o *PutV1ChannelsChannelIDArchiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the put v1 channels channel Id archive params
func (o *PutV1ChannelsChannelIDArchiveParams) WithChannelID(channelID string) *PutV1ChannelsChannelIDArchiveParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the put v1 channels channel Id archive params
func (o *PutV1ChannelsChannelIDArchiveParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1ChannelsChannelIDArchiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel_id
	if err := r.SetPathParam("channel_id", o.ChannelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
