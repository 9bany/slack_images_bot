CREATE TABLE "images" (
  "id" bigserial PRIMARY KEY,
  "photo" BYTEA NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
)
