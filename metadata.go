package tfd

import (
	"net/http"

	"github.com/DevZoneLabs/go-tfd-client/models"
)

const metaBasePath string = "/static/tfd/meta/"

// Retrieves descendant metadata
func (c *Client) GetDescendantsMetadata(language LanguageCode) ([]models.Descendant, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/descendant.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.Descendant{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetWeaponsMetadata(language LanguageCode) ([]models.Weapon, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/weapon.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.Weapon{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetModulesMetadata(language LanguageCode) ([]models.Module, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/module.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.Module{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetReactorsMetadata(language LanguageCode) ([]models.Reactor, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/reactor.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.Reactor{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetExternalComponentsMetadata(language LanguageCode) ([]models.ExternalComponent, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/external-component.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.ExternalComponent{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetRewardsMetadata(language LanguageCode) ([]models.Reward, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/reward.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.Reward{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetStatsMetadata(language LanguageCode) ([]models.StatMeta, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/stat.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.StatMeta{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetVoidBattlesMetadata(language LanguageCode) ([]models.VoidBattleResponse, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/void-battle.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.VoidBattleResponse{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetTitlesMetadata(language LanguageCode) ([]models.TitleResponse, error) {
	req, err := c.newRequest(http.MethodGet, false, metaBasePath+language.String()+"/title.json", nil)
	if err != nil {
		return nil, err
	}

	resp := []models.TitleResponse{}
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
