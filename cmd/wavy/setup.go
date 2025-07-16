package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"

	"whatsmeow-go/cmd/wavy/common"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up a WhatsApp connection using QR code",
	Run: func(cmd *cobra.Command, args []string) {
		runSetup()
	},
}

func runSetup() {
	// Create client
	client, needsSetup, err := common.CreateWAClient(true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}

	// Get data directory for QR code
	dataPath, err := common.GetDataPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting data path: %v\n", err)
		os.Exit(1)
	}
	
	// Path to QR code file
	qrPath := filepath.Join(dataPath, "whatsapp_qr_code.png")

	// Clean up existing QR code
	os.Remove(qrPath)

	// If already set up, show notice but continue anyway
	if !needsSetup {
		fmt.Println("Warning: WhatsApp is already set up, but you can scan QR again to re-authenticate.")
	}

	// Handle QR code
	qrChan, _ := client.GetQRChannel(context.Background())
	err = client.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}

	// Listen for Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nDisconnecting...")
		client.Disconnect()
		// Clean up QR code file if it exists
		os.Remove(qrPath)
		os.Exit(0)
	}()

	// Wait for QR code scan or successful authentication
	for evt := range qrChan {
		if evt.Event == "code" {
			// Generate QR code image file
			
			// Remove existing file if it exists
			os.Remove(qrPath)
			
			// Generate new QR code image file
			err := qrcode.WriteFile(evt.Code, qrcode.Medium, 256, qrPath)
			if err != nil {
				fmt.Printf("Failed to generate QR code image: %v\n", err)
				fmt.Printf("QR Code data (use an online QR generator): %s\n", evt.Code)
				continue
			}
			
			// Open the QR code image with default image viewer
			fmt.Println("Opening QR code image. Scan it with WhatsApp mobile app...")
			err = openFile(qrPath)
			if err != nil {
				fmt.Printf("Failed to open QR code image: %v\n", err)
				fmt.Printf("QR code saved to %s, please open it manually\n", qrPath)
			}
		} else if evt.Event == "success" {
			fmt.Println("Authentication successful!")
			// Clean up QR code file
			os.Remove(qrPath)
			break
		}
	}

	// Keep the connection alive
	fmt.Println("Setup complete! The connection is now authenticated.")
	fmt.Println("You can now use wavy commands to interact with WhatsApp.")
	fmt.Println("Press Ctrl+C to exit")
	
	// Block until Ctrl+C is pressed
	select {}
}

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