// Code generated by go-swagger; DO NOT EDIT.

package messages

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

// NewDeleteV1MessagesMessageIDParams creates a new DeleteV1MessagesMessageIDParams object
// with the default values initialized.
func NewDeleteV1MessagesMessageIDParams() *DeleteV1MessagesMessageIDParams {
	var ()
	return &DeleteV1MessagesMessageIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1MessagesMessageIDParamsWithTimeout creates a new DeleteV1MessagesMessageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1MessagesMessageIDParamsWithTimeout(timeout time.Duration) *DeleteV1MessagesMessageIDParams {
	var ()
	return &DeleteV1MessagesMessageIDParams{

		timeout: timeout,
	}
}

// NewDeleteV1MessagesMessageIDParamsWithContext creates a new DeleteV1MessagesMessageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1MessagesMessageIDParamsWithContext(ctx context.Context) *DeleteV1MessagesMessageIDParams {
	var ()
	return &DeleteV1MessagesMessageIDParams{

		Context: ctx,
	}
}

// NewDeleteV1MessagesMessageIDParamsWithHTTPClient creates a new DeleteV1MessagesMessageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV1MessagesMessageIDParamsWithHTTPClient(client *http.Client) *DeleteV1MessagesMessageIDParams {
	var ()
	return &DeleteV1MessagesMessageIDParams{
		HTTPClient: client,
	}
}

/*DeleteV1MessagesMessageIDParams contains all the parameters to send to the API endpoint
for the delete v1 messages message Id operation typically these are written to a http.Request
*/
type DeleteV1MessagesMessageIDParams struct {

	/*MessageID*/
	MessageID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v1 messages message Id params
func (o *DeleteV1MessagesMessageIDParams) WithTimeout(timeout time.Duration) *DeleteV1MessagesMessageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 messages message Id params
func (o *DeleteV1MessagesMessageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 messages message Id params
func (o *DeleteV1MessagesMessageIDParams) WithContext(ctx context.Context) *DeleteV1MessagesMessageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 messages message Id params
func (o *DeleteV1MessagesMessageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 messages message Id params
func (o *DeleteV1MessagesMessageIDParams) WithHTTPClient(client *http.Client) *DeleteV1MessagesMessageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 messages message Id params
func (o *DeleteV1MessagesMessageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessageID adds the messageID to the delete v1 messages message Id params
func (o *DeleteV1MessagesMessageIDParams) WithMessageID(messageID int32) *DeleteV1MessagesMessageIDParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the delete v1 messages message Id params
func (o *DeleteV1MessagesMessageIDParams) SetMessageID(messageID int32) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1MessagesMessageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param message_id
	if err := r.SetPathParam("message_id", swag.FormatInt32(o.MessageID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
