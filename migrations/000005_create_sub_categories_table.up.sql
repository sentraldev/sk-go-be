-- Create SubCategory table
CREATE TABLE IF NOT EXISTS sub_categories (
    uuid UUID PRIMARY KEY,
    category_uuid UUID NOT NULL,
    name VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT fk_category
        FOREIGN KEY (category_uuid)
        REFERENCES categories(uuid)
        ON DELETE CASCADE
);
