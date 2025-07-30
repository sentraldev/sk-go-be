ALTER TABLE products
ALTER COLUMN image_urls TYPE JSONB USING array_to_json(image_urls)::JSONB,
ALTER COLUMN image_urls SET DEFAULT '[]'::JSONB;