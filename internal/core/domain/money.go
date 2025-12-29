package domain

import "errors"

// ...existing code...
// This code defines a Money type with an amount and currency.
// It includes error handling for invalid amounts and currencies.
// The NewMoney function creates a new Money instance, ensuring valid input.
// GetAmount and GetCurrency methods retrieve the amount and currency respectively.
// ...existing code...
type Money struct {
	amount   int64
	currency string
}

var (
	ErrInvalidAmount    = errors.New("money: invalid amount ")
	ErrInvalidcurrency  = errors.New("currency: invalid currency")
	ErrCurrencyMismatch = errors.New("money: currency mismatch")
)

func NewMoney(amount int64, currency string) (Money, error) {

	if amount < 0 {

		return Money{}, ErrInvalidAmount
	} else if currency != "azn" && currency != "" {
		return Money{}, ErrCurrencyMismatch
	} else if currency == "" {
		return Money{}, ErrInvalidAmount
	}

	return Money{amount: amount, currency: currency}, nil

}

func (m Money) GetAmount() int64 {
	return m.amount
}

func (m Money) GetCurrency() string {
	return m.currency
}
