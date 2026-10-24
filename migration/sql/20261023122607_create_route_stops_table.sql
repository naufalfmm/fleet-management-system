-- +goose Up
CREATE TABLE IF NOT EXISTS route_stops (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    route_id BIGINT NOT NULL REFERENCES routes(id),
    direction VARCHAR(10) NOT NULL,
    name VARCHAR(255) NOT NULL,
    latitude DECIMAL(7,4) NOT NULL,
    longitude DECIMAL(7,4) NOT NULL,
    stop_order INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX uq_route_stops ON route_stops(route_id, direction, latitude, longitude);
CREATE INDEX IF NOT EXISTS idx_route_stops_route_direction ON route_stops(route_id, direction);

-- +goose Down
DROP TABLE IF EXISTS route_stops;
