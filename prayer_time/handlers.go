package prayer_time

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/mukmin-life/db"
)

type PrayerTimeOutput struct {
	Imsak   time.Time `json:"imsak"`
	Fajr    time.Time `json:"fajr"`
	Syuruk  time.Time `json:"syuruk"`
	Dhuhr   time.Time `json:"dhuhr"`
	Asr     time.Time `json:"asr"`
	Maghrib time.Time `json:"maghrib"`
	Isha    time.Time `json:"isha"`
}

// @Summary      List prayer times of the day
// @Description  get prayer time by date
// @Accept       json
// @Produce      json
// @Success      200  {object}  PrayerTimeOutput
// @Router       /prayer_time/ [get]
func GetPrayerTime(c *fiber.Ctx) error {

	ctx := context.Background()

	db :=  db.Connection

	queries := New(db)

	prayer_time, err := queries.GetPrayerTime(ctx, GetPrayerTimeParams{
		Date: time.Date(2023, 1, 1, 0,0,0,0,time.UTC), Zone: "JHR01",
	})
	if err != nil {
		return err
	}

	output := PrayerTimeOutput{
		Imsak: prayer_time.Imsak,
		Fajr: prayer_time.Fajr,
		Syuruk: prayer_time.Syuruk,
		Dhuhr: prayer_time.Dhuhr,
		Asr: prayer_time.Asr,
		Maghrib: prayer_time.Maghrib,
		Isha: prayer_time.Isha,
	}

	c.Set("Cache-Control", "public, max-age=300, s-maxage=600")
	return c.JSON(output)
}
