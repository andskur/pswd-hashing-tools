package cmd

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

// Prehash flag
var PreHashFlag bool

// Arguments fot interacting with commands
var Arguments = make(map[string]string, 2)

//TODO add viper package for bindings command line flags to config

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
	rootCmd.PersistentFlags().BoolVarP(&PreHashFlag, "prehash", "p", false, "Enable prehash SHA256 function")
	rootCmd.SetHelpTemplate(helpTemplate)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
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

// Prehash given string with sha256
func Prehash(string string) (hash string) {
	fmt.Println("Prehashing password with SHA256...")
	h := sha256.New()
	h.Write([]byte(string))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
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
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}

`
