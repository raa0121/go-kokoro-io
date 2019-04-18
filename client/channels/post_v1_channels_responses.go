// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/raa0121/go-kokoro-io/models"
)

// PostV1ChannelsReader is a Reader for the PostV1Channels structure.
type PostV1ChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostV1ChannelsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1ChannelsCreated creates a PostV1ChannelsCreated with default headers values
func NewPostV1ChannelsCreated() *PostV1ChannelsCreated {
	return &PostV1ChannelsCreated{}
}

/*PostV1ChannelsCreated handles this case with default header values.

Creates a new channel.
*/
type PostV1ChannelsCreated struct {
	Payload *models.ChannelEntity
}

func (o *PostV1ChannelsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/channels][%d] postV1ChannelsCreated  %+v", 201, o.Payload)
}

func (o *PostV1ChannelsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}