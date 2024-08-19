package sample3

import (
	"fmt"
	"strings"
)

func Validate() {
	email := ""
	fmt.Println("enter your email:")
	fmt.Scanln(&email)
	if strings.Contains(email, ".com") {
		fmt.Println("VALID EMAIL")
	} else {
		fmt.Println("inValid email")
	}
}
