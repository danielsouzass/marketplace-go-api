CREATE TABLE IF NOT EXISTS product_payment_methods (
    product_id UUID NOT NULL,
    payment_method_id UUID NOT NULL,
    PRIMARY KEY (product_id, payment_method_id),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    CONSTRAINT fk_payment FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS product_payment_methods; 
