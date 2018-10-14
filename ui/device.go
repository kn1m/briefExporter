package ui

import (
	"bufio"
	"fmt"
	"brief/briefExporter/common"
)

type Device struct {
	Manufacturer string
	Model string
	Classifier string
}

func GetDeviceToConnect(config *common.Config, reader *bufio.Reader) *Device {
	fmt.Print("Enter manufacturer: ")
	manufacturer, _ := reader.ReadString('\n')

	fmt.Print("Enter model: ")
	model, _ := reader.ReadString('\n')

	fmt.Print("Enter model classifier: ")
	classifier, _ := reader.ReadString('\n')

	device := &Device { Manufacturer: manufacturer, Model: model, Classifier: classifier }

	return device
}
