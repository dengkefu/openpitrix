// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetPasswordResetResponse openpitrix get password reset response
// swagger:model openpitrixGetPasswordResetResponse
type OpenpitrixGetPasswordResetResponse struct {

	// reset id
	ResetID string `json:"reset_id,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this openpitrix get password reset response
func (m *OpenpitrixGetPasswordResetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetPasswordResetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetPasswordResetResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetPasswordResetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
