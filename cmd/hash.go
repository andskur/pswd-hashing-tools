package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
)

func init() {
	rootCmd.AddCommand(hashCmd)
}

var hashCmd = &cobra.Command{
	Use:       "hash [password]",
	Short:     "Hash given string with specific algorithm",
	ValidArgs: []string{"password"},
	Args:      cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var password string
		if len(args) > 0 {
			password = args[0]
		}
		strToHash(algo, password)
	},
}

func strToHash(algo algorithms.Algorithm, password string) {

	if password == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter password to hash:")
		password, _ = reader.ReadString('\n')
	}

	hash := algo.DoHash(password)
	fmt.Println(hash)
}
