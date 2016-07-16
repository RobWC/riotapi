package riotapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Shard a League of Leguends shard data
type Shard struct {
	Hostname  string   `json:"hostname"`
	Locales   []string `json:"locales"`
	Name      string   `json:"name"`
	RegionTag string   `json:"region_tag"`
	Slug      string   `json:"slug"`
}

func (c *APIClient) Shards() (shards []*Shard, err error) {
	req, err := http.NewRequest("GET", "http://status.leagueoflegends.com/shards", nil)
	if err != nil {
		return shards, err
	}

	rc := make(chan struct {
		data []byte
		err  error
	})

	go func(req http.Request) {
		defer func() {
			<-c.tokens
		}()

		c.tokens <- struct{}{}
		resp, err := c.client.Do(&req)
		if err != nil {
			rc <- struct {
				data []byte
				err  error
			}{nil, err}
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			rc <- struct {
				data []byte
				err  error
			}{nil, err}
		}
		if resp.StatusCode != http.StatusOK {
			rc <- struct {
				data []byte
				err  error
			}{nil, fmt.Errorf("API Error %s", string(data))}
		}

		rc <- struct {
			data []byte
			err  error
		}{data, err}
	}(*req)

	r := <-rc
	if r.err != nil {
		return nil, err
	}

	err = json.Unmarshal(r.data, &shards)
	if err != nil {
		return shards, err
	}

	return shards, nil
}

func (c *APIClient) ShardStatus(name string) {

}
