package configuration

import (
	"briefExporter/common"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	NotesRetrieveUrl      string `json:"retrieve_url"`
	NotesSendUrl          string `json:"send_url"`
	LibraryCheckUrl       string `json:"library_check_url"`
	LibrarySyncUrl        string `json:"library_sync_url"`
	ScanFolder            string `json:"scan_folder"`
	ScanMountPathScript   string `json:"scan_mount_path_script"`
	DeviceAvailabilityUrl string `json:"device_availability_url"`
	CreateUserDeviceUrl   string `json:"create_user_device_url"`
	TokenRetrieveUrl      string `json:"token_retrieve_url"`
	PathToLocalDb         string `json:"path_to_local_db"`
}

func GetConfig(path string) (*Config, error) {
	var config *Config

	data, err := ioutil.ReadFile(path)
	common.Check(err)

	err = json.Unmarshal(data, &config)

	return config, err
}
