package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/fatih/color"
)

const apiKey = "21bb89bbc7a04aef82c53414242809"

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"` 
		}`json:"condition"`
	}`json:"current"`
}

func main() {
	if len(os.Args) <2 {
		fmt.Println("Usage: weather-cli [city]")
		return
	}

	city := os.Args[1]
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	color.Cyan("Weather in %s, %s", weather.Location.Name, weather.Location.Country)
	color.Green("Temperature: %.1f°C", weather.Current.TempC)
	color.Yellow("Condition: %s", weather.Current.Condition.Text)
}