-- +migrate Up
CREATE TABLE wishlist (
    uuid uuid PRIMARY KEY,
    user_uuid uuid NOT NULL REFERENCES users(uuid) ON DELETE CASCADE,
    product_uuid uuid NOT NULL REFERENCES products(uuid) ON DELETE CASCADE,
    UNIQUE (user_uuid, product_uuid),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);
