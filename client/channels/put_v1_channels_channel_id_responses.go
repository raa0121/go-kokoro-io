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

// PutV1ChannelsChannelIDReader is a Reader for the PutV1ChannelsChannelID structure.
type PutV1ChannelsChannelIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1ChannelsChannelIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutV1ChannelsChannelIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutV1ChannelsChannelIDOK creates a PutV1ChannelsChannelIDOK with default headers values
func NewPutV1ChannelsChannelIDOK() *PutV1ChannelsChannelIDOK {
	return &PutV1ChannelsChannelIDOK{}
}

/*PutV1ChannelsChannelIDOK handles this case with default header values.

Updates a channel.
*/
type PutV1ChannelsChannelIDOK struct {
	Payload *models.ChannelEntity
}

func (o *PutV1ChannelsChannelIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/channels/{channel_id}][%d] putV1ChannelsChannelIdOK  %+v", 200, o.Payload)
}

func (o *PutV1ChannelsChannelIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
