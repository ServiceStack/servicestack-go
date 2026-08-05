// Command hello calls the live https://test.servicestack.net Services with
// the ServiceStack Go Client:
//
//	go run ./examples/hello
package main

import (
	"fmt"

	ss "github.com/ServiceStack/servicestack-go"
)

// DTOs of the remote API, normally generated with:
//
//	npx get-dtos go https://test.servicestack.net

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

type HelloSecure struct {
	Name string `json:"name,omitempty"`
}

func (HelloSecure) CreateResponse() (r HelloResponse) { return }
func (HelloSecure) HttpMethod() string                { return ss.HttpGet }

func main() {
	client := ss.NewClient("https://test.servicestack.net")

	// Typed API Request, res is a HelloResponse
	res, err := ss.Send(client, Hello{Name: "World"})
	if err != nil {
		fmt.Println("Hello failed:", err)
		return
	}
	fmt.Println("Send:", res.Result)

	// Batched Requests
	responses, err := ss.SendAll(client, []Hello{{Name: "A"}, {Name: "B"}})
	if err != nil {
		fmt.Println("SendAll failed:", err)
		return
	}
	for _, response := range responses {
		fmt.Println("SendAll:", response.Result)
	}

	// Structured validation errors
	if _, err := ss.Send(client, ThrowValidation{}); err != nil {
		if webEx, ok := ss.AsWebServiceException(err); ok {
			fmt.Printf("ThrowValidation: %d %s\n", webEx.StatusCode, webEx.ErrorMessage())
			for _, fieldError := range webEx.FieldErrors() {
				fmt.Printf("  %s: %s\n", fieldError.FieldName, fieldError.Message)
			}
		}
	}

	// Authenticated Requests
	if _, err := client.Authenticate("test", "test"); err != nil {
		fmt.Println("Authenticate failed:", err)
		return
	}
	secure, err := ss.Send(client, HelloSecure{Name: "World"})
	if err != nil {
		fmt.Println("HelloSecure failed:", err)
		return
	}
	fmt.Println("HelloSecure:", secure.Result)
}
