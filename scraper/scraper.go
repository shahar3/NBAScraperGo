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
	Logger *logger.Logger
	Teams  map[string]*models.Team
}

//Init is responsible of the main logic behind the scraper
func (scraper *Scraper) Init() {
	scraper.Logger.Write("Initiating the NBA Scraper", logger.LogTypeDebug)
	scraper.Teams = make(map[string]*models.Team)

	scraper.getTeams()

	scraper.Logger.Write("Getting team rosters", logger.LogTypeHeader)
	for _, team := range scraper.Teams {
		team.GetRoster()
	}

	//Get the schedule
	scraper.getSchedule()
}

func (scraper *Scraper) getTeams() {
	startingUrl := fmt.Sprintf("%s%s", constants.BaseURL, constants.TeamsEndPoint)
	dom := utils.GetDocument(startingUrl)

	dom.Find(".team__list a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if ok {
			scraper.Teams[s.Text()] = &models.Team{
				Logger: scraper.Logger,
				Name:   s.Text(),
				Url:    link,
			}
		}
	})
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
		scraper.Logger.Write(fmt.Sprintf("Get the schedule of %s", val.Mscd.Name), logger.LogTypeHeader)
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
			//add game to both home and away teams
			scraper.Teams[homeTeam].Games = append(scraper.Teams[homeTeam].Games, &game)
			scraper.Teams[awayTeam].Games = append(scraper.Teams[awayTeam].Games, &game)

			scraper.Logger.Write(fmt.Sprintf("Added the game: %s (%s) - %s (%s)", homeTeam, game.HomeTeamStats.Points, awayTeam, game.AwayTeamStats.Points), logger.LogTypeDebug)
		}
	}
}
