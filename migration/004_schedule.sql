-- +goose Up
CREATE TABLE schedule (
    id SERIAL PRIMARY KEY,
    fk_company_id INT NOT NULL,
    fk_route_id INT NOT NULL,
    fk_vehicle_id INT NOT NULL,
    start_at TIMESTAMP NOT NULL,
    repeats_daily BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE schedule;