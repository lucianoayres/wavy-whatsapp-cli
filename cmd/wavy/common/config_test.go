package common

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandHomeDir(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantTilde bool
	}{
		{
			name:     "Path with tilde",
			path:     "~/some/path",
			wantTilde: true,
		},
		{
			name:     "Path without tilde",
			path:     "/absolute/path",
			wantTilde: false,
		},
		{
			name:     "Relative path",
			path:     "relative/path",
			wantTilde: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := expandHomeDir(tt.path)
			if err != nil {
				t.Fatalf("expandHomeDir(%q) returned error: %v", tt.path, err)
			}

			if tt.wantTilde {
				home, err := os.UserHomeDir()
				if err != nil {
					t.Fatalf("os.UserHomeDir() failed: %v", err)
				}
				
				expected := filepath.Join(home, tt.path[1:])
				if result != expected {
					t.Errorf("expandHomeDir(%q) = %q, want %q", tt.path, result, expected)
				}
				
				// Make sure tilde is expanded
				if result[0] == '~' {
					t.Errorf("tilde was not expanded in result: %q", result)
				}
			} else {
				// For paths without tilde, the result should be unchanged
				if result != tt.path {
					t.Errorf("expandHomeDir(%q) = %q, want %q", tt.path, result, tt.path)
				}
			}
		})
	}
}

func TestGetConfigPath(t *testing.T) {
	path, err := GetConfigPath()
	if err != nil {
		t.Fatalf("GetConfigPath() failed: %v", err)
	}
	
	// Config path should not have a tilde
	if len(path) > 0 && path[0] == '~' {
		t.Errorf("GetConfigPath() returned path with tilde: %q", path)
	}
	
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("os.UserHomeDir() failed: %v", err)
	}
	
	expected := filepath.Join(home, ConfigDir[1:])
	if path != expected {
		t.Errorf("GetConfigPath() = %q, want %q", path, expected)
	}
}

func TestGetDataPath(t *testing.T) {
	path, err := GetDataPath()
	if err != nil {
		t.Fatalf("GetDataPath() failed: %v", err)
	}
	
	// Data path should not have a tilde
	if len(path) > 0 && path[0] == '~' {
		t.Errorf("GetDataPath() returned path with tilde: %q", path)
	}
	
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("os.UserHomeDir() failed: %v", err)
	}
	
	expected := filepath.Join(home, DataDir[1:])
	if path != expected {
		t.Errorf("GetDataPath() = %q, want %q", path, expected)
	}
}

func TestGetDBPath(t *testing.T) {
	path, err := GetDBPath()
	if err != nil {
		t.Fatalf("GetDBPath() failed: %v", err)
	}
	
	// Check that it ends with "client.db"
	if filepath.Base(path) != "client.db" {
		t.Errorf("GetDBPath() should end with 'client.db', got: %q", path)
	}
	
	// Verify that the path starts with the data directory
	dataPath, err := GetDataPath()
	if err != nil {
		t.Fatalf("GetDataPath() failed: %v", err)
	}
	
	if filepath.Dir(path) != dataPath {
		t.Errorf("GetDBPath() directory = %q, want %q", filepath.Dir(path), dataPath)
	}
} 