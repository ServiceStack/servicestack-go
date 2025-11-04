package servicestack

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestRequest struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type TestResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func TestNewClient(t *testing.T) {
	client := NewClient("https://api.example.com")

	if client == nil {
		t.Fatal("Expected client to be created")
	}

	if client.BaseURL != "https://api.example.com" {
		t.Errorf("Expected BaseURL to be 'https://api.example.com', got '%s'", client.BaseURL)
	}

	if client.HTTPClient == nil {
		t.Error("Expected HTTPClient to be initialized")
	}

	if client.Headers == nil {
		t.Error("Expected Headers to be initialized")
	}
}

func TestSetHeader(t *testing.T) {
	client := NewClient("https://api.example.com")
	client.SetHeader("Authorization", "Bearer token123")

	if client.Headers["Authorization"] != "Bearer token123" {
		t.Errorf("Expected Authorization header to be 'Bearer token123', got '%s'", client.Headers["Authorization"])
	}
}

func TestGet(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header to be 'application/json', got '%s'", r.Header.Get("Accept"))
		}

		response := TestResponse{
			Message: "Success",
			Status:  "OK",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	var response TestResponse

	ctx := context.Background()
	err := client.Get(ctx, "/test", &response)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Message != "Success" {
		t.Errorf("Expected message 'Success', got '%s'", response.Message)
	}

	if response.Status != "OK" {
		t.Errorf("Expected status 'OK', got '%s'", response.Status)
	}
}

func TestPost(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}

		var req TestRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		if req.Name != "test" {
			t.Errorf("Expected name 'test', got '%s'", req.Name)
		}

		if req.Value != 42 {
			t.Errorf("Expected value 42, got %d", req.Value)
		}

		response := TestResponse{
			Message: "Created",
			Status:  "OK",
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	request := TestRequest{
		Name:  "test",
		Value: 42,
	}
	var response TestResponse

	ctx := context.Background()
	err := client.Post(ctx, "/test", request, &response)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Message != "Created" {
		t.Errorf("Expected message 'Created', got '%s'", response.Message)
	}
}

func TestPut(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		response := TestResponse{
			Message: "Updated",
			Status:  "OK",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	request := TestRequest{Name: "update", Value: 100}
	var response TestResponse

	ctx := context.Background()
	err := client.Put(ctx, "/test", request, &response)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Message != "Updated" {
		t.Errorf("Expected message 'Updated', got '%s'", response.Message)
	}
}

func TestDelete(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		response := TestResponse{
			Message: "Deleted",
			Status:  "OK",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	var response TestResponse

	ctx := context.Background()
	err := client.Delete(ctx, "/test", &response)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Message != "Deleted" {
		t.Errorf("Expected message 'Deleted', got '%s'", response.Message)
	}
}

func TestPatch(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH method, got %s", r.Method)
		}

		response := TestResponse{
			Message: "Patched",
			Status:  "OK",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	request := TestRequest{Name: "patch", Value: 50}
	var response TestResponse

	ctx := context.Background()
	err := client.Patch(ctx, "/test", request, &response)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Message != "Patched" {
		t.Errorf("Expected message 'Patched', got '%s'", response.Message)
	}
}

func TestCustomHeaders(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom-Header") != "custom-value" {
			t.Errorf("Expected X-Custom-Header to be 'custom-value', got '%s'", r.Header.Get("X-Custom-Header"))
		}

		response := TestResponse{
			Message: "Success",
			Status:  "OK",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	client.SetHeader("X-Custom-Header", "custom-value")
	var response TestResponse

	ctx := context.Background()
	err := client.Get(ctx, "/test", &response)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestErrorResponse(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	var response TestResponse

	ctx := context.Background()
	err := client.Get(ctx, "/test", &response)

	if err == nil {
		t.Fatal("Expected an error for 400 status code")
	}
}

func TestContextCancellation(t *testing.T) {
	// Create a test server with a delay
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		response := TestResponse{
			Message: "Success",
			Status:  "OK",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	var response TestResponse

	// Create a context with a very short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := client.Get(ctx, "/test", &response)

	if err == nil {
		t.Fatal("Expected an error due to context cancellation")
	}
}
