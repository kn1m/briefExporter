package net

import (
	"briefExporter/common"
	"briefExporter/configuration"
	"bytes"
	"io/ioutil"
)

func SendNotesToServer(notes *[]byte, config *configuration.Config, token *string) {
	headers := make(map[string]string)
	headers["Set-Type"] = "All"
	headers["Content-Type"] = "application/json"

	resp, err := executeRequest(config.NotesSendUrl, "POST", bytes.NewBuffer(*notes),
		getAuthorizationHeaders(headers, token))
	common.Check(err)

	body, _ := ioutil.ReadAll(resp.Body)

	logResponse(resp, body)
}
