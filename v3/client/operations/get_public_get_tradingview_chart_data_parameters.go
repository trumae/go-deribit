// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPublicGetTradingviewChartDataParams creates a new GetPublicGetTradingviewChartDataParams object
// with the default values initialized.
func NewGetPublicGetTradingviewChartDataParams() *GetPublicGetTradingviewChartDataParams {
	var ()
	return &GetPublicGetTradingviewChartDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicGetTradingviewChartDataParamsWithTimeout creates a new GetPublicGetTradingviewChartDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicGetTradingviewChartDataParamsWithTimeout(timeout time.Duration) *GetPublicGetTradingviewChartDataParams {
	var ()
	return &GetPublicGetTradingviewChartDataParams{

		timeout: timeout,
	}
}

// NewGetPublicGetTradingviewChartDataParamsWithContext creates a new GetPublicGetTradingviewChartDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicGetTradingviewChartDataParamsWithContext(ctx context.Context) *GetPublicGetTradingviewChartDataParams {
	var ()
	return &GetPublicGetTradingviewChartDataParams{

		Context: ctx,
	}
}

// NewGetPublicGetTradingviewChartDataParamsWithHTTPClient creates a new GetPublicGetTradingviewChartDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicGetTradingviewChartDataParamsWithHTTPClient(client *http.Client) *GetPublicGetTradingviewChartDataParams {
	var ()
	return &GetPublicGetTradingviewChartDataParams{
		HTTPClient: client,
	}
}

/*GetPublicGetTradingviewChartDataParams contains all the parameters to send to the API endpoint
for the get public get tradingview chart data operation typically these are written to a http.Request
*/
type GetPublicGetTradingviewChartDataParams struct {

	/*EndTimestamp
	  The most recent timestamp to return result for

	*/
	EndTimestamp int64
	/*InstrumentName
	  Instrument name

	*/
	InstrumentName string
	/*Resolution
	  Chart bars resolution given in full minutes or keyword `1D`

	*/
	Resolution string
	/*StartTimestamp
	  The earliest timestamp to return result for

	*/
	StartTimestamp int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) WithTimeout(timeout time.Duration) *GetPublicGetTradingviewChartDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) WithContext(ctx context.Context) *GetPublicGetTradingviewChartDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) WithHTTPClient(client *http.Client) *GetPublicGetTradingviewChartDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTimestamp adds the endTimestamp to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) WithEndTimestamp(endTimestamp int64) *GetPublicGetTradingviewChartDataParams {
	o.SetEndTimestamp(endTimestamp)
	return o
}

// SetEndTimestamp adds the endTimestamp to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) SetEndTimestamp(endTimestamp int64) {
	o.EndTimestamp = endTimestamp
}

// WithInstrumentName adds the instrumentName to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) WithInstrumentName(instrumentName string) *GetPublicGetTradingviewChartDataParams {
	o.SetInstrumentName(instrumentName)
	return o
}

// SetInstrumentName adds the instrumentName to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) SetInstrumentName(instrumentName string) {
	o.InstrumentName = instrumentName
}

// WithResolution adds the resolution to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) WithResolution(resolution string) *GetPublicGetTradingviewChartDataParams {
	o.SetResolution(resolution)
	return o
}

// SetResolution adds the resolution to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) SetResolution(resolution string) {
	o.Resolution = resolution
}

// WithStartTimestamp adds the startTimestamp to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) WithStartTimestamp(startTimestamp int64) *GetPublicGetTradingviewChartDataParams {
	o.SetStartTimestamp(startTimestamp)
	return o
}

// SetStartTimestamp adds the startTimestamp to the get public get tradingview chart data params
func (o *GetPublicGetTradingviewChartDataParams) SetStartTimestamp(startTimestamp int64) {
	o.StartTimestamp = startTimestamp
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicGetTradingviewChartDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param end_timestamp
	qrEndTimestamp := o.EndTimestamp
	qEndTimestamp := swag.FormatInt64(qrEndTimestamp)
	if qEndTimestamp != "" {
		if err := r.SetQueryParam("end_timestamp", qEndTimestamp); err != nil {
			return err
		}
	}

	// query param instrument_name
	qrInstrumentName := o.InstrumentName
	qInstrumentName := qrInstrumentName
	if qInstrumentName != "" {
		if err := r.SetQueryParam("instrument_name", qInstrumentName); err != nil {
			return err
		}
	}

	// query param resolution
	qrResolution := o.Resolution
	qResolution := qrResolution
	if qResolution != "" {
		if err := r.SetQueryParam("resolution", qResolution); err != nil {
			return err
		}
	}

	// query param start_timestamp
	qrStartTimestamp := o.StartTimestamp
	qStartTimestamp := swag.FormatInt64(qrStartTimestamp)
	if qStartTimestamp != "" {
		if err := r.SetQueryParam("start_timestamp", qStartTimestamp); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}