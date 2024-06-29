-- +goose Up
ALTER TABLE users
Add is_active bool
NOT NULL
DEFAULT true;

-- +goose Down
ALTER TABLE users
DROP COLUMN is_active;