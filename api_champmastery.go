package riotapi

// PlayerChampMasteryByChampion get a players mastery of a specific champion
func (c *APIClient) PlayerChampMasteryByChampion(playerid, champid int) (champ *ChampionMastery, err error) {
	return champ, err
}

// PlayerChampMasteryChampions get a players champions they have mastery with
func (c *APIClient) PlayerChampMasteryChampions(playerid int) (champs *ChampionMasteries, err error) {
	return champs, err
}

// PlayerChampMasteryScore get a players champ mastery score
func (c *APIClient) PlayerChampMasteryScore(playerid int) (score int, err error) {
	return score, err
}

// PlayerChampMasteryTopChampions get a players top champions
func (c *APIClient) PlayerChampMasteryTopChampions(playerid int) (champs *ChampionMasteries, err error) {
	return champs, err
}
