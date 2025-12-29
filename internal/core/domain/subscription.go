package domain

import (
	"errors"
	"time"
)

// Subscription represents a user's subscription with details like plan, price, period, and status.
type Subscription struct {
	id       string
	userID   string
	planName string
	price    Money
	period   SubscriptionPeriod
	status   Status
}

// NewSubscription creates a new Subscription instance with the provided details.
// It validates that the userID is not empty.
// Returns an error if the userID is empty.
func NewSubscription(id, userID, planName string, price Money, period SubscriptionPeriod) (*Subscription, error) {

	if userID == "" {
		return &Subscription{}, errors.New("userID cannot be empty")
	}
	return &Subscription{
		id:       id,
		userID:   userID,
		planName: planName,
		price:    price,
		period:   period,
		status:   Status{value: active},
	}, nil
}

// IsExpired checks if the subscription has expired based on the current time.
func (s *Subscription) IsExpired(now time.Time) bool {

	return !s.period.IsActive(now)
}

// Canceled cancels the subscription if it is not already expired.
// Returns an error if the subscription is already expired.
func (s *Subscription) Canceled(now time.Time) error {

	if s.IsExpired(now) {

		return errors.New("cannot cancel expired subscription")
	}

	s.status = Status{value: "canceled"}
	return nil

}

// Renew extends the subscription period by the specified number of months from the current end date.
// It updates the subscription's period accordingly.
func (s *Subscription) Renew(months int, now time.Time) {

	subPeriod := s.period.Extend(months, now)

	s.period = subPeriod

}
