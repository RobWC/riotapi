package riotapi

import (
	"os"
	"testing"
)

func TestRecentSummonerGames(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sr, err := c.RecentSummonerGames(46779953)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sr)
}
