-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ,
    email VARCHAR,
    first_name VARCHAR,
    last_name VARCHAR,
    password VARCHAR,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
