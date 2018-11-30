package main

import (
	"briefExporter/common"
	"briefExporter/configuration"
	"briefExporter/connectivity"
	"briefExporter/exporters"
	"briefExporter/libsync"
	"briefExporter/net"
	"briefExporter/ui"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	reader, user := ui.GetUserCredentials()

	var dataPath string
	flag.StringVar(&dataPath, "data_path", "", "path to data file")

	var configPath string
	flag.StringVar(&configPath, "config_path", "", "path to config file")

	logFlag := flag.Bool("log", false, "true if provide logs in output")
	flag.Parse()

	if dataPath == "" || configPath == "" {
		log.Fatalln("configPath and dataPath should been provided!")
		os.Exit(1)
	}

	var mem runtime.MemStats

	if *logFlag {
		runtime.ReadMemStats(&mem)
		log.Println(mem.Alloc)
		log.Println(mem.TotalAlloc)
		log.Println(mem.HeapAlloc)
		log.Println(mem.HeapSys)
	}

	config, err := configuration.GetConfig(configPath)
	common.Check(err)
	log.Println(config)

	device := ui.GetDeviceToConnect(config, reader)

	token, err := net.GetToken(config, user)
	common.Check(err)

	if net.CheckDeviceAvailability(device, config, token) {
		log.Fatalln("configPath and dataPath should been provided!")
		os.Exit(1)
	}

	kindleUsb := &connectivity.KindleUsbConnector{}

	devices, err := connectivity.GetAllCompatibleDevicesSerials("Amazon", "Amazon Kindle")
	for i := range devices {
		log.Printf("Found compatible device with serial: %s\n", *devices[i])
	}

	fmt.Println("Enter desired serial:")
	desiredSerial, _ := reader.ReadString('\n')

	var wg sync.WaitGroup
	wg.Add(2)

	var notes []*exporters.NoteRecord
	var mountPath string

	go func() {

		defer wg.Done()

		mountPath, err := kindleUsb.GetNotesFromDevice(strings.TrimRight(desiredSerial, "\n"), config)

		var matcher exporters.KindleExporter

		notes, err = matcher.GetNotes(mountPath + connectivity.DefaultKindleNotesFilePath)
		common.Check(err)
	}()

	go listStructure(&wg, mountPath)

	wg.Wait()

	for i := range notes {
		log.Printf("\n%d: %s %s %+v p: %d-%d l:%d-%d :: %s :: %s %s", i, notes[i].BookTitle,
			notes[i].BookOriginalName, notes[i].BookAuthor, notes[i].FirstPage, notes[i].SecondPage,
			notes[i].FirstLocation, notes[i].SecondLocation, notes[i].NoteTitle, notes[i].NoteText, notes[i].CreatedOn)
	}

	if *logFlag {
		runtime.ReadMemStats(&mem)
		log.Println(mem.Alloc)
		log.Println(mem.TotalAlloc)
		log.Println(mem.HeapAlloc)
		log.Println(mem.HeapSys)
	}
}

func listStructure(group *sync.WaitGroup, mountPath string) {

	defer group.Done()

	libDir := &libsync.Directory{Path: mountPath}
	libsync.CheckPath(libDir)
	libDir.PrintStructure(nil)
}
