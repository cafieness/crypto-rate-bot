package postgres

import (
	"context"
	"cryptobot/internal/rate"
	"database/sql"
	"errors"
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
	ctx context.Context,
	currency string,
	price float64,
) error {
	_, err := r.db.ExecContext(
		ctx,
		`
		INSERT INTO rates(currency, price)
		VALUES($1, $2)
		`,
		currency,
		price,
	)
	return err
}

func (r *RateRepository) GetDailyMinMax(ctx context.Context,
	currency string) (float64, float64, error) {
	var minPrice, maxPrice sql.NullFloat64
	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT min(price), max(price)
		from rates
		where currency=$1
		and fetched_at >= now() - interval '1 day';
		`,
		currency,
	).Scan(
		&minPrice,
		&maxPrice,
	)
	if err != nil {
		return 0, 0, err
	}
	return minPrice.Float64, maxPrice.Float64, nil
}

func (r *RateRepository) GetHourlyChange(ctx context.Context, currency string) (float64, error) {
	var old float64
	var current rate.Rate
	current, err := r.GetLatest(ctx, currency)
	if err != nil {
		return 0, fmt.Errorf("no data for latest rate: %w", err)
	}
	err = r.db.QueryRowContext(
		ctx,
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
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	if old == 0 {
		return 0, fmt.Errorf("old price is zero")
	}

	change := (current.Price - old) / old * 100
	return change, nil
}

func (r *RateRepository) GetLatest(ctx context.Context, currency string) (rate.Rate, error) {
	var rate rate.Rate
	err := r.db.QueryRowContext(
		ctx,
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
