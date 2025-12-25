package servicebus

import (
	"errors"
	"net/http"
	"time"
)

type Event struct {
	EventType    string `json:"eventType"`
	EntityId     string `json:"entityId"`
	EntityType   string `json:"entityType"`
	Timestamp    string `json:"timestamp"`
	Datetime     string `json:"datetime"`
	Version      string `json:"version"`
	Country      string `json:"country"`
	CustomerId   string `json:"customerId"`
	Channel      string `json:"channel"`
	Commerce     string `json:"commerce"`
	Domain       string `json:"domain"`
	Capability   string `json:"capability"`
	MimeType     string `json:"mimeType"`
	SourceSystem string `json:"sourceSystem"`
	Data         string `json:"data"`
}

const (
	OAUTH2_AUTH          = "Authorization"
	DEFAULT_API_TIME_OUT = 30
)

func PostEventAdapter(token string, req *http.Request) error {
	req.Header.Set(OAUTH2_AUTH, token)
	client := &http.Client{Timeout: DEFAULT_API_TIME_OUT * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}
	return nil
}
