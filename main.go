package main

import (
	"fmt"

	data "github.com/LiorMathan94/weather-analysis/data"
)

func main() {
	willItRain := data.WillItRain("Tel Aviv", 5)
	fmt.Printf("Chance of rain: %d%%", willItRain)
}
