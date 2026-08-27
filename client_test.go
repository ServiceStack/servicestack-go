// Unit tests for the Client, served by a mock HTTP Server.
package servicestack_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ss "github.com/ServiceStack/servicestack-go"

	// Typed DTOs generated from https://test.servicestack.net with:
	//     npx get-dtos go https://test.servicestack.net
	"github.com/ServiceStack/servicestack-go/dtos"
)

// newTestClient returns a Client and Server that records the last Request.
func newTestClient(t *testing.T, handler http.HandlerFunc) (*ss.Client, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	return ss.NewClient(server.URL), server
}

func writeJson(t *testing.T, w http.ResponseWriter, statusCode int, body any) {
	t.Helper()
	w.Header().Set(ss.HeaderContentType, ss.MimeTypeJson)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		t.Fatalf("failed to write response: %v", err)
	}
}

func TestSendGetRequestUsesQueryString(t *testing.T) {
	var gotUrl, gotMethod string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotUrl, gotMethod = r.URL.String(), r.Method
		writeJson(t, w, http.StatusOK, dtos.HelloVerbResponse{Result: "Hello, World!"})
	})

	res, err := ss.Send(client, dtos.HelloGet{Id: 1})
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got Result %q, want %q", res.Result, "Hello, World!")
	}
	if gotMethod != ss.HttpGet {
		t.Errorf("got method %q, want GET", gotMethod)
	}
	if gotUrl != "/api/HelloGet?id=1" {
		t.Errorf("got url %q, want /api/HelloGet?id=1", gotUrl)
	}
}

func TestSendPostRequestUsesJsonBody(t *testing.T) {
	var gotUrl, gotMethod, gotContentType, gotBody string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		gotUrl, gotMethod, gotBody = r.URL.String(), r.Method, string(body)
		gotContentType = r.Header.Get(ss.HeaderContentType)
		writeJson(t, w, http.StatusOK, dtos.HelloResponse{Result: "Hello, World!"})
	})

	if _, err := ss.Send(client, dtos.Hello{Name: "World", Title: "Mr"}); err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if gotMethod != ss.HttpPost {
		t.Errorf("got method %q, want POST", gotMethod)
	}
	if gotUrl != "/api/Hello" {
		t.Errorf("got url %q, want /api/Hello", gotUrl)
	}
	if gotContentType != ss.MimeTypeJson {
		t.Errorf("got Content-Type %q, want %q", gotContentType, ss.MimeTypeJson)
	}
	if gotBody != `{"name":"World","title":"Mr"}` {
		t.Errorf("got body %q, want %q", gotBody, `{"name":"World","title":"Mr"}`)
	}
}

func TestSendInfersGenericResponseType(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeJson(t, w, http.StatusOK, ss.QueryResponse[dtos.Booking]{
			Offset:  0,
			Total:   1,
			Results: []dtos.Booking{{Id: 1, Name: "First Booking"}},
		})
	})

	take := 10
	res, err := ss.Send(client, dtos.QueryBookings{QueryDb: ss.QueryDb{ss.QueryBase{Take: &take}}})
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if len(res.Results) != 1 || res.Results[0].Name != "First Booking" {
		t.Fatalf("got Results %+v, want 1 Booking", res.Results)
	}
}

func TestSendAppendsInheritedQueryStringArgs(t *testing.T) {
	var gotUrl string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotUrl = r.URL.String()
		writeJson(t, w, http.StatusOK, ss.QueryResponse[dtos.Booking]{})
	})

	skip, take := 5, 10
	_, err := ss.Send(client, dtos.QueryBookings{QueryDb: ss.QueryDb{ss.QueryBase{Skip: &skip, Take: &take, OrderBy: "id"}}})
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	for _, want := range []string{"skip=5", "take=10", "orderBy=id"} {
		if !strings.Contains(gotUrl, want) {
			t.Errorf("url %q missing %q", gotUrl, want)
		}
	}
}

func TestSendVoidIgnoresResponseBody(t *testing.T) {
	var gotMethod, gotUrl string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotUrl = r.Method, r.URL.String()
		w.WriteHeader(http.StatusOK)
	})

	if err := ss.SendVoid(client, dtos.DeleteBooking{Id: 1}); err != nil {
		t.Fatalf("SendVoid failed: %v", err)
	}
	if gotMethod != ss.HttpDelete {
		t.Errorf("got method %q, want DELETE", gotMethod)
	}
	if gotUrl != "/api/DeleteBooking?id=1" {
		t.Errorf("got url %q, want /api/DeleteBooking?id=1", gotUrl)
	}
}

func TestSendAsWithExplicitResponseType(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeJson(t, w, http.StatusOK, dtos.HelloResponse{Result: "Hello, World!"})
	})

	type UnTypedHello struct {
		Name string `json:"name,omitempty"`
	}
	res, err := ss.SendAs[dtos.HelloResponse](client, UnTypedHello{Name: "World"})
	if err != nil {
		t.Fatalf("SendAs failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got Result %q, want %q", res.Result, "Hello, World!")
	}
}

func TestErrorResponseReturnsWebServiceException(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeJson(t, w, http.StatusBadRequest, map[string]any{
			"responseStatus": ss.ResponseStatus{
				ErrorCode: "NotEmpty",
				Message:   "'Name' must not be empty.",
				Errors: []ss.ResponseError{
					{ErrorCode: "NotEmpty", FieldName: "Name", Message: "'Name' must not be empty."},
				},
			},
		})
	})

	_, err := ss.Send(client, dtos.Hello{})
	if err == nil {
		t.Fatal("expected an error")
	}

	var webEx *ss.WebServiceException
	if !errors.As(err, &webEx) {
		t.Fatalf("got %T, want *WebServiceException", err)
	}
	if webEx.StatusCode != http.StatusBadRequest {
		t.Errorf("got StatusCode %d, want 400", webEx.StatusCode)
	}
	if webEx.ErrorCode() != "NotEmpty" {
		t.Errorf("got ErrorCode %q, want NotEmpty", webEx.ErrorCode())
	}
	if webEx.FieldError("name") != "'Name' must not be empty." {
		t.Errorf("got FieldError %q", webEx.FieldError("name"))
	}
	if !webEx.IsValidationError() {
		t.Error("expected IsValidationError to be true")
	}
	if ss.GetStatusCode(err) != http.StatusBadRequest {
		t.Errorf("got GetStatusCode %d, want 400", ss.GetStatusCode(err))
	}
	if status := ss.GetResponseStatus(err); status == nil || status.ErrorCode != "NotEmpty" {
		t.Errorf("got GetResponseStatus %+v", status)
	}
}

func TestErrorResponseWithoutResponseStatus(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("<html>Not Found</html>"))
	})

	_, err := ss.Send(client, dtos.Hello{Name: "World"})
	webEx, ok := ss.AsWebServiceException(err)
	if !ok {
		t.Fatalf("got %T, want *WebServiceException", err)
	}
	if !webEx.IsNotFound() {
		t.Errorf("got StatusCode %d, want 404", webEx.StatusCode)
	}
	if webEx.ResponseBody != "<html>Not Found</html>" {
		t.Errorf("got ResponseBody %q", webEx.ResponseBody)
	}
	if !strings.Contains(webEx.Error(), "404") {
		t.Errorf("got Error() %q, want it to contain 404", webEx.Error())
	}
}

func TestApiReturnsErrorStatusInsteadOfError(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeJson(t, w, http.StatusBadRequest, map[string]any{
			"responseStatus": ss.ResponseStatus{
				ErrorCode: "NotEmpty",
				Message:   "'Name' must not be empty.",
				Errors: []ss.ResponseError{
					{ErrorCode: "NotEmpty", FieldName: "Name", Message: "'Name' must not be empty."},
				},
			},
		})
	})

	api := ss.Api(client, dtos.Hello{})
	if api.Succeeded() {
		t.Fatal("expected the Api Request to fail")
	}
	if api.ErrorCode() != "NotEmpty" {
		t.Errorf("got ErrorCode %q, want NotEmpty", api.ErrorCode())
	}
	if api.FieldError("Name") != "'Name' must not be empty." {
		t.Errorf("got FieldError %q", api.FieldError("Name"))
	}
}

func TestApiSucceeded(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeJson(t, w, http.StatusOK, dtos.HelloResponse{Result: "Hello, World!"})
	})

	api := ss.Api(client, dtos.Hello{Name: "World"})
	if api.Failed() {
		t.Fatalf("expected the Api Request to succeed, got %v", api.Error)
	}
	if api.Response.Result != "Hello, World!" {
		t.Errorf("got Result %q", api.Response.Result)
	}
	if api.FieldError("Name") != "" {
		t.Errorf("got FieldError %q, want empty", api.FieldError("Name"))
	}
}

func TestBearerTokenAndBasicAuthHeaders(t *testing.T) {
	var gotAuth string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get(ss.HeaderAuthorization)
		writeJson(t, w, http.StatusOK, dtos.HelloResponse{})
	})

	client.SetBearerToken("TOKEN")
	if _, err := ss.Send(client, dtos.Hello{Name: "World"}); err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if gotAuth != "Bearer TOKEN" {
		t.Errorf("got Authorization %q, want Bearer TOKEN", gotAuth)
	}

	client.SetBearerToken("").SetCredentials("user", "pass")
	if _, err := ss.Send(client, dtos.Hello{Name: "World"}); err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if gotAuth != "Basic dXNlcjpwYXNz" {
		t.Errorf("got Authorization %q, want Basic dXNlcjpwYXNz", gotAuth)
	}
}

func TestCustomHeadersAndRequestFilter(t *testing.T) {
	var gotHeader, gotFilterHeader string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotHeader = r.Header.Get("X-Api-Key")
		gotFilterHeader = r.Header.Get("X-Filter")
		writeJson(t, w, http.StatusOK, dtos.HelloResponse{})
	})

	client.SetHeader("X-Api-Key", "KEY")
	client.RequestFilter = func(req *http.Request) { req.Header.Set("X-Filter", "1") }

	if _, err := ss.Send(client, dtos.Hello{Name: "World"}); err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if gotHeader != "KEY" {
		t.Errorf("got X-Api-Key %q, want KEY", gotHeader)
	}
	if gotFilterHeader != "1" {
		t.Errorf("got X-Filter %q, want 1", gotFilterHeader)
	}
}

func TestRefreshTokenFetchesNewBearerTokenAndRetries(t *testing.T) {
	var requests []string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Method+" "+r.URL.Path)
		switch r.URL.Path {
		case "/api/GetAccessToken":
			writeJson(t, w, http.StatusOK, ss.GetAccessTokenResponse{AccessToken: "NEW_TOKEN"})
		default:
			if r.Header.Get(ss.HeaderAuthorization) != "Bearer NEW_TOKEN" {
				writeJson(t, w, http.StatusUnauthorized, map[string]any{
					"responseStatus": ss.ResponseStatus{ErrorCode: "Unauthorized", Message: "Unauthorized"},
				})
				return
			}
			writeJson(t, w, http.StatusOK, dtos.HelloResponse{Result: "Hello, World!"})
		}
	})

	client.SetRefreshToken("REFRESH_TOKEN")
	res, err := ss.Send(client, dtos.Hello{Name: "World"})
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got Result %q", res.Result)
	}
	if client.BearerToken() != "NEW_TOKEN" {
		t.Errorf("got BearerToken %q, want NEW_TOKEN", client.BearerToken())
	}
	want := []string{"POST /api/Hello", "POST /api/GetAccessToken", "POST /api/Hello"}
	if strings.Join(requests, ",") != strings.Join(want, ",") {
		t.Errorf("got requests %v, want %v", requests, want)
	}
}

func TestUnauthorizedWithoutRefreshTokenReturnsError(t *testing.T) {
	requests := 0
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		requests++
		writeJson(t, w, http.StatusUnauthorized, map[string]any{
			"responseStatus": ss.ResponseStatus{ErrorCode: "Unauthorized", Message: "Unauthorized"},
		})
	})

	_, err := ss.Send(client, dtos.Hello{Name: "World"})
	webEx, ok := ss.AsWebServiceException(err)
	if !ok {
		t.Fatalf("got %T, want *WebServiceException", err)
	}
	if !webEx.IsUnauthorized() {
		t.Errorf("got StatusCode %d, want 401", webEx.StatusCode)
	}
	if requests != 1 {
		t.Errorf("got %d requests, want 1", requests)
	}
}

func TestOnAuthenticationRequiredRetriesOnce(t *testing.T) {
	requests := 0
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Header.Get(ss.HeaderAuthorization) == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		writeJson(t, w, http.StatusOK, dtos.HelloResponse{Result: "Authenticated"})
	})

	client.OnAuthenticationRequired = func(c *ss.Client) error {
		c.SetBearerToken("TOKEN")
		return nil
	}

	res, err := ss.Send(client, dtos.Hello{Name: "World"})
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if res.Result != "Authenticated" {
		t.Errorf("got Result %q", res.Result)
	}
	if requests != 2 {
		t.Errorf("got %d requests, want 2", requests)
	}
}

func TestSendAllBatchesRequests(t *testing.T) {
	var gotUrl, gotBody string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		gotUrl, gotBody = r.URL.String(), string(body)
		writeJson(t, w, http.StatusOK, []dtos.HelloResponse{{Result: "Hello, A!"}, {Result: "Hello, B!"}})
	})

	responses, err := ss.SendAll(client, []dtos.Hello{{Name: "A", Title: "Mr"}, {Name: "B", Title: "Ms"}})
	if err != nil {
		t.Fatalf("SendAll failed: %v", err)
	}
	if len(responses) != 2 || responses[1].Result != "Hello, B!" {
		t.Fatalf("got %+v, want 2 responses", responses)
	}
	if gotUrl != "/api/Hello[]" {
		t.Errorf("got url %q, want /api/Hello[]", gotUrl)
	}
	if gotBody != `[{"name":"A","title":"Mr"},{"name":"B","title":"Ms"}]` {
		t.Errorf("got body %q", gotBody)
	}
}

func TestPublishSendsToOneWayUrl(t *testing.T) {
	var gotUrl string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotUrl = r.URL.String()
		w.WriteHeader(http.StatusOK)
	})

	if err := ss.Publish(client, dtos.Hello{Name: "World"}); err != nil {
		t.Fatalf("Publish failed: %v", err)
	}
	if gotUrl != "/api/Hello" {
		t.Errorf("got url %q, want /api/Hello", gotUrl)
	}
}

func TestUrlApiSendsToCustomPath(t *testing.T) {
	var gotUrl string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotUrl = r.URL.String()
		writeJson(t, w, http.StatusOK, dtos.HelloResponse{Result: "Hello, World!"})
	})

	res, err := ss.GetUrl[dtos.HelloResponse](client, "/hello/World", map[string]any{"detailed": true})
	if err != nil {
		t.Fatalf("GetUrl failed: %v", err)
	}
	if res.Result != "Hello, World!" {
		t.Errorf("got Result %q", res.Result)
	}
	if gotUrl != "/hello/World?detailed=true" {
		t.Errorf("got url %q, want /hello/World?detailed=true", gotUrl)
	}
}

func TestGetUrlAsStringAndBytes(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("plain text"))
	})

	str, err := ss.GetUrl[string](client, "/text")
	if err != nil {
		t.Fatalf("GetUrl[string] failed: %v", err)
	}
	if str != "plain text" {
		t.Errorf("got %q, want plain text", str)
	}

	data, err := ss.GetUrl[[]byte](client, "/text")
	if err != nil {
		t.Fatalf("GetUrl[[]byte] failed: %v", err)
	}
	if string(data) != "plain text" {
		t.Errorf("got %q, want plain text", string(data))
	}
}

func TestPostFilesWithRequest(t *testing.T) {
	var gotFileName, gotFileContent, gotField string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		mediaType, params, err := mime.ParseMediaType(r.Header.Get(ss.HeaderContentType))
		if err != nil || !strings.HasPrefix(mediaType, "multipart/") {
			t.Errorf("got Content-Type %q, want multipart", r.Header.Get(ss.HeaderContentType))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		reader := multipart.NewReader(r.Body, params["boundary"])
		for {
			part, err := reader.NextPart()
			if err != nil {
				break
			}
			content, _ := io.ReadAll(part)
			if part.FileName() != "" {
				gotFileName, gotFileContent = part.FileName(), string(content)
			} else if part.FormName() == "refId" {
				gotField = string(content)
			}
		}
		writeJson(t, w, http.StatusOK, dtos.TestFileUploadsResponse{
			Files: []dtos.UploadInfo{{FileName: "hello.txt"}},
		})
	})

	// Request DTO properties are sent as form fields alongside the uploaded files
	refId := "Holiday"
	res, err := ss.PostFileWithRequest(client, dtos.TestFileUploads{RefId: &refId}, ss.UploadFile{
		FieldName:   "file",
		FileName:    "hello.txt",
		ContentType: "text/plain",
		Reader:      strings.NewReader("file contents"),
	})
	if err != nil {
		t.Fatalf("PostFileWithRequest failed: %v", err)
	}
	if len(res.Files) != 1 || res.Files[0].FileName != "hello.txt" {
		t.Errorf("got Files %+v", res.Files)
	}
	if gotFileName != "hello.txt" || gotFileContent != "file contents" {
		t.Errorf("got file %q with contents %q", gotFileName, gotFileContent)
	}
	if gotField != "Holiday" {
		t.Errorf("got refId field %q, want Holiday", gotField)
	}
}

func TestContextCancellation(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := ss.SendCtx(ctx, client, dtos.Hello{Name: "World"})
	if err == nil {
		t.Fatal("expected an error")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("got %v, want context.Canceled", err)
	}
}

func TestBasePathConfiguration(t *testing.T) {
	client := ss.NewJsonServiceClient("https://example.org")
	if got := client.CreateUrlFromDto(ss.HttpPost, dtos.Hello{}); got != "https://example.org/json/reply/Hello" {
		t.Errorf("got %q", got)
	}
	if client.OneWayBaseUrl != "https://example.org/json/oneway" {
		t.Errorf("got OneWayBaseUrl %q", client.OneWayBaseUrl)
	}

	client = ss.NewClient("https://example.org/")
	if got := client.CreateUrlFromDto(ss.HttpPost, dtos.Hello{}); got != "https://example.org/api/Hello" {
		t.Errorf("got %q", got)
	}
	client.SetBasePath("custom/api")
	if got := client.CreateUrlFromDto(ss.HttpPost, dtos.Hello{}); got != "https://example.org/custom/api/Hello" {
		t.Errorf("got %q", got)
	}
}

func TestSessionCookiesArePreserved(t *testing.T) {
	var gotCookie string
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if cookie, err := r.Cookie("ss-id"); err == nil {
			gotCookie = cookie.Value
		}
		http.SetCookie(w, &http.Cookie{Name: "ss-id", Value: "SESSION_ID", Path: "/"})
		writeJson(t, w, http.StatusOK, dtos.HelloResponse{})
	})

	if _, err := ss.Send(client, dtos.Hello{Name: "1"}); err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if gotCookie != "" {
		t.Errorf("got cookie %q on first request, want none", gotCookie)
	}
	if _, err := ss.Send(client, dtos.Hello{Name: "2"}); err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	if gotCookie != "SESSION_ID" {
		t.Errorf("got cookie %q, want SESSION_ID", gotCookie)
	}
}

func TestAuthenticateUsesReturnedTokens(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeJson(t, w, http.StatusOK, ss.AuthenticateResponse{
			UserName:     "user",
			BearerToken:  "BEARER",
			RefreshToken: "REFRESH",
		})
	})

	res, err := client.Authenticate("user", "pass")
	if err != nil {
		t.Fatalf("Authenticate failed: %v", err)
	}
	if res.UserName != "user" {
		t.Errorf("got UserName %q", res.UserName)
	}
	if client.BearerToken() != "BEARER" || client.RefreshToken() != "REFRESH" {
		t.Errorf("got tokens %q/%q", client.BearerToken(), client.RefreshToken())
	}
}

func TestInvalidJsonResponseReturnsError(t *testing.T) {
	client, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("<html>not json</html>"))
	})

	_, err := ss.Send(client, dtos.Hello{Name: "World"})
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "failed to deserialize") {
		t.Errorf("got %v", err)
	}
}

func TestConnectionErrorReturnsWebServiceException(t *testing.T) {
	client := ss.NewClient("http://127.0.0.1:1")
	_, err := ss.Send(client, dtos.Hello{Name: "World"})
	webEx, ok := ss.AsWebServiceException(err)
	if !ok {
		t.Fatalf("got %T, want *WebServiceException", err)
	}
	if webEx.StatusCode != 0 {
		t.Errorf("got StatusCode %d, want 0", webEx.StatusCode)
	}
	if webEx.Unwrap() == nil {
		t.Error("expected the transport error to be wrapped")
	}
}
