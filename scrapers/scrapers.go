package scrapers

import (
	"fmt"
	"net/http"
)

type WeatherData struct {
	RainningChance int `json:"raining_chance"`
	LowTempValue   int `json:"low_temperature"`
	HighTempValue  int `json:"high_temperature"`
	Humidity       int `json:"humidity"`
	Wind           int `json:"wind"`
}

func Check(error error) {
	if error != nil {
		fmt.Println(error)
	}
}

func GetHtml(url string) *http.Response {
	response, error := http.Get(url)
	Check(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code: ", response.StatusCode)
	}

	return response
}
