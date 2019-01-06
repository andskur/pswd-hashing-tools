package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
)

func init() {
	rootCmd.AddCommand(hashCmd)
}

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
		strToHash(algo, Arguments)
	},
}

// strToHash hash given password string with specific algorithm
func strToHash(algo algorithms.Algorithm, args map[string]string) {
	password := BindArgument("password", args, "hash")
	hash := algo.DoHash(password)
	fmt.Println("Resulting:", hash)
}
