// Package servicestack provides a typed Go client for consuming ServiceStack APIs.
//
// Use the Go DTO generator to generate typed DTOs for a remote ServiceStack API:
//
//	npx get-dtos go https://blazor-vue.web-templates.io
//
// Then send them with the generic client functions, which infer the Response
// Type from the Request DTO:
//
//	client := servicestack.NewClient("https://blazor-vue.web-templates.io")
//	res, err := servicestack.Send(client, dtos.Hello{Name: "World"})
//	fmt.Println(res.Result)
package servicestack

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"sync"
	"time"
)

// DefaultBasePath is the base path of ServiceStack's pre-defined /api route,
// e.g. /api/{RequestDto}
const DefaultBasePath = "api"

// Global filters applied to all Clients, useful for logging and diagnostics.
var (
	GlobalRequestFilter  func(*http.Request)
	GlobalResponseFilter func(*http.Response)
)

// Client is a typed HTTP Client for consuming ServiceStack APIs.
//
// A Client is safe for concurrent use by multiple goroutines. Configure it
// before sharing it, mutating exported fields whilst in-flight Requests are
// being sent is not supported (use the Set* methods, which are guarded).
type Client struct {
	// BaseUrl of the remote ServiceStack instance, e.g. https://example.org
	BaseUrl string
	// ReplyBaseUrl is the base URL Request DTOs are sent to, e.g. https://example.org/api
	ReplyBaseUrl string
	// OneWayBaseUrl is the base URL one-way Requests are sent to.
	OneWayBaseUrl string
	// HttpClient used to send Requests, replace it to customize timeouts,
	// transports, proxies or to use a mocked client in tests.
	HttpClient *http.Client
	// Headers sent with every Request.
	Headers http.Header
	// UserAgent sent with every Request.
	UserAgent string
	// RequestFilter is invoked before each Request is sent.
	RequestFilter func(*http.Request)
	// ResponseFilter is invoked after each Response is received.
	ResponseFilter func(*http.Response)
	// OnAuthenticationRequired is invoked when a Request returns 401 Unauthorized,
	// letting the Client authenticate before the Request is retried once.
	OnAuthenticationRequired func(client *Client) error
	// RefreshTokenUri is the URL used to fetch a new Access Token from the
	// RefreshToken, defaults to the GetAccessToken API of this Client.
	RefreshTokenUri string

	mu           sync.RWMutex
	userName     string
	password     string
	bearerToken  string
	refreshToken string
}

// NewClient creates a Client that sends Requests to ServiceStack's pre-defined
// /api route, e.g. https://example.org/api/{RequestDto}
func NewClient(baseUrl string) *Client {
	jar, _ := cookiejar.New(nil)
	client := &Client{
		BaseUrl: strings.TrimSuffix(baseUrl, "/"),
		Headers: http.Header{
			HeaderAccept: []string{MimeTypeJson},
		},
		HttpClient: &http.Client{
			Timeout: 60 * time.Second,
			Jar:     jar,
		},
	}
	return client.SetBasePath(DefaultBasePath)
}

// NewJsonApiClient creates a Client that sends Requests to ServiceStack's
// pre-defined /api route. It's an alias of NewClient.
func NewJsonApiClient(baseUrl string) *Client { return NewClient(baseUrl) }

// NewJsonServiceClient creates a Client that sends Requests to ServiceStack's
// pre-defined /json/reply route, for compatibility with older ServiceStack
// instances that don't have the /api route enabled.
func NewJsonServiceClient(baseUrl string) *Client {
	return NewClient(baseUrl).SetBasePath("")
}

// SetBasePath changes the base path Request DTOs are sent to, e.g. "api".
// Use an empty basePath to send Requests to the /json/reply pre-defined routes.
func (c *Client) SetBasePath(basePath string) *Client {
	if basePath == "" {
		c.ReplyBaseUrl = CombineWith(c.BaseUrl, "json/reply")
		c.OneWayBaseUrl = CombineWith(c.BaseUrl, "json/oneway")
	} else {
		c.ReplyBaseUrl = CombineWith(c.BaseUrl, basePath)
		c.OneWayBaseUrl = CombineWith(c.BaseUrl, basePath)
	}
	return c
}

// SetCredentials sets the UserName and Password sent in the HTTP Basic Auth
// header of each Request.
func (c *Client) SetCredentials(userName, password string) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.userName, c.password = userName, password
	return c
}

// SetBearerToken sets the JWT or API Key sent in the Bearer Authorization header.
func (c *Client) SetBearerToken(token string) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bearerToken = token
	return c
}

// BearerToken returns the Bearer Token sent with each Request.
func (c *Client) BearerToken() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.bearerToken
}

// SetRefreshToken sets the Refresh Token used to fetch a new Bearer Token when
// a Request returns 401 Unauthorized.
func (c *Client) SetRefreshToken(token string) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.refreshToken = token
	return c
}

// RefreshToken returns the Refresh Token used to fetch new Bearer Tokens.
func (c *Client) RefreshToken() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.refreshToken
}

// SetHeader sets a HTTP Header sent with each Request.
func (c *Client) SetHeader(name, value string) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Headers == nil {
		c.Headers = http.Header{}
	}
	c.Headers.Set(name, value)
	return c
}

// SetTimeout sets the timeout of the underlying HTTP Client.
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.HttpClient.Timeout = timeout
	return c
}

// SetFollowRedirects configures whether the Client follows HTTP redirects,
// enabled by default. Disable it when Services that require Authentication
// redirect to a HTML sign in page instead of returning 401 Unauthorized.
func (c *Client) SetFollowRedirects(follow bool) *Client {
	if follow {
		c.HttpClient.CheckRedirect = nil
	} else {
		c.HttpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	return c
}

// ToAbsoluteUrl converts a relative path into an absolute URL of this Client.
func (c *Client) ToAbsoluteUrl(pathOrUrl string) string {
	return ToAbsoluteUrl(c.BaseUrl, pathOrUrl)
}

// CreateUrlFromDto returns the URL a Request DTO is sent to, appending the
// populated DTO properties to the QueryString for Requests without a Body.
func (c *Client) CreateUrlFromDto(method string, request any) string {
	requestUrl := CombineWith(c.ReplyBaseUrl, NameOf(request))
	if !HasRequestBody(method) {
		requestUrl = AppendQueryString(requestUrl, DtoToMap(request))
	}
	return requestUrl
}

// Authenticate signs in with the UserName and Password credentials, returning
// the AuthenticateResponse of a successful sign in.
//
// The Client's cookie jar maintains the authenticated Session for subsequent
// Requests, or when the Server returns a Bearer Token it's used instead.
func (c *Client) Authenticate(userName, password string) (AuthenticateResponse, error) {
	return c.AuthenticateCtx(context.Background(), userName, password)
}

// AuthenticateCtx is Authenticate with a context.Context.
func (c *Client) AuthenticateCtx(ctx context.Context, userName, password string) (AuthenticateResponse, error) {
	res, err := SendCtx(ctx, c, Authenticate{
		Provider: "credentials",
		UserName: userName,
		Password: password,
	})
	if err != nil {
		return res, err
	}
	if res.BearerToken != "" {
		c.SetBearerToken(res.BearerToken)
	}
	if res.RefreshToken != "" {
		c.SetRefreshToken(res.RefreshToken)
	}
	return res, nil
}

// ── Typed Client API ──

// Send sends a Request DTO with the HTTP Method it's annotated with,
// returning its typed Response.
func Send[T any](client *Client, request IReturn[T]) (T, error) {
	return SendCtx(context.Background(), client, request)
}

// SendCtx is Send with a context.Context.
func SendCtx[T any](ctx context.Context, client *Client, request IReturn[T]) (T, error) {
	return sendDto[T](ctx, client, ResolveHttpMethod(request), request, nil)
}

// Get sends a Request DTO with a GET Request, returning its typed Response.
func Get[T any](client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return GetCtx(context.Background(), client, request, args...)
}

// GetCtx is Get with a context.Context.
func GetCtx[T any](ctx context.Context, client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return sendDto[T](ctx, client, HttpGet, request, firstArg(args))
}

// Post sends a Request DTO with a POST Request, returning its typed Response.
func Post[T any](client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return PostCtx(context.Background(), client, request, args...)
}

// PostCtx is Post with a context.Context.
func PostCtx[T any](ctx context.Context, client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return sendDto[T](ctx, client, HttpPost, request, firstArg(args))
}

// Put sends a Request DTO with a PUT Request, returning its typed Response.
func Put[T any](client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return PutCtx(context.Background(), client, request, args...)
}

// PutCtx is Put with a context.Context.
func PutCtx[T any](ctx context.Context, client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return sendDto[T](ctx, client, HttpPut, request, firstArg(args))
}

// Patch sends a Request DTO with a PATCH Request, returning its typed Response.
func Patch[T any](client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return PatchCtx(context.Background(), client, request, args...)
}

// PatchCtx is Patch with a context.Context.
func PatchCtx[T any](ctx context.Context, client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return sendDto[T](ctx, client, HttpPatch, request, firstArg(args))
}

// Delete sends a Request DTO with a DELETE Request, returning its typed Response.
func Delete[T any](client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return DeleteCtx(context.Background(), client, request, args...)
}

// DeleteCtx is Delete with a context.Context.
func DeleteCtx[T any](ctx context.Context, client *Client, request IReturn[T], args ...map[string]any) (T, error) {
	return sendDto[T](ctx, client, HttpDelete, request, firstArg(args))
}

// SendVoid sends a Request DTO that doesn't return a Response Body.
func SendVoid(client *Client, request IReturnVoid, args ...map[string]any) error {
	return SendVoidCtx(context.Background(), client, request, args...)
}

// SendVoidCtx is SendVoid with a context.Context.
func SendVoidCtx(ctx context.Context, client *Client, request IReturnVoid, args ...map[string]any) error {
	_, err := sendDto[emptyResponse](ctx, client, ResolveHttpMethod(request), request, firstArg(args))
	return err
}

// SendAs sends a Request DTO that doesn't declare its Response Type, requiring
// the Response Type to be specified explicitly, e.g:
//
//	res, err := servicestack.SendAs[HelloResponse](client, Hello{Name: "World"})
func SendAs[T any](client *Client, request any, args ...map[string]any) (T, error) {
	return SendAsCtx[T](context.Background(), client, request, args...)
}

// SendAsCtx is SendAs with a context.Context.
func SendAsCtx[T any](ctx context.Context, client *Client, request any, args ...map[string]any) (T, error) {
	return sendDto[T](ctx, client, ResolveHttpMethod(request), request, firstArg(args))
}

// SendMethodAs sends a Request DTO with the specified HTTP Method, requiring
// the Response Type to be specified explicitly.
func SendMethodAs[T any](client *Client, method string, request any, args ...map[string]any) (T, error) {
	return sendDto[T](context.Background(), client, method, request, firstArg(args))
}

// ── Api (returns errors in ApiResult instead of error) ──

// ApiResult holds either the typed Response of a successful API Request or the
// structured ResponseStatus error of a failed one.
type ApiResult[T any] struct {
	Response T
	Error    *ResponseStatus
}

// Succeeded reports whether the API Request was successful.
func (r ApiResult[T]) Succeeded() bool { return r.Error == nil }

// Failed reports whether the API Request failed.
func (r ApiResult[T]) Failed() bool { return r.Error != nil }

// ErrorCode returns the ErrorCode of a failed API Request.
func (r ApiResult[T]) ErrorCode() string {
	if r.Error != nil {
		return r.Error.ErrorCode
	}
	return ""
}

// ErrorMessage returns the error message of a failed API Request.
func (r ApiResult[T]) ErrorMessage() string {
	if r.Error != nil {
		return r.Error.Message
	}
	return ""
}

// FieldError returns the validation error message for the specified field.
func (r ApiResult[T]) FieldError(fieldName string) string {
	return r.Error.FieldError(fieldName)
}

// Api sends a Request DTO, returning an ApiResult containing either the typed
// Response or the structured ResponseStatus error, e.g:
//
//	api := servicestack.Api(client, Hello{Name: "World"})
//	if api.Failed() {
//	    fmt.Println(api.ErrorMessage())
//	}
func Api[T any](client *Client, request IReturn[T]) ApiResult[T] {
	return ApiCtx(context.Background(), client, request)
}

// ApiCtx is Api with a context.Context.
func ApiCtx[T any](ctx context.Context, client *Client, request IReturn[T]) ApiResult[T] {
	response, err := SendCtx(ctx, client, request)
	return ApiResult[T]{Response: response, Error: toErrorStatus(err)}
}

// ── URL Client API ──

// SendUrl sends a Request to a custom relative path or absolute URL.
func SendUrl[T any](client *Client, method, path string, body any, args ...map[string]any) (T, error) {
	return SendUrlCtx[T](context.Background(), client, method, path, body, args...)
}

// SendUrlCtx is SendUrl with a context.Context.
func SendUrlCtx[T any](ctx context.Context, client *Client, method, path string, body any, args ...map[string]any) (T, error) {
	requestUrl := AppendQueryString(client.ToAbsoluteUrl(path), firstArg(args))
	return sendUrl[T](ctx, client, method, requestUrl, body)
}

// GetUrl sends a GET Request to a custom relative path or absolute URL.
func GetUrl[T any](client *Client, path string, args ...map[string]any) (T, error) {
	return SendUrlCtx[T](context.Background(), client, HttpGet, path, nil, args...)
}

// PostUrl sends a POST Request to a custom relative path or absolute URL.
func PostUrl[T any](client *Client, path string, body any, args ...map[string]any) (T, error) {
	return SendUrlCtx[T](context.Background(), client, HttpPost, path, body, args...)
}

// PutUrl sends a PUT Request to a custom relative path or absolute URL.
func PutUrl[T any](client *Client, path string, body any, args ...map[string]any) (T, error) {
	return SendUrlCtx[T](context.Background(), client, HttpPut, path, body, args...)
}

// PatchUrl sends a PATCH Request to a custom relative path or absolute URL.
func PatchUrl[T any](client *Client, path string, body any, args ...map[string]any) (T, error) {
	return SendUrlCtx[T](context.Background(), client, HttpPatch, path, body, args...)
}

// DeleteUrl sends a DELETE Request to a custom relative path or absolute URL.
func DeleteUrl[T any](client *Client, path string, args ...map[string]any) (T, error) {
	return SendUrlCtx[T](context.Background(), client, HttpDelete, path, nil, args...)
}

// ── Batched Requests ──

// SendAll sends multiple Request DTOs of the same Type in a single Request,
// returning all their Responses, e.g:
//
//	responses, err := servicestack.SendAll(client, []dtos.Hello{{Name:"A"},{Name:"B"}})
func SendAll[TRequest IReturn[TResponse], TResponse any](client *Client, requests []TRequest) ([]TResponse, error) {
	return SendAllCtx[TRequest, TResponse](context.Background(), client, requests)
}

// SendAllCtx is SendAll with a context.Context.
func SendAllCtx[TRequest IReturn[TResponse], TResponse any](ctx context.Context, client *Client, requests []TRequest) ([]TResponse, error) {
	if len(requests) == 0 {
		return []TResponse{}, nil
	}
	requestUrl := CombineWith(client.ReplyBaseUrl, NameOf(requests[0])+"[]")
	return sendUrl[[]TResponse](ctx, client, HttpPost, requestUrl, requests)
}

// Publish sends a Request DTO to a one-way endpoint, ignoring any Response.
func Publish(client *Client, request any) error {
	return PublishCtx(context.Background(), client, request)
}

// PublishCtx is Publish with a context.Context.
func PublishCtx(ctx context.Context, client *Client, request any) error {
	requestUrl := CombineWith(client.OneWayBaseUrl, NameOf(request))
	_, err := sendUrl[emptyResponse](ctx, client, HttpPost, requestUrl, request)
	return err
}

// PublishAll sends multiple Request DTOs of the same Type to a one-way endpoint.
func PublishAll[T any](client *Client, requests []T) error {
	return PublishAllCtx(context.Background(), client, requests)
}

// PublishAllCtx is PublishAll with a context.Context.
func PublishAllCtx[T any](ctx context.Context, client *Client, requests []T) error {
	if len(requests) == 0 {
		return nil
	}
	requestUrl := CombineWith(client.OneWayBaseUrl, NameOf(requests[0])+"[]")
	_, err := sendUrl[emptyResponse](ctx, client, HttpPost, requestUrl, requests)
	return err
}

// ── File Uploads ──

// UploadFile is a file uploaded in a multipart/form-data Request.
type UploadFile struct {
	// FieldName of the file in the multipart Request, defaults to "file".
	FieldName string
	// FileName of the uploaded file.
	FileName string
	// ContentType of the uploaded file, detected by the Server when omitted.
	ContentType string
	// Reader containing the file contents.
	Reader io.Reader
}

// PostFileWithRequest uploads a file with a Request DTO as a
// multipart/form-data Request, returning its typed Response.
func PostFileWithRequest[T any](client *Client, request IReturn[T], file UploadFile) (T, error) {
	return PostFilesWithRequestCtx(context.Background(), client, request, []UploadFile{file})
}

// PostFilesWithRequest uploads multiple files with a Request DTO as a
// multipart/form-data Request, returning its typed Response.
func PostFilesWithRequest[T any](client *Client, request IReturn[T], files []UploadFile) (T, error) {
	return PostFilesWithRequestCtx(context.Background(), client, request, files)
}

// PostFilesWithRequestCtx is PostFilesWithRequest with a context.Context.
func PostFilesWithRequestCtx[T any](ctx context.Context, client *Client, request IReturn[T], files []UploadFile) (T, error) {
	requestUrl := client.CreateUrlFromDto(HttpPost, request)
	return PostFilesWithRequestUrlCtx[T](ctx, client, requestUrl, request, files)
}

// PostFilesWithRequestUrlCtx uploads files with a Request DTO to a custom URL.
func PostFilesWithRequestUrlCtx[T any](ctx context.Context, client *Client, requestUrl string, request any, files []UploadFile) (T, error) {
	var to T
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for name, value := range DtoToMap(request) {
		if value == nil {
			continue
		}
		if err := writer.WriteField(name, QsValue(value)); err != nil {
			return to, fmt.Errorf("failed to write form field %s: %w", name, err)
		}
	}

	for _, file := range files {
		fieldName := file.FieldName
		if fieldName == "" {
			fieldName = "file"
		}
		part, err := createFormFile(writer, fieldName, file)
		if err != nil {
			return to, err
		}
		if file.Reader != nil {
			if _, err = io.Copy(part, file.Reader); err != nil {
				return to, fmt.Errorf("failed to write file %s: %w", file.FileName, err)
			}
		}
	}
	if err := writer.Close(); err != nil {
		return to, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	responseBytes, err := client.sendRequest(ctx, HttpPost, client.ToAbsoluteUrl(requestUrl),
		body.Bytes(), writer.FormDataContentType(), true)
	if err != nil {
		return to, err
	}
	return unmarshalAs[T](responseBytes)
}

func createFormFile(writer *multipart.Writer, fieldName string, file UploadFile) (io.Writer, error) {
	fileName := file.FileName
	if fileName == "" {
		fileName = "file"
	}
	if file.ContentType == "" {
		part, err := writer.CreateFormFile(fieldName, fileName)
		if err != nil {
			return nil, fmt.Errorf("failed to create form file %s: %w", fileName, err)
		}
		return part, nil
	}
	header := make(map[string][]string)
	header["Content-Disposition"] = []string{
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(fieldName), escapeQuotes(fileName)),
	}
	header[HeaderContentType] = []string{file.ContentType}
	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file %s: %w", fileName, err)
	}
	return part, nil
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { return quoteEscaper.Replace(s) }

// ── Internals ──

// emptyResponse is used for Requests that don't return a Response Body.
type emptyResponse struct{}

func firstArg(args []map[string]any) map[string]any {
	if len(args) > 0 {
		return args[0]
	}
	return nil
}

func sendDto[T any](ctx context.Context, client *Client, method string, request any, args map[string]any) (T, error) {
	requestUrl := AppendQueryString(client.CreateUrlFromDto(method, request), args)
	var body any
	if HasRequestBody(method) {
		body = request
	}
	return sendUrl[T](ctx, client, method, requestUrl, body)
}

func sendUrl[T any](ctx context.Context, client *Client, method, requestUrl string, body any) (T, error) {
	var to T
	bodyBytes, contentType, err := toRequestBody(body)
	if err != nil {
		return to, err
	}
	responseBytes, err := client.sendRequest(ctx, method, client.ToAbsoluteUrl(requestUrl), bodyBytes, contentType, true)
	if err != nil {
		return to, err
	}
	return unmarshalAs[T](responseBytes)
}

func toRequestBody(body any) ([]byte, string, error) {
	switch b := body.(type) {
	case nil:
		return nil, "", nil
	case []byte:
		return b, "", nil
	case string:
		return []byte(b), "", nil
	case io.Reader:
		bodyBytes, err := io.ReadAll(b)
		if err != nil {
			return nil, "", fmt.Errorf("failed to read request body: %w", err)
		}
		return bodyBytes, "", nil
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, "", fmt.Errorf("failed to serialize request body: %w", err)
	}
	return bodyBytes, MimeTypeJson, nil
}

func unmarshalAs[T any](responseBytes []byte) (T, error) {
	var to T
	switch ptr := any(&to).(type) {
	case *[]byte:
		*ptr = responseBytes
		return to, nil
	case *string:
		*ptr = string(responseBytes)
		return to, nil
	}
	if len(bytes.TrimSpace(responseBytes)) == 0 {
		return to, nil
	}
	if err := json.Unmarshal(responseBytes, &to); err != nil {
		return to, fmt.Errorf("failed to deserialize %T response: %w, body: %s", to, err, truncate(string(responseBytes), 500))
	}
	return to, nil
}

// sendRequest sends the HTTP Request, returning the Response Body or a
// *WebServiceException for error Responses.
func (c *Client) sendRequest(ctx context.Context, method, requestUrl string, body []byte, contentType string, retryOnAuthFailure bool) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(method), requestUrl, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s %s request: %w", method, requestUrl, err)
	}

	c.mu.RLock()
	for name, values := range c.Headers {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
	bearerToken, userName, password := c.bearerToken, c.userName, c.password
	c.mu.RUnlock()

	if req.Header.Get(HeaderAccept) == "" {
		req.Header.Set(HeaderAccept, MimeTypeJson)
	}
	if contentType != "" {
		req.Header.Set(HeaderContentType, contentType)
	}
	if c.UserAgent != "" {
		req.Header.Set(HeaderUserAgent, c.UserAgent)
	}
	if req.Header.Get(HeaderAuthorization) == "" {
		if bearerToken != "" {
			req.Header.Set(HeaderAuthorization, "Bearer "+bearerToken)
		} else if userName != "" || password != "" {
			credentials := base64.StdEncoding.EncodeToString([]byte(userName + ":" + password))
			req.Header.Set(HeaderAuthorization, "Basic "+credentials)
		}
	}

	if c.RequestFilter != nil {
		c.RequestFilter(req)
	}
	if GlobalRequestFilter != nil {
		GlobalRequestFilter(req)
	}

	httpClient := c.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, &WebServiceException{
			StatusCode:        0,
			StatusDescription: err.Error(),
			Inner:             err,
		}
	}
	defer res.Body.Close()

	if c.ResponseFilter != nil {
		c.ResponseFilter(res)
	}
	if GlobalResponseFilter != nil {
		GlobalResponseFilter(res)
	}

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &WebServiceException{
			StatusCode:        res.StatusCode,
			StatusDescription: res.Status,
			Inner:             fmt.Errorf("failed to read response body: %w", err),
		}
	}

	if res.StatusCode == http.StatusUnauthorized && retryOnAuthFailure {
		if err := c.onAuthenticationRequired(ctx); err == nil {
			return c.sendRequest(ctx, method, requestUrl, body, contentType, false)
		}
	}

	if res.StatusCode >= 400 {
		return nil, toWebServiceException(res, responseBytes)
	}

	return responseBytes, nil
}

// onAuthenticationRequired tries to re-authenticate the Client after a Request
// returned 401 Unauthorized, returning nil if the Request should be retried.
func (c *Client) onAuthenticationRequired(ctx context.Context) error {
	if c.RefreshToken() != "" {
		if err := c.fetchAccessToken(ctx); err != nil {
			return err
		}
		return nil
	}
	if c.OnAuthenticationRequired != nil {
		return c.OnAuthenticationRequired(c)
	}
	return fmt.Errorf("not authenticated")
}

// fetchAccessToken exchanges the Client's Refresh Token for a new Bearer Token.
func (c *Client) fetchAccessToken(ctx context.Context) error {
	refreshToken := c.RefreshToken()
	if refreshToken == "" {
		return fmt.Errorf("no refresh token")
	}
	requestUrl := c.RefreshTokenUri
	if requestUrl == "" {
		requestUrl = CombineWith(c.ReplyBaseUrl, NameOf(GetAccessToken{}))
	}
	body, err := json.Marshal(GetAccessToken{RefreshToken: refreshToken})
	if err != nil {
		return err
	}
	responseBytes, err := c.sendRequest(ctx, HttpPost, c.ToAbsoluteUrl(requestUrl), body, MimeTypeJson, false)
	if err != nil {
		return err
	}
	res, err := unmarshalAs[GetAccessTokenResponse](responseBytes)
	if err != nil {
		return err
	}
	if res.AccessToken == "" {
		return fmt.Errorf("could not fetch new access token")
	}
	c.SetBearerToken(res.AccessToken)
	return nil
}

func toWebServiceException(res *http.Response, responseBytes []byte) *WebServiceException {
	webEx := &WebServiceException{
		StatusCode:        res.StatusCode,
		StatusDescription: res.Status,
		ResponseBody:      string(responseBytes),
	}

	var errorResponse struct {
		ResponseStatus *ResponseStatus `json:"responseStatus"`
	}
	if err := json.Unmarshal(responseBytes, &errorResponse); err == nil && errorResponse.ResponseStatus != nil {
		webEx.ResponseStatus = errorResponse.ResponseStatus
		return webEx
	}

	// Fallback for Services that return a bare ResponseStatus
	var responseStatus ResponseStatus
	if err := json.Unmarshal(responseBytes, &responseStatus); err == nil &&
		(responseStatus.ErrorCode != "" || responseStatus.Message != "") {
		webEx.ResponseStatus = &responseStatus
		return webEx
	}

	webEx.ResponseStatus = &ResponseStatus{
		ErrorCode: strings.TrimSpace(http.StatusText(res.StatusCode)),
		Message:   strings.TrimSpace(res.Status),
	}
	return webEx
}

func truncate(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength] + "..."
}
