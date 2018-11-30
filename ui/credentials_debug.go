//+build debug

package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserCredentials() (*bufio.Reader, *User) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	password, _ := reader.ReadString('\n')

	return reader, &User{strings.TrimSpace(username), strings.TrimSpace(password)}
}
