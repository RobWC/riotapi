package riotapi

import "strconv"

// https://developer.riotgames.com/api/methods#!/1078/3718

// RecentSummonerGames get recent summoner games
func (c *APIClient) RecentSummonerGames(id int) (rg *RecentGames, err error) {
	err = c.makeRequest("GET", "v1.3", c.genURL([]string{"game", "by-summoner", strconv.Itoa(id), "recent"}), nil, true, rg)
	if err != nil {
		return rg, err
	}

	return rg, err
}
