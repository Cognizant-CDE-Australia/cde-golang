package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://samples.openweathermap.org/data/2.5/weather?q=London,uk&appid=b6907d289e10d714a6e88b30761fae22")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	weatherData := new(WeatherData)
	json.Unmarshal([]byte(body), &weatherData)

	fmt.Println("City ", weatherData.Name)
	fmt.Println("Coordinates Longitude ", weatherData.Coord.Lon)
	fmt.Println("Coordinates Latitude  ", weatherData.Coord.Lat)

	fmt.Println("Wind Speed information ", weatherData.Wind.Speed)
	fmt.Println("Wind Degree information ", weatherData.Wind.Deg)

	fmt.Println("Weather information")

	fmt.Println("Description 		", weatherData.Weather[0].Description)
	fmt.Println("Main information 	", weatherData.Weather[0].Main)
}

// WeatherData contains the weather information of the given coordinates
type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}
