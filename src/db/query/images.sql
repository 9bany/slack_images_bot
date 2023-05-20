
-- name: CreateImage :one
INSERT INTO images (
  photo
) VALUES (
  $1
)
RETURNING *;

-- name: GetImage :one
SELECT *
FROM images
ORDER BY random()
LIMIT 1;
