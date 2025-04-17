package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

func GenerateDefaultConfig(path string) error {
	cfg := Config{
		WebhookURL: "https://example.com/webhook",
		AvatarURL:  "",
		Username:   "Notifier",
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	configBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	_, err = file.Write(configBytes)
	if err != nil {
		return err
	}
	return nil
}
