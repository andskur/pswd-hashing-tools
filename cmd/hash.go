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

// hashCmd can receive password to hash with command line argument
// or u can leave argument nil and type password later with stdin
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

// strToHash hash given password string with specific algorithm
func strToHash(algo algorithms.Algorithm, password string) {

	// Ask user for type password if we don't receive it with command line argument
	if password == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter password to hash:")
		password, _ = reader.ReadString('\n')
	}

	hash := algo.DoHash(password)
	fmt.Println(hash)
}
