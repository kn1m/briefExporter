package net

import (
	"briefExporter/common"
	"briefExporter/configuration"
	"briefExporter/ui"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"
)

type HistoryRecord struct {
	SerialNumber string    `json:"serial_number"`
	Checksum     [16]byte  `json:"checksum"`
	CreatedOn    time.Time `json:"created_on"`
}

func GetPreviousHistoryRecord(serial string, config *configuration.Config) (*HistoryRecord, error) {
	resp, err := executeRequest(config.NotesRetrieveUrl+"/"+serial, "GET", nil, nil)

	var historyRecord *HistoryRecord

	body, _ := ioutil.ReadAll(resp.Body)
	logResponse(resp, body)

	err = json.Unmarshal(body, &historyRecord)

	return historyRecord, err
}

func SendNotesToServer(notes *[]byte, config *configuration.Config) {
	headers := make(map[string]string)
	headers["Set-Type"] = "All"
	headers["Content-Type"] = "application/json"

	resp, err := executeRequest(config.NotesSendUrl, "POST", bytes.NewBuffer(*notes), headers)
	common.Check(err)

	body, _ := ioutil.ReadAll(resp.Body)

	logResponse(resp, body)
}

func CheckDeviceAvailability(device *ui.Device, config *configuration.Config) bool {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	deviceJson, _ := json.Marshal(*device)

	resp, err := executeRequest(config.DeviceAvailabilityUrl, "GET", bytes.NewBuffer(deviceJson), headers)
	common.Check(err)

	if resp.StatusCode == 200 {
		return true
	}

	return false
}
