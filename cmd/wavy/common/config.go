package common

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

const (
	// XDG paths
	ConfigDir = "~/.config/wavy"
	DataDir   = "~/.local/share/wavy"
)

// GetConfigPath returns the expanded path to the config directory
func GetConfigPath() (string, error) {
	return expandHomeDir(ConfigDir)
}

// GetDataPath returns the expanded path to the data directory
func GetDataPath() (string, error) {
	return expandHomeDir(DataDir)
}

// GetDBPath returns the path to the client database file
func GetDBPath() (string, error) {
	dataPath, err := GetDataPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(dataPath, "client.db"), nil
}

// expandHomeDir expands the tilde in paths to the user's home directory
func expandHomeDir(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[1:]), nil
	}
	return path, nil
}

// EnsureDirectories creates necessary directories for wavy
func EnsureDirectories() error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	dataPath, err := GetDataPath()
	if err != nil {
		return err
	}

	// Create directories with appropriate permissions
	if err := os.MkdirAll(configPath, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.MkdirAll(dataPath, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	return nil
}

// CreateWAClient creates and connects a WhatsApp client
// Returns the client and a flag indicating if it needs setup
func CreateWAClient(debug bool) (*whatsmeow.Client, bool, error) {
	// Ensure directories exist
	if err := EnsureDirectories(); err != nil {
		return nil, false, err
	}

	// Get database path
	dbPath, err := GetDBPath()
	if err != nil {
		return nil, false, err
	}

	// Create database directory if it doesn't exist
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, false, fmt.Errorf("failed to create database directory: %w", err)
	}

	// SQLite connection string
	container := fmt.Sprintf("file:%s?_foreign_keys=on", dbPath)

	// Open the database
	var dbLog waLog.Logger
	if debug {
		dbLog = waLog.Stdout("Database", "DEBUG", true)
	}

	db, err := sqlstore.New(context.Background(), "sqlite3", container, dbLog)
	if err != nil {
		return nil, false, fmt.Errorf("failed to open database: %w", err)
	}

	// Get device store
	deviceStore, err := db.GetFirstDevice(context.Background())
	if err != nil {
		// No device found, need setup
		deviceStore = db.NewDevice()
	}

	// Create a client from the device
	var clientLog waLog.Logger
	if debug {
		clientLog = waLog.Stdout("Client", "DEBUG", true)
	}

	client := whatsmeow.NewClient(deviceStore, clientLog)

	// Check if setup is needed
	needsSetup := client.Store.ID == nil

	return client, needsSetup, nil
}
