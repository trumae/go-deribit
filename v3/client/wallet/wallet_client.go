// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new wallet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wallet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPrivateAddToAddressBook adds new entry to address book of given type
*/
func (a *Client) GetPrivateAddToAddressBook(params *GetPrivateAddToAddressBookParams) (*GetPrivateAddToAddressBookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateAddToAddressBookParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateAddToAddressBook",
		Method:             "GET",
		PathPattern:        "/private/add_to_address_book",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateAddToAddressBookReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateAddToAddressBookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateAddToAddressBook: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateCancelWithdrawal cancels withdrawal request
*/
func (a *Client) GetPrivateCancelWithdrawal(params *GetPrivateCancelWithdrawalParams) (*GetPrivateCancelWithdrawalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateCancelWithdrawalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateCancelWithdrawal",
		Method:             "GET",
		PathPattern:        "/private/cancel_withdrawal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateCancelWithdrawalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateCancelWithdrawalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateCancelWithdrawal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateCreateDepositAddress creates deposit address in currency
*/
func (a *Client) GetPrivateCreateDepositAddress(params *GetPrivateCreateDepositAddressParams) (*GetPrivateCreateDepositAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateCreateDepositAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateCreateDepositAddress",
		Method:             "GET",
		PathPattern:        "/private/create_deposit_address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateCreateDepositAddressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateCreateDepositAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateCreateDepositAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateGetAddressBook retrieves address book of given type
*/
func (a *Client) GetPrivateGetAddressBook(params *GetPrivateGetAddressBookParams) (*GetPrivateGetAddressBookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateGetAddressBookParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateGetAddressBook",
		Method:             "GET",
		PathPattern:        "/private/get_address_book",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateGetAddressBookReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateGetAddressBookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateGetAddressBook: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateGetCurrentDepositAddress retrieves deposit address for currency
*/
func (a *Client) GetPrivateGetCurrentDepositAddress(params *GetPrivateGetCurrentDepositAddressParams) (*GetPrivateGetCurrentDepositAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateGetCurrentDepositAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateGetCurrentDepositAddress",
		Method:             "GET",
		PathPattern:        "/private/get_current_deposit_address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateGetCurrentDepositAddressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateGetCurrentDepositAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateGetCurrentDepositAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateGetDeposits retrieves the latest users deposits
*/
func (a *Client) GetPrivateGetDeposits(params *GetPrivateGetDepositsParams) (*GetPrivateGetDepositsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateGetDepositsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateGetDeposits",
		Method:             "GET",
		PathPattern:        "/private/get_deposits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateGetDepositsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateGetDepositsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateGetDeposits: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateGetTransfers adds new entry to address book of given type
*/
func (a *Client) GetPrivateGetTransfers(params *GetPrivateGetTransfersParams) (*GetPrivateGetTransfersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateGetTransfersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateGetTransfers",
		Method:             "GET",
		PathPattern:        "/private/get_transfers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateGetTransfersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateGetTransfersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateGetTransfers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateGetWithdrawals retrieves the latest users withdrawals
*/
func (a *Client) GetPrivateGetWithdrawals(params *GetPrivateGetWithdrawalsParams) (*GetPrivateGetWithdrawalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateGetWithdrawalsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateGetWithdrawals",
		Method:             "GET",
		PathPattern:        "/private/get_withdrawals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateGetWithdrawalsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateGetWithdrawalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateGetWithdrawals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateRemoveFromAddressBook adds new entry to address book of given type
*/
func (a *Client) GetPrivateRemoveFromAddressBook(params *GetPrivateRemoveFromAddressBookParams) (*GetPrivateRemoveFromAddressBookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateRemoveFromAddressBookParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateRemoveFromAddressBook",
		Method:             "GET",
		PathPattern:        "/private/remove_from_address_book",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateRemoveFromAddressBookReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateRemoveFromAddressBookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateRemoveFromAddressBook: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateToggleDepositAddressCreation enables or disable deposit address creation
*/
func (a *Client) GetPrivateToggleDepositAddressCreation(params *GetPrivateToggleDepositAddressCreationParams) (*GetPrivateToggleDepositAddressCreationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateToggleDepositAddressCreationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateToggleDepositAddressCreation",
		Method:             "GET",
		PathPattern:        "/private/toggle_deposit_address_creation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateToggleDepositAddressCreationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateToggleDepositAddressCreationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateToggleDepositAddressCreation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPrivateWithdraw creates a new withdrawal request
*/
func (a *Client) GetPrivateWithdraw(params *GetPrivateWithdrawParams) (*GetPrivateWithdrawOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateWithdrawParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateWithdraw",
		Method:             "GET",
		PathPattern:        "/private/withdraw",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateWithdrawReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateWithdrawOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateWithdraw: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
