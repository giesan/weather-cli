package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "weather",
	Long:  "weather",
	Run: func(cmd *cobra.Command, args []string) {

		key := os.Getenv("WEATHERAPI_KEY")
		if key == "" {
			fmt.Println("The environment variable WEATHERAPI_KEY is not set!\nSet this variable and try again.\n\nFor more informations read https://www.weatherapi.com/docs.")
			os.Exit(1)
		}

		location, _ := cmd.Flags().GetString("l")
		res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + location)
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

	},
}

func ExecuteRootCmd() {
	rootCmd.Flags().String("l", "Berlin", "Location for which the weather is to be shown")
	// rootCmd.MarkFlagRequired("l")
	rootCmd.Execute()
}
