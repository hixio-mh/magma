// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

// NewGetLTENetworkIDGatewayPoolsGatewayPoolIDParams creates a new GetLTENetworkIDGatewayPoolsGatewayPoolIDParams object
// with the default values initialized.
func NewGetLTENetworkIDGatewayPoolsGatewayPoolIDParams() *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	var ()
	return &GetLTENetworkIDGatewayPoolsGatewayPoolIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewayPoolsGatewayPoolIDParamsWithTimeout creates a new GetLTENetworkIDGatewayPoolsGatewayPoolIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDGatewayPoolsGatewayPoolIDParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	var ()
	return &GetLTENetworkIDGatewayPoolsGatewayPoolIDParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewayPoolsGatewayPoolIDParamsWithContext creates a new GetLTENetworkIDGatewayPoolsGatewayPoolIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDGatewayPoolsGatewayPoolIDParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	var ()
	return &GetLTENetworkIDGatewayPoolsGatewayPoolIDParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewayPoolsGatewayPoolIDParamsWithHTTPClient creates a new GetLTENetworkIDGatewayPoolsGatewayPoolIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDGatewayPoolsGatewayPoolIDParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	var ()
	return &GetLTENetworkIDGatewayPoolsGatewayPoolIDParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDGatewayPoolsGatewayPoolIDParams contains all the parameters to send to the API endpoint
for the get LTE network ID gateway pools gateway pool ID operation typically these are written to a http.Request
*/
type GetLTENetworkIDGatewayPoolsGatewayPoolIDParams struct {

	/*GatewayPoolID
	  Gateway Pool ID

	*/
	GatewayPoolID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayPoolID adds the gatewayPoolID to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) WithGatewayPoolID(gatewayPoolID string) *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	o.SetGatewayPoolID(gatewayPoolID)
	return o
}

// SetGatewayPoolID adds the gatewayPoolId to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) SetGatewayPoolID(gatewayPoolID string) {
	o.GatewayPoolID = gatewayPoolID
}

// WithNetworkID adds the networkID to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateway pools gateway pool ID params
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewayPoolsGatewayPoolIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_pool_id
	if err := r.SetPathParam("gateway_pool_id", o.GatewayPoolID); err != nil {
		return err
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