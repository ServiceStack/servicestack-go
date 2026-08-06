# servicestack-go

Typed Go Client Library for consuming [ServiceStack](https://servicestack.net) APIs.

- Typed Request/Response DTOs, generated from any ServiceStack API
- Response Types inferred from the Request DTO — no type assertions, no `interface{}`
- Structured `ResponseStatus` errors with field validation errors
- Auth with Basic Auth, API Keys, JWT Bearer Tokens, Refresh Tokens and Session Cookies
- Batched Requests, one-way Requests and multipart file uploads
- `context.Context` support on every API
- Zero dependencies, only the Go standard library

## Install

```bash
go get github.com/ServiceStack/servicestack-go
```

Requires Go 1.21+.

## Generate Typed DTOs

Generate the Go DTOs of any ServiceStack API with the [get-dtos](https://www.npmjs.com/package/get-dtos) tool:

```bash
mkdir dtos && cd dtos
npx get-dtos go https://blazor-vue.web-templates.io
```

Which downloads a `dtos.go` in the `dtos` package containing the typed DTOs of the remote API:

```go
package dtos

import ss "github.com/ServiceStack/servicestack-go"

// @Route("/hello/{Name}")
type Hello struct {
    Name string `json:"name,omitempty"`
}

func (Hello) CreateResponse() (r HelloResponse) { return }
func (Hello) HttpMethod() string                { return "GET" }

type HelloResponse struct {
    Result         string             `json:"result"`
    ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}
```

The generated `CreateResponse()` method is what lets the client infer the Response Type
of each API, whilst `HttpMethod()` returns the HTTP Verb the API should be called with.

## Usage

```go
package main

import (
    "fmt"

    ss "github.com/ServiceStack/servicestack-go"
    "myapp/dtos"
)

func main() {
    client := ss.NewClient("https://blazor-vue.web-templates.io")

    res, err := ss.Send(client, dtos.Hello{Name: "World"}) // res is a dtos.HelloResponse
    if err != nil {
        panic(err)
    }
    fmt.Println(res.Result)
}
```

`Send` uses the HTTP Method the API is annotated with, use `Get`, `Post`, `Put`,
`Patch` or `Delete` to send a Request DTO with a specific HTTP Method:

```go
res, err := ss.Post(client, dtos.Hello{Name: "World"})
```

APIs that don't return a Response Body are sent with `SendVoid`:

```go
err := ss.SendVoid(client, dtos.DeleteBooking{Id: 1})
```

Additional QueryString params can be sent with any Request:

```go
res, err := ss.Get(client, dtos.Hello{Name: "World"}, map[string]any{"format": "json"})
```

### AutoQuery

AutoQuery APIs return a typed `QueryResponse[T]`:

```go
take := 5
res, err := ss.Send(client, dtos.QueryBookings{
    QueryDb: ss.QueryDb{QueryBase: ss.QueryBase{Take: &take, OrderByDesc: "id"}},
})
for _, booking := range res.Results { // booking is a dtos.Booking
    fmt.Println(booking.Id, booking.Name)
}
```

### Error Handling

Failed API Requests return a `*WebServiceException` containing the HTTP Status Code
and the API's structured `ResponseStatus` error:

```go
_, err := ss.Send(client, dtos.CreateBooking{})
if webEx, ok := ss.AsWebServiceException(err); ok {
    fmt.Println(webEx.StatusCode)          // 400
    fmt.Println(webEx.ErrorCode())         // "NotEmpty"
    fmt.Println(webEx.ErrorMessage())      // "'Name' must not be empty."
    fmt.Println(webEx.FieldError("Name"))  // "'Name' must not be empty."
    fmt.Println(webEx.IsUnauthorized())    // false
}
```

It also works with the standard `errors` package:

```go
var webEx *ss.WebServiceException
if errors.As(err, &webEx) { /* ... */ }
```

Alternatively `Api` returns errors in its result instead of a separate `error`:

```go
api := ss.Api(client, dtos.CreateBooking{})
if api.Failed() {
    fmt.Println(api.ErrorCode(), api.ErrorMessage(), api.FieldError("Name"))
} else {
    fmt.Println(api.Response.Id)
}
```

### Authentication

API Keys and JWTs are sent in the Bearer Token Authorization header:

```go
client.SetBearerToken("ak-87949de37e894627a9f6173154e7cafa")
```

HTTP Basic Auth credentials:

```go
client.SetCredentials("username", "password")
```

Sign in with ServiceStack's Authenticate API, which maintains the authenticated
Session in the Client's cookie jar and uses any Bearer Token the Server returns:

```go
authRes, err := client.Authenticate("username", "password")
```

When a Refresh Token is configured, expired Bearer Tokens are transparently
refreshed and the failed Request retried:

```go
client.SetRefreshToken(refreshToken)
```

Otherwise `OnAuthenticationRequired` can be used to re-authenticate before a
`401 Unauthorized` Request is retried:

```go
client.OnAuthenticationRequired = func(c *ss.Client) error {
    _, err := c.Authenticate("username", "password")
    return err
}
```

### Batched Requests

Send multiple Requests of the same Type in a single HTTP Request:

```go
responses, err := ss.SendAll(client, []dtos.Hello{{Name: "A"}, {Name: "B"}})
```

Or send them to a one-way endpoint that ignores their Responses:

```go
err := ss.PublishAll(client, []dtos.Hello{{Name: "A"}, {Name: "B"}})
```

### File Uploads

```go
file, _ := os.Open("photo.png")
defer file.Close()

res, err := ss.PostFileWithRequest(client, dtos.UploadPhoto{Album: "Holiday"}, ss.UploadFile{
    FieldName:   "file",
    FileName:    "photo.png",
    ContentType: "image/png",
    Reader:      file,
})
```

### Custom URLs

Send Requests to custom routes or external URLs:

```go
res, err := ss.GetUrl[dtos.HelloResponse](client, "/hello/World")
res, err := ss.PostUrl[dtos.HelloResponse](client, "/hello", dtos.Hello{Name: "World"})
```

Any Response Type can be requested, including `string` and `[]byte`:

```go
csv, err := ss.GetUrl[string](client, "/api/QueryBookings.csv")
```

### context.Context

Every API has a `*Ctx` variant accepting a `context.Context`:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

res, err := ss.SendCtx(ctx, client, dtos.Hello{Name: "World"})
```

### Client Configuration

```go
client := ss.NewClient("https://example.org")
client.SetHeader("X-Custom", "Value")
client.SetTimeout(10 * time.Second)
client.SetFollowRedirects(false)
client.UserAgent = "my-app/1.0"

// Inspect or modify each Request and Response
client.RequestFilter = func(req *http.Request) { log.Println(req.Method, req.URL) }
client.ResponseFilter = func(res *http.Response) { log.Println(res.Status) }

// Replace the underlying *http.Client to customize transports, proxies or TLS
client.HttpClient = &http.Client{Timeout: 30 * time.Second}
```

`NewClient` sends Requests to ServiceStack's pre-defined `/api` route. Use
`NewJsonServiceClient` for older ServiceStack instances that only have the
`/json/reply` routes enabled, or `SetBasePath` for a custom base path.

## Examples

- [examples/hello](examples/hello) — typed APIs, error handling and batched Requests
  against the live https://test.servicestack.net Services.

## Tests

```bash
go test ./...                    # unit tests
go test -tags integration ./...  # integration tests against test.servicestack.net
```

## Releasing

Releases are cut with npm scripts and published by the `release` GitHub Action:

```bash
npm run bump              # 0.1.0 -> 0.1.1 (also `-- minor`, `-- major`, `-- 1.2.3`)
# describe the release in CHANGELOG.md, then
npm run release
```

Or in a single step:

```bash
npm run release -- patch
```

`npm run release` tags the version, pushes it and creates the GitHub Release,
which triggers the workflow that runs the tests and publishes it to the Go module proxy.

## License

BSD-3-Clause. See [LICENSE](LICENSE).
