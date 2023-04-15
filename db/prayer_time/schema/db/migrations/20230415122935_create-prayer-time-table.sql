-- migrate:up
CREATE TABLE prayer_times (
  date DATE,
  zone VARCHAR (255),
  imsak timestamp NOT NULL,
  fajr timestamp NOT NULL,
  syuruk timestamp NOT NULL,
  dhuhr timestamp NOT NULL,
  asr timestamp NOT NULL,
  maghrib timestamp NOT NULL,
  isha timestamp NOT NULL,
  CONSTRAINT prayer_times_pk PRIMARY KEY (date, zone)
)

-- migrate:down
DROP TABLE prayer_times;
