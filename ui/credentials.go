//+build !debug

package ui

import (
	"bufio"
	"os"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"strings"
)

func GetUserCredentials() (*bufio.Reader, string, string){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
	password := string(bytePassword)

	return reader, strings.TrimSpace(username), strings.TrimSpace(password)
}
