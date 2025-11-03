package servicestack

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Test DTO types
type HelloRequest struct {
	Name string `json:"name"`
}

func (r *HelloRequest) ResponseType() interface{} {
	return &HelloResponse{}
}

type HelloResponse struct {
	Result string `json:"result"`
}

type AuthenticateRequest struct {
	Provider string `json:"provider"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (r *AuthenticateRequest) ResponseType() interface{} {
	return &AuthenticateResponse{}
}

type AuthenticateResponse struct {
	SessionId      string         `json:"sessionId"`
	UserName       string         `json:"userName"`
	BearerToken    string         `json:"bearerToken"`
	ResponseStatus ResponseStatus `json:"responseStatus,omitempty"`
}

func TestNewJsonServiceClient(t *testing.T) {
	client := NewJsonServiceClient("https://test.servicestack.net")

	if client == nil {
		t.Fatal("Expected client to be created")
	}

	if client.BaseURL != "https://test.servicestack.net" {
		t.Errorf("Expected BaseURL to be 'https://test.servicestack.net', got '%s'", client.BaseURL)
	}

	if client.httpClient == nil {
		t.Fatal("Expected httpClient to be initialized")
	}
}

func TestSetTimeout(t *testing.T) {
	client := NewJsonServiceClient("https://test.servicestack.net")
	client.SetTimeout(10 * time.Second)

	if client.httpClient.Timeout != 10*time.Second {
		t.Errorf("Expected timeout to be 10s, got %v", client.httpClient.Timeout)
	}
}

func TestSetBearerToken(t *testing.T) {
	client := NewJsonServiceClient("https://test.servicestack.net")
	client.SetBearerToken("test-token")

	if client.Headers["Authorization"] != "Bearer test-token" {
		t.Errorf("Expected Authorization header to be 'Bearer test-token', got '%s'", client.Headers["Authorization"])
	}
}

func TestPost(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type to be 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}

		// Parse request body
		var reqBody HelloRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		// Send response
		response := HelloResponse{
			Result: "Hello, " + reqBody.Name + "!",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client and send request
	client := NewJsonServiceClient(server.URL)
	request := &HelloRequest{Name: "World"}

	result, err := client.Post(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	response, ok := result.(*HelloResponse)
	if !ok {
		t.Fatalf("Expected response to be *HelloResponse, got %T", result)
	}

	if response.Result != "Hello, World!" {
		t.Errorf("Expected result to be 'Hello, World!', got '%s'", response.Result)
	}
}

func TestGet(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		// Check query parameters
		name := r.URL.Query().Get("name")
		if name != "World" {
			t.Errorf("Expected name parameter to be 'World', got '%s'", name)
		}

		// Send response
		response := HelloResponse{
			Result: "Hello, " + name + "!",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client and send request
	client := NewJsonServiceClient(server.URL)
	request := &HelloRequest{Name: "World"}

	result, err := client.Get(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	response, ok := result.(*HelloResponse)
	if !ok {
		t.Fatalf("Expected response to be *HelloResponse, got %T", result)
	}

	if response.Result != "Hello, World!" {
		t.Errorf("Expected result to be 'Hello, World!', got '%s'", response.Result)
	}
}

func TestErrorHandling(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		errorResponse := struct {
			ResponseStatus ResponseStatus `json:"responseStatus"`
		}{
			ResponseStatus: ResponseStatus{
				ErrorCode: "ValidationError",
				Message:   "Name is required",
				Errors: []ResponseError{
					{
						ErrorCode: "NotEmpty",
						FieldName: "Name",
						Message:   "Name cannot be empty",
					},
				},
			},
		}
		json.NewEncoder(w).Encode(errorResponse)
	}))
	defer server.Close()

	// Create client and send request
	client := NewJsonServiceClient(server.URL)
	request := &HelloRequest{Name: ""}

	_, err := client.Post(request)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	webEx, ok := err.(*WebServiceException)
	if !ok {
		t.Fatalf("Expected error to be *WebServiceException, got %T", err)
	}

	if webEx.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code to be 400, got %d", webEx.StatusCode)
	}

	if webEx.ResponseStatus.ErrorCode != "ValidationError" {
		t.Errorf("Expected error code to be 'ValidationError', got '%s'", webEx.ResponseStatus.ErrorCode)
	}

	if webEx.ResponseStatus.Message != "Name is required" {
		t.Errorf("Expected error message to be 'Name is required', got '%s'", webEx.ResponseStatus.Message)
	}

	if len(webEx.ResponseStatus.Errors) != 1 {
		t.Fatalf("Expected 1 field error, got %d", len(webEx.ResponseStatus.Errors))
	}

	if webEx.ResponseStatus.Errors[0].FieldName != "Name" {
		t.Errorf("Expected field name to be 'Name', got '%s'", webEx.ResponseStatus.Errors[0].FieldName)
	}
}

func TestBasicAuth(t *testing.T) {
	// Test basic auth encoding
	encoded := basicAuth("user", "pass")
	expected := "dXNlcjpwYXNz"

	if encoded != expected {
		t.Errorf("Expected base64 encoding to be '%s', got '%s'", expected, encoded)
	}
}

func TestSetCredentials(t *testing.T) {
	client := NewJsonServiceClient("https://test.servicestack.net")
	client.SetCredentials("user", "pass")

	authHeader := client.Headers["Authorization"]
	if !startsWith(authHeader, "Basic ") {
		t.Errorf("Expected Authorization header to start with 'Basic ', got '%s'", authHeader)
	}
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func TestToQueryString(t *testing.T) {
	request := &HelloRequest{Name: "World"}

	queryString, err := toQueryString(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if queryString != "name=World" {
		t.Errorf("Expected query string to be 'name=World', got '%s'", queryString)
	}
}

func TestGetRequestPath(t *testing.T) {
	client := NewJsonServiceClient("https://test.servicestack.net")
	request := &HelloRequest{Name: "World"}

	path := client.getRequestPath(request)
	if path != "/json/reply/HelloRequest" {
		t.Errorf("Expected path to be '/json/reply/HelloRequest', got '%s'", path)
	}
}

func TestPut(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		response := HelloResponse{Result: "Updated"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewJsonServiceClient(server.URL)
	request := &HelloRequest{Name: "Update"}

	result, err := client.Put(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	response := result.(*HelloResponse)
	if response.Result != "Updated" {
		t.Errorf("Expected result to be 'Updated', got '%s'", response.Result)
	}
}

func TestDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		response := HelloResponse{Result: "Deleted"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewJsonServiceClient(server.URL)
	request := &HelloRequest{Name: "Delete"}

	result, err := client.Delete(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	response := result.(*HelloResponse)
	if response.Result != "Deleted" {
		t.Errorf("Expected result to be 'Deleted', got '%s'", response.Result)
	}
}

func TestPatch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}

		response := HelloResponse{Result: "Patched"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewJsonServiceClient(server.URL)
	request := &HelloRequest{Name: "Patch"}

	result, err := client.Patch(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	response := result.(*HelloResponse)
	if response.Result != "Patched" {
		t.Errorf("Expected result to be 'Patched', got '%s'", response.Result)
	}
}

func TestCustomHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customHeader := r.Header.Get("X-Custom-Header")
		if customHeader != "custom-value" {
			t.Errorf("Expected X-Custom-Header to be 'custom-value', got '%s'", customHeader)
		}

		response := HelloResponse{Result: "OK"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewJsonServiceClient(server.URL)
	client.Headers["X-Custom-Header"] = "custom-value"

	request := &HelloRequest{Name: "Test"}
	_, err := client.Post(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestWebServiceExceptionError(t *testing.T) {
	webEx := &WebServiceException{
		StatusCode:        400,
		StatusDescription: "Bad Request",
		ResponseStatus: ResponseStatus{
			ErrorCode: "ValidationError",
			Message:   "Validation failed",
		},
	}

	errorMsg := webEx.Error()
	expected := "400 Bad Request: Validation failed"
	if errorMsg != expected {
		t.Errorf("Expected error message to be '%s', got '%s'", expected, errorMsg)
	}
}

func TestWebServiceExceptionErrorWithoutMessage(t *testing.T) {
	webEx := &WebServiceException{
		StatusCode:        500,
		StatusDescription: "Internal Server Error",
		ResponseStatus:    ResponseStatus{},
	}

	errorMsg := webEx.Error()
	expected := "500 Internal Server Error"
	if errorMsg != expected {
		t.Errorf("Expected error message to be '%s', got '%s'", expected, errorMsg)
	}
}

func TestNonServiceStackError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server error"))
	}))
	defer server.Close()

	client := NewJsonServiceClient(server.URL)
	request := &HelloRequest{Name: "Test"}

	_, err := client.Post(request)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	webEx, ok := err.(*WebServiceException)
	if !ok {
		t.Fatalf("Expected error to be *WebServiceException, got %T", err)
	}

	if webEx.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %d", webEx.StatusCode)
	}
}
