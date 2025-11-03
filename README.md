# ServiceStack Go Client

A Go client library for consuming ServiceStack services using typed DTOs.

## Features

- 🚀 Typed request/response DTOs
- 🔒 Built-in authentication support (Bearer Token & Basic Auth)
- 🔄 Support for all HTTP verbs (GET, POST, PUT, DELETE, PATCH)
- ⚠️ ServiceStack error handling with field-level validation errors
- 📦 Zero external dependencies (uses only Go standard library)
- ✅ Full test coverage

## Installation

```bash
go get github.com/ServiceStack/servicestack-go
```

## Quick Start

### 1. Define Your DTOs

Typically, DTOs are generated using ServiceStack's code generation tools. For Go, use:

```bash
x go
```

Or manually define your DTOs:

```go
package main

import "github.com/ServiceStack/servicestack-go"

// Request DTO
type HelloRequest struct {
    Name string `json:"name"`
}

// Implement IReturn interface to specify response type
func (r *HelloRequest) ResponseType() interface{} {
    return &HelloResponse{}
}

// Response DTO
type HelloResponse struct {
    Result string `json:"result"`
}
```

### 2. Create a Client

```go
client := servicestack.NewJsonServiceClient("https://your-service.com")
```

### 3. Make Requests

```go
// GET request
request := &HelloRequest{Name: "World"}
result, err := client.Get(request)
if err != nil {
    log.Fatal(err)
}
response := result.(*HelloResponse)
fmt.Println(response.Result)

// POST request
result, err = client.Post(request)

// PUT, DELETE, PATCH also supported
result, err = client.Put(request)
result, err = client.Delete(request)
result, err = client.Patch(request)
```

## Authentication

### Bearer Token

```go
client := servicestack.NewJsonServiceClient("https://your-service.com")
client.SetBearerToken("your-token-here")

// Now all requests include: Authorization: Bearer your-token-here
```

### Basic Authentication

```go
client := servicestack.NewJsonServiceClient("https://your-service.com")
client.SetCredentials("username", "password")

// Now all requests include: Authorization: Basic <base64-encoded-credentials>
```

### ServiceStack Authentication

```go
type AuthenticateRequest struct {
    Provider string `json:"provider"`
    UserName string `json:"userName"`
    Password string `json:"password"`
}

func (r *AuthenticateRequest) ResponseType() interface{} {
    return &AuthenticateResponse{}
}

type AuthenticateResponse struct {
    SessionId   string `json:"sessionId"`
    BearerToken string `json:"bearerToken"`
}

// Authenticate
authRequest := &AuthenticateRequest{
    Provider: "credentials",
    UserName: "user",
    Password: "pass",
}

result, err := client.Post(authRequest)
if err != nil {
    log.Fatal(err)
}

authResponse := result.(*AuthenticateResponse)
client.SetBearerToken(authResponse.BearerToken)
```

## Error Handling

ServiceStack errors include detailed validation information:

```go
result, err := client.Post(request)
if err != nil {
    if webEx, ok := err.(*servicestack.WebServiceException); ok {
        fmt.Printf("Error: %s - %s\n", 
            webEx.ResponseStatus.ErrorCode, 
            webEx.ResponseStatus.Message)
        
        // Handle field-level validation errors
        for _, fieldError := range webEx.ResponseStatus.Errors {
            fmt.Printf("Field '%s': %s\n", 
                fieldError.FieldName, 
                fieldError.Message)
        }
    } else {
        log.Fatal(err)
    }
}
```

## Configuration

### Custom Timeout

```go
client := servicestack.NewJsonServiceClient("https://your-service.com")
client.SetTimeout(60 * time.Second)
```

### Custom Headers

```go
client := servicestack.NewJsonServiceClient("https://your-service.com")
client.Headers["X-Custom-Header"] = "value"
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/ServiceStack/servicestack-go"
)

type HelloRequest struct {
    Name string `json:"name"`
}

func (r *HelloRequest) ResponseType() interface{} {
    return &HelloResponse{}
}

type HelloResponse struct {
    Result string `json:"result"`
}

func main() {
    // Create client
    client := servicestack.NewJsonServiceClient("https://test.servicestack.net")
    
    // Make request
    request := &HelloRequest{Name: "World"}
    result, err := client.Post(request)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use response
    response := result.(*HelloResponse)
    fmt.Println(response.Result)
}
```

## API Reference

### Client Methods

- `NewJsonServiceClient(baseURL string)` - Create a new client
- `Get(request IReturn)` - Send a GET request
- `Post(request IReturn)` - Send a POST request
- `Put(request IReturn)` - Send a PUT request
- `Delete(request IReturn)` - Send a DELETE request
- `Patch(request IReturn)` - Send a PATCH request
- `Send(method string, request interface{}, responseType interface{})` - Send with custom method
- `SetTimeout(timeout time.Duration)` - Set request timeout
- `SetBearerToken(token string)` - Set bearer token authentication
- `SetCredentials(username, password string)` - Set basic authentication

### Interfaces

- `IReturn` - Implemented by request DTOs that return a response
- `ResponseType() interface{}` - Returns the expected response type

### Types

- `ResponseStatus` - ServiceStack error response status
- `ResponseError` - Field-level validation error
- `WebServiceException` - ServiceStack service exception

## Running Tests

```bash
go test -v
```

## Running the Example

```bash
cd examples
go run main.go
```

## License

This library is released under the same license as ServiceStack.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
