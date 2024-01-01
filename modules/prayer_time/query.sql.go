// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package prayer_time

import (
	"context"
	"time"
)

const getPrayerTime = `-- name: GetPrayerTime :one
SELECT date, zone, hijri, imsak, fajr, syuruk, dhuhr, asr, maghrib, isha FROM prayer_times
WHERE date = $1 AND zone = $2 LIMIT 1
`

type GetPrayerTimeParams struct {
	Date time.Time
	Zone string
}

func (q *Queries) GetPrayerTime(ctx context.Context, arg GetPrayerTimeParams) (PrayerTime, error) {
	row := q.db.QueryRowContext(ctx, getPrayerTime, arg.Date, arg.Zone)
	var i PrayerTime
	err := row.Scan(
		&i.Date,
		&i.Zone,
		&i.Hijri,
		&i.Imsak,
		&i.Fajr,
		&i.Syuruk,
		&i.Dhuhr,
		&i.Asr,
		&i.Maghrib,
		&i.Isha,
	)
	return i, err
}

const insertPrayerTime = `-- name: InsertPrayerTime :exec

INSERT INTO prayer_times (date, zone, hijri, imsak, fajr, syuruk, dhuhr, asr, maghrib, isha) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type InsertPrayerTimeParams struct {
	Date    time.Time
	Zone    string
	Hijri   string
	Imsak   time.Time
	Fajr    time.Time
	Syuruk  time.Time
	Dhuhr   time.Time
	Asr     time.Time
	Maghrib time.Time
	Isha    time.Time
}

func (q *Queries) InsertPrayerTime(ctx context.Context, arg InsertPrayerTimeParams) error {
	_, err := q.db.ExecContext(ctx, insertPrayerTime,
		arg.Date,
		arg.Zone,
		arg.Hijri,
		arg.Imsak,
		arg.Fajr,
		arg.Syuruk,
		arg.Dhuhr,
		arg.Asr,
		arg.Maghrib,
		arg.Isha,
	)
	return err
}
