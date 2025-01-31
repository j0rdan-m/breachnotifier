package config

import (
    "os"

    "gopkg.in/yaml.v3"
)

type LoggerConfig struct {
    Type     string `yaml:"type"`
    FilePath string `yaml:"file_path"`
}

type CheckerConfig struct {
    Type string `yaml:"type"`
}

type NotifierConfig struct {
    Type         string `yaml:"type"`
    ApiURL       string `yaml:"api_url"`
    ApiKey       string `yaml:"api_key"`
    Organisation string `yaml:"organisation"`
}

type Config struct {
    Emails   []string       `yaml:"emails"`
    Logger   *LoggerConfig  `yaml:"logger"`   // Logger devient optionnel
    Checker  CheckerConfig  `yaml:"checker"`
    Notifier *NotifierConfig `yaml:"notifier"` // Notifier devient optionnel
}

func LoadConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var config Config
    decoder := yaml.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }
    return &config, nil
}
