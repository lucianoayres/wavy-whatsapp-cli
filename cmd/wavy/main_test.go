package main

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCmd(t *testing.T) {
	// Test that root command has expected values
	if rootCmd.Use != "wavy" {
		t.Errorf("Expected rootCmd.Use to be 'wavy', got %q", rootCmd.Use)
	}

	if rootCmd.Short != "WhatsApp CLI client" {
		t.Errorf("Expected rootCmd.Short to be 'WhatsApp CLI client', got %q", rootCmd.Short)
	}

	// Verify that the Run function is set
	if rootCmd.Run == nil {
		t.Error("Expected rootCmd.Run to be set, but it wasn't")
	}
}

func TestCommandRegistration(t *testing.T) {
	// Manual check that each command is defined with the correct name
	// Commands may include argument patterns in their Use field, so we check for starting pattern
	
	if !strings.HasPrefix(setupCmd.Use, "setup") {
		t.Errorf("Expected setupCmd.Use to start with 'setup', got %q", setupCmd.Use)
	}
	
	if !strings.HasPrefix(sendCmd.Use, "send") {
		t.Errorf("Expected sendCmd.Use to start with 'send', got %q", sendCmd.Use)
	}
	
	if !strings.HasPrefix(checkCmd.Use, "check") {
		t.Errorf("Expected checkCmd.Use to start with 'check', got %q", checkCmd.Use)
	}
	
	if !strings.HasPrefix(groupsCmd.Use, "groups") {
		t.Errorf("Expected groupsCmd.Use to start with 'groups', got %q", groupsCmd.Use)
	}
	
	// Verify each command has a meaningful description
	for _, cmd := range []*cobra.Command{setupCmd, sendCmd, checkCmd, groupsCmd} {
		if cmd.Short == "" {
			t.Errorf("Command %q is missing a Short description", cmd.Use)
		}
		
		// All commands should have a Run function
		if cmd.Run == nil {
			t.Errorf("Command %q is missing a Run function", cmd.Use)
		}
	}
} 