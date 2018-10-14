package connectivity

import (
	"log"
	"brief/briefExporter/common"
)

const (
	DefaultKindleNotesFilePath = "/documents/My Clippings.txt"

	manufacturerName = "Amazon"
	productName = "Amazon Kindle"
)

type KindleUsbConnector struct {}

func (c *KindleUsbConnector) GetNotesFromDevice(serialNumber string, config *common.Config) (string, error) {
	log.Println(serialNumber)
	deviceVerified := verifyDevice(manufacturerName, productName, serialNumber)
	if deviceVerified {
		mountPath, err := getDeviceMountPath(serialNumber, config)
		if err == nil {
			log.Printf("Mount path of device %s with serial number %s : %s", productName, serialNumber, mountPath)
			return mountPath, err
		}
	}

	return "", nil
}
