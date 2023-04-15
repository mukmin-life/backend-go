package prayer_time

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary      List prayer times of the day
// @Description  get prayer time by date
// @Accept       json
// @Produce      json
// @Success      200  {object}  PrayerTime
// @Router       /prayer_time/ [get]
func GetPrayerTime(c *fiber.Ctx) error {
	prayer_times := []PrayerTime{
		{
			Date: time.Now(),
			Zone: "JHR01",
			Imsak: time.Now(),
			Fajr: time.Now(), 
			Dhuhr: time.Now(), 
			Asr: time.Now(), 
			Maghrib: time.Now(), 
			Isha: time.Now(),
		},
	}

	return c.JSON(prayer_times)
}
