// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// InterestRate Interest rate used in implied volatility calculations (options only)
// swagger:model interest_rate
type InterestRate float64

// Validate validates this interest rate
func (m InterestRate) Validate(formats strfmt.Registry) error {
	return nil
}
