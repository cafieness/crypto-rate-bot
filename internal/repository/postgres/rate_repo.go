package postgres

import (
	"cryptobot/internal/domain"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type RateRepository struct {
	db *sql.DB
}

func NewRateRepository(db *sql.DB) *RateRepository {
	return &RateRepository{
		db: db,
	}
}

func (r *RateRepository) Save(
	currency string,
	price float64,
) error {
	_, err := r.db.Exec(
		`
		INSERT INTO rates(currency, price)
		VALUES($1, $2)
		`,
		currency,
		price,
	)
	return err
}

func (r *RateRepository) GetDailyMinMax(currency string) (float64, float64, error) {
	var min, max float64
	err := r.db.QueryRow(
		`
		SELECT min(price), max(price)
		from rates
		where currency=$1
		and fetched_at >= now() - interval '1 day';
		`,
		currency,
	).Scan(
		&min,
		&max,
	)
	return min, max, err
}

func (r *RateRepository) GetHourlyChange(currency string) (float64, error) {
	var old float64
	var current domain.Rate
	current, err := r.GetLatest(currency)
	if err != nil {
		return 0, fmt.Errorf("no data for latest rate: %w", err)
	}
	err = r.db.QueryRow(
		`
		SELECT price
		from rates
		where currency=$1
		and fetched_at <= now() - interval '1 hour'
		ORDER BY fetched_at DESC
		limit 1;
		`,
		currency,
	).Scan(
		&old,
	)
	if err != nil {
		return 0, err
	}
	if old == 0 {
		return 0, fmt.Errorf("old price is zero")
	}

	change := (current.Price - old) / old * 100
	return change, nil
}

func (r *RateRepository) GetLatest(currency string) (domain.Rate, error) {
	var rate domain.Rate
	err := r.db.QueryRow(
		`
		SELECT currency, price
		from rates
		where currency=$1
		order by fetched_at DESC
		limit 1;
		`,
		currency,
	).Scan(
		&rate.Currency,
		&rate.Price,
	)
	return rate, err
}
