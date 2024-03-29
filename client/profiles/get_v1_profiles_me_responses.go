// Code generated by go-swagger; DO NOT EDIT.

package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/raa0121/go-kokoro-io/models"
)

// GetV1ProfilesMeReader is a Reader for the GetV1ProfilesMe structure.
type GetV1ProfilesMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ProfilesMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1ProfilesMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1ProfilesMeOK creates a GetV1ProfilesMeOK with default headers values
func NewGetV1ProfilesMeOK() *GetV1ProfilesMeOK {
	return &GetV1ProfilesMeOK{}
}

/*GetV1ProfilesMeOK handles this case with default header values.

Returns current logged-in user's profile
*/
type GetV1ProfilesMeOK struct {
	Payload *models.ProfileEntity
}

func (o *GetV1ProfilesMeOK) Error() string {
	return fmt.Sprintf("[GET /v1/profiles/me][%d] getV1ProfilesMeOK  %+v", 200, o.Payload)
}

func (o *GetV1ProfilesMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProfileEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
