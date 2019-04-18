// Code generated by go-swagger; DO NOT EDIT.

package memberships

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

// NewPostV1MembershipsParams creates a new PostV1MembershipsParams object
// with the default values initialized.
func NewPostV1MembershipsParams() *PostV1MembershipsParams {
	var ()
	return &PostV1MembershipsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1MembershipsParamsWithTimeout creates a new PostV1MembershipsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostV1MembershipsParamsWithTimeout(timeout time.Duration) *PostV1MembershipsParams {
	var ()
	return &PostV1MembershipsParams{

		timeout: timeout,
	}
}

// NewPostV1MembershipsParamsWithContext creates a new PostV1MembershipsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostV1MembershipsParamsWithContext(ctx context.Context) *PostV1MembershipsParams {
	var ()
	return &PostV1MembershipsParams{

		Context: ctx,
	}
}

// NewPostV1MembershipsParamsWithHTTPClient creates a new PostV1MembershipsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostV1MembershipsParamsWithHTTPClient(client *http.Client) *PostV1MembershipsParams {
	var ()
	return &PostV1MembershipsParams{
		HTTPClient: client,
	}
}

/*PostV1MembershipsParams contains all the parameters to send to the API endpoint
for the post v1 memberships operation typically these are written to a http.Request
*/
type PostV1MembershipsParams struct {

	/*ChannelID*/
	ChannelID string
	/*Muted*/
	Muted *bool
	/*NotificationPolicy*/
	NotificationPolicy *string
	/*ReadStateTrackingPolicy*/
	ReadStateTrackingPolicy *string
	/*Visible*/
	Visible *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post v1 memberships params
func (o *PostV1MembershipsParams) WithTimeout(timeout time.Duration) *PostV1MembershipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 memberships params
func (o *PostV1MembershipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 memberships params
func (o *PostV1MembershipsParams) WithContext(ctx context.Context) *PostV1MembershipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 memberships params
func (o *PostV1MembershipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 memberships params
func (o *PostV1MembershipsParams) WithHTTPClient(client *http.Client) *PostV1MembershipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 memberships params
func (o *PostV1MembershipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the post v1 memberships params
func (o *PostV1MembershipsParams) WithChannelID(channelID string) *PostV1MembershipsParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the post v1 memberships params
func (o *PostV1MembershipsParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithMuted adds the muted to the post v1 memberships params
func (o *PostV1MembershipsParams) WithMuted(muted *bool) *PostV1MembershipsParams {
	o.SetMuted(muted)
	return o
}

// SetMuted adds the muted to the post v1 memberships params
func (o *PostV1MembershipsParams) SetMuted(muted *bool) {
	o.Muted = muted
}

// WithNotificationPolicy adds the notificationPolicy to the post v1 memberships params
func (o *PostV1MembershipsParams) WithNotificationPolicy(notificationPolicy *string) *PostV1MembershipsParams {
	o.SetNotificationPolicy(notificationPolicy)
	return o
}

// SetNotificationPolicy adds the notificationPolicy to the post v1 memberships params
func (o *PostV1MembershipsParams) SetNotificationPolicy(notificationPolicy *string) {
	o.NotificationPolicy = notificationPolicy
}

// WithReadStateTrackingPolicy adds the readStateTrackingPolicy to the post v1 memberships params
func (o *PostV1MembershipsParams) WithReadStateTrackingPolicy(readStateTrackingPolicy *string) *PostV1MembershipsParams {
	o.SetReadStateTrackingPolicy(readStateTrackingPolicy)
	return o
}

// SetReadStateTrackingPolicy adds the readStateTrackingPolicy to the post v1 memberships params
func (o *PostV1MembershipsParams) SetReadStateTrackingPolicy(readStateTrackingPolicy *string) {
	o.ReadStateTrackingPolicy = readStateTrackingPolicy
}

// WithVisible adds the visible to the post v1 memberships params
func (o *PostV1MembershipsParams) WithVisible(visible *bool) *PostV1MembershipsParams {
	o.SetVisible(visible)
	return o
}

// SetVisible adds the visible to the post v1 memberships params
func (o *PostV1MembershipsParams) SetVisible(visible *bool) {
	o.Visible = visible
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1MembershipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param channel_id
	frChannelID := o.ChannelID
	fChannelID := frChannelID
	if fChannelID != "" {
		if err := r.SetFormParam("channel_id", fChannelID); err != nil {
			return err
		}
	}

	if o.Muted != nil {

		// form param muted
		var frMuted bool
		if o.Muted != nil {
			frMuted = *o.Muted
		}
		fMuted := swag.FormatBool(frMuted)
		if fMuted != "" {
			if err := r.SetFormParam("muted", fMuted); err != nil {
				return err
			}
		}

	}

	if o.NotificationPolicy != nil {

		// form param notification_policy
		var frNotificationPolicy string
		if o.NotificationPolicy != nil {
			frNotificationPolicy = *o.NotificationPolicy
		}
		fNotificationPolicy := frNotificationPolicy
		if fNotificationPolicy != "" {
			if err := r.SetFormParam("notification_policy", fNotificationPolicy); err != nil {
				return err
			}
		}

	}

	if o.ReadStateTrackingPolicy != nil {

		// form param read_state_tracking_policy
		var frReadStateTrackingPolicy string
		if o.ReadStateTrackingPolicy != nil {
			frReadStateTrackingPolicy = *o.ReadStateTrackingPolicy
		}
		fReadStateTrackingPolicy := frReadStateTrackingPolicy
		if fReadStateTrackingPolicy != "" {
			if err := r.SetFormParam("read_state_tracking_policy", fReadStateTrackingPolicy); err != nil {
				return err
			}
		}

	}

	if o.Visible != nil {

		// form param visible
		var frVisible bool
		if o.Visible != nil {
			frVisible = *o.Visible
		}
		fVisible := swag.FormatBool(frVisible)
		if fVisible != "" {
			if err := r.SetFormParam("visible", fVisible); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
