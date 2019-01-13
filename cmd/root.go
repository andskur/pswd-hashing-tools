package cmd

import (
	"bufio"
	"fmt"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords"
)

// PaswordHasher flag vars
var (
	AlgoFlag string
	pswdAlgo hash_passwords.PaswordHasher
)

// Prehash flag vars
var (
	PreHashFlag string
	prehashAlgo hash.Hasher
)

// Arguments fot interacting with commands
var Arguments = make(map[string]string, 2)

//TODO add viper package for bindings command line flags to config

// Execute root command and binding flags
func Execute(pswdHahsers, hasher *algorithms.Algorithms) *cobra.Command {
	// rootCmd is a root command with general "algorithm" and "prehash"
	// command line flags with which can set execute hashing algorithm
	var rootCmd = &cobra.Command{
		Short: "Tools for hashing passwords and compare result with string",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			// Get password hashing algorithm from command line flag and set algorithm to use
			pswdHahsers.SetAlgorithm(AlgoFlag)
			pswdAlgo = pswdHahsers.Current
			AlgoFlag = algorithms.GetName(pswdAlgo)

			if PreHashFlag != "" {
				hasher.SetAlgorithm(PreHashFlag)
				prehashAlgo = hasher.Current
				PreHashFlag = algorithms.GetName(prehashAlgo)
			}

			fmt.Printf("Using %q hashing algorithm \n", strings.Title(AlgoFlag))
		},
	}

	// Add flags
	rootCmd.PersistentFlags().StringVarP(&AlgoFlag, "algorithm", "a", "bcrypt", "crypto algorithm to use")
	rootCmd.PersistentFlags().StringVarP(&PreHashFlag, "prehash", "p", "", "prehash algorithm to use")

	// Add help template
	rootCmd.SetHelpTemplate(helpTemplate)

	// Add commands
	rootCmd.AddCommand(hashCmd)
	rootCmd.AddCommand(compareCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	return rootCmd
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

// Help template for all commands
var helpTemplate = `{{ if .Long}}{{.Long}}{{else }}{{.Short}}{{end}}

Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command] [arguments] [flags]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .UseLine .NamePadding }}  - {{.Short}}{{end}}{{end}}

Password and hash arguments are optional, you can type it in stdin after command execution{{end}}

{{if .HasAvailableLocalFlags}}Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}

Available algorithms:
  bcrypt
  scrypt
  argon2
  
Available prehash algorithms:
  SHA2
  SHA3{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
