package mocks

import (
	"context"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

// MockClient is a mock implementation of whatsmeow.Client
type MockClient struct {
	ConnectCalled      bool
	DisconnectCalled   bool
	IsOnWhatsAppCalled bool

	// Store mock data
	Store struct {
		ID *types.JID
	}

	// Mock behaviors
	MockIsOnWhatsApp    func([]string) ([]types.IsOnWhatsAppResponse, error)
	MockGetJoinedGroups func() ([]types.GroupInfo, error)
}

// Connect mocks the Connect method
func (m *MockClient) Connect() error {
	m.ConnectCalled = true
	return nil
}

// Disconnect mocks the Disconnect method
func (m *MockClient) Disconnect() {
	m.DisconnectCalled = true
}

// IsOnWhatsApp mocks the IsOnWhatsApp method
func (m *MockClient) IsOnWhatsApp(numbers []string) ([]types.IsOnWhatsAppResponse, error) {
	m.IsOnWhatsAppCalled = true
	if m.MockIsOnWhatsApp != nil {
		return m.MockIsOnWhatsApp(numbers)
	}
	return []types.IsOnWhatsAppResponse{}, nil
}

// GetJoinedGroups mocks the GetJoinedGroups method
func (m *MockClient) GetJoinedGroups() ([]types.GroupInfo, error) {
	if m.MockGetJoinedGroups != nil {
		return m.MockGetJoinedGroups()
	}
	return []types.GroupInfo{}, nil
}

// GetQRChannel mocks the GetQRChannel method
func (m *MockClient) GetQRChannel(ctx context.Context) (chan string, error) {
	ch := make(chan string, 1)
	ch <- "mock-qr-code"
	return ch, nil
}

// AddEventHandler mocks the AddEventHandler method
func (m *MockClient) AddEventHandler(handler whatsmeow.EventHandler) {
	// Do nothing
}

// WaitForMessage mocks the WaitForMessage method
func (m *MockClient) WaitForMessage(msgID types.MessageID, timeout int) (*events.Message, error) {
	return &events.Message{
		Info: types.MessageInfo{
			ID: msgID,
		},
	}, nil
}

// NewMockClient creates a new mock client with defaults
func NewMockClient() *MockClient {
	m := &MockClient{}
	m.Store.ID = &types.JID{
		User:   "1234567890",
		Server: "s.whatsapp.net",
	}
	return m
}
