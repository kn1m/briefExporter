package net

import (
	"briefExporter/common"
	"briefExporter/configuration"
	"errors"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
}

func GetToken(configuration *configuration.Config, user *User) (string, error) {

	tokenData, err := executeRequest(configuration.TokenRetrieveUrl, "POST", nil, nil)
	common.Check(err)

	return tokenData, nil
}

func ApplyAuthorization(request *http.Request) error {
	if request == nil {
		return errors.New("No request provided!")
	}

	return nil
}
