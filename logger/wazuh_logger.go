package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// WazuhLogger écrit les logs dans un fichier JSON compatible avec Wazuh.
type WazuhLogger struct {
	FilePath string
}

// NewWazuhLogger crée un nouveau WazuhLogger.
func NewWazuhLogger(filePath string) *WazuhLogger {
	return &WazuhLogger{FilePath: filePath}
}

// WriteLog écrit une entrée de log dans un fichier JSON.
func (l *WazuhLogger) WriteLog(entry LogEntry) error {
	// Assurez-vous que le chemin du fichier est valide
	if l.FilePath == "" {
		return fmt.Errorf("log file path is empty")
	}

	// Ouvrir ou créer le fichier de log
	file, err := os.OpenFile(l.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	// Ajouter un champ "timestamp" à l'entrée pour Wazuh
	entryWithTimestamp := struct {
		Timestamp string `json:"timestamp"`
		LogEntry
	}{
		Timestamp: time.Now().Format(time.RFC3339), // Format ISO-8601 pour Wazuh
		LogEntry:  entry,
	}

	// Encoder l'entrée en JSON
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(entryWithTimestamp); err != nil {
		return fmt.Errorf("failed to encode log entry: %w", err)
	}

	return nil
}
