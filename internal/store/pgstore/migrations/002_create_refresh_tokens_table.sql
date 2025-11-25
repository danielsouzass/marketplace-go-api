CREATE TABLE IF NOT EXISTS refresh_tokens (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    expires_in INT NOT NULL,
    user_id UUID UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS refresh_tokens; 
