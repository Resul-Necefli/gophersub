package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	// Postgres driver-ini yükləyirik
	_ "github.com/lib/pq"

	"github.com/Resul-Necefli/gophersub/internal/core/domain"
)

// PostgresRepository - Bizim Postgres adapterimiz
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository - Bazaya qoşulmaq üçün konstruktor
func NewPostgresRepository(connStr string) (*PostgresRepository, error) {
	// men  burada connection pool  yaradcam   ve  db baglantilarina ev sahibliyi edecek
	//  o hecbir  baglanti  acmir sadece   hovuzu  yaradip  dayanir
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	//  burada men  databse   baglanti  atip onu test edirem  sonra herhansi bir problemi  varsa  onu  qaytariram  yoxdursa  ise  repository-ni qaytariram
	if err = db.Ping(); err != nil {

		return nil, err
	}

	return &PostgresRepository{db: db}, nil

}

func (p *PostgresRepository) GetPlanByName(name string) (*domain.Plan, error) {

	// ilkin olaraq men  abuneliyin novlerini saxladigim cedvele sorgu atacam permium ve ya basic    haqinda melumati cixarmag ucun
	// sorgumuzu yazaq
	query := `Select name,price_amount,price_currency,duration_days
    From plans
	WHERE name =$1`
	// 2 ci olaraq men  gelen datalari  deyiskenlere menimsedecem
	var planName string
	var priceAmount int64
	var priceCurrency string
	var durationDays int

	// burada artiq sorgunu ise salacam  ve databse  yollayacam
	// burada bizim databazaya go ile sorgu atmagimizin bir nece yolu var
	//QueryRow -   sorgudan tapdigi birinci setri tapir ve mene verir menim sorgu atdigim deyerler her setirde yalniz birdene oldugu ucn
	// yeni unikal oldugu ucun problem yoxdur
	//  Scan  gelend deyerleri  qabagcadan yaratdigim deyiskenlere  menimsedir   daha sonra men o deyerleri istfade ederek melumati oture bilirem

	err := p.db.QueryRow(query, name).Scan(&planName, &priceAmount, &priceCurrency, &durationDays)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("plan not found")
		}

		return nil, fmt.Errorf("database error %w", err)

	}

	// nurada domain qaydalrini qorumaq ucun  money value objectini  unexported etmisem onun ucun onu ayrica oz konsturctoru ile
	// hazir bir obyekt kimi hazirlayip verecem planin icine
	//DDD qaydalrina esasen hecbir xarici api  adapter onun value objectinin  fieldlarina cixis ederek deyisiklik ede bilmez
	moneyVO, err := domain.NewMoney(priceAmount, priceCurrency)
	if err != nil {
		return nil, fmt.Errorf("invalid money data in database: %w", err)
	}

	return &domain.Plan{
		Name:     planName,
		Price:    moneyVO,
		Duration: durationDays,
	}, nil

}

func (p *PostgresRepository) Save(sub *domain.Subscription) error {

	// ilkin olaraq upsert sorgusu yazacam
	// burada men  eger id artiq movcuddursa onun yalniz  abunelik novunu  ve tarixini yenileyecem

	query := `
	INSERT INTO  subscriptions (id , user_id,plan_name,start_date,end_date,price_amount,price_currency,status)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT(id)  DO UPDATE
	SET status = EXCLUDED.status,
	 end_date = EXCLUDED.end_date;
	`

	_, err := p.db.Exec(query,
		sub.ID(),
		sub.UserID(),
		sub.PlanName(),
		sub.PeriodStart(),
		sub.PeriodEnd(),
		sub.PriceAmount(),
		sub.PriceCurrency(),
	)

	if err != nil {
		// Əgər nəsə səhv getsə (məsələn, baza çöksə), xətanı qaytarırıq
		return fmt.Errorf("failed to save subscription to database: %w", err)
	}

	return nil

}

func (p *PostgresRepository) GetByID(id string) (*domain.Subscription, error) {

	query := `
		SELECT id, user_id, plan_name, start_date, end_date, price_amount, price_currency, status
		FROM subscriptions
		WHERE id = $1
	`

	var dbID string
	var userID string
	var planName string
	var startDate time.Time
	var endDate time.Time
	var priceAmount int64
	var priceCurrency string
	var status string

	err := p.db.QueryRow(query, id).Scan(
		&dbID, &userID, &planName, &startDate, &endDate, &priceAmount, &priceCurrency, &status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("subscription not found")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	moneyVO, err := domain.NewMoney(priceAmount, priceCurrency)
	if err != nil {
		return nil, fmt.Errorf("invalid money data in db: %w", err)
	}

	periodVO, err := domain.NewPeriod(startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid period data in db: %w", err)
	}

	domainStatus, err := domain.NewStatus(status)
	if err != nil {
		return nil, fmt.Errorf("invalid status data in db: %w", err)
	}

	sub := domain.RestoreSubscription(dbID, userID, planName, periodVO, moneyVO, domainStatus)

	return sub, nil
}

func (p *PostgresRepository) GetByUserID(userID string) ([]*domain.Subscription, error) {
	query := `
		SELECT id, user_id, plan_name, start_date, end_date, price_amount, price_currency, status
		FROM subscriptions
		WHERE user_id = $1
		ORDER BY start_date DESC
	`

	rows, err := p.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query user subscriptions: %w", err)
	}

	defer rows.Close()

	var subscriptions []*domain.Subscription

	for rows.Next() {
		var dbID, dbUserID, planName, priceCurrency, statusStr string
		var priceAmount int64
		var startDate, endDate time.Time

		err := rows.Scan(
			&dbID, &dbUserID, &planName, &startDate, &endDate, &priceAmount, &priceCurrency, &statusStr,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan subscription row: %w", err)
		}

		moneyVO, err := domain.NewMoney(priceAmount, priceCurrency)
		if err != nil {
			return nil, fmt.Errorf("invalid money data in db: %w", err)
		}

		periodVO, err := domain.NewPeriod(startDate, endDate)
		if err != nil {
			return nil, fmt.Errorf("invalid period data in db: %w", err)
		}

		domainStatus, err := domain.NewStatus(statusStr)
		if err != nil {
			return nil, fmt.Errorf("invalid status data in db: %w", err)
		}

		sub := domain.RestoreSubscription(dbID, dbUserID, planName, periodVO, moneyVO, domainStatus)
		subscriptions = append(subscriptions, sub)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return subscriptions, nil
}
