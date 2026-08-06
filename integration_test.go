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
	"github.com/ServiceStack/servicestack-go/dtos"
)

// Model the ChatCompletion integration test uses, available on test.servicestack.net
const chatModel = "openai/gpt-oss-120b"

func testUrl() string {
	if baseUrl := os.Getenv("SERVICESTACK_TEST_URL"); baseUrl != "" {
		return baseUrl
	}
	return "https://test.servicestack.net"
}

func TestIntegrationSend(t *testing.T) {
	client := ss.NewClient(testUrl())

	res, err := ss.Send(client, dtos.Hello{Name: "World", Title: "Mr"})
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if res.Result != "Hello, Mr. World!" {
		t.Errorf("got %q, want Hello, Mr. World!", res.Result)
	}
}

func TestIntegrationValidationErrors(t *testing.T) {
	client := ss.NewClient(testUrl())

	_, err := ss.Send(client, dtos.ThrowValidation{})
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

	_, err := ss.Send(client, dtos.ThrowType{Type: "NotFound", Message: "Not Here"})
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

	api := ss.Api(client, dtos.HelloSecure{Name: "World"})
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
	res, err := ss.Send(client, dtos.HelloSecure{Name: "World"})
	if err != nil {
		t.Fatalf("HelloSecure failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got %q, want Hello, World!", res.Result)
	}
}

func TestIntegrationSendAll(t *testing.T) {
	client := ss.NewClient(testUrl())

	responses, err := ss.SendAll(client, []dtos.Hello{{Name: "A", Title: "Mr"}, {Name: "B", Title: "Ms"}})
	if err != nil {
		t.Fatalf("SendAll failed: %v", err)
	}
	if len(responses) != 2 {
		t.Fatalf("got %d responses, want 2", len(responses))
	}
	if responses[0].Result != "Hello, Mr. A!" || responses[1].Result != "Hello, Ms. B!" {
		t.Errorf("got %+v", responses)
	}
}

func TestIntegrationCustomRoute(t *testing.T) {
	client := ss.NewClient(testUrl())

	res, err := ss.GetUrl[dtos.HelloResponse](client, "/hello/World")
	if err != nil {
		t.Fatalf("GetUrl failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got %q, want Hello, World!", res.Result)
	}
}

// Sends a Request to ServiceStack AI Chat's OpenAI-compatible ChatCompletion API
func TestIntegrationChatCompletion(t *testing.T) {
	client := ss.NewClient(testUrl())

	// The ChatCompletion API requires an authenticated User
	if _, err := client.Authenticate("test", "test"); err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}

	res, err := ss.Send(client, dtos.ChatCompletion{
		Model: chatModel,
		Messages: []dtos.AiMessage{
			{
				Role: "user",
				// Content parts are polymorphic, e.g. text, image_url or input_audio
				Content: []any{
					dtos.AiTextContent{
						AiContent: dtos.AiContent{Type: "text"},
						Text:      "Capital of France? Answer in 3 words",
					},
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
