// Code generated by go-swagger; DO NOT EDIT.

package bot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new bot API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bot API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostV1BotChannelsChannelIDMessages Creates a new message.
*/
func (a *Client) PostV1BotChannelsChannelIDMessages(params *PostV1BotChannelsChannelIDMessagesParams) (*PostV1BotChannelsChannelIDMessagesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1BotChannelsChannelIDMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postV1BotChannelsChannelIdMessages",
		Method:             "POST",
		PathPattern:        "/v1/bot/channels/{channel_id}/messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostV1BotChannelsChannelIDMessagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostV1BotChannelsChannelIDMessagesCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
