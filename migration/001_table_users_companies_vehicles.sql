-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50),
    phone VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    address TEXT,
    email VARCHAR(50),
    phone VARCHAR(20),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE vehicles (
    id SERIAL PRIMARY KEY,
    fk_owner_id INTEGER NOT NULL,
	fk_company_id INTEGER NOT NULL,
	name VARCHAR(20),
	temporary_name VARCHAR(20),
    license_number VARCHAR(20) UNIQUE,
    type INT NOT NULL,
    capacity INTEGER NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- +goose Down
DROP TABLE users;
DROP TABLE companies;
DROP TABLE vehicles;