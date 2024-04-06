-- +goose Up
-- +goose StatementBegin
CREATE TABLE token (
    id UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ,
    token VARCHAR,
    expires BIGINT,
    user_id UUID,
    PRIMARY KEY(id),
    CONSTRAINT fk_user_id
      FOREIGN KEY(user_id) 
	  REFERENCES user(id)
	  ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE token;
-- +goose StatementEnd
