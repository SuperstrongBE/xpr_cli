package gentypes

import (
	"flag"
	"fmt"
	key "xprnetwork/cli/core/keys"
)

const CMD_NAME = "gentypes"

func Parse(args []string) {
	fmt.Printf("trigger gentypes command")
	if len(args) == 0 {
		panic("Missing contract name")
	}
	cmdArgs := flag.NewFlagSet(CMD_NAME, flag.ExitOnError)
	argContract := args[0]
	argTestnet := cmdArgs.Bool("t", false, "Is contract on testnet")
	argRename := cmdArgs.String("r", "", "Rename contract type")
	cmdArgs.Parse(args[1:])

	fmt.Println("Contract name:", argContract)
	fmt.Println("Contract testnet:", *argTestnet)
	fmt.Println("Rename types:", *argRename)
	mnemonic, private, public := key.GenerateKeys()
	fmt.Println("Mnemonic:", mnemonic)
	fmt.Println("Private:", private)
	fmt.Println("Public:", public)

	//rpc.GetAbi(argContract, *argTestnet)

}
