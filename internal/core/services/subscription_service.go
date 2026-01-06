package services

import (
	"fmt"
	"log"
	"time"

	"github.com/Resul-Necefli/gophersub/internal/core/domain"
	"github.com/Resul-Necefli/gophersub/internal/core/ports"
)

type SubscriptionService struct {
	repo ports.SubscriptionRepository
}

func NewSubscriptionService(r ports.SubscriptionRepository) *SubscriptionService {

	return &SubscriptionService{repo: r}
}

func (s SubscriptionService) Subscribe(userID, planName string, amount int64, currency string) error {

	// birinci yoxlamaq ucun men  aggregate root  idare etdyi VO lardan  birine muraciet edecem bu
	// leqal  gorunur cunki servis  core domaini  gore biler  ve muracietde ede biler bu qaydalar uygun saylir
	// cunki biz daxile dogru importdan istfade edirik ve domain  tamamile tecrid edilip onun servis ve diger seylerden
	// xeberi bele  yoxdur
	subs, err := s.repo.GetByUserID(userID)
	if err != nil {
		log.Println("error fetching subscriptions:", err)
		return fmt.Errorf("error fetching subscriptions: %w", err)
	}

	// burada  geri qaytarilan siyahida subscribtion   domain obyektleri var onlar siyahi sekilnde qaytarilip sebebi ise
	// bir istfadecinin butun abunelik kecmisine baxa bilmeyimizdir  ve orada is active varmi deye yoxlamaq
	// bunu daha da optimallasdirmaq ucun sql terefinde  where   status = "active"  formasindada  yaza  bilerik
	// big O(n)  olacaq  performans  daha irelisi ucun  sql sorgusunu optimalsdira ve ya deyerler MAP da saxlaya bilerik
	now := time.Now()
	for _, sub := range subs {
		if sub.IsActive(now) {
			return nil
		}
	}

	// asagdaki koda men  VO  obyektlerini yartdim ki onlarla  bir yerde  domain obyektini yarada bilim men domaine  filedlarina
	//  uygun obyekt yaradip ona teqdim edirem yeni  burada men onun herhansi bir isine qarismiram
	money, err := domain.NewMoney(amount, currency)
	if err != nil {
		log.Println("error creating money value object:", err)
		return fmt.Errorf("error creating money value object: %w", err)
	}
	period, err := domain.NewPeriod(now, time.Now().AddDate(0, 1, 0))
	if err != nil {
		log.Println("error creating subscription:", err)
		return fmt.Errorf("error creating subscription: %w", err)
	}

	subscription, err := domain.NewSubscription("", userID, planName, money, period)
	if err != nil {

	}

	err = s.repo.Save(subscription)
	return err

}
