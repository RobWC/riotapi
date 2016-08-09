package riotapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// https://developer.riotgames.com/api/methods#!/1077

// ChampionStatus get the status of a current champion
// This is userful to see if the champ is free to play, enabled in ranked etc
func (c *APIClient) ChampionStatus() (cs *ChampionStatusList, err error) {
	var req *http.Request

	req, err = c.genRequest("GET", "v1.2", "champion", nil)
	if err != nil {
		return cs, err
	}
	data, err := c.do(req, true)
	if err != nil {
		return cs, err
	}

	err = json.Unmarshal(data, &cs)
	if err != nil {
		return nil, err
	}

	return cs, err
}

// ChampionStatusByID get the status of a current champion
// This is userful to see if the champ is free to play, enabled in ranked etc
func (c *APIClient) ChampionStatusByID(id int) (cs *ChampionStatus, err error) {
	var req *http.Request

	req, err = c.genRequest("GET", "v1.2", c.genURL([]string{"champion", strconv.Itoa(id)}), nil)
	if err != nil {
		return cs, err
	}
	data, err := c.do(req, true)
	if err != nil {
		return cs, err
	}

	err = json.Unmarshal(data, &cs)
	if err != nil {
		return nil, err
	}

	return cs, err
}
