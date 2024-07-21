package models

// Reactor represents information about a reactor
type Reactor struct {
	ID                     string       `json:"reactor_id"`
	Name                   string       `json:"reactor_name"`
	ImageURL               string       `json:"image_url"`
	Tier                   string       `json:"reactor_tier"`
	SkillPower             []SkillPower `json:"reactor_skill_power"`
	OptimizedConditionType string       `json:"optimized_condition_type"`
}

// SkillPower represents skill power by level information
type SkillPower struct {
	Level            int64             `json:"level"`
	AtkPower         float64           `json:"skill_atk_power"`
	SubSkillAtkPower float64           `json:"sub_skill_atk_power"`
	PowerCoeff       []SkillPowerBoost `json:"skill_power_coefficient"` // Skill Power Coefficient
	EnchantEffect    []Enchantment     `json:"enchant_effect"`
}

// SkillPowerBoost represents skill power boost ratio information
type SkillPowerBoost struct {
	CoefficientStatID    string  `json:"coefficient_stat_id"`
	CoefficientStatValue float64 `json:"coefficient_stat_value"`
}

// Enchantment represents enchantment effect by level information
type Enchantment struct {
	Level    int64   `json:"enchant_level"`
	StatType string  `json:"stat_type"`
	Value    float64 `json:"value"`
}
