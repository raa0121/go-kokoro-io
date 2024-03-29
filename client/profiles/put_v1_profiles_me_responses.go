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

// PutV1ProfilesMeReader is a Reader for the PutV1ProfilesMe structure.
type PutV1ProfilesMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1ProfilesMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutV1ProfilesMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutV1ProfilesMeOK creates a PutV1ProfilesMeOK with default headers values
func NewPutV1ProfilesMeOK() *PutV1ProfilesMeOK {
	return &PutV1ProfilesMeOK{}
}

/*PutV1ProfilesMeOK handles this case with default header values.

Updates your profile.
*/
type PutV1ProfilesMeOK struct {
	Payload *models.ProfileEntity
}

func (o *PutV1ProfilesMeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/profiles/me][%d] putV1ProfilesMeOK  %+v", 200, o.Payload)
}

func (o *PutV1ProfilesMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProfileEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
