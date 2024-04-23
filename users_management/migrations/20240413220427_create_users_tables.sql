-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Users" (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    "group" VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS "Credentials" (
    user_id BIGINT PRIMARY KEY REFERENCES "Users"(id),
    password_hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "Roles" (
    id BIGSERIAL PRIMARY KEY,
    role_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "UserRoles" (
    user_id BIGINT REFERENCES "Users"(id),
    role_id BIGINT REFERENCES "Roles"(id),
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE IF NOT EXISTS "JwtRefresh" (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES "Users"(id),
    refresh_token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Users" CASCADE;
DROP TABLE IF EXISTS "Credentials" CASCADE;
DROP TABLE IF EXISTS "Roles" CASCADE;
DROP TABLE IF EXISTS "UserRoles" CASCADE;
DROP TABLE IF EXISTS "JwtRefresh" CASCADE;
-- +goose StatementEnd
