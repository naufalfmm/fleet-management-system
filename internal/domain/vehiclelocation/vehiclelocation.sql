-- name: Create :one
INSERT INTO vehicle_locations(vehicle_id, latitude, longitude, timestamp)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetLatest :one
SELECT *
FROM vehicle_locations
WHERE vehicle_id = $1
ORDER BY timestamp DESC
LIMIT 1;

-- name: GetHistory :many
SELECT *
FROM vehicle_locations
WHERE vehicle_id = $1
    AND (timestamp >= sqlc.arg(timestamp_from)::bigint AND timestamp <= sqlc.arg(timestamp_to)::bigint)
ORDER BY timestamp DESC;