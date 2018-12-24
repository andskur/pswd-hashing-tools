package cmd

import (
	"github.com/spf13/cobra"
	"log"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms/bcrypt"
)

var algo = bcrypt.Bcrypt{}

var rootCmd = &cobra.Command{
	Short: "Tools for hashing passwords and compare result with string",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
