-- name: GetPrayerTime :one
SELECT * FROM prayer_times
WHERE date = $1 AND zone = $2 LIMIT 1;

-- name: InsertPrayerTime :exec

INSERT INTO prayer_times (date, zone, hijri, imsak, fajr, syuruk, dhuhr, asr, maghrib, isha) VALUES (@date, @zone, @hijri, @imsak, @fajr, @syuruk, @dhuhr, @asr, @maghrib, @isha);