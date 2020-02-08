package main

import (
	nbaScraper "NBAScraperGo/scraper"
	"NBAScraperGo/utils/logger"
)

func main() {
	scraper := nbaScraper.Scraper{
		Logger: &logger.Logger{
			Tag: "NBA Scraper 1.0",
		},
	}
	scraper.Init()
}
