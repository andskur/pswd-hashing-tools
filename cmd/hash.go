package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms/bcrypt"
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

func strToHash(algo bcrypt.Bcrypt) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter password to hash:")
	password, _ := reader.ReadString('\n')

	hash, err := algo.DoHash(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)
}
