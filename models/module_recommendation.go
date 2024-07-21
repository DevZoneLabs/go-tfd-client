package models

// ModuleRecommendationResponse represents the response object structure
type ModuleRecommendation struct {
	Descendant DescendantRecommendation `json:"descendant"`
	Weapon     WeaponRecommendation     `json:"weapon"`
}

// DescendantRecommendation represents descendant recommendation information
type DescendantRecommendation struct {
	DescendantID    string             `json:"descendant_id"`
	Recommendations []DescendantModule `json:"recommendation"`
}

// DescendantModule represents module recommendation information for a descendant
type DescendantModule struct {
	ModuleID string `json:"module_id"`
}

// WeaponRecommendation represents weapon recommendation information
type WeaponRecommendation struct {
	WeaponID        string         `json:"weapon_id"`
	Recommendations []WeaponModule `json:"recommendation"`
}

// WeaponModule represents module recommendation information for a weapon
type WeaponModule struct {
	ModuleID string `json:"module_id"`
}
