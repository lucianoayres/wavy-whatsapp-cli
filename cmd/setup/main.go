package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

// openFile opens the specified file with the default application
func openFile(path string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", path)
	case "darwin": // macOS
		cmd = exec.Command("open", path)
	default: // Linux and others
		cmd = exec.Command("xdg-open", path)
	}
	return cmd.Start()
}

func main() {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	clientLog := waLog.Stdout("Client", "DEBUG", true)

	// Open the database
	db, err := sqlstore.New(context.Background(), "sqlite3", "file:./client.db?_foreign_keys=on", dbLog)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Create a new device (WhatsApp session)
	device := db.NewDevice()
	
	client := whatsmeow.NewClient(device, clientLog)
	
	// Handle QR code
	qrChan, _ := client.GetQRChannel(context.Background())
	err = client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Listen for Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Disconnecting...")
		client.Disconnect()
		// Clean up QR code file if it exists
		os.Remove("whatsapp_qr_code.png")
		os.Exit(0)
	}()

	// Wait for QR code scan or successful authentication
	for evt := range qrChan {
		if evt.Event == "code" {
			// Generate QR code image file
			qrPath := filepath.Join(".", "whatsapp_qr_code.png")
			
			// Remove existing file if it exists
			os.Remove(qrPath)
			
			// Generate new QR code image file
			err := qrcode.WriteFile(evt.Code, qrcode.Medium, 256, qrPath)
			if err != nil {
				log.Printf("Failed to generate QR code image: %v", err)
				fmt.Printf("QR Code data (use an online QR generator): %s\n", evt.Code)
				continue
			}
			
			// Open the QR code image with default image viewer
			fmt.Println("Opening QR code image. Scan it with WhatsApp mobile app...")
			err = openFile(qrPath)
			if err != nil {
				log.Printf("Failed to open QR code image: %v", err)
				fmt.Printf("QR code saved to %s, please open it manually\n", qrPath)
			}
		} else {
			fmt.Println("Authentication successful!")
			// Clean up QR code file
			os.Remove("whatsapp_qr_code.png")
			break
		}
	}

	// Keep the connection alive
	fmt.Println("Setup complete! The connection is now authenticated.")
	fmt.Println("You can now use whatsapp-send to send messages.")
	fmt.Println("Press Ctrl+C to exit")
	
	// Block until Ctrl+C is pressed
	select {}
} 