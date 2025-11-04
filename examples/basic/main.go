package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ServiceStack/servicestack-go"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Result string `json:"result"`
}

func main() {
	// Create a new client
	client := servicestack.NewClient("https://example.servicestack.net")

	// Set custom headers if needed
	client.SetHeader("Authorization", "Bearer your-token-here")

	// Make a POST request
	request := HelloRequest{Name: "World"}
	var response HelloResponse

	ctx := context.Background()
	err := client.Post(ctx, "/hello", request, &response)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	fmt.Printf("Response: %s\n", response.Result)
}
