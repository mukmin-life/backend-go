-- migrate:up
CREATE TABLE prayer_times (
  date DATE PRIMARY KEY,
  zone VARCHAR (255) PRIMARY KEY,
  imsak timestamp NOT NULL,
  fajr timestamp NOT NULL,
  syuruk timestamp NOT NULL,
  dhuhr timestamp NOT NULL,
  asr timestamp NOT NULL,
  maghrib timestamp NOT NULL,
  isha timestamp NOT NULL
)