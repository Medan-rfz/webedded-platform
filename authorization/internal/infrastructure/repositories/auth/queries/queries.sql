-- name: GetById :one
SELECT *
FROM "Users" u
WHERE u.user_id = $1;

-- name: GetByEmail :one
SELECT *
FROM "Users" u
WHERE u.email = $1;

-- name: Create :exec
INSERT INTO "Users" (user_id, email, password_hash)
VALUES ($1,$2,$3);

-- name: Remove :exec
DELETE FROM "Users" u
WHERE u.user_id = $1; 
