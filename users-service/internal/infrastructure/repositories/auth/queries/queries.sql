-- name: GetByEmail :one
SELECT u.id, u.email, cr.password_hash 
FROM "Users" u
JOIN "Credentials" cr ON u.id = cr.user_id
WHERE u.email = $1;

-- name: CreateUser :one
INSERT INTO "Users" 
(email, first_name, last_name, "group")
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: InsertCredentials :exec
INSERT INTO "Credentials" (user_id, password_hash)
VALUES ($1, $2);

-- name: GetAuthById :one
SELECT u.id, u.email, cr.password_hash 
FROM "Users" u
JOIN "Credentials" cr ON u.id = cr.user_id
WHERE u.id = $1;

-- name: GetAuthByEmail :one
SELECT u.id, u.email, cr.password_hash 
FROM "Users" u
JOIN "Credentials" cr ON u.id = cr.user_id
WHERE u.email = $1;

-- name: GetRefreshToken :one
SELECT r.id, r.refresh_token, r.expires_at
FROM "JwtRefresh" r
WHERE r.refresh_token = $1;

-- name: AddRefreshToken :exec
INSERT INTO "JwtRefresh" 
(user_id, refresh_token, expires_at)
VALUES ($1, $2, $3);
