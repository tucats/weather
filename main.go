package main

import (
	"fmt"
	"os"

	"github.com/tucats/gopackages/app-cli/app"
	"github.com/tucats/gopackages/app-cli/cli"
)

func main() {

	app := app.New("weather: view weather for a given location")
	app.SetVersion(1, 0, 0)
	app.SetCopyright("(C) Copyright Tom Cole 2020")

	err := app.Parse(WeatherGrammar, os.Args, WeatherAction)

	// If something went wrong, report it to the user and force an exit
	// status from the error, else a default General error.
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		if e2, ok := err.(cli.ExitError); ok {
			os.Exit(e2.ExitStatus)
		}
		os.Exit(1)
	}
}
