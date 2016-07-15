package riotapi

// MasterPage a collection of summoners mastery pages
type MasteryPages map[string]MastereyPage

// MasteryPage a mastery page
type MastereyPage struct {
	Current   bool      `json:"current"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Masteries []Mastery `json:"masteries"`
}

type Mastery struct {
	ID   int `json:"id"`
	Rank int `json:"rank"`
}
