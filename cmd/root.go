package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms/argon2"
)

//var algo = bcrypt.Bcrypt{}
var algo = argon2.Argon2{}

var rootCmd = &cobra.Command{
	Short: "Tools for hashing passwords and compare result with string",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
