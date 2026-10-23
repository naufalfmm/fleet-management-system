-- name: GetByID :one
SELECT *
FROM route_stops
WHERE id = $1;

-- name: GetRest :many
SELECT *
FROM route_stops
WHERE route_id = $1
    AND direction = $2
    AND stop_order > $3
ORDER BY stop_order ASC;