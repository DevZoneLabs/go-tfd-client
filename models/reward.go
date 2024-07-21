package models

type RewardResponse struct {
	Reward []Reward `json:"reward"`
}

// Reward represents information about a reward
type Reward struct {
	MapID       string       `json:"map_id"`
	MapName     string       `json:"map_name"`
	BattleZones []BattleZone `json:"battle_zone"`
}

// BattleZone represents information about a battlefield
type BattleZone struct {
	BattleZoneID   string       `json:"battle_zone_id"`
	BattleZoneName string       `json:"battle_zone_name"`
	Rewards        []RewardInfo `json:"reward"`
}

// RewardInfo represents reward rotation information
type RewardInfo struct {
	Rotation           int64  `json:"rotation"`
	RewardType         string `json:"reward_type"`
	ReactorElementType string `json:"reactor_element_type"`
	WeaponRoundsType   string `json:"weapon_rounds_type"`
	ArcheType          string `json:"arche_type"`
}
