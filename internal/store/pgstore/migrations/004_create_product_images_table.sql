CREATE TABLE IF NOT EXISTS product_images (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    path TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS product_images; 
