-- name: GetById :one
SELECT *
FROM "Users" u
WHERE u.id = $1;