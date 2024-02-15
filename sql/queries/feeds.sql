-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE id = $1 LIMIT 1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedFetch :one
UPDATE feeds 
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM feeds
WHERE id = $1;