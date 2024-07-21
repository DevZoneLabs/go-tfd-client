package tfd

import (
	"net/http"
	"net/url"

	"github.com/DevZone-Labs/tfd-api/models"
)

const userBasePath string = "/tfd/v1"

// Retrieves the account identifier (OUID).
func (c *Client) GetAccountIdentifier(userName string) (*models.AccountIdentifier, error) {
	query := url.Values{"user_name": {userName}}

	req, err := c.newRequest(http.MethodGet, true, userBasePath+"/id", &query)
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
	query := url.Values{"ouid": {ouid}}

	req, err := c.newRequest(http.MethodGet, true, userBasePath+"/user/basic", &query)
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
	query := url.Values{"ouid": {ouid}}

	req, err := c.newRequest(http.MethodGet, true, userBasePath+"/user/descendant", &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserDescendant)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserWeapons(ouid string, language LanguageCode) (*models.UserWeapon, error) {
	query := url.Values{
		"language_code": {language.String()},
		"ouid":          {ouid},
	}

	req, err := c.newRequest(http.MethodGet, true, userBasePath+"/user/weapon", &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserWeapon)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetUserReactor(ouid string, language LanguageCode) (*models.UserReactor, error) {
	query := url.Values{
		"language_code": {language.String()},
		"ouid":          {ouid},
	}

	req, err := c.newRequest(http.MethodGet, true, userBasePath+"/user/reactor", &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserReactor)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetExternalComponent(ouid string, language LanguageCode) (*models.UserExternalComponent, error) {
	query := url.Values{
		"language_code": {language.String()},
		"ouid":          {ouid},
	}

	req, err := c.newRequest(http.MethodGet, true, userBasePath+"/user/external-component", &query)
	if err != nil {
		return nil, err
	}

	resp := new(models.UserExternalComponent)
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
