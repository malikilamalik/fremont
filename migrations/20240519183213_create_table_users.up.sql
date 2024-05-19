CREATE TABLE users (
  id VARCHAR (26) PRIMARY KEY,
  name VARCHAR (50) NOT NULL,
  password VARCHAR (64) NULL,
  created_at timestamptz NOT NULL,
  deleted_at timestamptz NULL
);

CREATE INDEX user_id ON users (id);
CREATE INDEX user_deleted_at_null ON users (deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX user_deleted_at_not_null ON users (deleted_at) WHERE deleted_at IS NOT NULL;