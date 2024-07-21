package tfd

import (
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

	resp := new(models.AccountIdentifier)

	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetBasicInformation(ouid string) (*models.UserBasic, error) {
	path := "/tfd/v1/user/basic"
	query := url.Values{"ouid": {ouid}}

	req, err := c.newRequest(http.MethodGet, true, path, &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserBasic)

	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserDescendant(ouid string) (*models.UserDescendant, error) {
	path := "/tfd/v1/user/descendant"
	query := url.Values{"ouid": {ouid}}

	req, err := c.newRequest(http.MethodGet, true, path, &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserDescendant)

	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserWeapon(ouid string, language string) (*models.UserWeapon, error) {
	path := "/tfd/v1/user/weapon"
	query := url.Values{"languague_code": {language}, "ouid": {ouid}}

	req, err := c.newRequest(http.MethodGet, true, path, &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserWeapon)

	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserReactor(ouid string, language string) (*models.UserReactor, error) {
	path := "/tfd/v1/user/reactor"
	query := url.Values{"languague_code": {language}, "ouid": {ouid}}

	req, err := c.newRequest(http.MethodGet, true, path, &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserReactor)

	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetExternalComponent(ouid string, language string) (*models.UserExternalComponent, error) {
	path := "/tfd/v1/user/external-component"
	query := url.Values{"languague_code": {language}, "ouid": {ouid}}

	req, err := c.newRequest(http.MethodGet, true, path, &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserExternalComponent)

	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
