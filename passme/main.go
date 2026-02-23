package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	const (
		jackName               = "jack"
		jackPassword           = "1888"
		inancName              = "inanc"
		inancPassword          = "1879"
		invalidUsernameMessage = "Access denied for %q\n"
		invalidPasswordMessage = "Invalid Password for %q\n"
		accessGrantMessage     = "Access granted to %q\n"
	)

	username := os.Args[1]
	password := os.Args[2]

	if username == jackName {
		if password == jackPassword {
			fmt.Printf(accessGrantMessage, jackName)
		} else {
			fmt.Printf(invalidPasswordMessage, jackName)
		}
	} else if username == inancName {
		if password == inancPassword {
			fmt.Printf(accessGrantMessage, inancName)
		} else {
			fmt.Printf(invalidPasswordMessage, inancName)
		}
	} else {
		fmt.Printf(invalidUsernameMessage, username)
	}
}
