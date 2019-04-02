// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AddressBookItem address book item
// swagger:model address_book_item
type AddressBookItem struct {

	// address
	// Required: true
	Address CurrencyAddress `json:"address"`

	// creation timestamp
	// Required: true
	CreationTimestamp Timestamp `json:"creation_timestamp"`

	// currency
	// Required: true
	Currency Currency `json:"currency"`

	// type
	Type AddressBookType `json:"type,omitempty"`
}

// Validate validates this address book item
func (m *AddressBookItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressBookItem) validateAddress(formats strfmt.Registry) error {

	if err := m.Address.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("address")
		}
		return err
	}

	return nil
}

func (m *AddressBookItem) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := m.CreationTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creation_timestamp")
		}
		return err
	}

	return nil
}

func (m *AddressBookItem) validateCurrency(formats strfmt.Registry) error {

	if err := m.Currency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("currency")
		}
		return err
	}

	return nil
}

func (m *AddressBookItem) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddressBookItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressBookItem) UnmarshalBinary(b []byte) error {
	var res AddressBookItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}