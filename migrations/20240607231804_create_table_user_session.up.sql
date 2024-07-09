CREATE TABLE IF NOT EXISTS sessions (
    id BIGINT PRIMARY KEY,
    token TEXT,
    expires_at BIGINT,
    ip_address TEXT,
    user_agent TEXT,
    user_id BIGINT,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
);