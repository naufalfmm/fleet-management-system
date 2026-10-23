CREATE TABLE IF NOT EXISTS vehicle_locations (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    vehicle_id VARCHAR(10) NOT NULL,
    latitude DECIMAL(7,4) NOT NULL,
    longitude DECIMAL(7,4) NOT NULL,
    timestamp BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS routes (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

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

CREATE TABLE IF NOT EXISTS vehicle_assignments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    vehicle_id VARCHAR(10) NOT NULL,
    assigned_route_id BIGINT NULL REFERENCES routes(id),
    assigned_direction VARCHAR(10) NULL,
    current_route_stop_id BIGINT NULL REFERENCES route_stops(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

