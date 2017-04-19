package riotapi

import (
	"encoding/json"
	"strconv"
)

// https://developer.riotgames.com/api/methods#!/1078/3718

// RecentSummonerGames get recent summoner games
func (c *APIClient) RecentSummonerGames(id int) (rg *RecentGames, err error) {
	data, err := c.makeRequest("GET", "v1.3", c.genURL([]string{"game", "by-summoner", strconv.Itoa(id), "recent"}), nil, true)
	if err != nil {
		return rg, err
	}

	err = json.Unmarshal(data, &rg)
	if err != nil {
		return nil, err
	}

	return rg, err
}
