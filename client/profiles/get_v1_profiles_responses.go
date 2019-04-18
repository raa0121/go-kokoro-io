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

// GetV1ProfilesReader is a Reader for the GetV1Profiles structure.
type GetV1ProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1ProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1ProfilesOK creates a GetV1ProfilesOK with default headers values
func NewGetV1ProfilesOK() *GetV1ProfilesOK {
	return &GetV1ProfilesOK{}
}

/*GetV1ProfilesOK handles this case with default header values.

Returns all user's profile except yours
*/
type GetV1ProfilesOK struct {
	Payload *models.ProfileEntity
}

func (o *GetV1ProfilesOK) Error() string {
	return fmt.Sprintf("[GET /v1/profiles][%d] getV1ProfilesOK  %+v", 200, o.Payload)
}

func (o *GetV1ProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProfileEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
