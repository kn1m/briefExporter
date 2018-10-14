package connectivity

import (
	"github.com/gotmc/libusb"
	"brief/briefExporter/common"
	"log"
	"os/exec"
)

type Connector interface {
	GetNotesFromDevice(serialNumber string, config *common.Config) (string, error)
}

type DeviceClassifier interface {
	GetAllDevicesSerials(manufacturer string, productName string) ([]string, error)
}

func getConnectedDevices() (*libusb.Context, []*libusb.Device, error) {
	ctx, err := libusb.NewContext()
	if err != nil {
		log.Fatal("Couldn't create USB context. Ending now.")
	}

	//defer ctx.Close()
	devices, err := ctx.GetDeviceList()
	if err != nil {
		log.Fatalf("Couldn't get devices")
	}

	log.Printf("Found %v USB devices.\n", len(devices))

	return ctx, devices, err
}

func verifyDeviceManufacturerAndProduct(manufacturerName string, productName string,
										handle *libusb.DeviceHandle, descriptor *libusb.DeviceDescriptor) bool {
	manufacturerConfirmed := false
	productConfirmed := false

	manufacturer, err := handle.GetStringDescriptorASCII(descriptor.ManufacturerIndex)
	if err == nil && manufacturer == manufacturerName {
		manufacturerConfirmed = true
	}
	product, err := handle.GetStringDescriptorASCII(descriptor.ProductIndex)
	if err == nil && product == productName {
		productConfirmed = true
	}

	return manufacturerConfirmed && productConfirmed
}

func GetAllCompatibleDevicesSerials(manufacturerName string, productName string) ([]*string, error){
	var serials []*string

	ctx, devices, err := getConnectedDevices()
	if err != nil {
		log.Fatalf("Aborting device verification")
		return nil, err
	}

	defer ctx.Close()

	for _, device := range devices {
		usbDeviceDescriptor, err := device.GetDeviceDescriptor()
		if err != nil {
			log.Printf("Error getting device descriptor: %s", err)
			continue
		}
		handle, err := device.Open()
		if err != nil {
			log.Printf("Error opening device: %s", err)
			continue
		}
		defer handle.Close()

		typeConfirmed := verifyDeviceManufacturerAndProduct(manufacturerName, productName, handle, usbDeviceDescriptor)

		if typeConfirmed {
			serialNumber, err := handle.GetStringDescriptorASCII(usbDeviceDescriptor.SerialNumberIndex)
			if err == nil {
				serials = append(serials, &serialNumber)
			}
		}
	}
	return serials, err
}

func verifyDevice(manufacturerName string, productName string, serialNumberToCheck string) bool {

	ctx, devices, err := getConnectedDevices()
	if err != nil {
		log.Fatalf("Aborting device verification")
		return false
	}

	defer ctx.Close()

	for _, device := range devices {
		usbDeviceDescriptor, err := device.GetDeviceDescriptor()
		if err != nil {
			log.Printf("Error getting device descriptor: %s", err)
			continue
		}
		handle, err := device.Open()
		if err != nil {
			log.Printf("Error opening device: %s", err)
			continue
		}
		defer handle.Close()

		typeConfirmed := verifyDeviceManufacturerAndProduct(manufacturerName, productName, handle, usbDeviceDescriptor)
		serialNumber, err := handle.GetStringDescriptorASCII(usbDeviceDescriptor.SerialNumberIndex)

		if err == nil && typeConfirmed && serialNumber == serialNumberToCheck {
			return true
		}
	}
	return false
}


func getDeviceMountPath(serialNumber string, config *common.Config) (string, error) {
	mountPath, err := exec.Command("sh", config.ScanMountPathScript).Output()
	return string(mountPath), err
}