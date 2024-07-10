CREATE TABLE IF NOT EXISTS credentials (
    id VARCHAR(36) PRIMARY KEY,
    active BOOLEAN,
    hash TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    user_id BIGINT,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
);