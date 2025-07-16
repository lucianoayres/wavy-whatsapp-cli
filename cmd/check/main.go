package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	// Enable debug logging
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	clientLog := waLog.Stdout("Client", "DEBUG", true)

	// Parse CLI flags
	checkPhone := flag.String("phone", "", "Phone number to check (optional)")
	flag.Parse()

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

	// Print own ID
	fmt.Printf("Connected as: %s\n", client.Store.ID)
	fmt.Printf("Your phone number: %s\n", client.Store.ID.User)

	// Check if a specific phone number was provided
	if *checkPhone != "" {
		fmt.Printf("\nChecking phone number: %s\n", *checkPhone)
		
		// Format the phone number
		phoneNumber := *checkPhone
		if phoneNumber[0] == '+' {
			phoneNumber = phoneNumber[1:]
		}
		
		// Create the JID
		jid := types.JID{
			User:   phoneNumber,
			Server: "s.whatsapp.net",
		}
		
		fmt.Printf("JID for this number: %s\n", jid.String())

		// Check if the user exists on WhatsApp
		fmt.Printf("Checking if %s exists on WhatsApp...\n", phoneNumber)
		
		exists, err := client.IsOnWhatsApp([]string{phoneNumber})
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
	}

	// Show some debugging info
	fmt.Println("\nConnection details:")
	fmt.Printf("Connected: %t\n", client.IsConnected())
	fmt.Printf("LoggedIn: %t\n", client.IsLoggedIn())

	fmt.Println("\nDiagnostic complete.")
	client.Disconnect()
} 