package riotapi

// https://developer.riotgames.com/api/methods#!/1077

// ChampionStatus get the status of a current champion
// This is userful to see if the champ is free to play, enabled in ranked etc
func (c *APIClient) ChampionStatus() (cs *ChampionStatusList, err error) {
	return cs, err
}

// ChampionStatusByID get the status of a current champion
// This is userful to see if the champ is free to play, enabled in ranked etc
func (c *APIClient) ChampionStatusByID(id int) (cs *ChampionStatus, err error) {
	return cs, err
}
