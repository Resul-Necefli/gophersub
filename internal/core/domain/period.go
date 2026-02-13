package domain

import (
	"errors"
	"time"
)

// SubscriptionPeriod represents the start and end dates of a subscription period.
type SubscriptionPeriod struct {
	start time.Time
	end   time.Time
}

// NewPeriod creates a new SubscriptionPeriod with the given start and end times.
// It validates that the end time is not before the start time and that the start time is not in the future.
// Returns an error if the validation fails.
func NewPeriod(start, end time.Time) (SubscriptionPeriod, error) {

	if end.Before(start) {

		return SubscriptionPeriod{}, errors.New("invalid subscription period: end date cannot be before start date")
	}

	if start.After(time.Now()) {
		return SubscriptionPeriod{}, errors.New("start date cannot be in the future")
	}

	return SubscriptionPeriod{
		start: start,
		end:   end,
	}, nil
}

// IsActive checks if the subscription period is currently active based on the provided time.
func (s SubscriptionPeriod) IsActive(now time.Time) bool {

	if (now.After(s.start) || now.Equal(s.start)) && now.Before(s.end) {

		return true
	}

	return false

}

// Extend extends the subscription period by the specified number of months from the current end date.
func (s SubscriptionPeriod) Extend(months int, now time.Time) SubscriptionPeriod {

	if s.end.After(now) {

		s.end = s.end.AddDate(0, months, 0)

	} else {
		s.end = now.AddDate(0, months, 0)
	}

	// return s
	return SubscriptionPeriod{
		start: s.start,
		end:   s.end,
	}

}
