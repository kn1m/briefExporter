//+build debug

package ui

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func GetUserCredentials() (*bufio.Reader, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	password, _ := reader.ReadString('\n')

	return reader, strings.TrimSpace(username), strings.TrimSpace(password)
}
