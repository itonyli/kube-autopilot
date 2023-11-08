package main

import (
	"autopilot/cmd/proxy/app"
	"os"
)

func main() {
	command := app.NewProxyServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
