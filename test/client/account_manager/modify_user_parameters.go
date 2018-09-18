// Code generated by go-swagger; DO NOT EDIT.

package account_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewModifyUserParams creates a new ModifyUserParams object
// with the default values initialized.
func NewModifyUserParams() *ModifyUserParams {
	var ()
	return &ModifyUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyUserParamsWithTimeout creates a new ModifyUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyUserParamsWithTimeout(timeout time.Duration) *ModifyUserParams {
	var ()
	return &ModifyUserParams{

		timeout: timeout,
	}
}

// NewModifyUserParamsWithContext creates a new ModifyUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyUserParamsWithContext(ctx context.Context) *ModifyUserParams {
	var ()
	return &ModifyUserParams{

		Context: ctx,
	}
}

// NewModifyUserParamsWithHTTPClient creates a new ModifyUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyUserParamsWithHTTPClient(client *http.Client) *ModifyUserParams {
	var ()
	return &ModifyUserParams{
		HTTPClient: client,
	}
}

/*ModifyUserParams contains all the parameters to send to the API endpoint
for the modify user operation typically these are written to a http.Request
*/
type ModifyUserParams struct {

	/*Body*/
	Body *models.OpenpitrixModifyUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify user params
func (o *ModifyUserParams) WithTimeout(timeout time.Duration) *ModifyUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify user params
func (o *ModifyUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify user params
func (o *ModifyUserParams) WithContext(ctx context.Context) *ModifyUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify user params
func (o *ModifyUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify user params
func (o *ModifyUserParams) WithHTTPClient(client *http.Client) *ModifyUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify user params
func (o *ModifyUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify user params
func (o *ModifyUserParams) WithBody(body *models.OpenpitrixModifyUserRequest) *ModifyUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify user params
func (o *ModifyUserParams) SetBody(body *models.OpenpitrixModifyUserRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
