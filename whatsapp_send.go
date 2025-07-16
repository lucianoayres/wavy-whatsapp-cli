package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
)

func main() {
	// Parse CLI flags
	to := flag.String("to", "", "Recipient JID, e.g. +5511999999999")
	msg := flag.String("msg", "", "Message text to send")
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
	client := whatsmeow.NewClient(deviceStore, nil)
	if client.Store.ID == nil {
		log.Fatalf("No WhatsApp session found; please run setup first")
	}
	
	err = client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Build and send the message
	recipient, err := types.ParseJID(*to)
	if err != nil {
		log.Fatalf("Invalid JID %s: %v", *to, err)
	}
	
	msgText := *msg
	message := &waProto.Message{
		Conversation: &msgText,
	}
	
	resp, err := client.SendMessage(context.Background(), recipient, message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
	
	fmt.Printf("Message sent, server response: %v\n", resp)
} 