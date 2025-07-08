-- Create SubCategory table
CREATE TABLE IF NOT EXISTS sub_categories (
    uuid UUID PRIMARY KEY,
    category_uuid UUID NOT NULL,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category
        FOREIGN KEY (category_uuid)
        REFERENCES categories(uuid)
        ON DELETE CASCADE
);
