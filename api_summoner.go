package riotapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Summoner struct {
	AccountID    int            `json:"accountId"`
	SummonerID   int            `json:"summonerId"`
	PlatformID   string         `json:"platformId"`
	AllyBadge    int            `json:"allyBadge"`
	CurrentUser  Account        `json:"currentUser"`
	ELO          map[string]int `json:"elo"`
	EnemeyBadge  int            `json:"enemyBadge"`
	HighestRank  string         `json:"highestRank"`
	SummonerName string         `json:"summonerName"`
	ProfileIcon  int            `json:"profileIcon"`
}

type SummonerList map[string]*Summoner

// SummonerByName v1.4 of get summoner by name
func (c *APIClient) SummonerByName(name string) (s *Summoner, err error) {

	var req *http.Request
	req, err = c.genRequest("GET", "v1.4", c.genURL([]string{"summoner", "by-name", name}), nil)
	if err != nil {
		return nil, err
	}
	data, err := c.do(req)
	if err != nil {
		return nil, err
	}

	sl := make(SummonerList)
	err = json.Unmarshal(data, &sl)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v", sl["status"])

	return sl[name], nil
}

func (c *APIClient) SummonerNameByID(id int) (name string, err error) {
	var req *http.Request
	req, err = c.genRequest("GET", "v1.4", c.genURL([]string{"summoner", strconv.Itoa(id), "name"}), nil)
	if err != nil {
		return name, err
	}
	data, err := c.do(req)
	if err != nil {
		return name, err
	}

	namem := make(map[string]string)
	err = json.Unmarshal(data, &namem)
	if err != nil {
		return name, err
	}

	log.Println(namem)

	return name, err
}

func (c *APIClient) SummonerByID(id string) (s *Summoner, err error) {

	return s, err
}

func (c *APIClient) SummonerMasteries(id string) (m *MasteryPages, err error) {
	return m, err
}
