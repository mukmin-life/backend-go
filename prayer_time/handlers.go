package prayer_time

import	"github.com/gofiber/fiber/v2"

// @Summary      List prayer times of the day
// @Description  get prayer time by date
// @Accept       json
// @Produce      json
// @Success      200  {object}  Author
// @Router       /prayer_time/ [get]
func GetPrayerTime(c *fiber.Ctx) error {
	prayer_times := []Author{{ID: 1, Name: "test"}}

	return c.JSON(prayer_times)
}
