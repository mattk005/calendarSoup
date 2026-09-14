-- name: AddEvent :exec
INSERT INTO events (
  id,
  summary,
  event_start,
  event_end,
  description,
  location,
  user_id
)
VALUES( 
  :id,
  :summary,
  :event_start,
  :event_end,
  :description,
  :location,
  :user_id
  );

-- name: GetEventsByUser :many
SELECT *
  FROM events
  WHERE user_id = ?;

-- name: GetEventsByWFMUUID :many
SELECT *
  FROM events
  JOIN users
  ON user_id = users.id
  WHERE users.wfm_external_calendar = ?
;

