--sql
-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, user_id, feed_id)
    VALUES (
    $1,
    $2,
    $3
)
    RETURNING *
)
SELECT
    inserted_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
JOIN feeds ON inserted_feed_follows.feed_id = feeds.id
JOIN users ON inserted_feed_follows.user_id = users.id;


-- sql
-- name: GetFeedFollowsForUser :many
SELECT
    ff.*,
    feeds.name AS feed_name,
    users.name AS user_name
    FROM feed_follows AS ff
    JOIN feeds ON ff.feed_id = feeds.id
    JOIN users ON ff.user_id = users.id
    WHERE ff.user_id = $1
    ORDER BY ff.created_at DESC;

-- sql
-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: Unfollow :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;