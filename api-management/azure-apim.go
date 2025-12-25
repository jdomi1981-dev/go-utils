package apimanagement

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

const (
	APIM_SUSCRIPTION      = "Ocp-Apim-Subscription-Key"
	X_API_VERSION         = "x-api-version"
	X_COMMERCE            = "x-commerce"
	HEADER_APP_JSON       = "application/json"
	CONTENT_TYPE          = "Content-Type"
	APIM_DEFAULT_COMMERCE = "IXC"
	SECRET_KEY            = "x-secret-key"
	DEFAULT_API_TIME_OUT  = 60
	EMPTY_STRING          = ""
)

type BearerToken struct {
	Token      string `json:"token"`
	UtcExpDate string `json:"utcExpDate"`
}

func CreateNewHeaderAPIM(method string, url string, apiVersion string, body io.Reader, apimKey string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(X_COMMERCE, APIM_DEFAULT_COMMERCE)
	req.Header.Set(X_API_VERSION, apiVersion)
	req.Header.Set(APIM_SUSCRIPTION, apimKey)
	req.Header.Set(CONTENT_TYPE, HEADER_APP_JSON)

	return req, nil
}

func CreateTokenOAuth2(req *http.Request, secretValue string) (string, error) {
	req.Header.Set(SECRET_KEY, secretValue)
	client := &http.Client{Timeout: DEFAULT_API_TIME_OUT * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return EMPTY_STRING, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return EMPTY_STRING, err
	}
	if resp.StatusCode != http.StatusOK {
		return EMPTY_STRING, errors.New(resp.Status)
	}
	var bearerToken BearerToken
	if err := json.Unmarshal(body, &bearerToken); err != nil {
		return EMPTY_STRING, err
	}
	return bearerToken.Token, nil
}
