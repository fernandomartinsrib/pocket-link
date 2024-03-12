-- name: CreateUrl :one
INSERT INTO urls (
    long_url,
    short_url,
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUrl :one
SELECT * FROM urls
WHERE short_url = $1 LIMIT 1;