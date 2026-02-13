package db

import (
	"errors"
	"sync"

	"github.com/Resul-Necefli/gophersub/internal/core/domain"
)

type InMemorySubscriptionRepository struct {
	memory map[string]*domain.Subscription
	mutex  sync.RWMutex
}

var (
	ErrNotFound = errors.New("subscriber not found")
)

func NewInMemorySubscriptionRepository() *InMemorySubscriptionRepository {
	return &InMemorySubscriptionRepository{
		map[string]*domain.Subscription{},
		sync.RWMutex{},
	}
}

/*
Save, GetByID, GetByUserID


	Save(sub *domain.Subscription) error
	GetByID(id string) (*domain.Subscription, error)
	GetByUserID(userID string) ([]*domain.Subscription, error)

*/

func (i *InMemorySubscriptionRepository) Save(sub *domain.Subscription) error {
	//   burada men  adi map istfade edirem deye struct  daxiline  thred  tehlukesizliyini temin etmek ucun
	//  sync.mutex elave etdim  burada  yazma emeliyatlari ustunluk  ve ya beraberlik teskil etdiyi ucun men burada  herhasnsi
	//bir  sync.Map  istfade etmedim    mutex ile map  birlesdirdim   cunki sync map icerisinde ikdene map saxlayir ve buradaki
	// veziyyetde o menim ucun performansima   ziddir
	id := sub.GetByID()
	i.mutex.Lock()
	i.memory[id] = sub
	i.mutex.Unlock()
	return nil
}

func (i *InMemorySubscriptionRepository) GetByID(id string) (*domain.Subscription, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	if val, ok := i.memory[id]; ok {
		return val, nil

	}
	return nil, ErrNotFound

}

func (i *InMemorySubscriptionRepository) GetByUserID(userID string) ([]*domain.Subscription, error) {

	//  burada men  make ile yaratdim ki  sliceni orade appende herdefe yeni slice yaradip  GC  ni  bezdirmesin bos referanslarla
	// ikincisi oxuma emeliyyatlari ucun  mentiq onsuz butun  metod  boyu irellilediyi ucun  burada R.Unluck defer ile gosterdim
	// onsuzda burada spesfik  bir   deyiskende deysiklik olunsaydida mutex kimi agir bir kilidlenmeni  yox  atomic istfade edecekdim

	// var result []*domain.Subscription
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	var result = make([]*domain.Subscription, 0, 10)
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
