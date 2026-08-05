package servicestack

import (
	"testing"
	"time"
)

func TestCombineWith(t *testing.T) {
	tests := []struct {
		base  string
		paths []string
		want  string
	}{
		{"https://example.org", []string{"api"}, "https://example.org/api"},
		{"https://example.org/", []string{"api"}, "https://example.org/api"},
		{"https://example.org/", []string{"/api/"}, "https://example.org/api"},
		{"https://example.org", []string{"api", "Hello"}, "https://example.org/api/Hello"},
		{"https://example.org", []string{"", "Hello"}, "https://example.org/Hello"},
		{"https://example.org", nil, "https://example.org"},
	}
	for _, test := range tests {
		if got := CombineWith(test.base, test.paths...); got != test.want {
			t.Errorf("CombineWith(%q, %v) = %q, want %q", test.base, test.paths, got, test.want)
		}
	}
}

func TestAppendQueryString(t *testing.T) {
	tests := []struct {
		url  string
		args map[string]any
		want string
	}{
		{"/api/Hello", nil, "/api/Hello"},
		{"/api/Hello", map[string]any{"name": "World"}, "/api/Hello?name=World"},
		{"/api/Hello?a=1", map[string]any{"name": "World"}, "/api/Hello?a=1&name=World"},
		{"/api/Hello", map[string]any{"b": 2, "a": 1}, "/api/Hello?a=1&b=2"},
		{"/api/Hello", map[string]any{"name": nil}, "/api/Hello"},
		{"/api/Hello", map[string]any{"name": "A B&C"}, "/api/Hello?name=A+B%26C"},
		{"/api/Hello", map[string]any{"ids": []int{1, 2, 3}}, "/api/Hello?ids=%5B1%2C2%2C3%5D"},
		{"/api/Hello", map[string]any{"enabled": true}, "/api/Hello?enabled=true"},
	}
	for _, test := range tests {
		if got := AppendQueryString(test.url, test.args); got != test.want {
			t.Errorf("AppendQueryString(%q, %v) = %q, want %q", test.url, test.args, got, test.want)
		}
	}
}

func TestQsValue(t *testing.T) {
	name := "World"
	tests := []struct {
		val  any
		want string
	}{
		{nil, ""},
		{"World", "World"},
		{&name, "World"},
		{(*string)(nil), ""},
		{true, "true"},
		{10, "10"},
		{1.5, "1.5"},
		{[]string{"a", "b"}, "[a,b]"},
		{map[string]int{"b": 2, "a": 1}, "{a:1,b:2}"},
		{time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC), "2001-01-01T00:00:00Z"},
		{90 * time.Minute, "PT1H30M"},
		{25 * time.Hour, "P1DT1H"},
		{0 * time.Second, "PT0S"},
	}
	for _, test := range tests {
		if got := QsValue(test.val); got != test.want {
			t.Errorf("QsValue(%v) = %q, want %q", test.val, got, test.want)
		}
	}
}

func TestNameOf(t *testing.T) {
	if got := NameOf(Hello{}); got != "Hello" {
		t.Errorf("got %q, want Hello", got)
	}
	if got := NameOf(&Hello{}); got != "Hello" {
		t.Errorf("got %q, want Hello", got)
	}
	if got := NameOf(nil); got != "" {
		t.Errorf("got %q, want empty", got)
	}
	if got := NameOf(QueryResponse[Booking]{}); got != "QueryResponse" {
		t.Errorf("got %q, want QueryResponse", got)
	}
}

func TestResolveHttpMethod(t *testing.T) {
	tests := []struct {
		request any
		want    string
	}{
		{Hello{}, HttpGet},          // declared with HttpMethod()
		{CreateHello{}, HttpPost},   // declared with HttpMethod()
		{DeleteHello{}, HttpDelete}, // declared with HttpMethod()
		{struct{}{}, HttpPost},      // default
		{GetApiKeys{}, HttpGet},     // declared with HttpMethod()
		{Authenticate{}, HttpPost},  // declared with HttpMethod()
	}
	for _, test := range tests {
		if got := ResolveHttpMethod(test.request); got != test.want {
			t.Errorf("ResolveHttpMethod(%T) = %q, want %q", test.request, got, test.want)
		}
	}

	// Fallback to inferring the Verb from the Request DTO name
	type QueryTodos struct{}
	type UpdateTodo struct{}
	type DeleteTodo struct{}
	type PatchTodo struct{}
	type SaveTodo struct{}
	nameTests := []struct {
		request any
		want    string
	}{
		{QueryTodos{}, HttpGet},
		{UpdateTodo{}, HttpPut},
		{DeleteTodo{}, HttpDelete},
		{PatchTodo{}, HttpPatch},
		{SaveTodo{}, HttpPost},
	}
	for _, test := range nameTests {
		if got := ResolveHttpMethod(test.request); got != test.want {
			t.Errorf("ResolveHttpMethod(%T) = %q, want %q", test.request, got, test.want)
		}
	}
}

func TestHasRequestBody(t *testing.T) {
	for _, method := range []string{HttpPost, HttpPut, HttpPatch} {
		if !HasRequestBody(method) {
			t.Errorf("HasRequestBody(%q) = false, want true", method)
		}
	}
	for _, method := range []string{HttpGet, HttpDelete, HttpHead, HttpOptions, "get"} {
		if HasRequestBody(method) {
			t.Errorf("HasRequestBody(%q) = true, want false", method)
		}
	}
}

func TestDtoToMapOmitsEmptyProperties(t *testing.T) {
	take := 10
	args := DtoToMap(QueryBookings{QueryDb: QueryDb{QueryBase{Take: &take}}})
	if len(args) != 1 {
		t.Fatalf("got %v, want only take", args)
	}
	if QsValue(args["take"]) != "10" {
		t.Errorf("got take %v, want 10", args["take"])
	}
}

func TestToAbsoluteUrl(t *testing.T) {
	if got := ToAbsoluteUrl("https://example.org", "/api/Hello"); got != "https://example.org/api/Hello" {
		t.Errorf("got %q", got)
	}
	if got := ToAbsoluteUrl("https://example.org", "https://other.org/api"); got != "https://other.org/api" {
		t.Errorf("got %q", got)
	}
}
