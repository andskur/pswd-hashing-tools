package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
)

// Algorithm flag vars
var (
	AlgoFlag string
	algo     algorithms.Algorithm
)

// Arguments fot interacting with commands
var Arguments = make(map[string]string, 2)

//TODO add viper package for bindings command line flags to config

//TODO improve help command with actual documentation

// rootCmd is a root command with general "algorithm" command line flag
// with which can set execute hashing algorithm
var rootCmd = &cobra.Command{
	Short: "Tools for hashing passwords and compare result with string",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Get password hashing algorithm from command line flag and set algorithm to use
		algo, AlgoFlag = algorithms.SetAlgorithm(AlgoFlag)
		fmt.Printf("Using %q hashing algorithm \n", strings.Title(AlgoFlag))
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
func BindArgument(check string, arguments map[string]string, cmd string) (argument string) {
	argument, exist := arguments[check]

	// Ask user for type argument if we don't receive it with command line
	if !exist {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter %s for %s:\n", check, cmd)
		argument, _ = reader.ReadString('\n')
	}
	return argument
}
