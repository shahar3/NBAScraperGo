package models

type BoxScoreJson struct {
	BasicGameData struct {
		AwayTeam TeamBoxScore `json:"vTeam"`
		HomeTeam TeamBoxScore `json:"hTeam"`
	} `json:"basicGameData"`
	Stats Stats `json:"stats"`
}

type TeamBoxScore struct {
	Win   string `json:"win"`
	Loss  string `json:"loss"`
	Score string `json:"score"`
	TeamId string `json:"teamId"`
}

type Stats struct {
	PlayersStats []*PlayerStats `json:"activePlayers"`
}

type PlayerStats struct {
	TeamId string `json:"teamId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    string `json:"points"`
	Min       string `json:"min"`
	FGM       string `json:"fgm"`
	FGA       string `json:"fga"`
	FGP       string `json:"fgp"`
	FTM       string `json:"ftm"`
	FTA       string `json:"fta"`
	FTP       string `json:"ftp"`
	TPM       string `json:"tpm"`
	TPA       string `json:"tpa"`
	TPP       string `json:"tpp"`
	OffReb    string `json:"offReb"`
	DefReb    string `json:"defReb"`
	TotReb    string `json:"totReb"`
	Assists   string `json:"assists"`
	PFouls    string `json:"pFouls"`
	Steals    string `json:"steals"`
	Turnovers string `json:"turnovers"`
	Blocks    string `json:"blocks"`
	PlusMinus string `json:"plusMinus"`
}