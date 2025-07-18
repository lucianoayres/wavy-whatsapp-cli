package main

import (
	"fmt"
	"os"

	"whatsmeow-go/cmd/wavy/common"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wavy",
	Short: "WhatsApp CLI client",
	Long:  `A command line interface to interact with WhatsApp.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Display the version of Wavy WhatsApp CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Wavy WhatsApp CLI v%s\n", common.GetVersion())
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(groupsCmd)
	rootCmd.AddCommand(versionCmd)
}
