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

// MembershipEntity Updates a membership.
// swagger:model MembershipEntity
type MembershipEntity struct {

	// 権限
	// Enum: [administrator maintainer member invited]
	Authority string `json:"authority,omitempty"`

	// チャンネル情報
	Channel *ChannelWithoutMembershipEntity `json:"channel,omitempty"`

	// 通知をオフにしているか（deprecated: use "notification_policy" instead.）
	DisableNotification bool `json:"disable_notification,omitempty"`

	// メンバーシップID
	ID string `json:"id,omitempty"`

	// 一番新しい既読メッセージ
	LatestReadMessageID int32 `json:"latest_read_message_id,omitempty"`

	// チャット画面のチャンネル一覧ペインにて未読数表示をオフにし、表示を薄くする
	Muted bool `json:"muted,omitempty"`

	// 通知ポリシー
	// Enum: [all_messages only_mentions nothing]
	NotificationPolicy string `json:"notification_policy,omitempty"`

	// プロフィール情報
	Profile *ProfileEntity `json:"profile,omitempty"`

	// 未読メッセージ表示ポリシー
	// Enum: [keep_latest consume_last consume_latest]
	ReadStateTrackingPolicy string `json:"read_state_tracking_policy,omitempty"`

	// 未読数（WIP）
	UnreadCount int32 `json:"unread_count,omitempty"`

	// チャット画面のチャンネル一覧ペインに表示する
	Visible bool `json:"visible,omitempty"`
}

// Validate validates this membership entity
func (m *MembershipEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotificationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadStateTrackingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var membershipEntityTypeAuthorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["administrator","maintainer","member","invited"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		membershipEntityTypeAuthorityPropEnum = append(membershipEntityTypeAuthorityPropEnum, v)
	}
}

const (

	// MembershipEntityAuthorityAdministrator captures enum value "administrator"
	MembershipEntityAuthorityAdministrator string = "administrator"

	// MembershipEntityAuthorityMaintainer captures enum value "maintainer"
	MembershipEntityAuthorityMaintainer string = "maintainer"

	// MembershipEntityAuthorityMember captures enum value "member"
	MembershipEntityAuthorityMember string = "member"

	// MembershipEntityAuthorityInvited captures enum value "invited"
	MembershipEntityAuthorityInvited string = "invited"
)

// prop value enum
func (m *MembershipEntity) validateAuthorityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, membershipEntityTypeAuthorityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MembershipEntity) validateAuthority(formats strfmt.Registry) error {

	if swag.IsZero(m.Authority) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthorityEnum("authority", "body", m.Authority); err != nil {
		return err
	}

	return nil
}

func (m *MembershipEntity) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

var membershipEntityTypeNotificationPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all_messages","only_mentions","nothing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		membershipEntityTypeNotificationPolicyPropEnum = append(membershipEntityTypeNotificationPolicyPropEnum, v)
	}
}

const (

	// MembershipEntityNotificationPolicyAllMessages captures enum value "all_messages"
	MembershipEntityNotificationPolicyAllMessages string = "all_messages"

	// MembershipEntityNotificationPolicyOnlyMentions captures enum value "only_mentions"
	MembershipEntityNotificationPolicyOnlyMentions string = "only_mentions"

	// MembershipEntityNotificationPolicyNothing captures enum value "nothing"
	MembershipEntityNotificationPolicyNothing string = "nothing"
)

// prop value enum
func (m *MembershipEntity) validateNotificationPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, membershipEntityTypeNotificationPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MembershipEntity) validateNotificationPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.NotificationPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateNotificationPolicyEnum("notification_policy", "body", m.NotificationPolicy); err != nil {
		return err
	}

	return nil
}

func (m *MembershipEntity) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

var membershipEntityTypeReadStateTrackingPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["keep_latest","consume_last","consume_latest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		membershipEntityTypeReadStateTrackingPolicyPropEnum = append(membershipEntityTypeReadStateTrackingPolicyPropEnum, v)
	}
}

const (

	// MembershipEntityReadStateTrackingPolicyKeepLatest captures enum value "keep_latest"
	MembershipEntityReadStateTrackingPolicyKeepLatest string = "keep_latest"

	// MembershipEntityReadStateTrackingPolicyConsumeLast captures enum value "consume_last"
	MembershipEntityReadStateTrackingPolicyConsumeLast string = "consume_last"

	// MembershipEntityReadStateTrackingPolicyConsumeLatest captures enum value "consume_latest"
	MembershipEntityReadStateTrackingPolicyConsumeLatest string = "consume_latest"
)

// prop value enum
func (m *MembershipEntity) validateReadStateTrackingPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, membershipEntityTypeReadStateTrackingPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MembershipEntity) validateReadStateTrackingPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ReadStateTrackingPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateReadStateTrackingPolicyEnum("read_state_tracking_policy", "body", m.ReadStateTrackingPolicy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MembershipEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MembershipEntity) UnmarshalBinary(b []byte) error {
	var res MembershipEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
