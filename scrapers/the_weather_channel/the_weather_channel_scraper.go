package the_weather_channel_scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/LiorMathan94/weather-analysis/scrapers"
	"github.com/PuerkitoBio/goquery"
)

var AllData = make(map[int]scrapers.WeatherData)

func initMap(doc *goquery.Document) {
	doc.Find("div.DetailsSummary--DetailsSummary--1DqhO").Each(func(i int, s *goquery.Selection) {
		AllData[i] = scrapers.WeatherData{}
	})
}

func scrapeRainingChance(doc *goquery.Document) {
	doc.Find("div.DetailsSummary--DetailsSummary--1DqhO>div.DetailsSummary--precip--1a98O").Each(func(i int, s *goquery.Selection) {
		span := s.Find("span")
		percentageString := span.Text()
		percentage, _ := strconv.Atoi(strings.Replace(percentageString, "%", "", -1))

		if entry, ok := AllData[i]; ok {
			entry.RainningChance = percentage
			AllData[i] = entry
		}
	})
}

func scrapeTemperature(doc *goquery.Document) {
	doc.Find("div.DetailsSummary--DetailsSummary--1DqhO>div.DetailsSummary--temperature--1kVVp").Each(func(i int, s *goquery.Selection) {
		highTempSpan := s.Find("span.DetailsSummary--highTempValue--3PjlX")
		highTempString := highTempSpan.Text()
		highTemp, _ := strconv.Atoi(strings.Replace(highTempString, "°", "", -1))

		if entry, ok := AllData[i]; ok {
			entry.HighTempValue = highTemp
			AllData[i] = entry
		}

		lowTempSpan := s.Find("span.DetailsSummary--lowTempValue--2tesQ")
		lowTempString := lowTempSpan.Text()
		lowTemp, _ := strconv.Atoi(strings.Replace(lowTempString, "°", "", -1))

		if entry, ok := AllData[i]; ok {
			entry.LowTempValue = lowTemp
			AllData[i] = entry
		}
	})
}

func scrapeHumidity(doc *goquery.Document) {
	doc.Find("div.DaypartDetails--DetailsTable--1zK4r>ul.DetailsTable--DetailsTable--3Bt2T").Each(func(i int, s *goquery.Selection) {
		s.Find("li.DetailsTable--listItem--Z-5Vi").Each(func(j int, li *goquery.Selection) {
			band, ok := li.Attr("data-testid")
			if ok {
				if band == "HumiditySection" {
					span := li.Find("span.DetailsTable--value--2YD0-")
					HumidityString := span.Text()
					humidity, _ := strconv.Atoi(strings.Replace(HumidityString, "%", "", -1))

					if entry, ok := AllData[i]; ok {
						entry.Humidity = humidity
						AllData[i] = entry
					}
				}
			}
		})
	})
}

func scrapeWind(doc *goquery.Document) {
	doc.Find("div.DetailsSummary--DetailsSummary--1DqhO>div.DetailsSummary--wind--1tv7t").Each(func(i int, s *goquery.Selection) {
		span := s.Find("span")
		windString := span.Text()
		re := regexp.MustCompile(`\d+`)
		match := re.FindStringSubmatch(windString)
		wind, _ := strconv.Atoi(match[0])

		if entry, ok := AllData[i]; ok {
			entry.Wind = wind
			AllData[i] = entry
		}
	})
}

func Scrape(city string, count int) map[int]scrapers.WeatherData {
	url := "https://weather.com/he-IL/weather/tenday/l/0bad9b388f3ea8468fd93c95fbfaf69c0e454a844692175e165c6b26a3cf8f62"

	response := scrapers.GetHtml(url)
	defer response.Body.Close()

	doc, error := goquery.NewDocumentFromReader(response.Body)
	scrapers.Check(error)

	initMap(doc)
	scrapeRainingChance(doc)
	scrapeTemperature(doc)
	scrapeHumidity(doc)
	scrapeWind(doc)
	fmt.Println(AllData)

	return AllData
}
