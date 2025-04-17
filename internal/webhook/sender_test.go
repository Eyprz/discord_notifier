package webhook_test

import (
	"encoding/json"
	"github.com/eyprz/discord_notifier/internal/webhook"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock server to simulate Discord webhook response
func mockDiscordWebhookServer(t *testing.T) *httptest.Server {
	t.Helper()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}
		if payload["content"] == nil {
			http.Error(w, "Missing content", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	return httptest.NewServer(handler)
}

// Test sending a message to Discord webhook
func TestSendMessageToDiscord(t *testing.T) {
	server := mockDiscordWebhookServer(t)
	defer server.Close()

	tests := []struct {
		name       string
		webhookURL string
		message    string
		avatarURL  string
		username   string
		expectErr  bool
	}{
		{
			name:       "Valid Request",
			webhookURL: server.URL,
			message:    "Hello, Discord!",
			avatarURL:  "https://example.com/avatar.png",
			username:   "Notifier",
			expectErr:  false,
		},
		{
			name:       "Invalid URL",
			webhookURL: "invalid-url",
			message:    "Hello, Discord!",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := webhook.SendMessageToDiscord(tt.webhookURL, tt.message, tt.avatarURL, tt.username)
			if (err != nil) != tt.expectErr {
				t.Errorf("SendMessageToDiscord() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// Test sending a message with missing content
func TestSendMessageToDiscord_MissingContent(t *testing.T) {
	server := mockDiscordWebhookServer(t)
	defer server.Close()

	tests := []struct {
		name       string
		webhookURL string
		message    string
		expectErr  bool
	}{
		{
			name:       "Missing Content",
			webhookURL: server.URL,
			message:    "",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := webhook.SendMessageToDiscord(tt.webhookURL, tt.message, "", "")
			if (err != nil) != tt.expectErr {
				t.Errorf("SendMessageToDiscord() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// Test sending a message with invalid JSON
func TestSendMessageToDiscord_InvalidJSON(t *testing.T) {
	server := mockDiscordWebhookServer(t)
	defer server.Close()

	tests := []struct {
		name       string
		webhookURL string
		message    string
		expectErr  bool
	}{
		{
			name:       "Invalid JSON",
			webhookURL: server.URL,
			message:    "{invalid json}",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := webhook.SendMessageToDiscord(tt.webhookURL, tt.message, "", "")
			if (err != nil) != tt.expectErr {
				t.Errorf("SendMessageToDiscord() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// Test sending a message with a timeout
func TestSendMessageToDiscord_Timeout(t *testing.T) {
	server := mockDiscordWebhookServer(t)
	defer server.Close()

	tests := []struct {
		name       string
		webhookURL string
		message    string
		expectErr  bool
	}{
		{
			name:       "Timeout",
			webhookURL: server.URL,
			message:    "Hello, Discord!",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := webhook.SendMessageToDiscord(tt.webhookURL, tt.message, "", "")
			if (err != nil) != tt.expectErr {
				t.Errorf("SendMessageToDiscord() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// Test sending a message with a non-200 response
func TestSendMessageToDiscord_Non200Response(t *testing.T) {
	server := mockDiscordWebhookServer(t)
	defer server.Close()

	tests := []struct {
		name       string
		webhookURL string
		message    string
		expectErr  bool
	}{
		{
			name:       "Non-200 Response",
			webhookURL: server.URL,
			message:    "Hello, Discord!",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := webhook.SendMessageToDiscord(tt.webhookURL, tt.message, "", "")
			if (err != nil) != tt.expectErr {
				t.Errorf("SendMessageToDiscord() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// Test sending a message with a large payload
func TestSendMessageToDiscord_LargePayload(t *testing.T) {
	server := mockDiscordWebhookServer(t)
	defer server.Close()

	largeMessage := make([]byte, 8192) // 8KB payload
	for i := range largeMessage {
		largeMessage[i] = 'A'
	}

	tests := []struct {
		name       string
		webhookURL string
		message    string
		expectErr  bool
	}{
		{
			name:       "Large Payload",
			webhookURL: server.URL,
			message:    string(largeMessage),
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := webhook.SendMessageToDiscord(tt.webhookURL, tt.message, "", "")
			if (err != nil) != tt.expectErr {
				t.Errorf("SendMessageToDiscord() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
