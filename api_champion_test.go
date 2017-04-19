package riotapi

import (
	"os"
	"testing"
)

func TestChampionStatus(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	cs, err := c.ChampionStatus()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cs)
}

func TestChampionStatusByID(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	cs, err := c.ChampionStatusByID(201)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cs)
}
