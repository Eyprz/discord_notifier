package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SendMessageToDiscord(webhookURL, message, avatarURL, username string) error {
	if webhookURL == "" || message == "" || username == "" {
		return fmt.Errorf("webhook URL, message, and username are required")
	}
	if len(message) > 2000 {
		return fmt.Errorf("message exceeds 2000 characters")
	}
	if len(avatarURL) > 2048 {
		return fmt.Errorf("avatar URL exceeds 2048 characters")
	}
	if len(username) > 80 {
		return fmt.Errorf("username exceeds 80 characters")
	}
	if len(webhookURL) > 2000 {
		return fmt.Errorf("webhook URL exceeds 2000 characters")
	}

	payload := map[string]interface{}{
		"content":    message,
		"avatar_url": avatarURL,
		"username":   username,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON payload: %w", err)
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("received non-200 response: %s", res.Status)
	}
	return nil
}
