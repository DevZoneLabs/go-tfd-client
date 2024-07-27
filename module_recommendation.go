package tfd

import (
	"net/http"
	"net/url"

	"github.com/DevZoneLabs/go-tfd-client/v2/models"
)

// Period represents the time period for the data.
type Period string

// Periods provides a set of predefined Period constants.
var Periods = struct {
	// Last 7 days period.
	SevenDays Period

	//  Last 30 days period.
	ThirtyDays Period

	// All-time period.
	AllTime Period
}{
	SevenDays:  "0",
	ThirtyDays: "1",
	AllTime:    "2",
}

type ModuleRecommendationOpts struct {
	DescendantID string
	WeaponID     string
	VoidBattleID string
	Period       Period
}

// GetModuleRecommendation retrieves module recommendations based on the given options.
func (c *Client) GetModuleRecommendation(opts ModuleRecommendationOpts) (*models.ModuleRecommendation, error) {
	query := url.Values{
		"descendant_id":  {opts.DescendantID},
		"weapon_id":      {opts.WeaponID},
		"void_battle_id": {opts.VoidBattleID},
		"period":         {string(opts.Period)},
	}

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
