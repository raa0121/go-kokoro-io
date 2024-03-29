// Code generated by go-swagger; DO NOT EDIT.

package bot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/raa0121/go-kokoro-io/models"
)

// PostV1BotChannelsChannelIDMessagesReader is a Reader for the PostV1BotChannelsChannelIDMessages structure.
type PostV1BotChannelsChannelIDMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1BotChannelsChannelIDMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostV1BotChannelsChannelIDMessagesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1BotChannelsChannelIDMessagesCreated creates a PostV1BotChannelsChannelIDMessagesCreated with default headers values
func NewPostV1BotChannelsChannelIDMessagesCreated() *PostV1BotChannelsChannelIDMessagesCreated {
	return &PostV1BotChannelsChannelIDMessagesCreated{}
}

/*PostV1BotChannelsChannelIDMessagesCreated handles this case with default header values.

Creates a new message.
*/
type PostV1BotChannelsChannelIDMessagesCreated struct {
	Payload *models.MessageEntity
}

func (o *PostV1BotChannelsChannelIDMessagesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/bot/channels/{channel_id}/messages][%d] postV1BotChannelsChannelIdMessagesCreated  %+v", 201, o.Payload)
}

func (o *PostV1BotChannelsChannelIDMessagesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
