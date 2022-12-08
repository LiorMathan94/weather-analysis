package data

import (
	the_weather_channel_scraper "github.com/LiorMathan94/weather-analysis/scrapers/the_weather_channel"
)

func WillItRain(city string, count int) int {
	weatherChannelData := the_weather_channel_scraper.Scrape("Tel Aviv", 5)

	rainningChanceSum := 0
	for i := 0; i < count; i++ {
		rainningChanceSum += weatherChannelData[i].RainningChance
	}

	return int(rainningChanceSum / count)
}
