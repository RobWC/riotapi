package riotapi

// Player a player within a game
type Player struct {
	ChampionID int `json:"championId"`
	SummonerID int `json:"summonerId"`
	TeamId     int `json:"teamId"`
}
