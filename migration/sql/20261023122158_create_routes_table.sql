-- +goose Up
CREATE TABLE IF NOT EXISTS routes (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX uq_routes_code ON routes(code);

-- +goose Down
DROP TABLE IF EXISTS routes;
