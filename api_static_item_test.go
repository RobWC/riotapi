package riotapi

import (
	"os"
	"testing"
)

func TestStaticItem(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.StaticItem("", ItemDataNil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestStaticItemByID(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.StaticItemByID(3748, "", ItemDataNil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestStaticItemAll(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.StaticItem("", ItemDataAll)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}

func TestStaticItemByIDAll(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	sl, err := c.StaticItemByID(3748, "", ItemDataAll)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sl)
}
