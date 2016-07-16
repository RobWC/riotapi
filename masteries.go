package riotapi

// MasterPage a collection of summoners mastery pages
type MasteryPages struct {
	Pages []*MasteryPage `json:"pages"`
}

type SummonerMasteryList map[string]MasteryPages

// MasteryPage a mastery page
type MasteryPage struct {
	Current   bool      `json:"current"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Masteries []Mastery `json:"masteries"`
}

type Mastery struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}
