// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNetworksNetworkIDPrometheusAlertConfigParams creates a new GetNetworksNetworkIDPrometheusAlertConfigParams object
// with the default values initialized.
func NewGetNetworksNetworkIDPrometheusAlertConfigParams() *GetNetworksNetworkIDPrometheusAlertConfigParams {
	var ()
	return &GetNetworksNetworkIDPrometheusAlertConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertConfigParamsWithTimeout creates a new GetNetworksNetworkIDPrometheusAlertConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworksNetworkIDPrometheusAlertConfigParamsWithTimeout(timeout time.Duration) *GetNetworksNetworkIDPrometheusAlertConfigParams {
	var ()
	return &GetNetworksNetworkIDPrometheusAlertConfigParams{

		timeout: timeout,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertConfigParamsWithContext creates a new GetNetworksNetworkIDPrometheusAlertConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworksNetworkIDPrometheusAlertConfigParamsWithContext(ctx context.Context) *GetNetworksNetworkIDPrometheusAlertConfigParams {
	var ()
	return &GetNetworksNetworkIDPrometheusAlertConfigParams{

		Context: ctx,
	}
}

// NewGetNetworksNetworkIDPrometheusAlertConfigParamsWithHTTPClient creates a new GetNetworksNetworkIDPrometheusAlertConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworksNetworkIDPrometheusAlertConfigParamsWithHTTPClient(client *http.Client) *GetNetworksNetworkIDPrometheusAlertConfigParams {
	var ()
	return &GetNetworksNetworkIDPrometheusAlertConfigParams{
		HTTPClient: client,
	}
}

/*GetNetworksNetworkIDPrometheusAlertConfigParams contains all the parameters to send to the API endpoint
for the get networks network ID prometheus alert config operation typically these are written to a http.Request
*/
type GetNetworksNetworkIDPrometheusAlertConfigParams struct {

	/*AlertName
	  Name of alert to be retrieved

	*/
	AlertName *string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) WithTimeout(timeout time.Duration) *GetNetworksNetworkIDPrometheusAlertConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) WithContext(ctx context.Context) *GetNetworksNetworkIDPrometheusAlertConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) WithHTTPClient(client *http.Client) *GetNetworksNetworkIDPrometheusAlertConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertName adds the alertName to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) WithAlertName(alertName *string) *GetNetworksNetworkIDPrometheusAlertConfigParams {
	o.SetAlertName(alertName)
	return o
}

// SetAlertName adds the alertName to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) SetAlertName(alertName *string) {
	o.AlertName = alertName
}

// WithNetworkID adds the networkID to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) WithNetworkID(networkID string) *GetNetworksNetworkIDPrometheusAlertConfigParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get networks network ID prometheus alert config params
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworksNetworkIDPrometheusAlertConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AlertName != nil {

		// query param alert_name
		var qrAlertName string
		if o.AlertName != nil {
			qrAlertName = *o.AlertName
		}
		qAlertName := qrAlertName
		if qAlertName != "" {
			if err := r.SetQueryParam("alert_name", qAlertName); err != nil {
				return err
			}
		}

	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
