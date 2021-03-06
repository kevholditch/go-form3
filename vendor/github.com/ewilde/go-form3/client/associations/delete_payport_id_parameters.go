// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePayportIDParams creates a new DeletePayportIDParams object
// with the default values initialized.
func NewDeletePayportIDParams() *DeletePayportIDParams {
	var ()
	return &DeletePayportIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePayportIDParamsWithTimeout creates a new DeletePayportIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePayportIDParamsWithTimeout(timeout time.Duration) *DeletePayportIDParams {
	var ()
	return &DeletePayportIDParams{

		timeout: timeout,
	}
}

// NewDeletePayportIDParamsWithContext creates a new DeletePayportIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePayportIDParamsWithContext(ctx context.Context) *DeletePayportIDParams {
	var ()
	return &DeletePayportIDParams{

		Context: ctx,
	}
}

// NewDeletePayportIDParamsWithHTTPClient creates a new DeletePayportIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePayportIDParamsWithHTTPClient(client *http.Client) *DeletePayportIDParams {
	var ()
	return &DeletePayportIDParams{
		HTTPClient: client,
	}
}

/*DeletePayportIDParams contains all the parameters to send to the API endpoint
for the delete payport ID operation typically these are written to a http.Request
*/
type DeletePayportIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete payport ID params
func (o *DeletePayportIDParams) WithTimeout(timeout time.Duration) *DeletePayportIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete payport ID params
func (o *DeletePayportIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete payport ID params
func (o *DeletePayportIDParams) WithContext(ctx context.Context) *DeletePayportIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete payport ID params
func (o *DeletePayportIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete payport ID params
func (o *DeletePayportIDParams) WithHTTPClient(client *http.Client) *DeletePayportIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete payport ID params
func (o *DeletePayportIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete payport ID params
func (o *DeletePayportIDParams) WithID(id strfmt.UUID) *DeletePayportIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete payport ID params
func (o *DeletePayportIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete payport ID params
func (o *DeletePayportIDParams) WithVersion(version int64) *DeletePayportIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete payport ID params
func (o *DeletePayportIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePayportIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
