CREATE TABLE IF NOT EXISTS records (
  date DATE PRIMARY KEY NOT NULL,
  executed TIMESTAMP NOT NULL,
  info JSON NOT NULL
)
