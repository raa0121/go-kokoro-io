// Code generated by go-swagger; DO NOT EDIT.

package memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/raa0121/go-kokoro-io/models"
)

// DeleteV1MembershipsIDReader is a Reader for the DeleteV1MembershipsID structure.
type DeleteV1MembershipsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1MembershipsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteV1MembershipsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1MembershipsIDOK creates a DeleteV1MembershipsIDOK with default headers values
func NewDeleteV1MembershipsIDOK() *DeleteV1MembershipsIDOK {
	return &DeleteV1MembershipsIDOK{}
}

/*DeleteV1MembershipsIDOK handles this case with default header values.

Delete a membership.
*/
type DeleteV1MembershipsIDOK struct {
	Payload *models.MembershipEntity
}

func (o *DeleteV1MembershipsIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/memberships/{id}][%d] deleteV1MembershipsIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1MembershipsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MembershipEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}