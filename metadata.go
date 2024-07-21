package tfd

import (
	"net/http"

	"github.com/DevZone-Labs/tfd-api/models"
)

const metaBasePath string = "/static/tfd/meta/"

// Retrieves descendant metadata
func (c *Client) GetDescendantsMetadata(language LanguageCode) (*models.DescendantResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/descendant.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.DescendantResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetWeaponsMetadata(language LanguageCode) (*models.WeaponResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/weapon.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.WeaponResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetModulesMetadata(language LanguageCode) (*models.ModuleResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/module.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.ModuleResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetReactorsMetadata(language LanguageCode) (*models.ReactorResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/reactor.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.ReactorResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetExternalComponentsMetadata(language LanguageCode) (*models.ExternalComponentResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/external-component.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.ExternalComponentResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetRewardsMetadata(language LanguageCode) (*models.RewardResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/reward.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.RewardResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetStatsMetadata(language LanguageCode) (*models.BaseStat, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/stat.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.BaseStat)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetVoidBattlesMetadata(language LanguageCode) (*models.VoidBattleResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/void-battle.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.VoidBattleResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetTitlesMetadata(language LanguageCode) (*models.TitleResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+language.String()+"/title.json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.TitleResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
