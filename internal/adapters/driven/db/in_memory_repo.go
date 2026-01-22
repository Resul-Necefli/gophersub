package db

import (
	"errors"

	"github.com/Resul-Necefli/gophersub/internal/core/domain"
)

type InMemorySubscriptionRepository struct {
	memory map[string]*domain.Subscription
}

var (
	ErrNotFound = errors.New("subscriber not found")
)

func NewInMemorySubscriptionRepository() *InMemorySubscriptionRepository {
	return &InMemorySubscriptionRepository{
		map[string]*domain.Subscription{},
	}
}

/*
Save, GetByID, GetByUserID


	Save(sub *domain.Subscription) error
	GetByID(id string) (*domain.Subscription, error)
	GetByUserID(userID string) ([]*domain.Subscription, error)

*/

func (i *InMemorySubscriptionRepository) Save(sub *domain.Subscription) error {

	id := sub.GetByID()

	i.memory[id] = sub
	return nil
}

func (i *InMemorySubscriptionRepository) GetByID(id string) (*domain.Subscription, error) {

	if val, ok := i.memory[id]; ok {
		return val, nil
	}
	return nil, ErrNotFound
}

func (i *InMemorySubscriptionRepository) GetByUserID(userID string) ([]*domain.Subscription, error) {

	var result []*domain.Subscription
	for _, sub := range i.memory {
		if sub.GetByUserID() == userID {
			result = append(result, sub)
		}
	}
	if len(result) > 0 {
		return result, nil
	}

	return nil, nil
}
