package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/argon2"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/bcrypt"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/scrypt"
)

// Algorithm flag vars
var (
	AlgoFlag string
	algo     algorithms.Algorithm
)

// Arguments fot interacting with commands
var Arguments = make(map[string]string, 2)

//TODO add viper package for bindings command line flags to config

// rootCmd is a root command with general "algorithm" command line flag
// with which can set execute hashing algorithm
var rootCmd = &cobra.Command{
	Short: "Tools for hashing passwords and compare result with string",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch AlgoFlag {
		case "argon2":
			algo = &argon2.Argon2{}
		case "bcrypt":
			algo = &bcrypt.Bcrypt{}
		case "scrypt":
			algo = &scrypt.Scrypt{}
		default:
			algo = &bcrypt.Bcrypt{}
		}
		fmt.Printf("Using %s hashing algorithm \n", strings.Title(AlgoFlag))
	},
}

// Execute root command and binding flags
func Execute() {
	rootCmd.PersistentFlags().StringVarP(&AlgoFlag, "algorithm", "a", "bcrypt", "Crypto algorithm to use")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// BindArgument binding argument dor use in next command iteration
// from user stdin or command line argument
func BindArgument(check string, arguments map[string]string) (argument string) {
	argument, exist := arguments[check]

	// Ask user for type argument if we don't receive it with command line
	if !exist {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter password to compare:")
		argument, _ = reader.ReadString('\n')
	}
	return
}
