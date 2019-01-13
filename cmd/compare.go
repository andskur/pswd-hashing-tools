package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash/sha2"
)

// compareCmd can receive password and hash with command line argument
// or u can leave arguments nil and type password later with stdin
var compareCmd = &cobra.Command{
	Use:       "compare [password] ['hash']",
	Short:     "Compare given string with a given hash",
	ValidArgs: []string{"password", "hash"},
	Args:      cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Arguments["password"] = args[0]
			if len(args) == 2 {
				Arguments["hash"] = args[1]
			}
		}
		comparePswdHash(pswdAlgo, Arguments)
	},
}

// comparePswdHash compare matching with given password and hash
func comparePswdHash(algo hash_passwords.PaswordHasher, args map[string]string) {
	password := BindArgument("password", args, "compare")

	if PreHashFlag {
		hasher := &sha2.Sha2{}
		password = hasher.DoHash(password)
	}

	hash := BindArgument("hash", args, "compare")
	result := algo.CheckHash(password, hash)

	switch result {
	case true:
		fmt.Println("Hash and password are matching")
	case false:
		fmt.Println("Hash and password aren't matching")
	}
}
