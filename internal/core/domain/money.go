package domain

import "errors"

// Money represents a monetary value with an amount and currency.
type Money struct {
	amount   int64
	currency string
}

// Error variables for money-related issues.
var (
	ErrInvalidAmount    = errors.New("money: invalid amount ")
	ErrInvalidCurrency  = errors.New("currency: invalid currency")
	ErrCurrencyMismatch = errors.New("money: currency mismatch")
)

// NewMoney creates a new Money instance with the given amount and currency.
// It validates that the amount is non-negative and the currency is either "azn".
// Returns an error if the amount is negative or if the currency is invalid.
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

// GetAmount returns the amount of money.
func (m Money) GetAmount() int64 {
	return m.amount
}

// GetCurrency returns the currency of money.
func (m Money) GetCurrency() string {
	return m.currency
}
