CREATE TABLE tokens (
  id VARCHAR (36) PRIMARY KEY,
  user_id VARCHAR (36) PRIMARY KEY,
  token VARCHAR (64) NULL,
  expired_at timestamptz NOT NULL
);

CREATE INDEX tokens_id ON tokens (id);
CREATE INDEX tokens_user_id ON tokens (user_id);