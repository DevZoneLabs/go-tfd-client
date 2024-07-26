package tfd

import (
	"net/http"
	"net/url"

	"github.com/DevZoneLabs/go-tfd-client/models"
)

const (
	userBasePath              = "/tfd/v1"
	accountIdentifierPath     = userBasePath + "/id"
	userBasicPath             = userBasePath + "/user/basic"
	userDescendantPath        = userBasePath + "/user/descendant"
	userWeaponsPath           = userBasePath + "/user/weapon"
	userReactorPath           = userBasePath + "/user/reactor"
	userExternalComponentPath = userBasePath + "/user/external-component"
)

// GetAccountIdentifier retrieves the account identifier (OUID) for a given username.
func (c *Client) GetAccountIdentifier(userName string) (*models.AccountIdentifier, error) {
	query := url.Values{"user_name": {userName}}
	result, err := c.getUserData(accountIdentifierPath, query, &models.AccountIdentifier{})
	if err != nil {
		return nil, err
	}
	return result.(*models.AccountIdentifier), nil
}

// GetBasicInformation retrieves basic information about the user identified by the given OUID.
func (c *Client) GetBasicInformation(ouid string) (*models.UserBasic, error) {
	query := url.Values{"ouid": {ouid}}
	result, err := c.getUserData(userBasicPath, query, &models.UserBasic{})
	if err != nil {
		return nil, err
	}
	return result.(*models.UserBasic), nil
}

// GetUserDescendant retrieves descendant information for the user identified by the given OUID.
func (c *Client) GetUserDescendant(ouid string) (*models.UserDescendant, error) {
	query := url.Values{"ouid": {ouid}}
	result, err := c.getUserData(userDescendantPath, query, &models.UserDescendant{})
	if err != nil {
		return nil, err
	}
	return result.(*models.UserDescendant), nil
}

// GetUserWeapons retrieves weapon information for the user identified by the given OUID and language.
func (c *Client) GetUserWeapons(ouid string, language LanguageCode) (*models.UserWeapon, error) {
	query := url.Values{
		"language_code": {string(language)},
		"ouid":          {ouid},
	}
	result, err := c.getUserData(userWeaponsPath, query, &models.UserWeapon{})
	if err != nil {
		return nil, err
	}
	return result.(*models.UserWeapon), nil
}

// GetUserReactor retrieves reactor information for the user identified by the given OUID and language.
func (c *Client) GetUserReactor(ouid string, language LanguageCode) (*models.UserReactor, error) {
	query := url.Values{
		"language_code": {string(language)},
		"ouid":          {ouid},
	}
	result, err := c.getUserData(userReactorPath, query, &models.UserReactor{})
	if err != nil {
		return nil, err
	}
	return result.(*models.UserReactor), nil
}

// GetExternalComponent retrieves external component information for the user identified by the given OUID and language.
func (c *Client) GetExternalComponent(ouid string, language LanguageCode) (*models.UserExternalComponent, error) {
	query := url.Values{
		"language_code": {string(language)},
		"ouid":          {ouid},
	}
	result, err := c.getUserData(userExternalComponentPath, query, &models.UserExternalComponent{})
	if err != nil {
		return nil, err
	}
	return result.(*models.UserExternalComponent), nil
}

// getUserData is a helper method that sends a request to the given path with the specified query parameters
// and unmarshals the response into the provided response model.
func (c *Client) getUserData(path string, query url.Values, respModel interface{}) (interface{}, error) {
	req, err := c.newRequest(http.MethodGet, true, path, &query)
	if err != nil {
		return nil, err
	}

	if err := c.do(req, respModel); err != nil {
		return nil, err
	}

	return respModel, nil
}