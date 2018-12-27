package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
)

func init() {
	rootCmd.AddCommand(compareCmd)
}

var compareCmd = &cobra.Command{
	Use:       "compare [password] ['hash']",
	Short:     "Compare string with a hash",
	ValidArgs: []string{"password", "hash"},
	Args:      cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var password, hash string
		if len(args) > 0 {
			password = args[0]
			if len(args) == 2 {
				hash = args[1]
			}
		}

		fmt.Println(args)
		fmt.Println(hash)

		compareStrHash(algo, password, hash)
	},
}

func compareStrHash(algo algorithms.Algorithm, password, hash string) {
	if password == "" || hash == "" {
		reader := bufio.NewReader(os.Stdin)
		if password == "" {
			fmt.Println("Enter password to compare:")
			password, _ = reader.ReadString('\n')
		}
		if hash == "" {
			fmt.Println("Enter hash to compare:")
			hash, _ = reader.ReadString('\n')
		}
	}

	result := algo.CheckHash(password, hash)

	switch result {
	case true:
		fmt.Println("Hash and password are matching")
	case false:
		fmt.Println("Hash and password are't matching")
	}
}
