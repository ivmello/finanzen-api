-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = ? LIMIT 1;

-- name: CreateTransaction :execresult
INSERT INTO transactions (
  id, amount, transaction_type, transaction_date, description, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);
