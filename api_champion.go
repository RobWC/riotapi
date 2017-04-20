package riotapi

import "strconv"

// https://developer.riotgames.com/api/methods#!/1077

// ChampionStatus get the status of a current champion
// This is useful to see if the champ is free to play, enabled in ranked etc
func (c *APIClient) ChampionStatus() (cl *ChampionStatusList, err error) {
	err = c.makeRequest("GET", "v1.2", "champion", nil, true, cl)
	if err != nil {
		return cl, err
	}

	return cl, err
}

// ChampionStatusByID get the status of a current champion
// This is useful to see if the champ is free to play, enabled in ranked etc
func (c *APIClient) ChampionStatusByID(id int) (cs *ChampionStatus, err error) {

	err = c.makeRequest("GET", "v1.2", c.genURL([]string{"champion", strconv.Itoa(id)}), nil, true, cs)
	if err != nil {
		return cs, err
	}

	return cs, err
}
