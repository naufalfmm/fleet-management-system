CREATE TABLE IF NOT EXISTS vehicle_locations (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    vehicle_id VARCHAR(10) NOT NULL,
    latitude DECIMAL(7,4) NOT NULL,
    longitude DECIMAL(7,4) NOT NULL,
    timestamp BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);