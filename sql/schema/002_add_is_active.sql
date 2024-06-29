-- +goose Up
ALTER TABLE users
Add is_active bool
DEFAULT true;

-- +goose Down
ALTER TABLE users
DROP COLUMN is_active;