CREATE TABLE rates (
    id serial primary key,
    currency varchar(20),
    price numeric(18,8),
    fetched_at TIMESTAMPTZ default now()
);

CREATE index idx_rates_currency_time on rates(currency, fetched_at desc);

CREATE TABLE subscriptions (
    id serial PRIMARY KEY,
    chat_id bigint not null,
    interval_minutes int not null,
    currency varchar(20) not null,
    is_active boolean not null default true,
    created_at TIMESTAMP default now(),
    last_sent_at      TIMESTAMPTZ,

    UNIQUE(chat_id, currency)
);

CREATE INDEX idx_subscription_active
ON subscriptions(is_active);