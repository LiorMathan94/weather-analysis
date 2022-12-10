package main

import (
	"fmt"

	data "github.com/LiorMathan94/weather-analysis/data"
)

func main() {
	city := "Tel Aviv"
	count := 5

	fmt.Printf("\n\n=========== Weather Analysis for %s: ===========", city)

	printRainingChance(city, count)
	printNextRainDay(city)
	printAverageTemperature(city, count)
	printTemperatureRange(city, count)
	printWeatherSummery(city, 0)
	printMostWindyDay(city)
	printMostHumidDay(city)
	printDriestDay(city)
	printHottestDay(city)
	printColdestDay(city)

	fmt.Println("\n\n======================================================")
}

func printRainingChance(city string, count int) {
	willItRain := data.WillItRain(city, count)
	fmt.Printf("\n\n> Chance of rain in %s in the next %d days: %d%%", city, count, willItRain)
}

func printNextRainDay(city string) {
	nextRainDay := data.NextRainDay(city)
	if nextRainDay >= 0 {
		fmt.Printf("\n\n> Expect rainy weather in %d days", nextRainDay)
	} else {
		fmt.Printf("\n\n> Hooray! No rainy days expected soon!")
	}
}

func printAverageTemperature(city string, count int) {
	averageTemperature := data.AverageTemperature(city, count)
	fmt.Printf("\n\n> Average temperature in %s in the next %d days: %d°", city, count, averageTemperature)
}

func printTemperatureRange(city string, count int) {
	tempRange := data.TemperatureRange(city, count)
	fmt.Printf("\n\n> Temperature range in %s in the next %d days: %v", city, count, tempRange)
}

func printWeatherSummery(city string, day int) {
	todaysWeatherSummery := data.WeatherSummery(city, day)
	fmt.Printf("\n\n> Today's weather summery in %s - minimum temp: %d°, maximum temp: %d°, raining chance: %d%%, humidity: %d%%, wind speed: %dKMH", city, todaysWeatherSummery.LowTempValue, todaysWeatherSummery.HighTempValue, todaysWeatherSummery.RainningChance, todaysWeatherSummery.Humidity, todaysWeatherSummery.Wind)
}

func printMostWindyDay(city string) {
	windSpeed, windyDay := data.MostWindyDay(city)
	fmt.Printf("\n\n> The day with the strongest winds in %s will occur in %d days. The wind speed will be %d km/h.", city, windyDay, windSpeed)
}

func printMostHumidDay(city string) {
	humidity, humidDay := data.MostHumidDay(city)
	fmt.Printf("\n\n> The day with the highest humidity in %s will occur in %d days. The humidity will be %d%%.", city, humidDay, humidity)
}

func printDriestDay(city string) {
	humidity, dryDay := data.DriestDay(city)
	fmt.Printf("\n\n> The driest day (with the lowest humidity) in %s will occur in %d days. The humidity will be %d%%.", city, dryDay, humidity)
}

func printHottestDay(city string) {
	temperature, hottestDay := data.HottestDay(city)
	fmt.Printf("\n\n> The hottest day in %s will occur in %d days. The temperature will be %d°.", city, hottestDay, temperature)
}

func printColdestDay(city string) {
	temperature, coldestDay := data.ColdestDay(city)
	fmt.Printf("\n\n> The coldest day in %s will occur in %d days. The temperature will be %d°.", city, coldestDay, temperature)
}
