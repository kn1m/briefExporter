package net

import (
	"briefExporter/common"
	"briefExporter/configuration"
	"briefExporter/ui"
	"encoding/json"
	"io/ioutil"
)

const authorizationHeaderName = "Authorization"

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	GrantType string `json:"grant_type"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func GetToken(configuration *configuration.Config, user *ui.User) (*string, error) {

	resp, err := executeRequest(configuration.TokenRetrieveUrl, "POST", nil, nil)
	common.Check(err)

	var tokenResponse *TokenResponse

	body, _ := ioutil.ReadAll(resp.Body)
	logResponse(resp, body)

	err = json.Unmarshal(body, &tokenResponse)

	return &tokenResponse.Token, err
}

func getAuthorizationHeaders(initialHeaders map[string]string, token *string) map[string]string {
	if initialHeaders != nil {
		initialHeaders[authorizationHeaderName] = *token

		return initialHeaders
	}

	return map[string]string{
		authorizationHeaderName: *token,
	}
}
