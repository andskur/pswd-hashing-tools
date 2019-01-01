package cmd

import (
	"fmt"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(compareCmd)
}

// compareCmd can receive password and hash with command line argument
// or u can leave arguments nil and type password later with stdin
var compareCmd = &cobra.Command{
	Use:       "compare [password] ['hash']",
	Short:     "Compare string with a hash",
	ValidArgs: []string{"password", "hash"},
	Args:      cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Arguments["password"] = args[0]
			if len(args) == 2 {
				Arguments["hash"] = args[1]
			}
		}
		comparePswdHash(algo, Arguments)
	},
}

// comparePswdHash compare matching with given password and hash
func comparePswdHash(algo algorithms.Algorithm, args map[string]string) {
	password := BindArgument("password", args, "compare")
	hash := BindArgument("hash", args, "compare")
	result := algo.CheckHash(password, hash)

	switch result {
	case true:
		fmt.Println("Hash and password are matching")
	case false:
		fmt.Println("Hash and password aren't matching")
	}
}
