package common

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func GetFileData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func Check(err error) {
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func GetSystemPathDelimiter() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "/"
	case "linux":
		return "/"
	default:
		return "\\"
	}
}
