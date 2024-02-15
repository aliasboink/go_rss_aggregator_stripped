-- +goose Up
CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  last_fetched_at TIMESTAMP,
  name TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE feeds;