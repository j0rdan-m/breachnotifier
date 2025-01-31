package notifier

import "fmt"

func GetNotifier(notifierType, apiURL, apiKey, organisation string) (Notifier, error) {
	switch notifierType {
	case "thehive":
		return NewTheHiveNotifier(apiURL, apiKey, organisation), nil
	default:
		return nil, fmt.Errorf("unknown notifier type: %s", notifierType)
	}
}
