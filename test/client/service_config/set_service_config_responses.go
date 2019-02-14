// Code generated by go-swagger; DO NOT EDIT.

package service_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// SetServiceConfigReader is a Reader for the SetServiceConfig structure.
type SetServiceConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetServiceConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetServiceConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetServiceConfigOK creates a SetServiceConfigOK with default headers values
func NewSetServiceConfigOK() *SetServiceConfigOK {
	return &SetServiceConfigOK{}
}

/*SetServiceConfigOK handles this case with default header values.

A successful response.
*/
type SetServiceConfigOK struct {
	Payload *models.OpenpitrixSetServiceConfigResponse
}

func (o *SetServiceConfigOK) Error() string {
	return fmt.Sprintf("[POST /v1/service_configs/set][%d] setServiceConfigOK  %+v", 200, o.Payload)
}

func (o *SetServiceConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixSetServiceConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}