package keys

import (
	"fmt"
	generate "xprnetwork/cli/commands/keys/commands"
)

const CMD_NAME = "keys"

func Parse(args []string) {
	cmd := args[1]
	tail := args[2:]
	fmt.Println("subcommand", cmd)
	switch cmd {
	case generate.CMD_NAME:
		generate.Parse(tail)
	}

}
