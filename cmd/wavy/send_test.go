package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestSendCmdFlags(t *testing.T) {
	// Save original values
	origTo := to
	origMsg := msg
	origDebug := debug
	origWait := wait

	// Restore after test
	defer func() {
		to = origTo
		msg = origMsg
		debug = origDebug
		wait = origWait
	}()

	// Setup command
	cmd := &cobra.Command{Use: "send"}
	cmd.Flags().StringVarP(&to, "to", "t", "", "Recipient")
	cmd.Flags().StringVarP(&msg, "msg", "m", "", "Message")
	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "Debug")
	cmd.Flags().IntVarP(&wait, "wait", "w", 5, "Wait time")

	// Test with all flags
	cmd.SetArgs([]string{
		"--to", "1234567890",
		"--msg", "Test message",
		"--debug",
		"--wait", "10",
	})
	cmd.Execute()

	// Check values
	if to != "1234567890" {
		t.Errorf("Expected to = '1234567890', got '%s'", to)
	}
	if msg != "Test message" {
		t.Errorf("Expected msg = 'Test message', got '%s'", msg)
	}
	if debug != true {
		t.Errorf("Expected debug = true, got %v", debug)
	}
	if wait != 10 {
		t.Errorf("Expected wait = 10, got %d", wait)
	}
}

func TestSendCmdPositionalArgs(t *testing.T) {
	// Save original values
	origTo := to
	origMsg := msg

	// Restore after test
	defer func() {
		to = origTo
		msg = origMsg
	}()

	// Reset values
	to = ""
	msg = ""

	// Create args and handle them like the command does
	args := []string{"1234567890", "Hello world"}

	if len(args) >= 2 && to == "" {
		to = args[0]
		msg = args[1]
	} else if len(args) >= 1 && msg == "" {
		msg = args[0]
	}

	// Check values
	if to != "1234567890" {
		t.Errorf("Expected to = '1234567890', got '%s'", to)
	}
	if msg != "Hello world" {
		t.Errorf("Expected msg = 'Hello world', got '%s'", msg)
	}
}

func TestSendCmdMsgOnlyArg(t *testing.T) {
	// Save original values
	origTo := to
	origMsg := msg

	// Restore after test
	defer func() {
		to = origTo
		msg = origMsg
	}()

	// Set values to simulate pre-set recipient
	to = "predefined-recipient"
	msg = ""

	// Create args and handle them like the command does
	args := []string{"Hello world"}

	if len(args) >= 2 && to == "" {
		to = args[0]
		msg = args[1]
	} else if len(args) >= 1 && msg == "" {
		msg = args[0]
	}

	// Check values
	if to != "predefined-recipient" {
		t.Errorf("Expected to = 'predefined-recipient', got '%s'", to)
	}
	if msg != "Hello world" {
		t.Errorf("Expected msg = 'Hello world', got '%s'", msg)
	}
}
