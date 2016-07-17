package riotapi

// SummonerRuneList a list of rune pages with sumnoner ID as the key
type SummonerRuneList map[string]*RunePages

// RunePages a set of rune pages for a summoner
type RunePages struct {
	Pages      []*RunePage `json:"pages"`
	SummonerID int         `json:"summonerId"`
}

// RunePage a page of runes
type RunePage struct {
	Current bool        `json:"current"`
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Slots   []*RuneSlot `json:"slots"`
}

// RuneSlot defines which rune is in which rune page slot
type RuneSlot struct {
	RuneID     int `json:"runeId"`
	RuneSlotID int `json:"runeSlot"`
}
