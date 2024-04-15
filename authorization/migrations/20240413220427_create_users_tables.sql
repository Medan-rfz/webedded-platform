-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Users" (
    email VARCHAR(50) PRIMARY KEY,
    user_id BIGINT,
    password_hash TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Users" CASCADE;
-- +goose StatementEnd
