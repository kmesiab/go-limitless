package go_limitless

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const BaseURL = "https://api.limitless.ai/v1"

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

// NewClient creates a new Limitless API client
func NewClient(apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    BaseURL,
		APIKey:     apiKey,
	}
}

// GetLifelogs retrieves a list of lifelogs
func (c *Client) GetLifelogs(ctx context.Context, params map[string]string) (*LifelogsResponse, error) {
	endpoint, err := url.Parse(fmt.Sprintf("%s/lifelogs", c.BaseURL))
	if err != nil {
		return nil, err
	}

	query := endpoint.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", c.APIKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	var lifelogsResponse LifelogsResponse
	if err := json.NewDecoder(resp.Body).Decode(&lifelogsResponse); err != nil {
		return nil, err
	}

	return &lifelogsResponse, nil
}

// GetLifelog retrieves details of a single lifelog entry by ID
func (c *Client) GetLifelog(ctx context.Context, lifelogID string) (*Lifelog, error) {
	endpoint := fmt.Sprintf("%s/lifelogs/%s", c.BaseURL, lifelogID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", c.APIKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("[Client.GetLifelog] Error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	var lifelog Lifelog
	if err := json.NewDecoder(resp.Body).Decode(&lifelog); err != nil {
		return nil, err
	}

	return &lifelog, nil
}
