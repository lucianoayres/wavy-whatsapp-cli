package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	// Enable debug logging
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	clientLog := waLog.Stdout("Client", "DEBUG", true)

	// Open the database
	db, err := sqlstore.New(context.Background(), "sqlite3", "file:./client.db?_foreign_keys=on", dbLog)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Get device store
	deviceStore, err := db.GetFirstDevice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get device: %v", err)
	}

	// Create a client from the device and connect
	client := whatsmeow.NewClient(deviceStore, clientLog)
	if client.Store.ID == nil {
		log.Fatalf("No WhatsApp session found; please run setup first")
	}

	// Connect to WhatsApp
	fmt.Println("Connecting to WhatsApp...")
	err = client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Get all joined groups
	groups, err := client.GetJoinedGroups()
	if err != nil {
		log.Fatalf("Failed to get groups: %v", err)
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
		fmt.Println("./whatsapp-send -to \"GROUP_ID\" -msg \"Hello group!\"")
		fmt.Println("\nExample:")
		if len(groups) > 0 {
			fmt.Printf("./whatsapp-send -to \"%s\" -msg \"Hello group!\"\n", groups[0].JID.String())
		}
	}

	client.Disconnect()
} 