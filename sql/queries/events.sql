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

-- name: DeleteEventsByUser :exec
DELETE FROM events
  WHERE user_id = ?;


-- name: DeleteEventsByDateRange :exec
DELETE FROM events
WHERE event_start BETWEEN :min_date AND :max_date ;
