package models

type DescendantResponse struct {
	Descendants []Descendant `json:"descendant"`
}

type Descendant struct {
	ID       string            `json:"id"`                   // Descendant identifier
	Name     string            `json:"descendant_name"`      // Descendant name
	ImageURL string            `json:"descendant_image_url"` // Descendant image path
	Stats    []DescendantStats `json:"descendant_stat"`      // Descendant stat information

}

type DescendantStats struct {
	Level   int64      `json:"level"`
	Details []BaseStat `json:"descendant_stat"`
	Skills  []Skill    `json:"descendant_skill"`
}

type Skill struct {
	Type             SkillType `json:"type"` // active / passive
	Name             string    `json:"name"`
	Element          string    `json:"element_type"`
	ArcheType        string    `json:"arche_type"`
	SkillImageUrl    string    `json:"skill_image_url"`
	SkillDescription string    `json:"skill_description"`
}

type SkillType string

var SkillTypes = struct {
	Pasive SkillType
	Active SkillType
}{"pasive", "active"}
