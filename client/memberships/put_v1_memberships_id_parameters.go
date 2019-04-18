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

// NewPutV1MembershipsIDParams creates a new PutV1MembershipsIDParams object
// with the default values initialized.
func NewPutV1MembershipsIDParams() *PutV1MembershipsIDParams {
	var ()
	return &PutV1MembershipsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1MembershipsIDParamsWithTimeout creates a new PutV1MembershipsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutV1MembershipsIDParamsWithTimeout(timeout time.Duration) *PutV1MembershipsIDParams {
	var ()
	return &PutV1MembershipsIDParams{

		timeout: timeout,
	}
}

// NewPutV1MembershipsIDParamsWithContext creates a new PutV1MembershipsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutV1MembershipsIDParamsWithContext(ctx context.Context) *PutV1MembershipsIDParams {
	var ()
	return &PutV1MembershipsIDParams{

		Context: ctx,
	}
}

// NewPutV1MembershipsIDParamsWithHTTPClient creates a new PutV1MembershipsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutV1MembershipsIDParamsWithHTTPClient(client *http.Client) *PutV1MembershipsIDParams {
	var ()
	return &PutV1MembershipsIDParams{
		HTTPClient: client,
	}
}

/*PutV1MembershipsIDParams contains all the parameters to send to the API endpoint
for the put v1 memberships Id operation typically these are written to a http.Request
*/
type PutV1MembershipsIDParams struct {

	/*ID*/
	ID string
	/*LatestReadMessageID*/
	LatestReadMessageID *int32
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

// WithTimeout adds the timeout to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithTimeout(timeout time.Duration) *PutV1MembershipsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithContext(ctx context.Context) *PutV1MembershipsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithHTTPClient(client *http.Client) *PutV1MembershipsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithID(id string) *PutV1MembershipsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetID(id string) {
	o.ID = id
}

// WithLatestReadMessageID adds the latestReadMessageID to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithLatestReadMessageID(latestReadMessageID *int32) *PutV1MembershipsIDParams {
	o.SetLatestReadMessageID(latestReadMessageID)
	return o
}

// SetLatestReadMessageID adds the latestReadMessageId to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetLatestReadMessageID(latestReadMessageID *int32) {
	o.LatestReadMessageID = latestReadMessageID
}

// WithMuted adds the muted to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithMuted(muted *bool) *PutV1MembershipsIDParams {
	o.SetMuted(muted)
	return o
}

// SetMuted adds the muted to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetMuted(muted *bool) {
	o.Muted = muted
}

// WithNotificationPolicy adds the notificationPolicy to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithNotificationPolicy(notificationPolicy *string) *PutV1MembershipsIDParams {
	o.SetNotificationPolicy(notificationPolicy)
	return o
}

// SetNotificationPolicy adds the notificationPolicy to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetNotificationPolicy(notificationPolicy *string) {
	o.NotificationPolicy = notificationPolicy
}

// WithReadStateTrackingPolicy adds the readStateTrackingPolicy to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithReadStateTrackingPolicy(readStateTrackingPolicy *string) *PutV1MembershipsIDParams {
	o.SetReadStateTrackingPolicy(readStateTrackingPolicy)
	return o
}

// SetReadStateTrackingPolicy adds the readStateTrackingPolicy to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetReadStateTrackingPolicy(readStateTrackingPolicy *string) {
	o.ReadStateTrackingPolicy = readStateTrackingPolicy
}

// WithVisible adds the visible to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) WithVisible(visible *bool) *PutV1MembershipsIDParams {
	o.SetVisible(visible)
	return o
}

// SetVisible adds the visible to the put v1 memberships Id params
func (o *PutV1MembershipsIDParams) SetVisible(visible *bool) {
	o.Visible = visible
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1MembershipsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.LatestReadMessageID != nil {

		// form param latest_read_message_id
		var frLatestReadMessageID int32
		if o.LatestReadMessageID != nil {
			frLatestReadMessageID = *o.LatestReadMessageID
		}
		fLatestReadMessageID := swag.FormatInt32(frLatestReadMessageID)
		if fLatestReadMessageID != "" {
			if err := r.SetFormParam("latest_read_message_id", fLatestReadMessageID); err != nil {
				return err
			}
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
