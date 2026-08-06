//go:build integration

// Integration tests run against the live https://test.servicestack.net Services:
//
//	go test -tags integration ./...
//
// Use SERVICESTACK_TEST_URL to run them against a different ServiceStack instance.
package servicestack_test

import (
	"os"
	"strings"
	"testing"

	ss "github.com/ServiceStack/servicestack-go"
)

// Model the ChatCompletion integration test uses, available on test.servicestack.net
const chatModel = "openai/gpt-oss-120b"

func testUrl() string {
	if baseUrl := os.Getenv("SERVICESTACK_TEST_URL"); baseUrl != "" {
		return baseUrl
	}
	return "https://test.servicestack.net"
}

// Hand-written DTOs matching the remote Services, generated DTOs have the same shape.

type Hello struct {
	Name string `json:"name,omitempty"`
}

func (Hello) CreateResponse() (r HelloResponse) { return }
func (Hello) HttpMethod() string                { return ss.HttpGet }

type HelloResponse struct {
	Result string `json:"result,omitempty"`
}

type ThrowValidation struct {
	Age   int    `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
}

func (ThrowValidation) CreateResponse() (r ThrowValidationResponse) { return }
func (ThrowValidation) HttpMethod() string                          { return ss.HttpPost }

type ThrowValidationResponse struct {
	Age            int                `json:"age,omitempty"`
	Email          string             `json:"email,omitempty"`
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type ThrowType struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

func (ThrowType) CreateResponse() (r ThrowTypeResponse) { return }
func (ThrowType) HttpMethod() string                    { return ss.HttpGet }

type ThrowTypeResponse struct {
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type HelloSecure struct {
	Name string `json:"name,omitempty"`
}

func (HelloSecure) CreateResponse() (r HelloResponse) { return }
func (HelloSecure) HttpMethod() string                { return ss.HttpGet }

func TestIntegrationSend(t *testing.T) {
	client := ss.NewClient(testUrl())

	res, err := ss.Send(client, Hello{Name: "World"})
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got %q, want Hello, World!", res.Result)
	}
}

func TestIntegrationValidationErrors(t *testing.T) {
	client := ss.NewClient(testUrl())

	_, err := ss.Send(client, ThrowValidation{})
	webEx, ok := ss.AsWebServiceException(err)
	if !ok {
		t.Fatalf("got %T, want *WebServiceException", err)
	}
	if webEx.StatusCode != 400 {
		t.Errorf("got StatusCode %d, want 400", webEx.StatusCode)
	}
	if !webEx.IsValidationError() {
		t.Fatalf("expected field errors, got %+v", webEx.ResponseStatus)
	}
	if !strings.Contains(webEx.FieldError("Age"), "must be between 1 and 120") {
		t.Errorf("got Age FieldError %q", webEx.FieldError("Age"))
	}
	if webEx.FieldError("Email") == "" {
		t.Error("expected an Email field error")
	}
}

func TestIntegrationErrorStatusCodes(t *testing.T) {
	client := ss.NewClient(testUrl())

	_, err := ss.Send(client, ThrowType{Type: "NotFound", Message: "Not Here"})
	webEx, ok := ss.AsWebServiceException(err)
	if !ok {
		t.Fatalf("got %T, want *WebServiceException", err)
	}
	if !webEx.IsNotFound() {
		t.Errorf("got StatusCode %d, want 404", webEx.StatusCode)
	}
	if webEx.ErrorCode() != "NotFound" || webEx.ErrorMessage() != "Not Here" {
		t.Errorf("got %q: %q", webEx.ErrorCode(), webEx.ErrorMessage())
	}
}

func TestIntegrationUnauthorized(t *testing.T) {
	client := ss.NewClient(testUrl())

	api := ss.Api(client, HelloSecure{Name: "World"})
	if api.Succeeded() {
		t.Fatal("expected HelloSecure to require authentication")
	}
	if api.ErrorCode() != "Unauthorized" {
		t.Errorf("got ErrorCode %q, want Unauthorized", api.ErrorCode())
	}
}

func TestIntegrationAuthenticateThenCallSecureService(t *testing.T) {
	client := ss.NewClient(testUrl())

	authRes, err := client.Authenticate("test", "test")
	if err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}
	if authRes.UserName != "test" {
		t.Errorf("got UserName %q, want test", authRes.UserName)
	}

	// The authenticated Session is maintained in the Client's cookie jar
	res, err := ss.Send(client, HelloSecure{Name: "World"})
	if err != nil {
		t.Fatalf("HelloSecure failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got %q, want Hello, World!", res.Result)
	}
}

func TestIntegrationSendAll(t *testing.T) {
	client := ss.NewClient(testUrl())

	responses, err := ss.SendAll(client, []Hello{{Name: "A"}, {Name: "B"}})
	if err != nil {
		t.Fatalf("SendAll failed: %v", err)
	}
	if len(responses) != 2 {
		t.Fatalf("got %d responses, want 2", len(responses))
	}
	if responses[0].Result != "Hello, A!" || responses[1].Result != "Hello, B!" {
		t.Errorf("got %+v", responses)
	}
}

func TestIntegrationCustomRoute(t *testing.T) {
	client := ss.NewClient(testUrl())

	res, err := ss.GetUrl[HelloResponse](client, "/hello/World")
	if err != nil {
		t.Fatalf("GetUrl failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got %q, want Hello, World!", res.Result)
	}
}

// ── AI Chat ──

// DTOs of ServiceStack's AI Chat ChatCompletion API, an OpenAI-compatible
// Chat Completions endpoint.

type AiTextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type AiMessage struct {
	Role    string `json:"role"`
	Content []any  `json:"content,omitempty"`
}

type ChatCompletion struct {
	Model    string      `json:"model"`
	Messages []AiMessage `json:"messages"`
}

func (ChatCompletion) CreateResponse() (r ChatResponse) { return }
func (ChatCompletion) HttpMethod() string               { return ss.HttpPost }

type ChoiceMessage struct {
	Role      string `json:"role,omitempty"`
	Content   string `json:"content,omitempty"`
	Reasoning string `json:"reasoning,omitempty"`
}

type Choice struct {
	Index        int           `json:"index,omitempty"`
	FinishReason string        `json:"finish_reason,omitempty"`
	Message      ChoiceMessage `json:"message"`
}

type ChatResponse struct {
	Id      string   `json:"id,omitempty"`
	Model   string   `json:"model,omitempty"`
	Choices []Choice `json:"choices,omitempty"`
}

func TestIntegrationChatCompletion(t *testing.T) {
	client := ss.NewClient(testUrl())

	// The ChatCompletion API requires an authenticated User
	if _, err := client.Authenticate("test", "test"); err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}

	res, err := ss.Send(client, ChatCompletion{
		Model: chatModel,
		Messages: []AiMessage{
			{
				Role: "user",
				Content: []any{
					AiTextContent{Type: "text", Text: "Capital of France? Answer in 3 words"},
				},
			},
		},
	})
	if err != nil {
		// A shared LLM can be rate limited or temporarily unavailable
		if webEx, ok := ss.AsWebServiceException(err); ok && webEx.IsAny(429, 502, 503, 504) {
			t.Skipf("ChatCompletion unavailable: %v", err)
		}
		t.Fatalf("ChatCompletion failed: %v", err)
	}

	if len(res.Choices) == 0 {
		t.Fatalf("got no choices, response: %+v", res)
	}
	if res.Choices[0].Message.Content == "" {
		t.Errorf("got an empty message, response: %+v", res)
	}
	if res.Model != chatModel {
		t.Errorf("got model %q, want %q", res.Model, chatModel)
	}
}
