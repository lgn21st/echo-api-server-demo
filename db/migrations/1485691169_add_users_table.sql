-- +migrate Up
CREATE TABLE users (
  id serial NOT NULL,
  email character varying NOT NULL,
  encrypted_password character varying NOT NULL,
  name character varying,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  deleted_at timestamp without time zone DEFAULT NULL
);

-- +migrate Down
DROP TABLE users;
