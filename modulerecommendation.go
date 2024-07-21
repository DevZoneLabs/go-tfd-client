package tfd

import (
	"net/http"

	"github.com/DevZoneLabs/go-tfd-api/models"
)

// Retrieves descendant metadata
func (c *Client) GetModuleRecommendation() (*models.ModuleRecommendationResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, "tfd/v1/recommendation/module", nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.ModuleRecommendationResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
