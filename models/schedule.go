package models

type Schedule struct {
	MonthSchedule []MonthSchedule `json:"lscd"`
}

type MonthSchedule struct {
	Mscd struct {
		Name  string     `json:"mon"`
		Games []GameJson `json:"g"`
	} `json:"mscd"`
}

type GameJson struct {
	GameId   string    `json:"gid"`
	Date     string    `json:"gdte"`
	AwayTeam TeamSched `json:"v"`
	HomeTeam TeamSched `json:"h"`
	GameWeek int       `json:"gweek"`
}

type TeamSched struct {
	City     string `json:"tc"`
	Nickname string `json:"tn"`
}