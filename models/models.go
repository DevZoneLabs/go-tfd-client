package models

type AccountIdentifier struct {
	OUID string `json:"ouid"` // OUID
}

type UserBasic struct {
	AccountIdentifier
	UserName         string `json:"user_name"`          // Nickname
	PlatformType     string `json:"platform_type"`      // Platform type
	MasteryRankLevel int64  `json:"mastery_rank_level"` // Mastery Rank
	MasteryRankExp   int64  `json:"mastery_rank_exp"`   // Mastery EXP
	TitlePrefix      string `json:"title_prefix_id"`    // Prefix title identifier (Refer to /meta/title API)
	TitleSuffix      string `json:"title_suffix_id"`    // Suffix title identifier (Refer to /meta/title API)
	OSLanguage       string `json:"os_language"`        // OS language settings
	GameLanguage     string `json:"game_language"`      // Game language settings
}

type UserDescendant struct {
	AccountIdentifier
	UserName          string       `json:"user_name"`           // Nickname
	ID                string       `json:"descendant_id"`       // Equipped descendant identifier
	SlotID            string       `json:"descendant_slot_id"`  // Descendant slot identifier
	Level             int64        `json:"descendant_level"`    // Equipped descendant level
	ModuleMaxCapacity int64        `json:"module_max_capacity"` // Max. equippable module capacity
	ModuleCapacity    int64        `json:"module_capacity"`     // Equipped capacity
	Modules           []ModuleInfo `json:"module"`              // Module information
}

type ModuleInfo struct {
	SlotID       string `json:"module_slot_id"`       // Module slot identifier
	ID           string `json:"module_id"`            // Module identifier (Refer to /meta/module API)
	EnchantLevel int64  `json:"module_enchant_level"` // Module enchantment level
}

type UserWeapon struct {
	AccountIdentifier
	UserName string       `json:"user_name"` // Nickname
	Weapons  []WeaponInfo `json:"weapon"`    // Weapon information
}

type WeaponInfo struct {
	ModuleMaxCapacity           int64                   `json:"module_max_capacity"`        // Max. equippable module capacity
	ModuleCapacity              int64                   `json:"module_capacity"`            // Equipped capacity
	SlotID                      string                  `json:"weapon_slot_id"`             // Weapon slot identifier
	ID                          string                  `json:"weapon_id"`                  // Weapon identifier (Refer to /meta/weapon API)
	Level                       int64                   `json:"weapon_level"`               // Weapon level
	PerkAbilityEnchantmentLevel int64                   `json:"perk_ability_enchant_level"` // Weapon Unique Ability enchantment level
	AdditionalStats             []WeaponAdditionalStats `json:"weapon_additional_stat"`     // Weapon random option information
	Module                      []ModuleInfo            `json:"module"`                     // Module information
}

type WeaponAdditionalStats struct {
	AdditionalStats     string `json:"additional_stat_name"`  // Weapon random option name
	AdditionalStatValue string `json:"additional_stat_value"` // Weapon random option value
}

type UserReactorInfo struct {
	AccountIdentifier
	UserName        string                   `json:"user_name"`                // Nickname
	ID              string                   `json:"reactor_id"`               // Reactor identifier (Refer to /meta/reactor API)
	SlotID          string                   `json:"reactor_slot_id"`          // Reactor slot identifier
	Level           int64                    `json:"reactor_level"`            // Reactor level
	AdditionalStats []ReactorAdditionalStats `json:"reactor_additional_stats"` // Reactor information
	EnchantLevel    int64                    `json:"reactor_enchant_level"`    // Reactor enchantment level
}

type ReactorAdditionalStats struct {
	AdditionalStats     string `json:"additional_stat_name"`  // Reactor random option name
	AdditionalStatValue string `json:"additional_stat_value"` // Reactor random option value
}

type UserExternalComponent struct {
	AccountIdentifier
	UserName           string                  `json:"user_name"`          // Nickname
	ExternalComponents []ExternalComponentInfo `json:"external_component"` // External component information
}

type ExternalComponentInfo struct {
	SlotID          string                             `json:"external_component_slot_id"`         // External component slot identifier
	ID              string                             `json:"external_component_id"`              // External component identifier (Refer to /meta/external-component API)
	Level           int64                              `json:"external_component_level"`           // External component level
	AdditionalStats []ExternalComponentAdditionalStats `json:"external_component_additional_stat"` // External component random option information
}

type ExternalComponentAdditionalStats struct {
	AdditionalStats      string `json:"additional_stat_name"`  // External component random option name
	AdditionalStatsValue string `json:"additional_stat_value"` // External component random option value
}

type ModuleRecommendation struct {
	Descendant []DescendantRec
	Weapon     []WeaponRec
}

type DescendantRec struct {
	DescendantID   string `json:"descendant_id"`
	Recommendation []Recommendation
}

type Recommendation struct {
	ModuleID string `json:"module_id"`
}

type WeaponRec struct {
	WeaponID       string `json:"weapon_id"`
	Recommendation []Recommendation
}

type BaseStat struct {
	ID    string  `json:"stat_id"`
	Value float64 `json:"stat_value"`
}

type ErrorMessage struct {
	Error struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	} `json:"error"`
}
