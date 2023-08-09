package main

import (
	"flag"
	"fmt"

	"os"
)

func main() {

	ExecuteRootCmd()

	var location *string
	location = flag.String("l", "", "Location for which the weather is to be shown")
	flag.Parse()

	if *location == "" {
		fmt.Println("The location is not set!")
		os.Exit(0)
	}

}
