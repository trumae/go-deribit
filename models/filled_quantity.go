// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// FilledQuantity The number of contracts to be traded.
// swagger:model filled_quantity
type FilledQuantity float64

// Validate validates this filled quantity
func (m FilledQuantity) Validate(formats strfmt.Registry) error {
	return nil
}