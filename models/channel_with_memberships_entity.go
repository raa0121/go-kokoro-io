// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChannelWithMembershipsEntity Returns a channel
// swagger:model ChannelWithMembershipsEntity
type ChannelWithMembershipsEntity struct {

	// アーカイブ済か
	Archived bool `json:"archived,omitempty"`

	// チャンネル名
	ChannelName string `json:"channel_name,omitempty"`

	// チャンネル説明
	Description string `json:"description,omitempty"`

	// チャンネルID
	ID string `json:"id,omitempty"`

	// チャンネルタイプ
	// Enum: [public_channel private_channel direct_message]
	Kind string `json:"kind,omitempty"`

	// 最新メッセージ
	LatestMessageID int32 `json:"latest_message_id,omitempty"`

	// 最新メッセージ投稿日時
	// Format: date-time
	LatestMessagePublishedAt strfmt.DateTime `json:"latest_message_published_at,omitempty"`

	// メンバーシップ情報
	Membership *MembershipEntity `json:"membership,omitempty"`

	// memberships
	Memberships *MembershipWithoutChannelEntity `json:"memberships,omitempty"`

	// メッセージ数
	MessagesCount int32 `json:"messages_count,omitempty"`
}

// Validate validates this channel with memberships entity
func (m *ChannelWithMembershipsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestMessagePublishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembership(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var channelWithMembershipsEntityTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public_channel","private_channel","direct_message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		channelWithMembershipsEntityTypeKindPropEnum = append(channelWithMembershipsEntityTypeKindPropEnum, v)
	}
}

const (

	// ChannelWithMembershipsEntityKindPublicChannel captures enum value "public_channel"
	ChannelWithMembershipsEntityKindPublicChannel string = "public_channel"

	// ChannelWithMembershipsEntityKindPrivateChannel captures enum value "private_channel"
	ChannelWithMembershipsEntityKindPrivateChannel string = "private_channel"

	// ChannelWithMembershipsEntityKindDirectMessage captures enum value "direct_message"
	ChannelWithMembershipsEntityKindDirectMessage string = "direct_message"
)

// prop value enum
func (m *ChannelWithMembershipsEntity) validateKindEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, channelWithMembershipsEntityTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ChannelWithMembershipsEntity) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *ChannelWithMembershipsEntity) validateLatestMessagePublishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LatestMessagePublishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("latest_message_published_at", "body", "date-time", m.LatestMessagePublishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChannelWithMembershipsEntity) validateMembership(formats strfmt.Registry) error {

	if swag.IsZero(m.Membership) { // not required
		return nil
	}

	if m.Membership != nil {
		if err := m.Membership.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("membership")
			}
			return err
		}
	}

	return nil
}

func (m *ChannelWithMembershipsEntity) validateMemberships(formats strfmt.Registry) error {

	if swag.IsZero(m.Memberships) { // not required
		return nil
	}

	if m.Memberships != nil {
		if err := m.Memberships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memberships")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelWithMembershipsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelWithMembershipsEntity) UnmarshalBinary(b []byte) error {
	var res ChannelWithMembershipsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
