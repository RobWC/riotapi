package riotapi

// ChampionMasteries a list of champion masteries
type ChampionMasteries []*ChampionMastery

// ChampionMastery information about a player's champion mastery
type ChampionMastery struct {
	ChampionID                   int  `json:"championId"`
	ChampionLevel                int  `json:"championLevel"`
	ChampionPoints               int  `json:"championPoints"`
	ChampionPointsSinceLastLevel int  `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int  `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool `json:"chestGranted"`
	LastPlayTime                 int  `json:"lastPlayTime"`
	PlayerID                     int  `json:"playerId"`
}
