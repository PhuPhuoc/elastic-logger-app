-- name: CreateAccount :execresult
INSERT INTO account (id, name, email, password, status)
VALUES (?, ?, ?, ?, ?);

-- name: GetAccountByEmail :one
SELECT id, name, email, password, status, created_at
FROM account
WHERE email = ? LIMIT 1;

