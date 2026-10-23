-- name: GetByVehicleID :one
SELECT *
FROM vehicle_assignments
WHERE vehicle_id = $1;

-- name: UpdateCurrentStop :one
UPDATE vehicle_assignments
SET current_route_stop_id = $1
WHERE id = $2
RETURNING *;