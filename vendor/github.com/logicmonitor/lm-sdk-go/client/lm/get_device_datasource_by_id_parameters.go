// Code generated by go-swagger; DO NOT EDIT.

package lm

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

// NewGetDeviceDatasourceByIDParams creates a new GetDeviceDatasourceByIDParams object
// with the default values initialized.
func NewGetDeviceDatasourceByIDParams() *GetDeviceDatasourceByIDParams {
	var ()
	return &GetDeviceDatasourceByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceByIDParamsWithTimeout creates a new GetDeviceDatasourceByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceDatasourceByIDParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceByIDParams {
	var ()
	return &GetDeviceDatasourceByIDParams{

		timeout: timeout,
	}
}

// NewGetDeviceDatasourceByIDParamsWithContext creates a new GetDeviceDatasourceByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceDatasourceByIDParamsWithContext(ctx context.Context) *GetDeviceDatasourceByIDParams {
	var ()
	return &GetDeviceDatasourceByIDParams{

		Context: ctx,
	}
}

// NewGetDeviceDatasourceByIDParamsWithHTTPClient creates a new GetDeviceDatasourceByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceDatasourceByIDParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceByIDParams {
	var ()
	return &GetDeviceDatasourceByIDParams{
		HTTPClient: client,
	}
}

/*GetDeviceDatasourceByIDParams contains all the parameters to send to the API endpoint
for the get device datasource by Id operation typically these are written to a http.Request
*/
type GetDeviceDatasourceByIDParams struct {

	/*DeviceID*/
	DeviceID int32
	/*Fields*/
	Fields *string
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) WithContext(ctx context.Context) *GetDeviceDatasourceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) WithFields(fields *string) *GetDeviceDatasourceByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) WithID(id int32) *GetDeviceDatasourceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device datasource by Id params
func (o *GetDeviceDatasourceByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}