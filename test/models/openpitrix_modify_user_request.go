// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyUserRequest openpitrix modify user request
// swagger:model openpitrixModifyUserRequest
type OpenpitrixModifyUserRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this openpitrix modify user request
func (m *OpenpitrixModifyUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyUserRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
