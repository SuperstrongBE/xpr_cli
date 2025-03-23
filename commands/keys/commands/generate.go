package generate

import (
	"fmt"
	keyutils "xprnetwork/cli/core/keys"
)

const CMD_NAME = "generate"

func Parse(args []string) {
	mnemonic, private, public := keyutils.GenerateKeys()
	fmt.Println(mnemonic)
	fmt.Println(private)
	fmt.Println(public)

}
