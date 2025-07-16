package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	// Enable debug logging
	clientLog := waLog.Stdout("Client", "DEBUG", true)

	// Parse CLI flags
	to := flag.String("to", "", "Recipient JID, e.g. +5511999999999")
	msg := flag.String("msg", "", "Message text to send")
	debug := flag.Bool("debug", false, "Enable verbose debug output")
	wait := flag.Int("wait", 5, "Seconds to wait for message confirmation")
	flag.Parse()

	if *to == "" || *msg == "" {
		flag.Usage()
		return
	}

	// Open or create the SQLite store (reuses credentials)
	db, err := sqlstore.New(context.Background(), "sqlite3", "file:./client.db?_foreign_keys=on", nil)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	
	// Get device store
	deviceStore, err := db.GetFirstDevice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get device: %v", err)
	}

	// Create a client from the device and connect
	var client *whatsmeow.Client
	if *debug {
		client = whatsmeow.NewClient(deviceStore, clientLog)
	} else {
		client = whatsmeow.NewClient(deviceStore, nil)
	}
	
	if client.Store.ID == nil {
		log.Fatalf("No WhatsApp session found; please run setup first")
	}
	
	err = client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	// Print own ID for debugging
	if *debug {
		fmt.Printf("Connected as JID: %s\n", client.Store.ID)
	}

	// Process phone number to ensure proper format
	phoneNumber := *to
	phoneNumber = strings.TrimSpace(phoneNumber)
	if strings.HasPrefix(phoneNumber, "+") {
		phoneNumber = phoneNumber[1:] // Remove the '+' prefix
	}
	
	// First verify the number is on WhatsApp
	exists, err := client.IsOnWhatsApp([]string{phoneNumber})
	if err != nil {
		log.Printf("Warning: Error checking if number exists on WhatsApp: %v", err)
	} else if len(exists) > 0 && exists[0].IsIn {
		// Use the exact JID returned by the WhatsApp server
		recipient := exists[0].JID
		
		if *debug {
			fmt.Printf("Verified recipient JID: %s\n", recipient.String())
			fmt.Printf("Message content: %s\n", *msg)
		}
		
		msgText := *msg
		message := &waProto.Message{
			Conversation: &msgText,
		}
		
		// Send message with context and timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*wait)*time.Second)
		defer cancel()
		
		fmt.Printf("Sending message to %s...\n", recipient.String())
		resp, err := client.SendMessage(ctx, recipient, message)
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
		
		fmt.Printf("Message sent to %s, server response: %v\n", recipient.String(), resp)
	} else {
		log.Fatalf("Error: Phone number %s not found on WhatsApp", phoneNumber)
	}
	
	// Disconnect client after sending
	client.Disconnect()
} 