package riotapi

import (
	"os"
	"testing"
)

func TestSummonerByName(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	s, err := c.SummonerByName("MrCrapper")
	if err != nil {
		t.Fatal(s)
	}
	t.Log(s)
}

func TestSummonerNameByID(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	t.Log(apiKeyEnv)
	c := NewAPIClient("na", apiKeyEnv)
	s, err := c.SummonerByName("MrCrapper")
	if err != nil {
		t.Fatal(s)
	}

	name, err := c.SummonerNameByID(s.AccountID)
	if err != nil {
		t.Fatal(name)
	}

	t.Log(s)
}
