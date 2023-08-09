package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Condition struct {
	Text string `json:"text"`
}

type Current struct {
	TempC     float64 `json:"temp_c"`
	Condition `json:"condition"`
}

type Weather struct {
	Location `json:"location"`
	Current  `json:"current"`
}

func printWeather(res *http.Response) {

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current := weather.Location, weather.Current

	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)
}
