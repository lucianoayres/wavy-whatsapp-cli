package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	//nolint:staticcheck // Using deprecated package for compatibility
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"

	"whatsmeow-go/cmd/wavy/common"
)

var (
	to     string
	msg    string
	debug  bool
	wait   int
)

var sendCmd = &cobra.Command{
	Use:   "send [recipient] [message]",
	Short: "Send a WhatsApp message",
	Long:  `Send a WhatsApp message to a contact or group.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Handle positional arguments if provided
		if len(args) >= 2 && to == "" {
			to = args[0]
			msg = args[1]
		} else if len(args) >= 1 && msg == "" {
			msg = args[0]
		}

		if to == "" || msg == "" {
			cmd.Help()
			os.Exit(1)
		}

		runSend()
	},
}

func init() {
	sendCmd.Flags().StringVarP(&to, "to", "t", "", "Recipient (phone number or group ID)")
	sendCmd.Flags().StringVarP(&msg, "msg", "m", "", "Message text to send")
	sendCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable verbose debugging")
	sendCmd.Flags().IntVarP(&wait, "wait", "w", 5, "Seconds to wait for message confirmation")
}

func runSend() {
	// Create client
	client, needsSetup, err := common.CreateWAClient(debug)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}

	if needsSetup {
		fmt.Fprintf(os.Stderr, "No WhatsApp session found. Please run 'wavy setup' first.\n")
		os.Exit(1)
	}

	// Connect to WhatsApp
	err = client.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}

	// Print own ID for debugging
	if debug {
		fmt.Printf("Connected as JID: %s\n", client.Store.ID)
	}

	// Determine recipient type and parse the JID
	var recipient types.JID

	// Check if this is a group JID (contains "@g.us")
	if strings.Contains(to, "@g.us") {
		// Parse directly as a group JID
		if strings.Count(to, "@") != 1 {
			fmt.Fprintf(os.Stderr, "Invalid group ID format. Should be 'number@g.us'\n")
			os.Exit(1)
		}

		parts := strings.Split(to, "@")
		recipient = types.JID{
			User:   parts[0],
			Server: "g.us",
		}

		if debug {
			fmt.Printf("Sending to group: %s\n", recipient.String())
		}
	} else {
		// Handle as individual contact
		phoneNumber := to
		phoneNumber = strings.TrimSpace(phoneNumber)
		if strings.HasPrefix(phoneNumber, "+") {
			phoneNumber = phoneNumber[1:] // Remove the '+' prefix
		}

		// First verify the number is on WhatsApp
		exists, err := client.IsOnWhatsApp([]string{phoneNumber})
		if err != nil {
			fmt.Printf("Warning: Error checking if number exists on WhatsApp: %v\n", err)

			// If we can't verify, try to construct the JID anyway
			recipient = types.JID{
				User:   phoneNumber,
				Server: "s.whatsapp.net",
			}
		} else if len(exists) > 0 && exists[0].IsIn {
			// Use the exact JID returned by the WhatsApp server
			recipient = exists[0].JID
		} else {
			fmt.Fprintf(os.Stderr, "Error: Phone number %s not found on WhatsApp\n", phoneNumber)
			os.Exit(1)
		}

		if debug {
			fmt.Printf("Sending to individual contact: %s\n", recipient.String())
		}
	}

	// Prepare the message
	message := &waProto.Message{
		Conversation: &msg,
	}

	// Send message with context and timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(wait)*time.Second)
	defer cancel()

	fmt.Printf("Sending message to %s...\n", recipient.String())
	resp, err := client.SendMessage(ctx, recipient, message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Message sent successfully to %s, server response: %v\n", recipient.String(), resp)

	// Disconnect client after sending
	client.Disconnect()
} 