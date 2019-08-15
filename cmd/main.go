package main

import (
	"fmt"

	"github.com/fujimaru-lab/FetchWeather/internal/logo"
	"github.com/fujimaru-lab/FetchWeather/pkg/console"
)

func main() {
	logo.PrintLogo()

	reader := console.NewUserInputReader(">>Where is the city Do you want to know how the weather is ?:")
	reader.Prompt()
	cityName := reader.ReadUserInput()

	fmt.Println(cityName)
}
