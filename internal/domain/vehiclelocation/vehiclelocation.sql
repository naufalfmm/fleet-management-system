-- name: Create :one
INSERT INTO vehicle_locations(vehicle_id, latitude, longitude, timestamp)
VALUES ($1, $2, $3, $4)
RETURNING *;