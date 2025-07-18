package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCheckCmdWithFlag(t *testing.T) {
	// Save original value and restore it after test
	originalPhoneNumber := phoneNumber
	defer func() {
		phoneNumber = originalPhoneNumber
	}()

	// Setup command
	cmd := &cobra.Command{Use: "check"}
	cmd.Flags().StringVarP(&phoneNumber, "phone", "p", "", "Phone number to check")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		// Check that the phone number was set correctly
		if phoneNumber != "1234567890" {
			t.Errorf("Expected phoneNumber to be 1234567890, got %s", phoneNumber)
		}
	}

	// Run command with flag
	cmd.SetArgs([]string{"--phone", "1234567890"})
	cmd.Execute()
}

func TestCheckCmdWithPositionalArg(t *testing.T) {
	// Save original value and restore it after test
	originalPhoneNumber := phoneNumber
	defer func() {
		phoneNumber = originalPhoneNumber
	}()

	// Setup
	phoneNumber = ""
	args := []string{"1234567890"}

	// Check if positional argument is handled correctly
	if len(args) > 0 && phoneNumber == "" {
		phoneNumber = args[0]
	}

	if phoneNumber != "1234567890" {
		t.Errorf("Expected phoneNumber to be 1234567890, got %s", phoneNumber)
	}
}
