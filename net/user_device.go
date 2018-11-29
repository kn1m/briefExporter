package net

import (
	"briefExporter/common"
	"briefExporter/configuration"
	"briefExporter/ui"
	"bytes"
	"encoding/json"
	"net/http"
)

type UserDevice struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Serial string `json:"Serial"`
}

func CheckDeviceAvailability(device *ui.Device, config *configuration.Config, token *string) bool {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	deviceJson, err := json.Marshal(*device)
	common.Check(err)

	resp, err := executeRequest(config.DeviceAvailabilityUrl, "GET", bytes.NewBuffer(deviceJson),
		getAuthorizationHeaders(nil, token))
	common.Check(err)

	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func CreateUserDevice(deviceId *string, deviceName *string, deviceSerial *string,
	config *configuration.Config, token *string) bool {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	userDevice := map[string]interface{}{"DeviceId": *deviceId, "DeviceName": *deviceName, "DeviceSerial": *deviceSerial}

	userDeviceJson, err := json.Marshal(userDevice)
	common.Check(err)

	resp, err := executeRequest(config.CreateUserDeviceUrl, "POST", bytes.NewBuffer(userDeviceJson),
		getAuthorizationHeaders(headers, token))
	common.Check(err)

	if resp.StatusCode == http.StatusCreated {
		return true
	}

	return false
}
