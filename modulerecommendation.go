package tfd

import (
	"net/http"
	"net/url"

	"github.com/DevZoneLabs/go-tfd-api/models"
)

type ModuleRecommendationOpts struct {
	DescendantID string
	WeaponID     string
	VoidBattleID string
	Period       Period
}

type Period string

var Periods = struct {
	SevenDays  Period // Last 7 days
	ThirtyDays Period // Last 30 days
	AllTime    Period // All Time
}{"0", "1", "2"}

// Retrieves descendant metadata
func (c *Client) GetModuleRecommendation(opts ModuleRecommendationOpts) (*models.ModuleRecommendation, error) {
	query := url.Values{
		"descendant_id":  {opts.DescendantID},
		"weapon_id":      {opts.WeaponID},
		"void_battle_id": {opts.VoidBattleID},
		"period":         {string(opts.Period)}}

	req, err := c.newRequest(http.MethodGet, true, "tfd/v1/recommendation/module", &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.ModuleRecommendation)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
