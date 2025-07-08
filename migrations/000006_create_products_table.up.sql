-- Create Product table
CREATE TABLE IF NOT EXISTS products (
    uuid UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    category_uuid UUID NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    description TEXT,
    image_urls TEXT[],
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category
        FOREIGN KEY (category_uuid)
        REFERENCES categories(uuid)
        ON DELETE CASCADE
);
