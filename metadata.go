package tfd

import (
	"fmt"
	"net/http"
)

// Metadata paths
const metadataBasePath string = "/static/tfd/meta/"

var metadataPaths = map[string]string{
	"descendants":        "/descendant.json",
	"weapons":            "/weapon.json",
	"modules":            "/module.json",
	"reactors":           "/reactor.json",
	"externalComponents": "/external-component.json",
	"rewards":            "/reward.json",
	"stats":              "/stat.json",
	"voidBattles":        "/void-battle.json",
	"titles":             "/title.json",
}

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
	Type  string  `json:"stat_type"`
	Value float64 `json:"stat_value"`
}

type Skill struct {
	Type             string `json:"skill_type"` // active / passive (respective LanguageCode)
	Name             string `json:"skill_name"`
	Element          string `json:"element_type"`
	ArcheType        string `json:"arche_type"`
	SkillImageUrl    string `json:"skill_image_url"`
	SkillDescription string `json:"skill_description"`
}

// GetDescendantsMetadata retrieves metadata for descendants in the specified language.
// It returns a slice of Descendant models and an error if the request fails.
func (c *Client) GetDescendantsMetadata(language LanguageCode) ([]Descendant, error) {
	var resp []Descendant
	if err := c.getMetadata(language, "descendants", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ******************************************

type Weapon struct {
	Name               string       `json:"weapon_name"`
	ID                 string       `json:"weapon_id"`
	ImageURL           string       `json:"image_url"`
	Type               string       `json:"weapon_type"`
	Tier               string       `json:"weapon_tier"`
	Rounds             string       `json:"weapon_rounds_type"`
	BaseStat           []BaseStat   `json:"base_stat"`
	FireArmATK         []FireArmATK `json:"firearm_atk"`
	PerkAbilityName    string       `json:"weapon_perk_ability_name"`
	AbilityDescription string       `json:"weapon_perk_ability_description"`
	AbilityImageURL    string       `json:"weapon_perk_ability_image_url"`
}

type FireArmATK struct {
	Level   int64     `json:"level"`
	FireArm []FireArm `json:"firearm"`
}

type FireArm struct {
	FireArmType  string  `json:"firearm_atk_type"`
	FireArmValue float64 `json:"firearm_atk_value"`
}

// GetWeaponsMetadata retrieves metadata for weapons in the specified language.
// It returns a slice of Weapon models and an error if the request fails.
func (c *Client) GetWeaponsMetadata(language LanguageCode) ([]Weapon, error) {
	var resp []Weapon
	if err := c.getMetadata(language, "weapons", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ******************************************

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

// GetExternalComponentsMetadata retrieves metadata for external components in the specified language.
// It returns a slice of Weapon models and an error if the request fails.
func (c *Client) GetExternalComponentsMetadata(language LanguageCode) ([]ExternalComponent, error) {
	var resp []ExternalComponent
	if err := c.getMetadata(language, "externalComponents", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ******************************************

type Module struct {
	Name       string       `json:"module_name"`
	ID         string       `json:"module_id"`
	ImageURL   string       `json:"image_url"`
	Type       string       `json:"module_type"`
	Tier       string       `json:"module_tier"`
	SocketType string       `json:"module_socket_type"`
	Class      string       `json:"module_class"`
	Stat       []ModuleStat `json:"module_stat"`
}

type ModuleStat struct {
	Level          int64  `json:"level"`
	ModuleCapacity int64  `json:"module_capacity"`
	Value          string `json:"value"`
}

// GetModulesMetadata retrieves metadata for modules in the specified language.
// It returns a slice of Module models and an error if the request fails.
func (c *Client) GetModulesMetadata(language LanguageCode) ([]Module, error) {
	var resp []Module
	if err := c.getMetadata(language, "modules", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ******************************************

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
	Level            int64               `json:"level"`
	AtkPower         float64             `json:"skill_atk_power"`
	SubSkillAtkPower float64             `json:"sub_skill_atk_power"`
	PowerCoefficient []SkillPowerBoost   `json:"skill_power_coefficient"`
	EnchantEffect    []EnchantmentEffect `json:"enchant_effect"`
}

// SkillPowerBoost represents skill power boost ratio information
type SkillPowerBoost struct {
	CoefficientStatID    string  `json:"coefficient_stat_id"`
	CoefficientStatValue float64 `json:"coefficient_stat_value"`
}

// Enchantment represents enchantment effect by level information
type EnchantmentEffect struct {
	Level    int64   `json:"enchant_level"`
	StatType string  `json:"stat_type"`
	Value    float64 `json:"value"`
}

// GetReactorsMetadata retrieves metadata for reactors in the specified language.
// It returns a slice of Reactor models and an error if the request fails.
func (c *Client) GetReactorsMetadata(language LanguageCode) ([]Reactor, error) {
	var resp []Reactor
	if err := c.getMetadata(language, "reactors", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ******************************************

// Reward represents information about a reward
type Reward struct {
	MapID       string       `json:"map_id"`
	MapName     string       `json:"map_name"`
	BattleZones []BattleZone `json:"battle_zone"`
}

// BattleZone represents information about a battlefield
type BattleZone struct {
	ID      string       `json:"battle_zone_id"`
	Name    string       `json:"battle_zone_name"`
	Rewards []RewardInfo `json:"reward"`
}

// RewardInfo represents reward rotation information
type RewardInfo struct {
	Rotation           int64  `json:"rotation"`
	Type               string `json:"reward_type"` // Reward Type
	ReactorElementType string `json:"reactor_element_type"`
	WeaponRoundsType   string `json:"weapon_rounds_type"`
	ArcheType          string `json:"arche_type"`
}

// GetRewardsMetadata retrieves metadata for rewards in the specified language.
// It returns a slice of Reward models and an error if the request fails.
func (c *Client) GetRewardsMetadata(language LanguageCode) ([]Reward, error) {
	var resp []Reward
	if err := c.getMetadata(language, "rewards", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ******************************************

type StatMeta struct {
	ID   string `json:"stat_id"`
	Name string `json:"stat_name"`
}

// GetStatsMetadata retrieves metadata for stats in the specified language.
// It returns a slice of StatMeta models and an error if the request fails.
func (c *Client) GetStatsMetadata(language LanguageCode) ([]StatMeta, error) {
	var resp []StatMeta
	if err := c.getMetadata(language, "stats", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type VoidBattleResponse struct {
	ID   string `json:"void_battle_id"`
	Name string `json:"void_battle_name"`
}

// GetVoidBattlesMetadata retrieves metadata for void battles in the specified language.
// It returns a slice of VoidBattleResponse models and an error if the request fails.
func (c *Client) GetVoidBattlesMetadata(language LanguageCode) ([]VoidBattleResponse, error) {
	var resp []VoidBattleResponse
	if err := c.getMetadata(language, "voidBattles", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ******************************************

type TitleResponse struct {
	ID   string `json:"title_id"`
	Name string `json:"title_name"`
}

// GetTitlesMetadata retrieves metadata for titles in the specified language.
// It returns a slice of TitleResponse models and an error if the request fails.
func (c *Client) GetTitlesMetadata(language LanguageCode) ([]TitleResponse, error) {
	var resp []TitleResponse
	if err := c.getMetadata(language, "titles", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ************ Helper Function *************

// getMetadata is a generic method to retrieve metadata for a given type and language.
// It takes a LanguageCode, the type of metadata to fetch, and a response interface to unmarshal the data.
// It returns an error if the request or unmarshalling fails.
func (c *Client) getMetadata(language LanguageCode, metadataType string, response interface{}) error {
	path, ok := metadataPaths[metadataType]
	if !ok {
		return fmt.Errorf("unknown metadata type: %s", metadataType)
	}

	req, err := c.newRequest(http.MethodGet, false, fmt.Sprintf("%s%s%s", metadataBasePath, language, path), nil)
	if err != nil {
		return err
	}

	if err := c.do(req, response); err != nil {
		return err
	}

	return nil
}
