
-- name: CreateImage :one
INSERT INTO images (
  photo
) VALUES (
  $1
)
RETURNING *;

-- name: GetRanDomImage :one
SELECT *
FROM images
ORDER BY random()
LIMIT 1;
