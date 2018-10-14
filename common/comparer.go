package common

import (
	"crypto/md5"
	"log"
)

func GetFileChecksum(byte []byte) [16]byte {
	hash := md5.Sum(byte)

	log.Printf("Current file hash: %x\n", hash)
	return hash
}