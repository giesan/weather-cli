package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	key := os.Getenv("WEATHERAPI_KEY")
	if key == "" {
		fmt.Println("The environment variable WEATHERAPI_KEY is not set!\nSet this variable and try again.\n\nFor more informations read https://www.weatherapi.com/docs.")
		os.Exit(1)
	}

	var location *string
	location = flag.String("l", "", "Location for which the weather is to be shown")
	flag.Parse()

	if *location == "" {
		fmt.Println("The location is not set!")
		os.Exit(0)
	}

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + *location)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 && res.StatusCode <= 410 {
		printAPIError(res)
		os.Exit(0)
	}

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	printWeather(res)

}
