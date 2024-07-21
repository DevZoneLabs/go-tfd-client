package models

// ExternalComponentResponse represents the response structure containing external component information

// ExternalComponent represents the information of an external component
type ExternalComponent struct {
	ID            string `json:"external_component_id"`             // External component identifier
	Name          string `json:"external_component_name"`           // External component name
	ImageURL      string `json:"image_url"`                         // External component image path
	EquipmentType string `json:"external_component_equipment_type"` // External component equipment part
	Tier          string `json:"external_component_tier"`           // External component tier
	BaseStat      []struct {
		BaseStat
		Level int64 `json:"level"` // External component level
	} `json:"base_stat"` // External component effect by level information
	SetOptionDetail []SetOptionDetail `json:"set_option_detail"` // External component set option information
}

// SetOptionDetail represents the set option information for an external component
type SetOptionDetail struct {
	Option       string `json:"set_option"`        // External component set option type
	Count        int64  `json:"set_count"`         // Number of external component sets
	OptionEffect string `json:"set_option_effect"` // External component set option set effect
}
