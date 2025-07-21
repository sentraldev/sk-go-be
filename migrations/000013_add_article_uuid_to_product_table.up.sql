ALTER TABLE products ADD COLUMN article_uuid UUID;
ALTER TABLE products
ADD CONSTRAINT fk_article
    FOREIGN KEY (article_uuid)
    REFERENCES articles(uuid);