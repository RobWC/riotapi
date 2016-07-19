package riotapi

// ChampionStatusList a list of all the currently available champions
type ChampionStatusList []*ChampionStatus

// ChampionStatus the current status of a Champion
type ChampionStatus struct {
	Active            bool `json:"active"`
	BotEnabled        bool `json:"botEnabled"`
	BotMmEnabled      bool `json:"botMmEnabled"`
	FreeToPlay        bool `json:"freeToPlay"`
	ID                int  `json:"id"`
	RankedPlayEnabled bool `json:"rankedPlayEnabled"`
}
