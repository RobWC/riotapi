package riotapi

// RawStats raw stats from a game
type RawStats struct {
	Assists                         int `json:"assists"`
	BarracksKilled                  int `json:"barracksKilled"` // ean	Flag specifying if the summoner got the killing blow on the nexus.
	BountyLevel                     int
	ChampionsKilled                 int
	CombatPlayerScore               int
	ConsumablesPurchased            int
	DamageDealtPlayer               int
	DoubleKills                     int
	FirstBlood                      int
	Gold                            int
	GoldEarned                      int
	GoldSpent                       int
	Item0                           int
	Item1                           int
	Item2                           int
	Item3                           int
	Item4                           int
	Item5                           int
	Item6                           int
	ItemsPurchased                  int
	KillingSprees                   int
	LargestCriticalStrike           int
	LargestKillingSpree             int
	LargestMultiKill                int
	LegendaryItemsCreated           int // 	Number of tier 3 items built.
	Level                           int
	MagicDamageDealtPlayer          int
	MagicDamageDealtToChampions     int
	MagicDamageTaken                int
	MinionsDenied                   int
	MinionsKilled                   int
	NeutralMinionsKilled            int
	NeutralMinionsKilledEnemyJungle int
	NeutralMinionsKilledYourJungle  int
	NexusKilled                     bool // 	Flag specifying if the summoner got the killing blow on the nexus
	NodeCapture                     int
	NodeCaptureAssist               int
	NodeNeutralize                  int
	NodeNeutralizeAssist            int
	NumDeaths                       int
	NumItemsBought                  int
	ObjectivePlayerScore            int
	PentaKills                      int
	PhysicalDamageDealtPlayer       int
	PhysicalDamageDealtToChampions  int
	PhysicalDamageTaken             int
	PlayerPosition                  int // Player position (Legal values: TOP(1), MIDDLE(2), JUNGLE(3), BOT(4))
	PlayerRole                      int // Player role (Legal values: DUO(1), SUPPORT(2), CARRY(3), SOLO(4))
	PlayerScore0                    int
	PlayerScore1                    int
	PlayerScore2                    int
	PlayerScore3                    int
	PlayerScore4                    int
	PlayerScore5                    int
	PlayerScore6                    int
	PlayerScore7                    int
	PlayerScore8                    int
	PlayerScore9                    int
	QuadraKills                     int
	SightWardsBought                int
	Spell1Cast                      int
	Spell2Cast                      int
	Spell3Cast                      int
	Spell4Cast                      int
	SummonSpell1Cast                int
	SummonSpell2Cast                int
	SuperMonsterKilled              int
	Team                            int
	TeamObjective                   int
	TimePlayed                      int
	TotalDamageDealt                int
	TotalDamageDealtToChampions     int
	TotalDamageTaken                int
	TotalHeal                       int
	TotalPlayerScore                int
	TotalScoreRank                  int
	TotalTimeCrowdControlDealt      int
	TotalUnitsHealed                int
	TripleKills                     int
	TrueDamageDealtPlayer           int
	TrueDamageDealtToChampions      int
	TrueDamageTaken                 int
	TurretsKilled                   int
	UnrealKills                     int
	VictoryPointTotal               int
	VisionWardsBought               int
	wardKilled                      int
	WardPlaced                      int
	Win                             bool
}
