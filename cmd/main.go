package main

import (
	"fmt"

	"github.com/fujimaru-lab/FetchWeather/internal/logo"
	"github.com/fujimaru-lab/FetchWeather/pkg/console"
)

func main() {
	// プログラムのロゴを表示
	logo.PrintItalicLogo()

	// ユーザーの標準入力により都市名を取得
	reader := console.NewUserInputReader(">>Where is the city Do you want to know how the weather is ?:")
	reader.Prompt()
	cityName := reader.ReadUserInput()

	fmt.Println(cityName)
}
