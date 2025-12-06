-- +goose Up
CREATE TABLE counter (
    id SERIAL PRIMARY KEY,
    fk_company_id INT NOT NULL,
    fk_assign_user_id INT,
    name VARCHAR(50) NOT NULL,
    address VARCHAR(200),
    lat VARCHAR(50),
    long VARCHAR(50),
    note VARCHAR(200),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE route (
    id  SERIAL PRIMARY KEY,
    fk_company_id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    note VARCHAR(200),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);
CREATE TABLE route_counter (
    route_id INT NOT NULL,
    counter_id INT NOT NULL,
    fk_company_id INT NOT NULL,
    duration INT,
    cost DECIMAL,
    serial INT NOT NULL,
    PRIMARY KEY (route_id, counter_id)
);

-- +goose Down
DROP TABLE counter;
DROP TABLE route;
DROP TABLE route_counter;