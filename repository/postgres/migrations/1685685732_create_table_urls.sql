-- +migrate Up
CREATE TABLE urls (
  id SERIAL PRIMARY KEY,
  key TEXT NOT NULL UNIQUE,
  url TEXT NOT NULL
);

-- +migrate Down 
Drop table urls;
