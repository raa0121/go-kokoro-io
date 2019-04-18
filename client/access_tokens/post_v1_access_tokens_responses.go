// Code generated by go-swagger; DO NOT EDIT.

package access_tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/raa0121/go-kokoro-io/models"
)

// PostV1AccessTokensReader is a Reader for the PostV1AccessTokens structure.
type PostV1AccessTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1AccessTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostV1AccessTokensCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1AccessTokensCreated creates a PostV1AccessTokensCreated with default headers values
func NewPostV1AccessTokensCreated() *PostV1AccessTokensCreated {
	return &PostV1AccessTokensCreated{}
}

/*PostV1AccessTokensCreated handles this case with default header values.

Creates new access token
*/
type PostV1AccessTokensCreated struct {
	Payload *models.AccessTokenEntity
}

func (o *PostV1AccessTokensCreated) Error() string {
	return fmt.Sprintf("[POST /v1/access_tokens][%d] postV1AccessTokensCreated  %+v", 201, o.Payload)
}

func (o *PostV1AccessTokensCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessTokenEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
