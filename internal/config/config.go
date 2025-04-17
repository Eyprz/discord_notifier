package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type Config struct {
	WebhookURL string `yaml:"webhook_url"`
	AvatarURL  string `yaml:"avatar_url"`
	Username   string `yaml:"username"`
}

func LoadConfig(filePath string) (*Config, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := GenerateDefaultConfig(filePath)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	configBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}
	if config.WebhookURL == "" {
		return nil, errors.New("webhook URL is required")
	}
	return config, nil
}
