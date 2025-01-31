package checker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LeakCheck struct{}

func NewLeakCheck() *LeakCheck {
	return &LeakCheck{}
}

func (lc *LeakCheck) CheckEmail(email string) (*Response, error) {
	const apiURL = "https://leakcheck.net/api/public?check="

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", apiURL, email), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "BreachNotifier/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
