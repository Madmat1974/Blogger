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

