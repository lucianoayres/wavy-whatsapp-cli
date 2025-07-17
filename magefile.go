//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

const (
	binaryName = "wavy"
	binDir     = "./bin"
	installDir = "/usr/local/bin"
	configDir  = "~/.config/wavy"
	dataDir    = "~/.local/share/wavy"
)

// Build builds the wavy binary
func Build() error {
	fmt.Println("Building wavy binary...")

	// Create bin directory if it doesn't exist
	if err := os.MkdirAll(binDir, 0755); err != nil {
		return fmt.Errorf("failed to create bin directory: %w", err)
	}

	// Build the binary
	cmd := exec.Command("go", "build", "-o", filepath.Join(binDir, binaryName), "./cmd/wavy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	fmt.Printf("Binary built successfully: %s\n", filepath.Join(binDir, binaryName))
	return nil
}

// Install installs the wavy binary to /usr/local/bin
func Install() error {
	fmt.Println("Installing wavy...")

	// Build first
	if err := Build(); err != nil {
		return err
	}

	// Create config and data directories
	configDirExpanded, err := homePath(configDir)
	if err != nil {
		return err
	}

	dataDirExpanded, err := homePath(dataDir)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDirExpanded, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.MkdirAll(dataDirExpanded, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	// Copy binary to install directory using sudo
	srcPath := filepath.Join(binDir, binaryName)
	destPath := filepath.Join(installDir, binaryName)

	fmt.Println("Copying binary to system directory. You may be prompted for your password...")
	
	// Use sudo to copy the binary
	cmd := exec.Command("sudo", "cp", srcPath, destPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Connect to terminal for password prompt
	cmd.Stdin = os.Stdin
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to copy binary to %s: %w", destPath, err)
	}

	// Set executable permissions using sudo
	chmodCmd := exec.Command("sudo", "chmod", "755", destPath)
	chmodCmd.Stdout = os.Stdout
	chmodCmd.Stderr = os.Stderr
	if err := chmodCmd.Run(); err != nil {
		return fmt.Errorf("failed to set permissions: %w", err)
	}

	fmt.Printf("Installed %s to %s\n", binaryName, destPath)
	fmt.Printf("Created config directory: %s\n", configDirExpanded)
	fmt.Printf("Created data directory: %s\n", dataDirExpanded)
	return nil
}

// Uninstall removes all wavy files from the system
func Uninstall() error {
	fmt.Println("Uninstalling wavy...")

	// Remove binary using sudo
	binaryPath := filepath.Join(installDir, binaryName)
	fmt.Println("Removing binary from system directory. You may be prompted for your password...")
	
	// Check if binary exists before trying to remove it
	if _, err := os.Stat(binaryPath); err == nil {
		// Binary exists, use sudo to remove it
		rmCmd := exec.Command("sudo", "rm", binaryPath)
		rmCmd.Stdout = os.Stdout
		rmCmd.Stderr = os.Stderr
		// Connect to terminal for password prompt
		rmCmd.Stdin = os.Stdin
		
		if err := rmCmd.Run(); err != nil {
			return fmt.Errorf("failed to remove binary %s: %w", binaryPath, err)
		}
		fmt.Printf("Removed binary: %s\n", binaryPath)
	} else if os.IsNotExist(err) {
		fmt.Printf("Binary not found at %s, skipping removal\n", binaryPath)
	} else {
		return fmt.Errorf("failed to check if binary exists: %w", err)
	}

	// Remove config directory
	configDirExpanded, err := homePath(configDir)
	if err != nil {
		return err
	}
	if err := os.RemoveAll(configDirExpanded); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove config directory %s: %w", configDirExpanded, err)
	}
	fmt.Printf("Removed config directory: %s\n", configDirExpanded)

	// Remove data directory
	dataDirExpanded, err := homePath(dataDir)
	if err != nil {
		return err
	}
	if err := os.RemoveAll(dataDirExpanded); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove data directory %s: %w", dataDirExpanded, err)
	}
	fmt.Printf("Removed data directory: %s\n", dataDirExpanded)

	fmt.Println("Wavy has been completely removed from your system")
	return nil
}

// Clean removes build artifacts
func Clean() error {
	fmt.Println("Cleaning build artifacts...")
	return os.RemoveAll(binDir)
}

// Test runs tests
func Test() error {
	fmt.Println("Running tests...")
	return sh.Run("go", "test", "./...")
}

// Check runs linters and static analysis
func Check() error {
	fmt.Println("Running linters and static analysis...")
	if err := sh.Run("go", "fmt", "./..."); err != nil {
		return err
	}
	if err := sh.Run("go", "vet", "./..."); err != nil {
		return err
	}
	return nil
}

// Helper function to expand home directory (~)
func homePath(path string) (string, error) {
	if path[:2] == "~/" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		return filepath.Join(home, path[2:]), nil
	}
	return path, nil
} 