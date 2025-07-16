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

	// Copy binary to install directory
	srcPath := filepath.Join(binDir, binaryName)
	destPath := filepath.Join(installDir, binaryName)

	if err := sh.Copy(destPath, srcPath); err != nil {
		return fmt.Errorf("failed to copy binary to %s: %w", destPath, err)
	}

	// Set executable permissions
	if err := os.Chmod(destPath, 0755); err != nil {
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

	// Remove binary
	binaryPath := filepath.Join(installDir, binaryName)
	if err := os.Remove(binaryPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove binary %s: %w", binaryPath, err)
	}
	fmt.Printf("Removed binary: %s\n", binaryPath)

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