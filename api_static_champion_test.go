package riotapi

import (
	"os"
	"testing"
)

func TestStaticChampion(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.StaticChampion("", ChampDataNil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestStaticChampionByID(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.StaticChampionByID(41, "", ChampDataNil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestStaticChampionAll(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.StaticChampion("", ChampDataAll)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestStaticChampionByIDAll(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.StaticChampionByID(41, "", ChampDataAll)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}
