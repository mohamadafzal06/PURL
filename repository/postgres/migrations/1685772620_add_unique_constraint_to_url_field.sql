-- +migrate Up
ALTER TABLE urls ADD UNIQUE (url);
