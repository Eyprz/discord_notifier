package config_test

import (
	"github.com/eyprz/discord_notifier/internal/config"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

func TestGenerateDefaultConfig(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
	}{
		{
			name:     "Generate Default Config",
			filePath: "test_config.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := config.GenerateDefaultConfig(tt.filePath)
			if err != nil {
				t.Errorf("GenerateDefaultConfig() error = %v", err)
				return
			}
			defer os.Remove(tt.filePath)
		})
	}
}

// Test loading a config file with valid data
func TestLoadConfig(t *testing.T) {
	testConfig := config.Config{
		WebhookURL: "https://example.com/webhook",
		AvatarURL:  "https://example.com/avatar.png",
		Username:   "Notifier",
	}
	testConfigFile := "test_config.yaml"
	file, err := os.Create(testConfigFile)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	defer os.Remove(testConfigFile)
	defer file.Close()

	configBytes, err := yaml.Marshal(testConfig)
	if err != nil {
		t.Fatalf("Failed to marshal test config: %v", err)
	}
	_, err = file.Write(configBytes)
	if err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}
	loadedConfig, err := config.LoadConfig(testConfigFile)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	if loadedConfig.WebhookURL != testConfig.WebhookURL {
		t.Errorf("Expected WebhookURL %s, got %s", testConfig.WebhookURL, loadedConfig.WebhookURL)
	}
	if loadedConfig.AvatarURL != testConfig.AvatarURL {
		t.Errorf("Expected AvatarURL %s, got %s", testConfig.AvatarURL, loadedConfig.AvatarURL)
	}
	if loadedConfig.Username != testConfig.Username {
		t.Errorf("Expected Username %s, got %s", testConfig.Username, loadedConfig.Username)
	}

	// Test loading a config file with missing required fields
	invalidConfig := `
	webhook_url: ""
	avatar_url: "https://example.com/avatar.png"
	username: "Notifier"
	`
	invalidConfigFile := "invalid_config.yaml"
	file, err = os.Create(invalidConfigFile)
	if err != nil {
		t.Fatalf("Failed to create invalid config file: %v", err)
	}
	defer os.Remove(invalidConfigFile)
	defer file.Close()
	_, err = file.Write([]byte(invalidConfig))

	if err != nil {
		t.Fatalf("Failed to write invalid config: %v", err)
	}
	_, err = config.LoadConfig(invalidConfigFile)
	if err == nil {
		t.Error("Expected error loading config with missing required fields, got nil")
	}

	// Test loading a config file with invalid YAML
	invalidYAML := `
	webhook_url: "https://example.com/webhook"
	avatar_url: "https://example.com/avatar.png"
	username: "Notifier"
	invalid_field: "Invalid"
	`
	invalidYAMLFile := "invalid_yaml_config.yaml"
	file, err = os.Create(invalidYAMLFile)
	if err != nil {
		t.Fatalf("Failed to create invalid YAML config file: %v", err)
	}
	defer os.Remove(invalidYAMLFile)
	defer file.Close()
	_, err = file.Write([]byte(invalidYAML))
	if err != nil {
		t.Fatalf("Failed to write invalid YAML config: %v", err)
	}
	_, err = config.LoadConfig(invalidYAMLFile)
	if err == nil {
		t.Error("Expected error loading config with invalid YAML, got nil")
	}
}
