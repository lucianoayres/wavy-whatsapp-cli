package main

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestGroupsCmd(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "groups",
		Short: "List all your WhatsApp groups",
		Run: func(cmd *cobra.Command, args []string) {
			// Just verify that the command structure is correct
			if cmd.Use != "groups" {
				t.Errorf("Expected cmd.Use to be 'groups', got %q", cmd.Use)
			}
			if cmd.Short != "List all your WhatsApp groups" {
				t.Errorf("Expected cmd.Short to be 'List all your WhatsApp groups', got %q", cmd.Short)
			}
		},
	}

	// Execute the command
	cmd.Execute()
}

func TestGroupsCommand(t *testing.T) {
	// This test verifies that the actual groupsCmd structure matches what we expect
	if groupsCmd.Use != "groups" {
		t.Errorf("Expected groupsCmd.Use to be 'groups', got %q", groupsCmd.Use)
	}

	if !strings.Contains(groupsCmd.Short, "List all your WhatsApp groups") {
		t.Errorf("Expected groupsCmd.Short to contain 'List all your WhatsApp groups', got %q", groupsCmd.Short)
	}

	if !strings.Contains(groupsCmd.Long, "Display information about all the WhatsApp groups") {
		t.Errorf("Expected groupsCmd.Long to contain 'Display information about all the WhatsApp groups', got %q", groupsCmd.Long)
	}

	// Verify that the Run function is set
	if groupsCmd.Run == nil {
		t.Error("Expected groupsCmd.Run to be set, but it wasn't")
	}
} 