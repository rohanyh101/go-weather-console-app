package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC float64 `json:"temp_c"`

		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				Time_epoch int64   `json:"time_epoch"`
				TempC      float64 `json:"temp_c"`

				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`

				ChanceOfRain float64 `json:"chance_of_rain"`
			} `'json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: %w", err)
	}

	WEATHER_API := os.Getenv("WEATHER_API")
	if WEATHER_API == "" {
		log.Fatal("API KEY NOT PRESENT IN .env...")
	}

	city := "Jamkhandi"
	if len(os.Args) >= 2 {
		city = os.Args[1]
	}

	apiRequestURL := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=np&alerts=no", WEATHER_API, city)
	resp, err := http.Get(apiRequestURL)
	if err != nil {
		log.Fatalf("Error in API URL..., %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Error while fetching data: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error while reading body: %v", err)
	}

	// fmt.Println(string(body))

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("Error while unmarshalling... %v", err)
	}

	// fmt.Println(weather)

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Country,
		location.Name,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.Time_epoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		if hour.ChanceOfRain >= 40 {
			color.Red(
				"%s - %.0fC, %.0f, %s\n",
				date.Format("15:04"),
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text,
			)
			continue
		}

		color.Cyan(
			"%s - %.0fC, %.0f, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
	}

	// fmt.Println()
}
