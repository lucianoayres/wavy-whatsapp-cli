package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"whatsmeow-go/cmd/wavy/common"
)

var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List all your WhatsApp groups",
	Long:  `Display information about all the WhatsApp groups you're a member of.`,
	Run: func(cmd *cobra.Command, args []string) {
		runGroups()
	},
}

func runGroups() {
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

	// Get all joined groups
	groups, err := client.GetJoinedGroups()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get groups: %v\n", err)
		os.Exit(1)
	}

	// Print the list of groups
	if len(groups) == 0 {
		fmt.Println("You are not a member of any groups")
	} else {
		fmt.Println("\n===== YOUR WHATSAPP GROUPS =====")
		fmt.Println("Count:", len(groups))
		fmt.Println("----------------------------------")

		for i, group := range groups {
			fmt.Printf("%d. Group Name: %s\n", i+1, group.Name)
			fmt.Printf("   Group ID: %s\n", group.JID.String())
			fmt.Printf("   Member Count: %d\n", len(group.Participants))
			fmt.Println("----------------------------------")
		}

		fmt.Println("\nTo send a message to a group, use:")
		fmt.Println("wavy send -to \"GROUP_ID\" -msg \"Hello group!\"")
		fmt.Println("\nExample:")
		if len(groups) > 0 {
			fmt.Printf("wavy send -to \"%s\" -msg \"Hello group!\"\n", groups[0].JID.String())
		}
	}

	client.Disconnect()
}
