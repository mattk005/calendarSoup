-- name: CreateUser :one
INSERT INTO users (
  name,
  wfm_id,
  wfm_external_calendar
) VALUES ( 
  ?,
  ?,
  ?
)RETURNING *;

-- name: GetUsers :many
SELECT *
  FROM users;

-- name: GetUserBywfmID :one
SELECT * 
  FROM users
  WHERE wfm_external_calendar = ?;
