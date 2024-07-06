-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    default_location VARCHAR(100),
    preferred_unit VARCHAR(1) CHECK (preferred_unit IN ('C', 'F')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;