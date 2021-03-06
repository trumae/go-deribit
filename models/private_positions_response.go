// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivatePositionsResponse private positions response
// swagger:model private_positions_response
type PrivatePositionsResponse struct {

	// result
	// Required: true
	Result *PrivatePositionsResponseResult `json:"result"`
}

// Validate validates this private positions response
func (m *PrivatePositionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivatePositionsResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivatePositionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivatePositionsResponse) UnmarshalBinary(b []byte) error {
	var res PrivatePositionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivatePositionsResponseResult private positions response result
// swagger:model PrivatePositionsResponseResult
type PrivatePositionsResponseResult struct {

	// average price for the position
	// Required: true
	AveragePrice *float64 `json:"averagePrice"`

	// The base currency of the instrument
	// Required: true
	Currency *string `json:"currency"`

	// The position delta
	// Required: true
	Delta *float64 `json:"delta"`

	// The direction of the position. Can be "buy" (long) or "sell" (short)
	// Required: true
	Direction *string `json:"direction"`

	// Estimated liquidation price
	// Required: true
	EstLiqPrice *float64 `json:"estLiqPrice"`

	// floating PnL
	// Required: true
	FloatingPl *float64 `json:"floatingPl"`

	// index price
	// Required: true
	IndexPrice *float64 `json:"indexPrice"`

	// initial margin
	// Required: true
	InitialMargin *float64 `json:"initialMargin"`

	// name of the instrument
	// Required: true
	Instrument *string `json:"instrument"`

	// The type of instrument. "future" or "option"
	// Required: true
	Kind *string `json:"kind"`

	// maintenance margin
	// Required: true
	MaintenanceMargin *float64 `json:"maintenanceMargin"`

	// mark price
	// Required: true
	MarkPrice *float64 `json:"markPrice"`

	// The margin used to back the position
	// Required: true
	OpenOrderMargin *float64 `json:"openOrderMargin"`

	// The PnL for the position
	// Required: true
	ProfitLoss *float64 `json:"profitLoss"`

	// realized PnL
	// Required: true
	RealizedPl *float64 `json:"realizedPl"`

	// The settlement price for the instrument
	// Required: true
	SettlementPrice *float64 `json:"settlementPrice"`

	// The position size in contracts. Can be negative (short) or positive (long)
	// Required: true
	Size *float64 `json:"size"`

	// position size in BTC (0 if currency <>`BTC`)
	// Required: true
	SizeBtc *float64 `json:"sizeBtc"`

	// position size in the base currency for the instrument
	// Required: true
	SizeCurrency *float64 `json:"sizeCurrency"`
}

// Validate validates this private positions response result
func (m *PrivatePositionsResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAveragePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstLiqPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloatingPl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialMargin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceMargin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenOrderMargin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfitLoss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRealizedPl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeBtc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivatePositionsResponseResult) validateAveragePrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"averagePrice", "body", m.AveragePrice); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateDelta(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"delta", "body", m.Delta); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateEstLiqPrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"estLiqPrice", "body", m.EstLiqPrice); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateFloatingPl(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"floatingPl", "body", m.FloatingPl); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateIndexPrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"indexPrice", "body", m.IndexPrice); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateInitialMargin(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"initialMargin", "body", m.InitialMargin); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateInstrument(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"instrument", "body", m.Instrument); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateMaintenanceMargin(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"maintenanceMargin", "body", m.MaintenanceMargin); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateMarkPrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"markPrice", "body", m.MarkPrice); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateOpenOrderMargin(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"openOrderMargin", "body", m.OpenOrderMargin); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateProfitLoss(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"profitLoss", "body", m.ProfitLoss); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateRealizedPl(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"realizedPl", "body", m.RealizedPl); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateSettlementPrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"settlementPrice", "body", m.SettlementPrice); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateSizeBtc(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"sizeBtc", "body", m.SizeBtc); err != nil {
		return err
	}

	return nil
}

func (m *PrivatePositionsResponseResult) validateSizeCurrency(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"sizeCurrency", "body", m.SizeCurrency); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivatePositionsResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivatePositionsResponseResult) UnmarshalBinary(b []byte) error {
	var res PrivatePositionsResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
