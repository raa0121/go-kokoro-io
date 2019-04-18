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

// MembershipWithoutChannelEntity membership without channel entity
// swagger:model MembershipWithoutChannelEntity
type MembershipWithoutChannelEntity struct {

	// 権限
	// Enum: [administrator maintainer member invited]
	Authority string `json:"authority,omitempty"`

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

// Validate validates this membership without channel entity
func (m *MembershipWithoutChannelEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthority(formats); err != nil {
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

var membershipWithoutChannelEntityTypeAuthorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["administrator","maintainer","member","invited"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		membershipWithoutChannelEntityTypeAuthorityPropEnum = append(membershipWithoutChannelEntityTypeAuthorityPropEnum, v)
	}
}

const (

	// MembershipWithoutChannelEntityAuthorityAdministrator captures enum value "administrator"
	MembershipWithoutChannelEntityAuthorityAdministrator string = "administrator"

	// MembershipWithoutChannelEntityAuthorityMaintainer captures enum value "maintainer"
	MembershipWithoutChannelEntityAuthorityMaintainer string = "maintainer"

	// MembershipWithoutChannelEntityAuthorityMember captures enum value "member"
	MembershipWithoutChannelEntityAuthorityMember string = "member"

	// MembershipWithoutChannelEntityAuthorityInvited captures enum value "invited"
	MembershipWithoutChannelEntityAuthorityInvited string = "invited"
)

// prop value enum
func (m *MembershipWithoutChannelEntity) validateAuthorityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, membershipWithoutChannelEntityTypeAuthorityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MembershipWithoutChannelEntity) validateAuthority(formats strfmt.Registry) error {

	if swag.IsZero(m.Authority) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthorityEnum("authority", "body", m.Authority); err != nil {
		return err
	}

	return nil
}

var membershipWithoutChannelEntityTypeNotificationPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all_messages","only_mentions","nothing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		membershipWithoutChannelEntityTypeNotificationPolicyPropEnum = append(membershipWithoutChannelEntityTypeNotificationPolicyPropEnum, v)
	}
}

const (

	// MembershipWithoutChannelEntityNotificationPolicyAllMessages captures enum value "all_messages"
	MembershipWithoutChannelEntityNotificationPolicyAllMessages string = "all_messages"

	// MembershipWithoutChannelEntityNotificationPolicyOnlyMentions captures enum value "only_mentions"
	MembershipWithoutChannelEntityNotificationPolicyOnlyMentions string = "only_mentions"

	// MembershipWithoutChannelEntityNotificationPolicyNothing captures enum value "nothing"
	MembershipWithoutChannelEntityNotificationPolicyNothing string = "nothing"
)

// prop value enum
func (m *MembershipWithoutChannelEntity) validateNotificationPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, membershipWithoutChannelEntityTypeNotificationPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MembershipWithoutChannelEntity) validateNotificationPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.NotificationPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateNotificationPolicyEnum("notification_policy", "body", m.NotificationPolicy); err != nil {
		return err
	}

	return nil
}

func (m *MembershipWithoutChannelEntity) validateProfile(formats strfmt.Registry) error {

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

var membershipWithoutChannelEntityTypeReadStateTrackingPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["keep_latest","consume_last","consume_latest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		membershipWithoutChannelEntityTypeReadStateTrackingPolicyPropEnum = append(membershipWithoutChannelEntityTypeReadStateTrackingPolicyPropEnum, v)
	}
}

const (

	// MembershipWithoutChannelEntityReadStateTrackingPolicyKeepLatest captures enum value "keep_latest"
	MembershipWithoutChannelEntityReadStateTrackingPolicyKeepLatest string = "keep_latest"

	// MembershipWithoutChannelEntityReadStateTrackingPolicyConsumeLast captures enum value "consume_last"
	MembershipWithoutChannelEntityReadStateTrackingPolicyConsumeLast string = "consume_last"

	// MembershipWithoutChannelEntityReadStateTrackingPolicyConsumeLatest captures enum value "consume_latest"
	MembershipWithoutChannelEntityReadStateTrackingPolicyConsumeLatest string = "consume_latest"
)

// prop value enum
func (m *MembershipWithoutChannelEntity) validateReadStateTrackingPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, membershipWithoutChannelEntityTypeReadStateTrackingPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MembershipWithoutChannelEntity) validateReadStateTrackingPolicy(formats strfmt.Registry) error {

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
func (m *MembershipWithoutChannelEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MembershipWithoutChannelEntity) UnmarshalBinary(b []byte) error {
	var res MembershipWithoutChannelEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
