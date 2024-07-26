package tfd

import (
	"fmt"
	"net/http"

	"github.com/DevZoneLabs/go-tfd-client/models"
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

// GetDescendantsMetadata retrieves metadata for descendants in the specified language.
// It returns a slice of Descendant models and an error if the request fails.
func (c *Client) GetDescendantsMetadata(language LanguageCode) ([]models.Descendant, error) {
	var resp []models.Descendant
	if err := c.getMetadata(language, "descendants", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetWeaponsMetadata retrieves metadata for weapons in the specified language.
// It returns a slice of Weapon models and an error if the request fails.
func (c *Client) GetWeaponsMetadata(language LanguageCode) ([]models.Weapon, error) {
	var resp []models.Weapon
	if err := c.getMetadata(language, "weapons", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetModulesMetadata retrieves metadata for modules in the specified language.
// It returns a slice of Module models and an error if the request fails.
func (c *Client) GetModulesMetadata(language LanguageCode) ([]models.Module, error) {
	var resp []models.Module
	if err := c.getMetadata(language, "modules", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetReactorsMetadata retrieves metadata for reactors in the specified language.
// It returns a slice of Reactor models and an error if the request fails.
func (c *Client) GetReactorsMetadata(language LanguageCode) ([]models.Reactor, error) {
	var resp []models.Reactor
	if err := c.getMetadata(language, "reactors", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetRewardsMetadata retrieves metadata for rewards in the specified language.
// It returns a slice of Reward models and an error if the request fails.
func (c *Client) GetRewardsMetadata(language LanguageCode) ([]models.Reward, error) {
	var resp []models.Reward
	if err := c.getMetadata(language, "rewards", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetStatsMetadata retrieves metadata for stats in the specified language.
// It returns a slice of StatMeta models and an error if the request fails.
func (c *Client) GetStatsMetadata(language LanguageCode) ([]models.StatMeta, error) {
	var resp []models.StatMeta
	if err := c.getMetadata(language, "stats", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetVoidBattlesMetadata retrieves metadata for void battles in the specified language.
// It returns a slice of VoidBattleResponse models and an error if the request fails.
func (c *Client) GetVoidBattlesMetadata(language LanguageCode) ([]models.VoidBattleResponse, error) {
	var resp []models.VoidBattleResponse
	if err := c.getMetadata(language, "voidBattles", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTitlesMetadata retrieves metadata for titles in the specified language.
// It returns a slice of TitleResponse models and an error if the request fails.
func (c *Client) GetTitlesMetadata(language LanguageCode) ([]models.TitleResponse, error) {
	var resp []models.TitleResponse
	if err := c.getMetadata(language, "titles", &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

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
