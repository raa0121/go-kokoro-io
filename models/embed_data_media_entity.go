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

// EmbedDataMediaEntity embed data media entity
// swagger:model EmbedDataMediaEntity
type EmbedDataMediaEntity struct {

	// Raw resource URL of media.
	RawURL string `json:"raw_url,omitempty"`

	// Restriction policy
	// Enum: [Unknown Safe NotSafe]
	RestrictionPolicy string `json:"restriction_policy,omitempty"`

	// Thumbnail image
	Thumbnail *EmbedDataImageInfoEntity `json:"thumbnail,omitempty"`

	// Media type
	// Enum: [Image Video Audio]
	Type string `json:"type,omitempty"`
}

// Validate validates this embed data media entity
func (m *EmbedDataMediaEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestrictionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbnail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var embedDataMediaEntityTypeRestrictionPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Safe","NotSafe"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		embedDataMediaEntityTypeRestrictionPolicyPropEnum = append(embedDataMediaEntityTypeRestrictionPolicyPropEnum, v)
	}
}

const (

	// EmbedDataMediaEntityRestrictionPolicyUnknown captures enum value "Unknown"
	EmbedDataMediaEntityRestrictionPolicyUnknown string = "Unknown"

	// EmbedDataMediaEntityRestrictionPolicySafe captures enum value "Safe"
	EmbedDataMediaEntityRestrictionPolicySafe string = "Safe"

	// EmbedDataMediaEntityRestrictionPolicyNotSafe captures enum value "NotSafe"
	EmbedDataMediaEntityRestrictionPolicyNotSafe string = "NotSafe"
)

// prop value enum
func (m *EmbedDataMediaEntity) validateRestrictionPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, embedDataMediaEntityTypeRestrictionPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EmbedDataMediaEntity) validateRestrictionPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.RestrictionPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateRestrictionPolicyEnum("restriction_policy", "body", m.RestrictionPolicy); err != nil {
		return err
	}

	return nil
}

func (m *EmbedDataMediaEntity) validateThumbnail(formats strfmt.Registry) error {

	if swag.IsZero(m.Thumbnail) { // not required
		return nil
	}

	if m.Thumbnail != nil {
		if err := m.Thumbnail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thumbnail")
			}
			return err
		}
	}

	return nil
}

var embedDataMediaEntityTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Image","Video","Audio"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		embedDataMediaEntityTypeTypePropEnum = append(embedDataMediaEntityTypeTypePropEnum, v)
	}
}

const (

	// EmbedDataMediaEntityTypeImage captures enum value "Image"
	EmbedDataMediaEntityTypeImage string = "Image"

	// EmbedDataMediaEntityTypeVideo captures enum value "Video"
	EmbedDataMediaEntityTypeVideo string = "Video"

	// EmbedDataMediaEntityTypeAudio captures enum value "Audio"
	EmbedDataMediaEntityTypeAudio string = "Audio"
)

// prop value enum
func (m *EmbedDataMediaEntity) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, embedDataMediaEntityTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EmbedDataMediaEntity) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmbedDataMediaEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmbedDataMediaEntity) UnmarshalBinary(b []byte) error {
	var res EmbedDataMediaEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
