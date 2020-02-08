package models

import (
	"NBAScraperGo/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Game struct {
	HomeTeam *Team
	HomeTeamStats *TeamStats
	AwayTeam *Team
	AwayTeamStats *TeamStats
}

type TeamStats struct {
	Points string
	PlayersStats []*PlayerStats
}

func (game *Game) ExtractGame(gameId string, gameDate string) {
	boxScoreUrl := utils.BuildBoxScoreUrl(gameDate, gameId)
	resp, err := http.Get(boxScoreUrl)
	if err != nil {
		log.Fatalln(err)
	}

	boxScore := BoxScoreJson{}
	err = json.NewDecoder(resp.Body).Decode(&boxScore)
	if err != nil {
		log.Fatalln(err)
	}

	teams := make(map[string]*Team)
	fmt.Println(teams)

	game.HomeTeamStats.Points = boxScore.BasicGameData.HomeTeam.Score
	game.AwayTeamStats.Points = boxScore.BasicGameData.AwayTeam.Score
}