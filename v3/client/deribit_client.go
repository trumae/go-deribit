// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/adampointer/go-deribit/v3/client/account_management"
	"github.com/adampointer/go-deribit/v3/client/internal"
	"github.com/adampointer/go-deribit/v3/client/market_data"
	"github.com/adampointer/go-deribit/v3/client/private"
	"github.com/adampointer/go-deribit/v3/client/public"
	"github.com/adampointer/go-deribit/v3/client/supporting"
	"github.com/adampointer/go-deribit/v3/client/trading"
	"github.com/adampointer/go-deribit/v3/client/wallet"
	"github.com/adampointer/go-deribit/v3/client/websocket_only"
)

// Default deribit HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:8082"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v2"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new deribit HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Deribit {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new deribit HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Deribit {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new deribit client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Deribit {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Deribit)
	cli.Transport = transport

	cli.AccountManagement = account_management.New(transport, formats)

	cli.Internal = internal.New(transport, formats)

	cli.MarketData = market_data.New(transport, formats)

	cli.Private = private.New(transport, formats)

	cli.Public = public.New(transport, formats)

	cli.Supporting = supporting.New(transport, formats)

	cli.Trading = trading.New(transport, formats)

	cli.Wallet = wallet.New(transport, formats)

	cli.WebsocketOnly = websocket_only.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Deribit is a client for deribit
type Deribit struct {
	AccountManagement *account_management.Client

	Internal *internal.Client

	MarketData *market_data.Client

	Private *private.Client

	Public *public.Client

	Supporting *supporting.Client

	Trading *trading.Client

	Wallet *wallet.Client

	WebsocketOnly *websocket_only.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Deribit) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.AccountManagement.SetTransport(transport)

	c.Internal.SetTransport(transport)

	c.MarketData.SetTransport(transport)

	c.Private.SetTransport(transport)

	c.Public.SetTransport(transport)

	c.Supporting.SetTransport(transport)

	c.Trading.SetTransport(transport)

	c.Wallet.SetTransport(transport)

	c.WebsocketOnly.SetTransport(transport)

}
