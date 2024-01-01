package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/mukmin-life/db"
	p "github.com/mukmin-life/modules/prayer_time"
)

type Esolat struct {
	PrayerTime []struct {
		Hijri   string `json:"hijri"`
		Date    string `json:"date"`
		Day     string `json:"day"`
		Imsak   string `json:"imsak"`
		Fajr    string `json:"fajr"`
		Syuruk  string `json:"syuruk"`
		Dhuhr   string `json:"dhuhr"`
		Asr     string `json:"asr"`
		Maghrib string `json:"maghrib"`
		Isha    string `json:"isha"`
	} `json:"prayerTime"`
	Zone string `json:"zone"`
}

// type Scraper interface {
// 	// db sql.DB
// 	scrape()
// 	insertPrayerTimeData(Esolat, *p.Queries)
// }

func parseTime(timeString string) time.Time {
	result, err := time.Parse("02-Jan-2006 15:04:05", timeString)
	if err != nil {
		fmt.Println(err)
		panic("Cannot parse time")
	}

	return result
}

func parseDate(timeString string) time.Time {
	result, err := time.Parse("02-Jan-2006", timeString)
	if err != nil {
		fmt.Println(err)
		panic("Cannot parse time")
	}

	return result
}


func constructUrl(zone string) string {
	url, err := url.Parse("https://www.e-solat.gov.my/index.php")
	url.RawQuery = "r=esolatApi/takwimsolat&period=duration&zone=" + zone
	if err != nil {
		fmt.Println("Failed to construct URL")
	}
	return url.String()
}

func fetchEsolatData(year int) Esolat {
	data := url.Values{
		"datestart": {"2024-01-01"},
		"dateend":   {"2024-01-02"},
	}
	url := constructUrl("WLY01")
	fmt.Println(url)
	res, err := http.PostForm(constructUrl("WLY01"), data)
	if err != nil {
		fmt.Println("failed request")
	}
	defer res.Body.Close()

	var esolatData Esolat
	err = json.NewDecoder(res.Body).Decode(&esolatData)
	if err != nil {
		fmt.Println("failed to decode")
	}

	return esolatData
}

func insertPrayerTimeData(esolatData Esolat, queries *p.Queries) {
	ctx := context.Background()
	for _, day := range esolatData.PrayerTime {
		fmt.Println(parseDate(day.Date))
		queries.InsertPrayerTime(ctx, p.InsertPrayerTimeParams{
			Zone:    esolatData.Zone,
			Hijri:   day.Hijri,
			Date:    parseDate(day.Date),
			Imsak:   parseTime(day.Date + " " + day.Imsak).Add(-8 * time.Hour),
			Fajr:    parseTime(day.Date + " " + day.Fajr).Add(-8 * time.Hour),
			Syuruk:  parseTime(day.Date + " " + day.Syuruk).Add(-8 * time.Hour),
			Dhuhr:   parseTime(day.Date + " " + day.Dhuhr).Add(-8 * time.Hour),
			Asr:     parseTime(day.Date + " " + day.Asr).Add(-8 * time.Hour),
			Maghrib: parseTime(day.Date + " " + day.Maghrib).Add(-8 * time.Hour),
			Isha:    parseTime(day.Date + " " + day.Isha).Add(-8 * time.Hour),
		})
	}
}

func main() {
	esolatData := fetchEsolatData(2024)
	err := db.Connect()
	if err != nil {
		fmt.Println("failed to connect to DB")
	}
	defer db.Connection.Close()
	db := db.Connection
	queries := p.New(db)
	insertPrayerTimeData(esolatData, queries)
	fmt.Println("Finish updating data")
}
