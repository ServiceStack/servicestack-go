# servicestack-go

ServiceStack Client Go Library

A Go HTTP client library for consuming ServiceStack services.

## Installation

```bash
go get github.com/ServiceStack/servicestack-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/ServiceStack/servicestack-go"
)

func main() {
    // Create a new client
    client := servicestack.NewClient("https://api.example.com")
    
    // Set custom headers if needed
    client.SetHeader("Authorization", "Bearer your-token")
    
    // Define your request and response types
    type HelloRequest struct {
        Name string `json:"name"`
    }
    
    type HelloResponse struct {
        Result string `json:"result"`
    }
    
    // Make a POST request
    request := HelloRequest{Name: "World"}
    var response HelloResponse
    
    ctx := context.Background()
    err := client.Post(ctx, "/hello", request, &response)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(response.Result)
}
```

## Features

- Support for all HTTP methods: GET, POST, PUT, DELETE, PATCH
- Automatic JSON serialization/deserialization
- Context support for cancellation and timeouts
- Custom headers support
- Simple and idiomatic Go API

## API

### Creating a Client

```go
client := servicestack.NewClient("https://api.example.com")
```

### Setting Custom Headers

```go
client.SetHeader("Authorization", "Bearer token")
client.SetHeader("X-Custom-Header", "value")
```

### Making Requests

#### GET Request
```go
var response MyResponse
err := client.Get(ctx, "/endpoint", &response)
```

#### POST Request
```go
request := MyRequest{...}
var response MyResponse
err := client.Post(ctx, "/endpoint", request, &response)
```

#### PUT Request
```go
request := MyRequest{...}
var response MyResponse
err := client.Put(ctx, "/endpoint", request, &response)
```

#### DELETE Request
```go
var response MyResponse
err := client.Delete(ctx, "/endpoint", &response)
```

#### PATCH Request
```go
request := MyRequest{...}
var response MyResponse
err := client.Patch(ctx, "/endpoint", request, &response)
```

## License

See [LICENSE](LICENSE) for details.
