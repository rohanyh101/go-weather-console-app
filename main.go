// Weather console application that fetches and displays weather data
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

// Weather struct maps JSON response from WeatherAPI
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
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	// Load API key from .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: %w", err)
	}

	WEATHER_API := os.Getenv("WEATHER_API")
	if WEATHER_API == "" {
		log.Fatal("API KEY NOT PRESENT IN .env...")
	}

	// Default city or use command line argument
	city := "Hubli"
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

	// Parse JSON response into Weather struct
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatalf("Error while unmarshalling... %v", err)
	}

	// Extract data for easier access
	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	// Display current weather
	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Country,
		location.Name,
		current.TempC,
		current.Condition.Text,
	)

	// Display hourly forecast with color coding
	for _, hour := range hours {
		date := time.Unix(hour.Time_epoch, 0)

		// Skip past hours
		if date.Before(time.Now()) {
			continue
		}

		// Red for high rain probability (≥40%), cyan for normal weather
		if hour.ChanceOfRain >= 40 {
			color.Red(
				"%s - %.0fC, %.0f%%, %s\n",
				date.Format("15:04"),
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text,
			)
			continue
		}

		color.Cyan(
			"%s - %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
	}
}
