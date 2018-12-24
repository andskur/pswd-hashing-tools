package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/argon2"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/bcrypt"
)

var (
	AlgoFlag string
	algo     algorithms.Algorithm
)

//TODO add viper package for bindings command line flags to config

var rootCmd = &cobra.Command{
	Short: "Tools for hashing passwords and compare result with string",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch AlgoFlag {
		case "argon2":
			algo = &argon2.Argon2{}
		case "bcrypt":
			algo = &bcrypt.Bcrypt{}
		default:
			algo = &bcrypt.Bcrypt{}
		}
		fmt.Printf("Using %s hashing algorithm \n", strings.Title(AlgoFlag))
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&AlgoFlag, "algorithm", "a", "bcrypt", "Crypto algorithm to use")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
