-- +goose Up
CREATE TABLE IF NOT EXISTS tasks(
  id              UUID    PRIMARY KEY,
  title           TEXT    NOT NULL,
  description     TEXT    NOT NULL,
  created_date    DATE    NOT NULL,
  completed_date  DATE,
  deleted_date    DATE
);

-- +goose Down
DROP TABLE IF EXISTS tasks;
