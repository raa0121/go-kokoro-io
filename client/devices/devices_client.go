// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new devices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for devices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteV1DevicesDeviceIdentifier Delete a device
*/
func (a *Client) DeleteV1DevicesDeviceIdentifier(params *DeleteV1DevicesDeviceIdentifierParams) (*DeleteV1DevicesDeviceIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1DevicesDeviceIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteV1DevicesDeviceIdentifier",
		Method:             "DELETE",
		PathPattern:        "/v1/devices/{device_identifier}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteV1DevicesDeviceIdentifierReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteV1DevicesDeviceIdentifierOK), nil

}

/*
GetV1Devices Returns user's devices
*/
func (a *Client) GetV1Devices(params *GetV1DevicesParams) (*GetV1DevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DevicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getV1Devices",
		Method:             "GET",
		PathPattern:        "/v1/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DevicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetV1DevicesOK), nil

}

/*
PostV1Devices Creates new device
*/
func (a *Client) PostV1Devices(params *PostV1DevicesParams) (*PostV1DevicesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1DevicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postV1Devices",
		Method:             "POST",
		PathPattern:        "/v1/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostV1DevicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostV1DevicesCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
