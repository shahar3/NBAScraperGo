package scraper

import (
	"NBAScraperGo/constants"
	"NBAScraperGo/models"
	"NBAScraperGo/utils"
	"NBAScraperGo/utils/logger"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Scraper struct {
	logger *logger.Logger
	Teams map[string]*models.Team
}

func (scraper *Scraper) Init() {
	scraper.logger.Write("Initiating the NBA Scraper", logger.LogTypeDebug)
	scraper.Teams = make(map[string]*models.Team)

	startingUrl := fmt.Sprintf("%s%s", constants.BaseURL, constants.TeamsEndPoint)
	dom := utils.GetDocument(startingUrl)

	dom.Find(".team__list a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if ok {
			scraper.Teams[s.Text()] = &models.Team{
				Logger: scraper.logger,
				Name: s.Text(),
				Url:  link,
			}
		}
	})

	scraper.logger.Write("Getting team rosters", logger.LogTypeHeader)
	for _, team := range scraper.Teams {
		team.GetRoster()
	}

	//Get the schedule
	scraper.getSchedule()
}

func (scraper *Scraper) getSchedule() {
	resp, err := http.Get(constants.FullSchedule)
	if err != nil {
		log.Fatalln(err)
	}

	sched := models.Schedule{}
	err = json.NewDecoder(resp.Body).Decode(&sched)
	if err != nil {
		log.Fatalln(err)
	}

	//iterate on schedule json and extract games
	for _, val := range sched.MonthSchedule {
		fmt.Println(val.Mscd.Name)
		for _, gameVal := range val.Mscd.Games {
			if gameVal.GameWeek == 0 {
				continue //skip pre season games
			}
			awayTeam := fmt.Sprintf("%s %s", gameVal.AwayTeam.City, gameVal.AwayTeam.Nickname)
			homeTeam := fmt.Sprintf("%s %s", gameVal.HomeTeam.City, gameVal.HomeTeam.Nickname)
			game := models.Game{
				HomeTeam: scraper.Teams[homeTeam],
				AwayTeam: scraper.Teams[awayTeam],
			}

			gameId := gameVal.GameId
			gameDate := gameVal.Date
			gameDate = strings.Replace(gameDate, "-", "", -1)

			game.ExtractGame(gameId, gameDate)
			fmt.Println("ok")
		}
	}
}