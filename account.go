package tfd

import (
	"net/http"
	"net/url"
)

const (
	userBasePath              = "/tfd/v1"
	accountIdentifierPath     = userBasePath + "/id"
	userBasicPath             = userBasePath + "/user/basic"
	userDescendantPath        = userBasePath + "/user/descendant"
	userWeaponsPath           = userBasePath + "/user/weapon"
	userReactorPath           = userBasePath + "/user/reactor"
	userExternalComponentPath = userBasePath + "/user/external-component"
)

type AccountIdentifier struct {
	OUID string `json:"ouid"` // OUID
}

// GetAccountIdentifier retrieves the account identifier (OUID) for a given username.
func (c *Client) GetAccountIdentifier(userName string) (*AccountIdentifier, error) {
	query := url.Values{"user_name": {userName}}
	result, err := c.getUserData(accountIdentifierPath, query, &AccountIdentifier{})
	if err != nil {
		return nil, err
	}
	return result.(*AccountIdentifier), nil
}

// ******************************************

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

// GetBasicInformation retrieves basic information about the user identified by the given OUID.
func (c *Client) GetUserInfo(ouid string) (*UserBasic, error) {
	query := url.Values{"ouid": {ouid}}
	result, err := c.getUserData(userBasicPath, query, &UserBasic{})
	if err != nil {
		return nil, err
	}
	return result.(*UserBasic), nil
}

// ******************************************

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

// GetUserDescendant retrieves descendant information for the user identified by the given OUID.
func (c *Client) GetUserDescendant(ouid string) (*UserDescendant, error) {
	query := url.Values{"ouid": {ouid}}
	result, err := c.getUserData(userDescendantPath, query, &UserDescendant{})
	if err != nil {
		return nil, err
	}
	return result.(*UserDescendant), nil
}

// ******************************************

type UserWeapon struct {
	AccountIdentifier
	UserName string       `json:"user_name"` // Nickname
	Weapons  []WeaponInfo `json:"weapon"`    // Weapon information
}

type WeaponInfo struct {
	ModuleMaxCapacity           int64                   `json:"module_max_capacity"`        // Max. equipable module capacity
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

// GetUserWeapons retrieves weapon information for the user identified by the given OUID and language.
func (c *Client) GetUserWeapons(ouid string, language LanguageCode) (*UserWeapon, error) {
	query := url.Values{
		"language_code": {string(language)},
		"ouid":          {ouid},
	}
	result, err := c.getUserData(userWeaponsPath, query, &UserWeapon{})
	if err != nil {
		return nil, err
	}
	return result.(*UserWeapon), nil
}

// ******************************************

type UserReactor struct {
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

// GetUserReactor retrieves reactor information for the user identified by the given OUID and language.
func (c *Client) GetUserReactor(ouid string, language LanguageCode) (*UserReactor, error) {
	query := url.Values{
		"language_code": {string(language)},
		"ouid":          {ouid},
	}
	result, err := c.getUserData(userReactorPath, query, &UserReactor{})
	if err != nil {
		return nil, err
	}
	return result.(*UserReactor), nil
}

// ******************************************

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

// GetExternalComponent retrieves external component information for the user identified by the given OUID and language.
func (c *Client) GetExternalComponent(ouid string, language LanguageCode) (*UserExternalComponent, error) {
	query := url.Values{
		"language_code": {string(language)},
		"ouid":          {ouid},
	}
	result, err := c.getUserData(userExternalComponentPath, query, &UserExternalComponent{})
	if err != nil {
		return nil, err
	}
	return result.(*UserExternalComponent), nil
}

// ********* Helper Function **************

// getUserData is a helper method that sends a request to the given path with the specified query parameters
// and unmarshals the response into the provided response model.
func (c *Client) getUserData(path string, query url.Values, respModel interface{}) (interface{}, error) {
	req, err := c.newRequest(http.MethodGet, true, path, &query)
	if err != nil {
		return nil, err
	}

	if err := c.do(req, respModel); err != nil {
		return nil, err
	}

	return respModel, nil
}
