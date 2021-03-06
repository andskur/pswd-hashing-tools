package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords"
)

// hashCmd can receive password to hash with command line argument
// or u can leave argument nil and type password later with stdin
var hashCmd = &cobra.Command{
	Use:       "hash [password]",
	Short:     "Create hash from given string",
	ValidArgs: []string{"password"},
	Args:      cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Arguments["password"] = args[0]
		}
		strToHash(pswdAlgo, Arguments)
	},
}

// strToHash hash given password string with specific algorithm
func strToHash(algo hash_passwords.PaswordHasher, args map[string]string) {
	password := BindArgument("password", args, "hash")

	if PreHashFlag != "" {
		fmt.Printf("Prehashing password with %q...\n", strings.ToUpper(PreHashFlag))
		password = prehashAlgo.DoHash(password)
		fmt.Println("Prehash:", password)
	}

	hash := algo.DoHash(password)
	fmt.Println("Result:", hash)
}
