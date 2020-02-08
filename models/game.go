package models

import (
	"NBAScraperGo/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Game struct {
	HomeTeam      *Team
	HomeTeamStats *TeamStats
	AwayTeam      *Team
	AwayTeamStats *TeamStats
}

type TeamStats struct {
	Points       string
	PlayersStats map[string]*PlayerStats
}

func (game *Game) ExtractGame(gameId string, gameDate string) {
	boxScoreUrl := utils.BuildBoxScoreUrl(gameDate, gameId) //Get the json url with the match details
	resp, err := http.Get(boxScoreUrl)
	if err != nil {
		log.Fatalln(err)
	}

	boxScore := BoxScoreJson{}
	err = json.NewDecoder(resp.Body).Decode(&boxScore)
	if err != nil {
		log.Fatalln(err)
	}

	teamsConversionMap := make(map[string]string)
	teamsConversionMap[boxScore.BasicGameData.HomeTeam.TeamId] = game.HomeTeam.Name
	teamsConversionMap[boxScore.BasicGameData.AwayTeam.TeamId] = game.AwayTeam.Name

	//initialize TeamStats struct
	game.HomeTeamStats = &TeamStats{
		PlayersStats: make(map[string]*PlayerStats),
	}
	game.AwayTeamStats = &TeamStats{
		PlayersStats: make(map[string]*PlayerStats),
	}

	game.HomeTeamStats.Points = boxScore.BasicGameData.HomeTeam.Score
	game.AwayTeamStats.Points = boxScore.BasicGameData.AwayTeam.Score

	//add player stats
	for _, stats := range boxScore.Stats.PlayersStats {
		team := teamsConversionMap[stats.TeamId]
		fullName := fmt.Sprintf("%s %s", stats.FirstName, stats.LastName)

		if game.HomeTeam.Name == team {
			if game.HomeTeam.Roster[fullName] != nil {
				game.addPlayerStats(fullName, stats, game.HomeTeamStats, game.HomeTeam)
			}
		} else if game.AwayTeam.Name == team {
			if game.AwayTeam.Roster[fullName] != nil {
				game.addPlayerStats(fullName, stats, game.AwayTeamStats, game.AwayTeam)
			}
		} else {
			fmt.Printf("Didnt find the team %s\n", team)
		}
	}
}

func (game *Game) addPlayerStats(playerName string, stats *PlayerStats, teamStats *TeamStats, team *Team) {
	teamStats.PlayersStats[playerName] = stats
	playerGame := &PlayerGame{
		HomeTeam:           game.HomeTeam,
		AwayTeam:           game.AwayTeam,
		HomeScore:          game.HomeTeamStats.Points,
		AwayScore:          game.AwayTeamStats.Points,
		PlayerStatsForGame: stats,
	}

	if team.Roster[playerName].Games == nil {
		team.Roster[playerName].Games = make([]*PlayerGame, 0)
	}

	team.Roster[playerName].Games = append(team.Roster[playerName].Games, playerGame)
}
