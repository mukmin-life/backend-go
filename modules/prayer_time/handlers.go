package prayer_time

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/mukmin-life/db"
)

type PrayerTimeOutput struct {
	Hijri   string	 `json:"hijri"`
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
// @Param date path string true "Date"
// @Param zone path string true "Zone"
// @Accept       json
// @Produce      json
// @Success      200  {object}  PrayerTimeOutput
// @Router       /prayer_time/{zone}/{date} [get]
func GetPrayerTime(c *fiber.Ctx) error {

	date, err := time.Parse("2006-01-02", c.Params("date"))
	if err != nil {
		return err
	}

	zone := c.Params("zone")

	ctx := context.Background()

	db :=  db.Connection
	queries := New(db)
	prayer_time, err := queries.GetPrayerTime(ctx, GetPrayerTimeParams{
		Date: date, Zone: zone,
	})
	if err != nil {
		return err
	}

	output := PrayerTimeOutput{
		Hijri: prayer_time.Hijri,
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
