package logger

import (
	"fmt"
)

type ELKLogger struct{}

func NewELKLogger() *ELKLogger {
	return &ELKLogger{}
}

func (l *ELKLogger) WriteLog(entry LogEntry) error {
	// Simule un envoi vers un système ELK
	fmt.Printf("Envoi vers ELK : %+v\n", entry)
	return nil
}
