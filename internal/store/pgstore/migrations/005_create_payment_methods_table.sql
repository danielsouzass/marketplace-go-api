CREATE TABLE IF NOT EXISTS payment_methods (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    key TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS payment_methods; 
