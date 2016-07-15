package riotapi

import (
	"encoding/json"
	"strconv"
	"strings"
)

type SummonerChampionStats struct {
	Champions        string
	LastRankedChange int
	User             Account
	Version          string
}

func (a *APIClient) ChampionStatsBySummoner(id int) (*ChampionStats, error) {

	apiString := strings.Join([]string{"stats", "by-summoner", strconv.Itoa(id)}, "/")
	data, err := a.doRequest("GET", "v2.2", apiString, nil)
	if err != nil {
		return nil, err
	}

	cs := &ChampionStats{}

	err = json.Unmarshal(data, &cs)
	if err != nil {
		return nil, err
	}

	return cs, nil
}
