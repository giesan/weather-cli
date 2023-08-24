package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	key := os.Getenv("WEATHERAPI_KEY")
	if key == "" {
		slog.Error("The environment variable WEATHERAPI_KEY is not set!\nSet this variable and try again.\n\nFor more informations read https://www.weatherapi.com/docs.")
		return
	}

	var location *string
	location = flag.String("l", "", "Location for which the weather is to be shown")
	flag.Parse()

	if *location == "" {
		slog.Error("The location is not set!")
		return
	}

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + *location)
	if err != nil {
		slog.Error(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 && res.StatusCode <= 410 {
		printAPIError(res)
		return
	}

	if res.StatusCode != 200 {
		slog.Error("Weather API not available")
	}

	printWeather(res)

}
