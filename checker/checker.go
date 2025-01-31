package checker

type Checker interface {
	CheckEmail(email string) (*Response, error)
}

type Response struct {
	Found   int      `json:"found"`
	Fields  []string `json:"fields"`
	Sources []Source `json:"sources"`
}

type Source struct {
	Name string `json:"name"`
	Date string `json:"date"`
}
