// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MessageEntity Creates a new message.
// swagger:model MessageEntity
type MessageEntity struct {

	// 発言時のアバターURL
	Avatar string `json:"avatar,omitempty"`

	// avatars
	Avatars []*AvatarEntity `json:"avatars"`

	// 発言があったチャンネル
	Channel *ChannelEntity `json:"channel,omitempty"`

	// 発言内容(Deprecated. Use 'html_content' property instead of this.)
	Content string `json:"content,omitempty"`

	// 発言時の表示名
	DisplayName string `json:"display_name,omitempty"`

	// 埋め込みコンテンツ
	EmbedContents []*EmbedContentEntity `json:"embed_contents"`

	// 埋め込みURL
	EmbeddedUrls []string `json:"embedded_urls"`

	// URLが表すコンテンツを展開するかどうか
	ExpandEmbedContents bool `json:"expand_embed_contents,omitempty"`

	// 発言内容
	HTMLContent string `json:"html_content,omitempty"`

	// メッセージID
	ID int32 `json:"id,omitempty"`

	// 冪等性ID(Version 4 UUID)
	IdempotentKey string `json:"idempotent_key,omitempty"`

	// NSFW(Not suitable for work)かどうか
	Nsfw bool `json:"nsfw,omitempty"`

	// プレインテキスト形式に変換した発言内容
	PlaintextContent string `json:"plaintext_content,omitempty"`

	// 発言者情報
	Profile *ProfileEntity `json:"profile,omitempty"`

	// 発言日時
	// Format: date-time
	PublishedAt strfmt.DateTime `json:"published_at,omitempty"`

	// 発言内容（プレインテキスト）
	RawContent string `json:"raw_content,omitempty"`

	// 発言の状態
	// Enum: [active deleted_by_publisher deleted_by_another_member]
	Status string `json:"status,omitempty"`
}

// Validate validates this message entity
func (m *MessageEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvatars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmbedContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageEntity) validateAvatars(formats strfmt.Registry) error {

	if swag.IsZero(m.Avatars) { // not required
		return nil
	}

	for i := 0; i < len(m.Avatars); i++ {
		if swag.IsZero(m.Avatars[i]) { // not required
			continue
		}

		if m.Avatars[i] != nil {
			if err := m.Avatars[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("avatars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessageEntity) validateChannel(formats strfmt.Registry) error {

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

func (m *MessageEntity) validateEmbedContents(formats strfmt.Registry) error {

	if swag.IsZero(m.EmbedContents) { // not required
		return nil
	}

	for i := 0; i < len(m.EmbedContents); i++ {
		if swag.IsZero(m.EmbedContents[i]) { // not required
			continue
		}

		if m.EmbedContents[i] != nil {
			if err := m.EmbedContents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("embed_contents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessageEntity) validateProfile(formats strfmt.Registry) error {

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

func (m *MessageEntity) validatePublishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("published_at", "body", "date-time", m.PublishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var messageEntityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","deleted_by_publisher","deleted_by_another_member"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageEntityTypeStatusPropEnum = append(messageEntityTypeStatusPropEnum, v)
	}
}

const (

	// MessageEntityStatusActive captures enum value "active"
	MessageEntityStatusActive string = "active"

	// MessageEntityStatusDeletedByPublisher captures enum value "deleted_by_publisher"
	MessageEntityStatusDeletedByPublisher string = "deleted_by_publisher"

	// MessageEntityStatusDeletedByAnotherMember captures enum value "deleted_by_another_member"
	MessageEntityStatusDeletedByAnotherMember string = "deleted_by_another_member"
)

// prop value enum
func (m *MessageEntity) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, messageEntityTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MessageEntity) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageEntity) UnmarshalBinary(b []byte) error {
	var res MessageEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
