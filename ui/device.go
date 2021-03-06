package ui

import (
	"briefExporter/configuration"
	"bufio"
	"fmt"
)

func GetDeviceToConnect(config *configuration.Config, reader *bufio.Reader) *Device {
	fmt.Print("Enter manufacturer: ")
	manufacturer, _ := reader.ReadString('\n')

	fmt.Print("Enter model: ")
	model, _ := reader.ReadString('\n')

	fmt.Print("Enter model classifier: ")
	classifier, _ := reader.ReadString('\n')

	device := &Device{Manufacturer: manufacturer, Model: model, Classifier: classifier}

	return device
}
