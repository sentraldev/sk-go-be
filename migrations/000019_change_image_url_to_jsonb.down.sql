ALTER TABLE products
    ALTER COLUMN image_urls TYPE TEXT USING image_urls::TEXT,
    ALTER COLUMN image_urls SET DEFAULT ''::TEXT;
