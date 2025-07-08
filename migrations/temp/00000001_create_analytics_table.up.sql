-- +migrate Up
CREATE TABLE analytics (
    uuid uuid PRIMARY KEY,
    metric_id SERIAL,
    type VARCHAR(100) NOT NULL,
    value NUMERIC(12,2) NOT NULL,
    date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);
