package servicestack

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IReturn interface for typed requests that return a response
type IReturn interface {
	ResponseType() interface{}
}

// IReturnVoid interface for requests that don't return a response
type IReturnVoid interface {
	CreateResponse() interface{}
}

// IVerb interface for requests that specify HTTP verb
type IVerb interface {
	GetMethod() string
}

// IGet interface for GET requests
type IGet interface {
	IReturn
}

// IPost interface for POST requests
type IPost interface {
	IReturn
}

// IPut interface for PUT requests
type IPut interface {
	IReturn
}

// IDelete interface for DELETE requests
type IDelete interface {
	IReturn
}

// IPatch interface for PATCH requests
type IPatch interface {
	IReturn
}

// ResponseStatus represents ServiceStack error response status
type ResponseStatus struct {
	ErrorCode  string            `json:"errorCode,omitempty"`
	Message    string            `json:"message,omitempty"`
	StackTrace string            `json:"stackTrace,omitempty"`
	Errors     []ResponseError   `json:"errors,omitempty"`
	Meta       map[string]string `json:"meta,omitempty"`
}

// ResponseError represents field-level validation errors
type ResponseError struct {
	ErrorCode string            `json:"errorCode,omitempty"`
	FieldName string            `json:"fieldName,omitempty"`
	Message   string            `json:"message,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
}

// WebServiceException represents ServiceStack service errors
type WebServiceException struct {
	StatusCode        int
	StatusDescription string
	ResponseStatus    ResponseStatus
}

func (e *WebServiceException) Error() string {
	if e.ResponseStatus.Message != "" {
		return fmt.Sprintf("%d %s: %s", e.StatusCode, e.StatusDescription, e.ResponseStatus.Message)
	}
	return fmt.Sprintf("%d %s", e.StatusCode, e.StatusDescription)
}

// JsonServiceClient is the main client for making ServiceStack API requests
type JsonServiceClient struct {
	BaseURL    string
	httpClient *http.Client
	Headers    map[string]string
}

// NewJsonServiceClient creates a new JsonServiceClient
func NewJsonServiceClient(baseURL string) *JsonServiceClient {
	return &JsonServiceClient{
		BaseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Headers: make(map[string]string),
	}
}

// SetTimeout sets the HTTP client timeout
func (c *JsonServiceClient) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// SetBearerToken sets the Bearer token for authentication
func (c *JsonServiceClient) SetBearerToken(token string) {
	c.Headers["Authorization"] = "Bearer " + token
}

// SetCredentials sets basic authentication credentials
func (c *JsonServiceClient) SetCredentials(username, password string) {
	c.Headers["Authorization"] = "Basic " + basicAuth(username, password)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Get sends a GET request
func (c *JsonServiceClient) Get(request IReturn) (interface{}, error) {
	return c.Send("GET", request, request.ResponseType())
}

// Post sends a POST request
func (c *JsonServiceClient) Post(request IReturn) (interface{}, error) {
	return c.Send("POST", request, request.ResponseType())
}

// Put sends a PUT request
func (c *JsonServiceClient) Put(request IReturn) (interface{}, error) {
	return c.Send("PUT", request, request.ResponseType())
}

// Delete sends a DELETE request
func (c *JsonServiceClient) Delete(request IReturn) (interface{}, error) {
	return c.Send("DELETE", request, request.ResponseType())
}

// Patch sends a PATCH request
func (c *JsonServiceClient) Patch(request IReturn) (interface{}, error) {
	return c.Send("PATCH", request, request.ResponseType())
}

// Send sends a request with the specified HTTP method
func (c *JsonServiceClient) Send(method string, request interface{}, responseType interface{}) (interface{}, error) {
	// Determine the request path
	requestPath := c.getRequestPath(request)
	requestURL := c.BaseURL + requestPath

	var body io.Reader
	var err error

	// For GET and DELETE, add query string parameters
	if method == "GET" || method == "DELETE" {
		params, err := toQueryString(request)
		if err != nil {
			return nil, err
		}
		if params != "" {
			requestURL += "?" + params
		}
	} else {
		// For POST, PUT, PATCH, send JSON body
		jsonData, err := json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %w", err)
		}
		body = bytes.NewReader(jsonData)
	}

	// Create HTTP request
	req, err := http.NewRequest(method, requestURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, c.parseError(resp.StatusCode, resp.Status, respBody)
	}

	// Parse successful response
	if responseType != nil {
		// responseType is already a pointer to a new instance from ResponseType()
		if err := json.Unmarshal(respBody, responseType); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return responseType, nil
	}

	return nil, nil
}

// getRequestPath extracts the request path from the request type name
func (c *JsonServiceClient) getRequestPath(request interface{}) string {
	// Get the type name and use it as the path
	typeName := fmt.Sprintf("%T", request)

	// Remove package prefix if present
	parts := strings.Split(typeName, ".")
	if len(parts) > 1 {
		typeName = parts[len(parts)-1]
	}

	// Remove pointer prefix if present
	typeName = strings.TrimPrefix(typeName, "*")

	return "/json/reply/" + typeName
}

// toQueryString converts a struct to URL query string parameters
func toQueryString(v interface{}) (string, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", err
	}

	params := url.Values{}
	for key, value := range data {
		if value != nil {
			params.Add(key, fmt.Sprintf("%v", value))
		}
	}

	return params.Encode(), nil
}

// parseError parses ServiceStack error response
func (c *JsonServiceClient) parseError(statusCode int, statusDescription string, body []byte) error {
	var errorResponse struct {
		ResponseStatus ResponseStatus `json:"responseStatus"`
	}

	if err := json.Unmarshal(body, &errorResponse); err != nil {
		// If we can't parse as ServiceStack error, return generic error
		return &WebServiceException{
			StatusCode:        statusCode,
			StatusDescription: statusDescription,
			ResponseStatus: ResponseStatus{
				Message: string(body),
			},
		}
	}

	return &WebServiceException{
		StatusCode:        statusCode,
		StatusDescription: statusDescription,
		ResponseStatus:    errorResponse.ResponseStatus,
	}
}
