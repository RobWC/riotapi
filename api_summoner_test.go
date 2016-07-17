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
	if err != nil && s != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestSummonersByAccountID(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.SummonersByAccountID([]int{46779953})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestSummonersByID(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.SummonersByID([]int{46779953})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestSummonersMasteries(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.SummonerMasteries([]int{46779953})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestSummonersRunes(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sr, err := c.SummonerRunes([]int{46779953})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sr)
}
