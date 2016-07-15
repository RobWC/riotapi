package riotapi

import (
	"os"
	"testing"
)

func TestChampionStatsBySummoner(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	cs, err := c.ChampionStatsBySummoner(46779953)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cs)
}
