-- name: GetPrayerTime :one
SELECT * FROM prayer_times
WHERE date = $1 AND zone = $2 LIMIT 1;
