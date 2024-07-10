CREATE TABLE IF NOT EXISTS sessions (
    id VARCHAR(36) PRIMARY KEY,
    token TEXT,
    created_at TIMESTAMPTZ,
    ip_address TEXT,
    user_agent TEXT,
    user_id BIGINT,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
);