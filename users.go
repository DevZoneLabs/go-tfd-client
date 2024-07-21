package tfd

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/DevZone-Labs/tfd-api/models"
)

// Retrieves the account identifier (OUID).
func (c *Client) GetAccountIdentifier(userName string) (*models.AccountIdentifier, error) {

	path := "tfd/v1/id"
	query := url.Values{"user_name": {userName}}

	req, err := c.newRequest(http.MethodGet, true, path, &query)
	if err != nil {
		return nil, err
	}

	fmt.Println(req.URL.String())

	resp := new(models.AccountIdentifier)

	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
