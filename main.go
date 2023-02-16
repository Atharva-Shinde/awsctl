package main

import (
	"main.go/authenticator"
	"main.go/cmd"
)

func main() {
	cmd.Execute()

	// TODO: better way to avoid import cycle
	authenticator.GetCredentials()
}
