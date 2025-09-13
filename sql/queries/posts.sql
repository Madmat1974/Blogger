-- name: CreatePost :one
INSERT INTO posts (id, feed_id, created_at, updated_at, title, url, description, published_at)
VALUES (
    $1,
    $2,
    now(),
    now(),
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetPosts :many
SELECT posts.id, posts.title, posts.url, posts.published_at, feeds.name AS feed_name
FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
JOIN feeds ON posts.feed_id = feeds.id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;