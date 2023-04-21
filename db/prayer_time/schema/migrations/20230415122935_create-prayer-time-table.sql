-- migrate:up
CREATE TABLE prayer_times (
  date DATE NOT NULL,
  zone VARCHAR (255) NOT NULL,
  hijri VARCHAR (255) NOT NULL,
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
