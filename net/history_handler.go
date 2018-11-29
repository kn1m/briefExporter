package net

import (
	"briefExporter/configuration"
	"encoding/json"
	"io/ioutil"
	"time"
)

type HistoryRecord struct {
	SerialNumber string    `json:"serial_number"`
	Checksum     [16]byte  `json:"checksum"`
	CreatedOn    time.Time `json:"created_on"`
}

func GetPreviousHistoryRecord(serial string, config *configuration.Config, token *string) (*HistoryRecord, error) {

	resp, err := executeRequest(config.NotesRetrieveUrl+"/"+serial, "GET", nil,
		getAuthorizationHeaders(nil, token))

	var historyRecord *HistoryRecord

	body, _ := ioutil.ReadAll(resp.Body)
	logResponse(resp, body)

	err = json.Unmarshal(body, &historyRecord)

	return historyRecord, err
}
