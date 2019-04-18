// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EmbedDataImageInfoEntity embed data image info entity
// swagger:model EmbedDataImageInfoEntity
type EmbedDataImageInfoEntity struct {

	// Image Height
	Height int32 `json:"height,omitempty"`

	// Image URL
	URL string `json:"url,omitempty"`

	// Image Width
	Width int32 `json:"width,omitempty"`
}

// Validate validates this embed data image info entity
func (m *EmbedDataImageInfoEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmbedDataImageInfoEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmbedDataImageInfoEntity) UnmarshalBinary(b []byte) error {
	var res EmbedDataImageInfoEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}