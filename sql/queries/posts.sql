-- name: CreatePost :one
WITH 
  inserted_post AS (
    INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *
  )

SELECT inserted_post.*;

-- name: GetPostsForUser :many
SELECT feed_follows.feed_id AS feed_id, posts.id AS post_id, posts.title AS post_title, posts.url AS post_url, posts.published_at AS post_published_at
FROM feed_follows
INNER JOIN posts ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;