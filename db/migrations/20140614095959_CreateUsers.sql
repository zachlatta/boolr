
-- +goose Up
CREATE TABLE users (
  id serial primary key,
  created timestamp not null,
  updated timestamp not null,
  username text not null unique,
  password text not null
);


-- +goose Down
DROP TABLE users;

