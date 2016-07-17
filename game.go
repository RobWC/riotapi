package riotapi

// RecentGames recent games played by a summoner
type RecentGames struct {
	Games      []*Game `json:"games"`
	SummonerID int     `json:"summonerid"`
}

// Game a game record
type Game struct {
	ChampionID    int       `json:"championId"`
	CreateDate    int       `json:"createDate"`
	FellowPlayers []*Player `json:"fellowPlayers"`
	GameID        int       `json:"gameId"`
	GameMode      string    `json:"gameMode"`
	GameType      string    `json:"gameType"`
	Invalid       bool      `json:"invalid"`
	IPEarned      int       `json:"ipEarned"`
	Level         int       `json:"level"`
	MapID         int       `json:"mapId"`
	Spell1        int       `json:"spell1"`
	Spell2        int       `json:"spell2"`
	Stats         *RawStats `json:"stats"`
	SubType       string    `json:"subType"`
	TeamId        int       `json:"teamId"`
}
