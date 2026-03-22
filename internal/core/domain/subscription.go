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

// demeeli burada bele bir method domaine  elave edirem sebeb odur ki bizim servisimze data nece sualinin sizmasinin qarsini aliriq
// ve men   bunu  biznese  yegane giris qapisi olan agregate roota verecemki o idare etsin amma davranis eele statusn oz icndede olsun
// oz mesuliyyetini ozu  idare etsin
//  now  kimi time colden verirem sebeb odur ki  bu test zamani mene problem yaratmayacaq
// umumiyetle domain   icersinde bir datanin   indiki veziyyetini saxlamasi  duzgun deyil o data ile ne edeceyine qerar vermeldir datani
// colden almalidir  bura   atomun icidir sehv bir  davranis bizi partlada biler !
func (s *Subscription) IsActive(now time.Time) bool {

	if s.status.IsActive() && !s.period.IsActive(now) {
		return true
	}
	return false
}

//  burada  getter metodlar elave edirem ki  repository  ve ya servis terefinden  bu datalara
func (s *Subscription) UserID() string {
	return s.userID
}

func (s *Subscription) ID() string {
	return s.id
}

func (s *Subscription) PlanName() string {
	return s.planName
}

func (s *Subscription) PeriodStart() time.Time {
	return s.period.start
}

func (s *Subscription) PeriodEnd() time.Time {
	return s.period.end
}

func (s *Subscription) PriceAmount() int64 {

	return s.price.amount
}

func (s *Subscription) PriceCurrency() string {

	return s.price.currency
}

// sub.ID(),                  // $1: Abunəliyin unikal ID-si++++++++++++++
// 	sub.UserID(),              // $2: İstifadəçinin ID-si +++++++++++++++++++++++
// 	sub.PlanName(),            // $3: "premium" və ya "basic"++++++++++++++++++++++++++++++++
// 	sub.Period().Start,        // $4: Başlanğıc tarixi (Period VO-dan gəlir)
// 	sub.Period().End,          // $5: Bitiş tarixi (Period VO-dan gəlir)
// 	sub.Price().Amount,        // $6: Məbləğ (Money VO-dan gəlir)
// 	sub.Price().Currency,      // $7: Valyuta (Money VO-dan gəlir)
// 	sub.Status(),
