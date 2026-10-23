-- +goose Up
CREATE TABLE IF NOT EXISTS vehicle_assignments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    vehicle_id VARCHAR(10) NOT NULL,
    assigned_route_id BIGINT NULL REFERENCES routes(id),
    assigned_direction VARCHAR(10) NULL,
    current_route_stop_id BIGINT NULL REFERENCES route_stops(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX uq_vehicle_assignment_vehicle_id ON vehicle_assignments(vehicle_id);

-- +goose Down
DROP TABLE IF EXISTS vehicle_assignments;
