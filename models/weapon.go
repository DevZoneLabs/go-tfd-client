package models

type WeaponResponse struct {
	Weapons []Weapon
}

type Weapon struct {
	Name       string `json:"weapon_name"`
	ID         string `json:"weapon_id"`
	ImageURL   string `json:"image_url"`
	Type       string `json:"weapon_type"`
	Tier       string `json:"weapon_tier"`
	Rounds     string `json:"weapon_rounds_type"`
	BaseStat   []BaseStat
	FireArmATK []struct {
		Level   int64 `json:"level"`
		FireArm []FireArm
	}
	PerkAbilityName    string `json:"weapon_perk_ability_name"`
	AbilityDescription string `json:"weapon_perk_ability_description"`
	AbilityImageURL    string `json:"weapon_perk_ability_image_url"`
}

type FireArm struct {
	FireArmType  string  `json:"firearm_atk_type"`
	FireArmValue float64 `json:"firearm_atk_value"`
}
