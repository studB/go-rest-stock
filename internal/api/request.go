package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rest-stock/internal/config"
	"time"
)

type HttpClient struct {
	client *http.Client
}

type AuthClient struct {
	http      *HttpClient
	url       string
	appkey    string
	secretkey string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *HttpClient) Post(url string, headers map[string]string, body interface{}) (string, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("JSON marshal error: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("request build error: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("API error: %d - %s", resp.StatusCode, string(data))
	}

	return string(data), nil
}

func NewAuthClient() *AuthClient {
	cfg := config.Load()

	return &AuthClient{
		http:      NewHttpClient(),
		url:       cfg.ApiEndpoint + "/oauth2/tokenP",
		appkey:    cfg.AppKey,
		secretkey: cfg.SecretKey,
	}
}

func (a *AuthClient) RequestToken() (string, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "text/plain",
		"charset":      "UTF-8",
	}

	body := map[string]string{
		"grant_type": "client_credentials",
		"appkey":     a.appkey,
		"appsecret":  a.secretkey,
	}

	resp, err := a.http.Post(a.url, headers, body)
	if err != nil {
		return "", fmt.Errorf("token request failed: %w", err)
	}

	return resp, nil
}
