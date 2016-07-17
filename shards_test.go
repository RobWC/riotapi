package riotapi

import (
	"os"
	"testing"
)

func TestShards(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	s, err := c.Shards()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestShardStatus(t *testing.T) {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		t.Fatal("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	ss, err := c.ShardStatus("na")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}
