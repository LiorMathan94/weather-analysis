package time_and_date_scraper

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/LiorMathan94/weather-analysis/scrapers"
	"github.com/PuerkitoBio/goquery"
)

var AllData = make(map[int]scrapers.WeatherData)

const temperatureColumnIndex = 1
const windColumnIndex = 4
const humidityColumnIndex = 6
const rainingChanceColumnIndex = 7

func initMap(doc *goquery.Document) {
	doc.Find("table.zebra>tbody>tr").Each(func(i int, s *goquery.Selection) {
		AllData[i] = scrapers.WeatherData{}
	})
}

func scrapeRainingChance(doc *goquery.Document) {
	doc.Find("table.zebra>tbody>tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(j int, td *goquery.Selection) {
			if j == rainingChanceColumnIndex {
				percentageString := td.Text()
				percentage, _ := strconv.Atoi(strings.Replace(percentageString, "%", "", -1))

				if entry, ok := AllData[i]; ok {
					entry.RainningChance = percentage
					AllData[i] = entry
				}
			}
		})
	})
}

func scrapeTemperature(doc *goquery.Document) {
	doc.Find("table.zebra>tbody>tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(j int, td *goquery.Selection) {
			if j == temperatureColumnIndex {
				tempString := td.Text()

				re := regexp.MustCompile(`[0-9]+`)
				match := re.FindAllString(tempString, -1)
				minTemp, _ := strconv.Atoi(match[1])
				maxTemp, _ := strconv.Atoi(match[0])

				if entry, ok := AllData[i]; ok {
					entry.LowTempValue = minTemp
					entry.HighTempValue = maxTemp
					AllData[i] = entry
				}
			}
		})
	})
}

func scrapeHumidity(doc *goquery.Document) {
	doc.Find("table.zebra>tbody>tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(j int, td *goquery.Selection) {
			if j == humidityColumnIndex {
				humidityString := td.Text()
				humidity, _ := strconv.Atoi(strings.Replace(humidityString, "%", "", -1))

				if entry, ok := AllData[i]; ok {
					entry.Humidity = humidity
					AllData[i] = entry
				}
			}
		})
	})
}

func scrapeWind(doc *goquery.Document) {
	doc.Find("table.zebra>tbody>tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(j int, td *goquery.Selection) {
			if j == windColumnIndex {
				windString := td.Text()
				re := regexp.MustCompile(`\d+`)
				match := re.FindStringSubmatch(windString)
				wind, _ := strconv.Atoi(match[0])

				if entry, ok := AllData[i]; ok {
					entry.Wind = wind
					AllData[i] = entry
				}
			}
		})
	})
}

func Scrape(city string) map[int]scrapers.WeatherData {
	url := "https://www.timeanddate.com/weather/israel/tel-aviv/ext"

	response := scrapers.GetHtml(url)
	defer response.Body.Close()

	doc, error := goquery.NewDocumentFromReader(response.Body)
	scrapers.Check(error)

	initMap(doc)
	scrapeRainingChance(doc)
	scrapeTemperature(doc)
	scrapeHumidity(doc)
	scrapeWind(doc)

	return AllData
}
