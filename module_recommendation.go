package tfd

import (
	"net/http"
	"net/url"

	"github.com/DevZoneLabs/go-tfd-client/models"
)

// Period represents the time period for the data.
type Period string

const (
	// SevenDays represents the last 7 days period.
	SevenDays Period = "0"

	// ThirtyDays represents the last 30 days period.
	ThirtyDays Period = "1"

	// AllTime represents all-time period.
	AllTime Period = "2"
)

// Periods provides a set of predefined Period constants.
var Periods = struct {
	SevenDays  Period
	ThirtyDays Period
	AllTime    Period
}{
	SevenDays:  SevenDays,
	ThirtyDays: ThirtyDays,
	AllTime:    AllTime,
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
