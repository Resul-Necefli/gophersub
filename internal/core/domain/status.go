package domain

import "errors"

const (
	active   string = "active"
	canceled string = "canceled"
	expired  string = "expired"
)

// Status represents the status of a subscription.
type Status struct {
	value string
}

// NewStatus creates a new Status instance with the given value.
// It validates that the provided value is one of the allowed statuses: "active", "canceled", or "expired".
// Returns an error if the validation fails.
func NewStatus(v string) (Status, error) {

	status := Status{value: v}
	if !status.IsValid() {
		return Status{}, errors.New("invalid status value")
	}

	return status, nil
}

// IsValid checks if the status value is one of the allowed statuses.
func (s Status) IsValid() bool {
	switch s.value {
	case active, canceled, expired:
		return true
	default:
		return false
	}
}

func (s Status) IsActive() bool {
	return s.value == active
}
