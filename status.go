package riotapi

import (
	"encoding/json"
	"fmt"
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

// Shards get a listing for all of the LoL shards
func (c *APIClient) Shards() (shards []*Shard, err error) {
	req, err := http.NewRequest("GET", "http://status.leagueoflegends.com/shards", nil)

	var data []byte
	data, err = c.do(req, false)
	if err != nil {
		return shards, err
	}

	err = json.Unmarshal(data, &shards)
	if err != nil {
		return shards, err
	}

	return shards, nil
}

// ShardStatus a report of the current status of a shard
type ShardStatus struct {
	Hostname string     `json:"hostname"`
	Locales  []string   `json:"locales"`
	Name     string     `json:"region_tag"`
	Services []*Service `json:"services"`
	Slug     string     `json:"slug,omitempty"`
}

// Service a shard service
type Service struct {
	Incidents []*Incident `json:"incidents"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug,omitempty"`
	Status    string      `json:"status"`
}

// Incident an incident report
type Incident struct {
	Active    bool               `json:"active"`
	CreatedAt string             `json:"created_at"`
	ID        int                `json:"id"`
	Updates   []*IncidentMessage `json:"updates"`
}

// IncidentMessage a message about an incident
type IncidentMessage struct {
	Author       string                `json:"author"`
	Content      string                `json:"content"`
	CreatedAt    string                `json:"created_at"`
	ID           string                `json:"id"`
	Severity     string                `json:"severity"`
	Translations []*MessageTranslation `json:"translations"`
	UpdatedAt    string                `json:"update_at"`
}

// MessageTranslation a translation of a message into another locale
type MessageTranslation struct {
	Content   string `json:"content"`
	Locale    string `json:"locale"`
	UpdatedAt string `json:"updated_at"`
}

// ShardStatus get the status of a specified shard
func (c *APIClient) ShardStatus(name string) (ss *ShardStatus, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://status.leagueoflegends.com/shards/%s", name), nil)

	var data []byte
	data, err = c.do(req, false)
	if err != nil {
		return ss, err
	}

	err = json.Unmarshal(data, &ss)
	if err != nil {
		return ss, err
	}

	return ss, err
}
