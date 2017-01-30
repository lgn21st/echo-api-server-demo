-- +migrate Up
CREATE UNIQUE INDEX users_email_idx ON users (email);

-- +migrate Down
DROP INDEX users_email_idx;
