-- +goose Up
CREATE TABLE company_users (
    user_id INT PRIMARY KEY,
    company_id INT NOT NULL,
    role INT DEFAULT 1
);
CREATE TABLE company_sub_users (
    company_id INT NOT NULL,
    user_id INT NOT NULL,
    role INT NOT NULL,
    PRIMARY KEY (company_id, user_id, role)
);

-- +goose Down
DROP TABLE company_users;
DROP TABLE company_sub_users;