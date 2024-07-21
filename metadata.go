package tfd

import (
	"net/http"

	"github.com/DevZone-Labs/tfd-api/models"
)

const metaBasePath string = "/static/tfd/meta/"

// Retrieves descendant metadata
func (c *Client) GetDescendantMetadata(language LanguageCode) (*models.DescendantResponse, error) {
	req, err := c.newRequest(http.MethodGet, true, metaBasePath+string(language), nil)
	if err != nil {
		return nil, err
	}

	resp := new(models.DescendantResponse)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
