package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms/bcrypt"
)

func init() {
	rootCmd.AddCommand(compareCmd)
}

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare string with a hash",
	Run: func(cmd *cobra.Command, args []string) {
		compareStrHash(algo)
	},
}

func compareStrHash(algo bcrypt.Bcrypt) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter password to compare:")
	password, _ := reader.ReadString('\n')

	fmt.Println("Enter hash to compare:")
	hash, _ := reader.ReadString('\n')

	result := algo.CheckHash(password, hash)

	switch result {
	case true:
		fmt.Println("Hash and password are matching")
	case false:
		fmt.Println("Hash and password are't matching")
	}
}
