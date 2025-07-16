package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"go.mau.fi/whatsmeow/types"

	"whatsmeow-go/cmd/wavy/common"
)

var (
	phoneNumber string
)

var checkCmd = &cobra.Command{
	Use:   "check [phoneNumber]",
	Short: "Check if a phone number is on WhatsApp",
	Long:  `Verify if a phone number is registered on WhatsApp.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Handle positional argument if provided
		if len(args) > 0 && phoneNumber == "" {
			phoneNumber = args[0]
		}

		if phoneNumber == "" {
			cmd.Help()
			os.Exit(1)
		}

		runCheck()
	},
}

func init() {
	checkCmd.Flags().StringVarP(&phoneNumber, "phone", "p", "", "Phone number to check")
}

func runCheck() {
	// Create client
	client, needsSetup, err := common.CreateWAClient(true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}

	if needsSetup {
		fmt.Fprintf(os.Stderr, "No WhatsApp session found. Please run 'wavy setup' first.\n")
		os.Exit(1)
	}

	// Connect to WhatsApp
	fmt.Println("Connecting to WhatsApp...")
	err = client.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}

	// Print own ID
	fmt.Printf("Connected as: %s\n", client.Store.ID)
	fmt.Printf("Your phone number: %s\n", client.Store.ID.User)

	// Format the phone number
	formattedNumber := strings.TrimSpace(phoneNumber)
	if strings.HasPrefix(formattedNumber, "+") {
		formattedNumber = formattedNumber[1:]
	}

	fmt.Printf("\nChecking phone number: %s\n", phoneNumber)

	// Create the JID
	jid := types.JID{
		User:   formattedNumber,
		Server: "s.whatsapp.net",
	}

	fmt.Printf("JID for this number: %s\n", jid.String())

	// Check if the user exists on WhatsApp
	fmt.Printf("Checking if %s exists on WhatsApp...\n", formattedNumber)

	exists, err := client.IsOnWhatsApp([]string{formattedNumber})
	if err != nil {
		fmt.Printf("Error checking if user exists: %v\n", err)
	} else {
		for _, user := range exists {
			if user.IsIn {
				fmt.Printf("✅ %s is on WhatsApp (JID: %s)\n", user.Query, user.JID)
			} else {
				fmt.Printf("❌ %s is NOT on WhatsApp\n", user.Query)
			}
		}
	}

	// Show some debugging info
	fmt.Println("\nConnection details:")
	fmt.Printf("Connected: %t\n", client.IsConnected())
	fmt.Printf("LoggedIn: %t\n", client.IsLoggedIn())

	fmt.Println("\nDiagnostic complete.")
	client.Disconnect()
} 