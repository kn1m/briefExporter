package net

import "briefExporter/common"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetToken(configuration *common.Config, user *User) (string, error) {
	return "", nil
}
