-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    iff.*,
    f.name AS feed_name,
    u.name AS user_name
FROM
    inserted_feed_follow AS iff
INNER JOIN users u ON iff.user_id = u.id
INNER JOIN feeds f ON iff.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT u.name AS user_name, f.name AS feed_name, iff.*
FROM users u
INNER JOIN feed_follows iff ON u.id=iff.user_id
INNER JOIN feeds f ON f.id=iff.feed_id
WHERE u.id = $1;
