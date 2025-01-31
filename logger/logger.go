package logger

type Logger interface {
	WriteLog(entry LogEntry) error
}

type LogEntry struct {
	Email     string      `json:"email"`
	Found     int         `json:"found"`
	Fields    []string    `json:"fields"`
	Sources   interface{} `json:"sources"`
	Timestamp string      `json:"timestamp"`
}
