package notifier

type NotificationEntry struct {
	Email     string   `json:"email"`
	Found     int      `json:"found"`
	Fields    []string `json:"fields"`
	Sources   []string `json:"sources"` // Type spécifique
	Timestamp string   `json:"timestamp"`
}

type Notifier interface {
	Notify(entry NotificationEntry) error
}
