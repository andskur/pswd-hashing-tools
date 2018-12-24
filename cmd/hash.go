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
	Use:   "hash",
	Short: "Hash given string with specific algorithm",
	Run: func(cmd *cobra.Command, args []string) {
		strToHash(algo)
	},
}

func strToHash(algo algorithms.Algorithm) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter password to hash:")
	password, _ := reader.ReadString('\n')

	hash := algo.DoHash(password)
	fmt.Println(hash)
}
