package postgres

import (
	"context"
	"cryptobot/internal/domain"
	"database/sql"
)

type SubscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

func (s *SubscriptionRepository) Create(
	ctx context.Context,
	sub domain.Subscription,
) error {

	_, err := s.db.ExecContext(
		ctx,
		`
		insert into subscriptions(
			chat_id,
			interval_minutes,
			currency,
			is_active
		)
		values($1,$2, $3, true)
		on conflict(chat_id, currency)
		do update set
			interval_minutes = excluded.interval_minutes,
			is_active = true;
		`,
		sub.ChatID,
		sub.IntervalMinutes,
		sub.Currency,
	)
	return err

}
func (s *SubscriptionRepository) Deactivate(
	ctx context.Context,
	chatID int64,
) error {
	_, err := s.db.ExecContext(
		ctx,
		`
		update subscriptions
		set is_active = false
		where chat_id = $1;
		`,
		chatID,
	)
	return err
}
func (s *SubscriptionRepository) GetActive(ctx context.Context) (
	[]domain.Subscription,
	error) {
	rows, err := s.db.QueryContext(
		ctx,
		`
		select chat_id, interval_minutes, currency, is_active, last_sent_at
		from subscriptions
		where is_active = true
		`,
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			return
		}
	}()

	var subscriptions []domain.Subscription

	for rows.Next() {
		var sub domain.Subscription

		var lastSentAt sql.NullTime
		err := rows.Scan(
			&sub.ChatID,
			&sub.IntervalMinutes,
			&sub.Currency,
			&sub.Active,
			&lastSentAt,
		)
		if err != nil {
			return nil, err
		}
		if lastSentAt.Valid {
			sub.LastSentAt = &lastSentAt.Time
		}

		subscriptions = append(subscriptions, sub)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return subscriptions, nil

}

func (s *SubscriptionRepository) UpdateLastSent(
	ctx context.Context,
	chatID int64,
	currency string,
) error {

	_, err := s.db.ExecContext(
		ctx,
		`
		update subscriptions
		set last_sent_at = now()
		where chat_id = $1
		and currency = $2;
		`,
		chatID,
		currency,
	)

	return err
}

func (s *SubscriptionRepository) Exists(

	ctx context.Context,
	chatID int64,
	currency string,
) (bool, error) {

	var exists bool

	err := s.db.QueryRowContext(
		ctx,
		`
		SELECT EXISTS(
			SELECT 1
			FROM subscriptions
			WHERE chat_id=$1
			AND currency=$2
			AND is_active=true
		)
		`,
		chatID,
		currency,
	).Scan(&exists)

	return exists, err
}
