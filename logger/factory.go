package logger

import "fmt"

func GetLogger(loggerType, filePath string) (Logger, error) {
	switch loggerType {
	case "wazuh":
		return NewWazuhLogger(filePath), nil
	case "elk":
		return NewELKLogger(), nil
	default:
		return nil, fmt.Errorf("type de logger inconnu : %s", loggerType)
	}
}
