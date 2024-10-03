-- +goose Up
ALTER TABLE users
ADD COLUMN is_chirpy_red Boolean
DEFAULT 'false';

-- +goose Down
ALTER TABLE users
DROP COLUMN hashed_password;