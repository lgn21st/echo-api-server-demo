-- +migrate Up
CREATE TABLE users (
  id integer NOT NULL,
  email character varying NOT NULL,
  encrypted_password character varying NOT NULL
);

-- +migrate Down
DROP TABLE users;
