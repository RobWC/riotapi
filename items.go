package riotapi

// Item item data
type Item struct {
	Colloq           string      `json:"colloq"`
	ConsumeOnFull    bool        `json:"consumeOnFull"`
	Consumed         bool        `json:"consumed"`
	Depth            int         `json:"depth"`
	Description      string      `json:"description"`
	From             []string    `json:"from"`
	Gold             ItemCost    `json:"gold"`
	Group            string      `json:"group"`
	HideFromAll      bool        `json:"hideFromAll"`
	ID               int         `json:"id"`
	Image            Image       `json:"image"`
	InStore          bool        `json:"inStore"`
	Into             []string    `json:"into"`
	Maps             ItemMapList `json:"maps"`
	Name             string      `json:"name"`
	PlainText        string      `json:"plaintext"`
	RequiredChampion string      `json:"requiredChampion"`
	Rune             RuneData    `json:"rune"`
	Stats            ItemStats   `json:"stats"`
	Tags             []string    `json:"tags"`
}

// RuneData rune metadata
type RuneData struct {
	IsRune bool   `json:"isRune"`
	Tier   string `json:"tier"`
	Type   string `json:"type"`
}

// ItemMapList list of maps an item can be used on
type ItemMapList struct {
	Map1  bool `json:"1"`
	Map8  bool `json:"8"`
	Map10 bool `json:"10"`
	Map12 bool `json:"12"`
}

// ItemCost the cost of an item
type ItemCost struct {
	Base        int  `json:"base"`
	Purchasable bool `json:"purchasable"`
	Total       int  `json:"total"`
	Sell        int  `json:"sell"`
}

// ItemStats item stats and data
type ItemStats struct {
	FlatHPPoolMod                       float32 `json:"FlatHPPoolMod"`
	rFlatHPModPerLevel                  float32 `json:"rFlatHPModPerLevel"`
	FlatMPPoolMod                       float32 `json:"FlatMPPoolMod"`
	rFlatMPModPerLevel                  float32 `json:"rFlatMPModPerLevel"`
	PercentHPPoolMod                    float32 `json:"PercentHPPoolMod"`
	PercentMPPoolMod                    float32 `json:"PercentMPPoolMod"`
	FlatHPRegenMod                      float32 `json:"FlatHPRegenMod"`
	rFlatHPRegenModPerLevel             float32 `json:"rFlatHPRegenModPerLevel"`
	PercentHPRegenMod                   float32 `json:"PercentHPRegenMod"`
	FlatMPRegenMod                      float32 `json:"FlatMPRegenMod"`
	rFlatMPRegenModPerLevel             float32 `json:"rFlatMPRegenModPerLevel"`
	PercentMPRegenMod                   float32 `json:"PercentMPRegenMod"`
	FlatArmorMod                        float32 `json:"FlatArmorMod"`
	rFlatArmorModPerLevel               float32 `json:"rFlatArmorModPerLevel"`
	PercentArmorMod                     float32 `json:"PercentArmorMod"`
	rFlatArmorPenetrationMod            float32 `json:"rFlatArmorPenetrationMod"`
	rFlatArmorPenetrationModPerLevel    float32 `json:"rFlatArmorPenetrationModPerLevel"`
	rPercentArmorPenetrationMod         float32 `json:"rPercentArmorPenetrationMod"`
	rPercentArmorPenetrationModPerLevel float32 `json:"rPercentArmorPenetrationModPerLevel"`
	FlatPhysicalDamageMod               float32 `json:"FlatPhysicalDamageMod"`
	rFlatPhysicalDamageModPerLevel      float32 `json:"rFlatPhysicalDamageModPerLevel"`
	PercentPhysicalDamageMod            float32 `json:"PercentPhysicalDamageMod"`
	FlatMagicDamageMod                  float32 `json:"FlatMagicDamageMod"`
	rFlatMagicDamageModPerLevel         float32 `json:"rFlatMagicDamageModPerLevel"`
	PercentMagicDamageMod               float32 `json:"PercentMagicDamageMod"`
	FlatMovementSpeedMod                float32 `json:"FlatMovementSpeedMod"`
	rFlatMovementSpeedModPerLevel       float32 `json:"rFlatMovementSpeedModPerLevel"`
	PercentMovementSpeedMod             float32 `json:"PercentMovementSpeedMod"`
	rPercentMovementSpeedModPerLevel    float32 `json:"rPercentMovementSpeedModPerLevel"`
	FlatAttackSpeedMod                  float32 `json:"FlatAttackSpeedMod"`
	PercentAttackSpeedMod               float32 `json:"PercentAttackSpeedMod"`
	rPercentAttackSpeedModPerLevel      float32 `json:"rPercentAttackSpeedModPerLevel"`
	rFlatDodgeMod                       float32 `json:"rFlatDodgeMod"`
	rFlatDodgeModPerLevel               float32 `json:"rFlatDodgeModPerLevel"`
	PercentDodgeMod                     float32 `json:"PercentDodgeMod"`
	FlatCritChanceMod                   float32 `json:"FlatCritChanceMod"`
	rFlatCritChanceModPerLevel          float32 `json:"rFlatCritChanceModPerLevel"`
	PercentCritChanceMod                float32 `json:"PercentCritChanceMod"`
	FlatCritDamageMod                   float32 `json:"FlatCritDamageMod"`
	rFlatCritDamageModPerLevel          float32 `json:"rFlatCritDamageModPerLevel"`
	PercentCritDamageMod                float32 `json:"PercentCritDamageMod"`
	FlatBlockMod                        float32 `json:"FlatBlockMod"`
	PercentBlockMod                     float32 `json:"PercentBlockMod"`
	FlatSpellBlockMod                   float32 `json:"FlatSpellBlockMod"`
	rFlatSpellBlockModPerLevel          float32 `json:"rFlatSpellBlockModPerLevel"`
	PercentSpellBlockMod                float32 `json:"PercentSpellBlockMod"`
	FlatEXPBonus                        float32 `json:"FlatEXPBonus"`
	PercentEXPBonus                     float32 `json:"PercentEXPBonus"`
	rPercentCooldownMod                 float32 `json:"rPercentCooldownMod"`
	rPercentCooldownModPerLevel         float32 `json:"rPercentCooldownModPerLevel"`
	rFlatTimeDeadMod                    float32 `json:"rFlatTimeDeadMod"`
	rFlatTimeDeadModPerLevel            float32 `json:"rFlatTimeDeadModPerLevel"`
	rPercentTimeDeadMod                 float32 `json:"rPercentTimeDeadMod"`
	rPercentTimeDeadModPerLevel         float32 `json:"rPercentTimeDeadModPerLevel"`
	rFlatGoldPer10Mod                   float32 `json:"rFlatGoldPer10Mod"`
	rFlatMagicPenetrationMod            float32 `json:"rFlatMagicPenetrationMod"`
	rFlatMagicPenetrationModPerLevel    float32 `json:"rFlatMagicPenetrationModPerLevel"`
	rPercentMagicPenetrationMod         float32 `json:"rPercentMagicPenetrationMod"`
	rPercentMagicPenetrationModPerLevel float32 `json:"rPercentMagicPenetrationModPerLevel"`
	FlatEnergyRegenMod                  float32 `json:"FlatEnergyRegenMod"`
	rFlatEnergyRegenModPerLevel         float32 `json:"rFlatEnergyRegenModPerLevel"`
	FlatEnergyPoolMod                   float32 `json:"FlatEnergyPoolMod"`
	rFlatEnergyModPerLevel              float32 `json:"rFlatEnergyModPerLevel"`
	PercentLifeStealMod                 float32 `json:"PercentLifeStealMod"`
	PercentSpellVampMod                 float32 `json:"PercentSpellVampMod"`
}

// ItemEffect effects that an item can posess
type ItemEffect struct {
	Effect1Amount string `json:"Effect1Amount"`
	Effect2Amount string `json:"Effect2Amount"`
	Effect3Amount string `json:"Effect3Amount"`
	Effect4Amount string `json:"Effect4Amount"`
	Effect5Amount string `json:"Effect5Amount"`
	Effect6Amount string `json:"Effect6Amount"`
	Effect7Amount string `json:"Effect7Amount"`
	Effect8Amount string `json:"Effect8Amount"`
}

// ItemData a set of items
type ItemData struct {
	Type    string          `json:"type"`
	Format  string          `json:"format"`
	Version string          `json:"version"`
	Basic   string          `json:"-"`
	Items   map[string]Item `json:"data"`
}
