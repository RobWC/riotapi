package riotapi

import (
	"encoding/json"
	"strconv"
)

// https://developer.riotgames.com/api/methods#!/1077

// ChampionStatus get the status of a current champion
// This is useful to see if the champ is free to play, enabled in ranked etc
func (c *APIClient) ChampionStatus() (cl *ChampionStatusList, err error) {
	data, err := c.makeRequest("GET", "v1.2", "champion", nil, true)
	if err != nil {
		return cl, err
	}

	err = json.Unmarshal(data, &cl)
	if err != nil {
		return nil, err
	}

	return cl, err
}

// ChampionStatusByID get the status of a current champion
// This is useful to see if the champ is free to play, enabled in ranked etc
func (c *APIClient) ChampionStatusByID(id int) (cs *ChampionStatus, err error) {
	data, err := c.makeRequest("GET", "v1.2", c.genURL([]string{"champion", strconv.Itoa(id)}), nil, true)
	if err != nil {
		return cs, err
	}

	err = json.Unmarshal(data, &cs)
	if err != nil {
		return nil, err
	}

	return cs, err
}
