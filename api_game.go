package riotapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// https://developer.riotgames.com/api/methods#!/1078/3718

// RecentSummonerGames get recent summoner games
func (c *APIClient) RecentSummonerGames(id int) (rg *RecentGames, err error) {
	var req *http.Request

	req, err = c.genRequest("GET", "v1.3", c.genURL([]string{"game", "by-summoner", strconv.Itoa(id), "recent"}), nil)
	if err != nil {
		return rg, err
	}
	data, err := c.do(req, true)
	if err != nil {
		return rg, err
	}

	err = json.Unmarshal(data, &rg)
	if err != nil {
		return nil, err
	}

	return rg, err
}
