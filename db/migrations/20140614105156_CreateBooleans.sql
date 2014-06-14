
-- +goose Up
CREATE TABLE booleans (
  id serial primary key,
  created timestamp not null,
  updated timestamp not null,
  user_id integer references users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  label text not null,
  bool boolean not null,
  switch_count integer not null
);


-- +goose Down
DROP TABLE booleans;

