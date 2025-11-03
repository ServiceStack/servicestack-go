// Package servicestack provides a Go client library for ServiceStack services
package servicestack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client is the ServiceStack HTTP client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Headers    map[string]string
}

// NewClient creates a new ServiceStack client with the given base URL
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Headers: make(map[string]string),
	}
}

// SetHeader sets a custom header for all requests
func (c *Client) SetHeader(key, value string) {
	c.Headers[key] = value
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, path string, response interface{}) error {
	return c.doRequest(ctx, http.MethodGet, path, nil, response)
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, request, response interface{}) error {
	return c.doRequest(ctx, http.MethodPost, path, request, response)
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, path string, request, response interface{}) error {
	return c.doRequest(ctx, http.MethodPut, path, request, response)
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string, response interface{}) error {
	return c.doRequest(ctx, http.MethodDelete, path, nil, response)
}

// Patch performs a PATCH request
func (c *Client) Patch(ctx context.Context, path string, request, response interface{}) error {
	return c.doRequest(ctx, http.MethodPatch, path, request, response)
}

// doRequest performs the actual HTTP request
func (c *Client) doRequest(ctx context.Context, method, path string, request, response interface{}) error {
	// Build full URL
	fullURL, err := url.JoinPath(c.BaseURL, path)
	if err != nil {
		return fmt.Errorf("failed to build URL: %w", err)
	}

	// Prepare request body
	var body io.Reader
	if request != nil {
		jsonData, err := json.Marshal(request)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	if request != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	// Execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Unmarshal response
	if response != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, response); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
