package models

import "NBAScraperGo/utils"

type Player struct {
	FullName     string
	Team         string
	JerseyNumber int
	Position     []utils.Position
	Height       int
	Weight       int
	Games        []*PlayerGame
	Stats        PlayerStats
}

type PlayerGame struct {
	HomeTeam           *Team
	AwayTeam           *Team
	HomeScore          string
	AwayScore          string
	PlayerStatsForGame *PlayerStats
}
