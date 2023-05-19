
-- name: CreateImage :one
INSERT INTO images (
  photo
) VALUES (
  $1
)
RETURNING *;
