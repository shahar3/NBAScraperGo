package models

import "NBAScraperGo/utils"

type Player struct {
	FullName     string
	Team         string
	JerseyNumber int
	Position     []utils.Position
	Height       int
	Weight       int
}