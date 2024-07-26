package models

type Descendant struct {
	ID       string            `json:"descendant_id"`        // Descendant identifier
	Name     string            `json:"descendant_name"`      // Descendant name
	ImageURL string            `json:"descendant_image_url"` // Descendant image path
	Stats    []DescendantStats `json:"descendant_stat"`      // Descendant stat information
	Skills   []Skill           `json:"descendant_skill"`     // Descendant skill
}

type DescendantStats struct {
	Level   int64        `json:"level"`
	Details []StatDetail `json:"stat_detail"`
}

type StatDetail struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Skill struct {
	Type             SkillType `json:"skill_type"` // active / passive
	Name             string    `json:"skill_name"`
	Element          string    `json:"element_type"`
	ArcheType        string    `json:"arche_type"`
	SkillImageUrl    string    `json:"skill_image_url"`
	SkillDescription string    `json:"skill_description"`
}

type SkillType string

var SkillTypes = struct {
	Passive SkillType
	Active  SkillType
}{"passive", "active"}
