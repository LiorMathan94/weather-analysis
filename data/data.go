package data

import (
	"math"

	"github.com/LiorMathan94/weather-analysis/scrapers"
	the_weather_channel_scraper "github.com/LiorMathan94/weather-analysis/scrapers/the_weather_channel"
	time_and_date_scraper "github.com/LiorMathan94/weather-analysis/scrapers/time_and_date"
)

type TempRange struct {
	Min int
	Max int
}

const numWebsites = 2

func WillItRain(city string, count int) int {
	weatherChannelData := the_weather_channel_scraper.Scrape(city)
	timeAndDateData := time_and_date_scraper.Scrape(city)

	rainningChanceSum := 0
	for i := 0; i < count; i++ {
		rainningChanceSum += weatherChannelData[i].RainningChance
		rainningChanceSum += timeAndDateData[i].RainningChance
	}

	return int(rainningChanceSum / (count * numWebsites))
}

func NextRainDay(city string) int {
	weatherChannelData := the_weather_channel_scraper.Scrape(city)
	timeAndDateData := time_and_date_scraper.Scrape(city)

	for i := 0; i < len(weatherChannelData); i++ {
		if weatherChannelData[i].RainningChance >= 50 || timeAndDateData[i].RainningChance >= 50 {
			return i
		}
	}

	return -1
}

func AverageTemperature(city string, count int) int {
	weatherChannelData := the_weather_channel_scraper.Scrape(city)
	timeAndDateData := time_and_date_scraper.Scrape(city)

	TemperatureSum := 0
	for i := 0; i < count; i++ {
		TemperatureSum += (weatherChannelData[i].LowTempValue + weatherChannelData[i].HighTempValue) / 2
		TemperatureSum += (timeAndDateData[i].LowTempValue + timeAndDateData[i].HighTempValue) / 2
	}

	return int(TemperatureSum / (count * numWebsites))
}

func TemperatureRange(city string, count int) map[int]TempRange {
	weatherChannelData := the_weather_channel_scraper.Scrape(city)
	timeAndDateData := time_and_date_scraper.Scrape(city)

	var tempRange = make(map[int]TempRange)
	for i := 0; i < count; i++ {
		minimumTemperatureAvg := int((weatherChannelData[i].LowTempValue + timeAndDateData[i].LowTempValue) / numWebsites)
		maximumTemperatureAvg := int((weatherChannelData[i].HighTempValue + timeAndDateData[i].HighTempValue) / numWebsites)
		tempRange[i] = TempRange{Min: minimumTemperatureAvg, Max: maximumTemperatureAvg}
	}

	return tempRange
}

func WeatherSummery(city string, day int) scrapers.WeatherData {
	weatherChannelData := the_weather_channel_scraper.Scrape(city)
	timeAndDateData := time_and_date_scraper.Scrape(city)

	var representative = scrapers.WeatherData{}
	representative.RainningChance = int((weatherChannelData[day].RainningChance + timeAndDateData[day].RainningChance) / numWebsites)
	representative.LowTempValue = int((weatherChannelData[day].LowTempValue + timeAndDateData[day].LowTempValue) / numWebsites)
	representative.HighTempValue = int((weatherChannelData[day].HighTempValue + timeAndDateData[day].HighTempValue) / numWebsites)
	representative.Humidity = int((weatherChannelData[day].Humidity + timeAndDateData[day].Humidity) / numWebsites)
	representative.Wind = int((weatherChannelData[day].Wind + timeAndDateData[day].Wind) / numWebsites)

	return representative
}

func MostWindyDay(city string) (int, int) {
	weatherChannelData := the_weather_channel_scraper.Scrape(city)
	timeAndDateData := time_and_date_scraper.Scrape(city)

	maxWindSpeed := 0
	maxWindSpeedIndex := 0

	iterations := int(math.Min(float64(len(weatherChannelData)), float64(len(timeAndDateData))))

	for i := 0; i < iterations; i++ {
		averageWindSpeed := (weatherChannelData[i].Wind + timeAndDateData[i].Wind) / numWebsites
		if averageWindSpeed > maxWindSpeed {
			maxWindSpeed = averageWindSpeed
			maxWindSpeedIndex = i
		}
	}

	return maxWindSpeed, maxWindSpeedIndex
}

func MostHumidDay(city string) (int, int) {
	weatherChannelData := the_weather_channel_scraper.Scrape(city)
	timeAndDateData := time_and_date_scraper.Scrape(city)

	maxHumidity := 0
	maxHumidityIndex := 0

	iterations := int(math.Min(float64(len(weatherChannelData)), float64(len(timeAndDateData))))
	
	for i := 0; i < iterations; i++ {
		averageHumidity := (weatherChannelData[i].Humidity + timeAndDateData[i].Humidity) / numWebsites
		if averageHumidity > maxHumidity {
			maxHumidity = averageHumidity
			maxHumidityIndex = i
		}
	}

	return maxHumidity, maxHumidityIndex
}
