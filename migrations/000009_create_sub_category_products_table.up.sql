-- Create SubCategoryProduct join table
CREATE TABLE IF NOT EXISTS sub_category_products (
    product_uuid UUID NOT NULL,
    sub_category_uuid UUID NOT NULL,
    PRIMARY KEY (product_uuid, sub_category_uuid),
    CONSTRAINT fk_product
        FOREIGN KEY (product_uuid)
        REFERENCES products(uuid)
        ON DELETE CASCADE,
    CONSTRAINT fk_sub_category
        FOREIGN KEY (sub_category_uuid)
        REFERENCES sub_categories(uuid)
        ON DELETE CASCADE
);