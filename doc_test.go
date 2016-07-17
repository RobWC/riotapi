package riotapi

import "os"

// Return a summoner by name
func ExampleSummonerByName() {
	apiKeyEnv := os.Getenv("RIOTKEY")

	if apiKeyEnv == "" {
		panic("API Key Not Specified")
	}

	c := NewAPIClient("na", apiKeyEnv)
	s, err := c.SummonerByName("MrCrapper")
	if err != nil && s != nil {
		panic(err)
	}
}
