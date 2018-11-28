package net

import (
	"briefExporter/common"
	"briefExporter/configuration"
	"encoding/json"
	"io/ioutil"
)

const authorizationHeaderName = "Authorization"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func GetToken(configuration *configuration.Config, user *User) (*string, error) {

	resp, err := executeRequest(configuration.TokenRetrieveUrl, "POST", nil, nil)
	common.Check(err)

	var tokenResponse *TokenResponse

	body, _ := ioutil.ReadAll(resp.Body)
	logResponse(resp, body)

	err = json.Unmarshal(body, &tokenResponse)

	return &tokenResponse.Token, err
}

func GetAuthorizationHeaders(initialHeaders map[string]string, token *string) map[string]string {
	if initialHeaders != nil {
		initialHeaders[authorizationHeaderName] = *token

		return initialHeaders
	}

	return map[string]string{
		authorizationHeaderName: *token,
	}
}
