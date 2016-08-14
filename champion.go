package riotapi

// ChampionData champion list data
type ChampionData struct {
	Type      string              `json:"type"`
	Format    string              `json:"format"`
	Version   string              `json:"version"`
	Champions map[string]Champion `json:"data"`
	Keys      map[string]string   `json:"keys"`
}

// Champion champion details
type Champion struct {
	ID          int                   `json:"id"`
	Key         string                `json:"key"`
	Name        string                `json:"name"`
	Title       string                `json:"title"`
	Blurb       string                `json:"blurb"`
	AllyTips    []string              `json:"allytips"`
	EnemyTips   []string              `json:"enemytips"`
	Spells      []ChampionSpell       `json:"spells"`
	Passive     ChampionPassive       `json:"passive"`
	Info        ChampionInfo          `json:"info"`
	Image       Image                 `json:"image"`
	Skins       []ChampionSkin        `json:"skins"`
	Lore        string                `json:"lore"`
	Tags        []string              `json:"tags"`
	ParType     string                `json:"partype"`
	Stats       ChampionStats         `json:"stats"`
	Recommended []ChampionRecommended `json:"recommended"`
}

// ChampionSkin jdetails about a champion skin
type ChampionSkin struct {
	ID   int    `json:"id"`
	Num  int    `json:"num"`
	Name string `json:"name"`
}

// ChampionSpell champion abilitiy
type ChampionSpell struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	ToolTip      string      `json:"tooltip"`
	LevelTip     LevelTip    `json:"leveltip"`
	MaxRank      int         `json:"maxrank"`
	Cooldown     []float32   `json:"cooldown"`
	CooldownBurn string      `json:"cooldownBurn"`
	Cost         []int       `json:"cost"`
	CostBurn     string      `json:"costBurn"`
	Effect       [][]float32 `json:"effect"` //TODO: What is this really?
	EffectBurn   []string    `json:"effectburn"`
	Vars         []SpellVar  `json:"vars"`
	CostType     string      `json:"costType"`
	Range        []float32   `json:"range"`
	RangeBurn    string      `json:"rangeBurn"`
	Image        Image       `json:"image"`
	AltImages    Image       `json:"altimages"`
	Resource     string      `json:"resource"`
}

// SpellVar spell data
type SpellVar struct {
	Link  string    `json:"link"`
	Coeff []float32 `json:"coeff"`
	Key   string    `json:"key"`
}

// ChampionInfo
type ChampionInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty" gorethink:"difficulty"`
}

// ChampionStats champion stats details
type ChampionStats struct {
	HP                   float32 `json:"hp"`
	HPPerLevel           float32 `json:"hpperlevel"`
	MP                   float32 `json:"mp"`
	MPPerLevel           float32 `json:"mpperlevel"`
	MoveSpeed            float32 `json:"movespeed"`
	Armor                float32 `json:"armor"`
	ArmorPerLevel        float32 `json:"armorperlevel"`
	SpellBlock           float32 `json:"spellblock"`
	SpellBlockPerLevel   float32 `json:"spellblockperlevel"`
	AttackRange          float32 `json:"attackrange"`
	HPRegen              float32 `json:"hpregen"`
	HPRegenPerLevel      float32 `json:"hpregenperlevel"`
	MPRegen              float32 `json:"mpregen"`
	MPRegenPerLevel      float32 `json:"mpregenperlevel"`
	Crit                 float32 `json:"crit"`
	CritPerLevel         float32 `json:"critperlevel"`
	AttackDamage         float32 `json:"attackdamage"`
	AttackDamagePerLevel float32 `json:"attackdamageperlevel"`
	AttackSpeedOffset    float32 `json:"attackspeedoffset"`
	AttackSpeedPerLevel  float32 `json:"attackspeedperlevel"`
}

// ChampionPassive champion passive ability detais
type ChampionPassive struct {
	Description          string `json:"description"`
	SanitizedDescription string `json:"sanitizedDescription"`
	Name                 string `json:"name"`
	Image                Image  `json:"image"`
}

// ChampionRecommended recommended champion options
type ChampionRecommended struct {
	Blocks []*BlockData
}

// BlockData champion recommended block data
type BlockData struct {
	Items []*BlockItem `json:"items"`
}

// BlockItem champion recommended block item data
type BlockItem struct {
	Count int `json:"count"`
	ID    int `json:"id"`
}

// LevelTip champipn level tip data
type LevelTip struct {
	Effect []*string `json:"effect"`
	Label  []*string `json:"label"`
}
