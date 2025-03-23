package contract

import (
	"fmt"
	"xprnetwork/cli/commands/contract/gentypes"
)

const CMD_NAME = "contract"

func Parse(args []string) {
	cmd := args[1]
	tail := args[2:]
	fmt.Println("subcommand", cmd)
	switch cmd {
	case gentypes.CMD_NAME:
		gentypes.Parse(tail)
	}

}
