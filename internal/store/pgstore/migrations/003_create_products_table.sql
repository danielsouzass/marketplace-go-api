CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    is_new BOOLEAN NOT NULL,
    price NUMERIC NOT NULL,
    accept_trade BOOLEAN NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS products; 
