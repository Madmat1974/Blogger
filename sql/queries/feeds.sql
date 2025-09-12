-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, now(), now(), $2, $3, $4)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name AS feeds_name, feeds.url AS feeds_url, users.name AS user_name
FROM feeds
INNER JOIN
    users ON feeds.user_id = users.id
ORDER BY feeds.id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id AS id, url AS url, name AS name, user_id AS user_id, created_at AS created_at, updated_at AS updated_at, last_fetched_at AS last_fetched_at
FROM feeds
ORDER BY last_fetched_at NULLS FIRST, id
LIMIT 1;