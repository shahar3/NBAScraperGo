package models

import (
	"NBAScraperGo/constants"
	"NBAScraperGo/utils"
	"NBAScraperGo/utils/logger"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
)

type Team struct {
	Logger *logger.Logger
	Name   string
	Url    string
	Roster map[string]*Player
	Games  []*Game
}

func (team *Team) GetRoster() {
	team.Logger.Write(fmt.Sprintf("Collecting the roster of %s", team.Name), logger.LogTypeDebug)
	teamUrl := fmt.Sprintf("%s%s", constants.BaseURL, team.Url)
	dom := utils.GetDocument(teamUrl)

	team.Roster = make(map[string]*Player)
	//extract the players from .nba-player-index class
	dom.Find(".nba-player-index__trending-item").Each(func(i int, s *goquery.Selection) {
		//Create player object and add to the roster
		player := &Player{
			Team: team.Name,
		}

		//extract jersey number
		jerseyNumSel := s.Find(".nba-player-trending-item__number")
		if jerseyNumSel.Nodes != nil {
			var err error
			player.JerseyNumber, err = strconv.Atoi(jerseyNumSel.Text())
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			fmt.Println("No jersey number")
		}

		//extract player name
		playerNameSel := s.Find("a").First()
		if playerNameSel.Nodes != nil {
			playerFullName, ok := playerNameSel.Attr("title")
			if ok {
				player.FullName = playerFullName
			}
		} else {
			fmt.Println("No player name")
		}

		// ".nba-player-index__details"
		s.Find(".nba-player-index__details span").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				//get position
				positions := utils.GetPosition(s.Text())
				player.Position = positions
			case 1:
				//get height and weight
				heightInCm, weightInKg := utils.GetHeightAndWeight(s.Text())
				player.Height = heightInCm
				player.Weight = weightInKg
			}
		})

		team.Roster[player.FullName] = player
	})
}