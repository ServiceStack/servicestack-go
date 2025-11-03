package main

import (
	"fmt"
	"log"

	"github.com/ServiceStack/servicestack-go"
)

// Example DTOs - typically these would be generated using ServiceStack's Go code generation
// from your ServiceStack services using: x go

// HelloRequest is a sample request DTO
type HelloRequest struct {
	Name string `json:"name"`
}

// ResponseType returns the type of response expected
func (r *HelloRequest) ResponseType() interface{} {
	return &HelloResponse{}
}

// HelloResponse is a sample response DTO
type HelloResponse struct {
	Result string `json:"result"`
}

// AuthenticateRequest is a request to authenticate
type AuthenticateRequest struct {
	Provider string `json:"provider"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (r *AuthenticateRequest) ResponseType() interface{} {
	return &AuthenticateResponse{}
}

// AuthenticateResponse contains authentication result
type AuthenticateResponse struct {
	SessionId      string                      `json:"sessionId"`
	UserName       string                      `json:"userName"`
	BearerToken    string                      `json:"bearerToken"`
	ResponseStatus servicestack.ResponseStatus `json:"responseStatus,omitempty"`
}

func main() {
	// Create a new JsonServiceClient
	client := servicestack.NewJsonServiceClient("https://test.servicestack.net")

	// Example 1: Simple GET request
	fmt.Println("Example 1: GET request")
	getRequest := &HelloRequest{Name: "World"}
	result, err := client.Get(getRequest)
	if err != nil {
		if webEx, ok := err.(*servicestack.WebServiceException); ok {
			log.Printf("Error: %s - %s\n", webEx.ResponseStatus.ErrorCode, webEx.ResponseStatus.Message)
		} else {
			log.Printf("Error: %v\n", err)
		}
	} else {
		response := result.(*HelloResponse)
		fmt.Printf("GET Response: %s\n\n", response.Result)
	}

	// Example 2: Simple POST request
	fmt.Println("Example 2: POST request")
	postRequest := &HelloRequest{Name: "ServiceStack"}
	result, err = client.Post(postRequest)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		response := result.(*HelloResponse)
		fmt.Printf("POST Response: %s\n\n", response.Result)
	}

	// Example 3: Using authentication
	fmt.Println("Example 3: Authentication")
	authRequest := &AuthenticateRequest{
		Provider: "credentials",
		UserName: "test",
		Password: "test",
	}
	result, err = client.Post(authRequest)
	if err != nil {
		log.Printf("Auth Error: %v\n", err)
	} else {
		authResponse := result.(*AuthenticateResponse)
		fmt.Printf("Authenticated as: %s\n", authResponse.UserName)
		
		// Set bearer token for subsequent requests
		if authResponse.BearerToken != "" {
			client.SetBearerToken(authResponse.BearerToken)
			fmt.Println("Bearer token set for future requests\n")
		}
	}

	// Example 4: Using basic authentication
	fmt.Println("Example 4: Basic Authentication")
	client2 := servicestack.NewJsonServiceClient("https://test.servicestack.net")
	client2.SetCredentials("username", "password")
	fmt.Println("Basic auth credentials set\n")

	// Example 5: Error handling
	fmt.Println("Example 5: Error Handling")
	invalidRequest := &HelloRequest{Name: ""} // Assuming empty name might cause validation error
	_, err = client.Post(invalidRequest)
	if err != nil {
		if webEx, ok := err.(*servicestack.WebServiceException); ok {
			fmt.Printf("Service Error: %s\n", webEx.ResponseStatus.Message)
			if len(webEx.ResponseStatus.Errors) > 0 {
				fmt.Println("Validation Errors:")
				for _, fieldError := range webEx.ResponseStatus.Errors {
					fmt.Printf("  - %s: %s\n", fieldError.FieldName, fieldError.Message)
				}
			}
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
