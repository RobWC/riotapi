package riotapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Summoner struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ProfileIcon  int    `json:"profileIconId"`
	Level        int    `json:"summonerLevel"`
	RevisionDate int    `json:"revisionDate"`
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
	} else if sl[strings.ToLower(name)] == nil {
		return nil, fmt.Errorf("User does not exist")
	}

	return sl[strings.ToLower(name)], nil
}

// SummonerByIDs Get up to 40 summoners by ID
func (c *APIClient) SummonersByAccountID(ids []int) (sl SummonerList, err error) {
	var req *http.Request
	var sids []string
	sl = make(SummonerList)

	if len(ids) > 40 {
		return sl, fmt.Errorf("Exceeds maximum of 40 ids")
	}

	for i := range ids {
		sids = append(sids, strconv.Itoa(ids[i]))
	}

	req, err = c.genRequest("GET", "v1.4", c.genURL([]string{"summoner", "by-account", strings.Join(sids, ",")}), nil)
	if err != nil {
		return sl, err
	}
	data, err := c.do(req)
	if err != nil {
		return sl, err
	}

	err = json.Unmarshal(data, &sl)
	if err != nil {
		return nil, err
	}

	return sl, nil
}

func (c *APIClient) SummonersByID(ids []int) (sl SummonerList, err error) {
	var req *http.Request
	var sids []string
	sl = make(SummonerList)

	if len(ids) > 40 {
		return sl, fmt.Errorf("Exceeds maximum of 40 ids")
	}

	for i := range ids {
		sids = append(sids, strconv.Itoa(ids[i]))
	}

	req, err = c.genRequest("GET", "v1.4", c.genURL([]string{"summoner", strings.Join(sids, ",")}), nil)
	if err != nil {
		return sl, err
	}
	data, err := c.do(req)
	if err != nil {
		return sl, err
	}

	err = json.Unmarshal(data, &sl)
	if err != nil {
		return nil, err
	}

	return sl, nil
}

func (c *APIClient) SummonerMasteries(ids []int) (sml SummonerMasteryList, err error) {
	var req *http.Request
	var sids []string
	sml = make(SummonerMasteryList)

	if len(ids) > 40 {
		return sml, fmt.Errorf("Exceeds maximum of 40 ids")
	}

	for i := range ids {
		sids = append(sids, strconv.Itoa(ids[i]))
	}

	req, err = c.genRequest("GET", "v1.4", c.genURL([]string{"summoner", strings.Join(sids, ","), "masteries"}), nil)
	if err != nil {
		return sml, err
	}
	data, err := c.do(req)
	if err != nil {
		return sml, err
	}

	err = json.Unmarshal(data, &sml)
	if err != nil {
		return nil, err
	}

	return sml, err
}
