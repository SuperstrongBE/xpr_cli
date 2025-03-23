package main

import (
	"fmt"
	"os"
	"xprnetwork/cli/commands/contract"
	"xprnetwork/cli/commands/keys"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Subcommands expected")
		os.Exit(1)
	}

	switch os.Args[1] {

	case contract.CMD_NAME:
		contract.Parse(os.Args[1:])
	case keys.CMD_NAME:
		keys.Parse(os.Args[1:])
	default:
		fmt.Println("Sub command expected")
		os.Exit(1)
	}
}
