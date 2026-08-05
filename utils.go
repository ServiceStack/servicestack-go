package servicestack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// HTTP Verbs used by ServiceStack APIs.
const (
	HttpGet     = "GET"
	HttpPost    = "POST"
	HttpPut     = "PUT"
	HttpPatch   = "PATCH"
	HttpDelete  = "DELETE"
	HttpOptions = "OPTIONS"
	HttpHead    = "HEAD"
)

// MIME Types and Headers used by ServiceStack APIs.
const (
	MimeTypeJson        = "application/json"
	HeaderAccept        = "Accept"
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
	HeaderUserAgent     = "User-Agent"
)

// HasRequestBody reports whether the HTTP Method sends the Request DTO in the
// Request Body, otherwise it's sent in the QueryString.
func HasRequestBody(method string) bool {
	switch strings.ToUpper(method) {
	case HttpGet, HttpDelete, HttpHead, HttpOptions:
		return false
	}
	return true
}

// NameOf returns the Request DTO Type Name used in ServiceStack's pre-defined routes.
func NameOf(dto any) string {
	if dto == nil {
		return ""
	}
	t := reflect.TypeOf(dto)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := t.Name()
	// Strip type args from generic Request DTOs, e.g. QueryDb[Booking] -> QueryDb
	if pos := strings.IndexByte(name, '['); pos >= 0 {
		name = name[:pos]
	}
	return name
}

// ResolveHttpMethod returns the HTTP Method a Request DTO should be sent with.
//
// It uses the Verb declared by generated DTOs (IVerb), otherwise falls back to
// inferring it from the Request DTO name.
func ResolveHttpMethod(request any) string {
	if verb, ok := request.(IVerb); ok {
		if method := strings.ToUpper(verb.HttpMethod()); method != "" {
			return method
		}
	}
	name := NameOf(request)
	switch {
	case strings.HasPrefix(name, "Get"), strings.HasPrefix(name, "Query"), strings.HasPrefix(name, "Find"), strings.HasPrefix(name, "Search"):
		return HttpGet
	case strings.HasPrefix(name, "Create"):
		return HttpPost
	case strings.HasPrefix(name, "Update"), strings.HasPrefix(name, "Replace"):
		return HttpPut
	case strings.HasPrefix(name, "Patch"):
		return HttpPatch
	case strings.HasPrefix(name, "Delete"), strings.HasPrefix(name, "Remove"):
		return HttpDelete
	}
	return HttpPost
}

// CombineWith joins URL path segments with a single "/" separator.
func CombineWith(basePath string, paths ...string) string {
	sb := strings.Builder{}
	sb.WriteString(strings.TrimSuffix(basePath, "/"))
	for _, path := range paths {
		if path == "" {
			continue
		}
		if !strings.HasPrefix(path, "/") {
			sb.WriteString("/")
		}
		sb.WriteString(strings.TrimSuffix(path, "/"))
	}
	ret := sb.String()
	if ret == "" && basePath == "/" {
		return "/"
	}
	return ret
}

// DtoToMap converts a Request DTO into the map of populated properties to send
// in the QueryString, using the DTO's json tags as the property names.
func DtoToMap(dto any) map[string]any {
	if dto == nil {
		return nil
	}
	jsonBytes, err := json.Marshal(dto)
	if err != nil {
		return nil
	}
	to := map[string]any{}
	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))
	decoder.UseNumber()
	if err := decoder.Decode(&to); err != nil {
		return nil
	}
	return to
}

// AppendQueryString appends the args to the URL's QueryString, ignoring nil
// values. Args are sorted by name so the same args always produce the same URL.
func AppendQueryString(requestUrl string, args map[string]any) string {
	if len(args) == 0 {
		return requestUrl
	}
	keys := make([]string, 0, len(args))
	for key := range args {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	sb := strings.Builder{}
	sb.WriteString(requestUrl)
	sep := "?"
	if strings.Contains(requestUrl, "?") {
		sep = "&"
	}
	for _, key := range keys {
		val := args[key]
		if val == nil {
			continue
		}
		sb.WriteString(sep)
		sep = "&"
		sb.WriteString(url.QueryEscape(key))
		sb.WriteString("=")
		sb.WriteString(url.QueryEscape(QsValue(val)))
	}
	return sb.String()
}

// QsValue converts a value into its ServiceStack QueryString representation.
func QsValue(val any) string {
	switch v := val.(type) {
	case nil:
		return ""
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case json.Number:
		return v.String()
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case time.Time:
		return v.UTC().Format(time.RFC3339)
	case time.Duration:
		return formatTimeSpan(v)
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface:
		if rv.IsNil() {
			return ""
		}
		return QsValue(rv.Elem().Interface())
	case reflect.Slice, reflect.Array:
		vals := make([]string, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			vals[i] = QsValue(rv.Index(i).Interface())
		}
		return "[" + strings.Join(vals, ",") + "]"
	case reflect.Map:
		keys := rv.MapKeys()
		vals := make([]string, 0, len(keys))
		for _, key := range keys {
			vals = append(vals, fmt.Sprintf("%v:%s", key.Interface(), QsValue(rv.MapIndex(key).Interface())))
		}
		sort.Strings(vals)
		return "{" + strings.Join(vals, ",") + "}"
	case reflect.Struct:
		if jsonBytes, err := json.Marshal(val); err == nil {
			return string(jsonBytes)
		}
	}
	return fmt.Sprintf("%v", val)
}

// formatTimeSpan formats a Duration in the XSD duration format used by
// ServiceStack's TimeSpan properties, e.g. PT1H30M.
func formatTimeSpan(d time.Duration) string {
	sb := strings.Builder{}
	if d < 0 {
		sb.WriteString("-")
		d = -d
	}
	sb.WriteString("P")
	days := int64(d / (24 * time.Hour))
	if days > 0 {
		sb.WriteString(strconv.FormatInt(days, 10) + "D")
		d -= time.Duration(days) * 24 * time.Hour
	}
	if d > 0 {
		sb.WriteString("T")
		hours := int64(d / time.Hour)
		if hours > 0 {
			sb.WriteString(strconv.FormatInt(hours, 10) + "H")
			d -= time.Duration(hours) * time.Hour
		}
		mins := int64(d / time.Minute)
		if mins > 0 {
			sb.WriteString(strconv.FormatInt(mins, 10) + "M")
			d -= time.Duration(mins) * time.Minute
		}
		if d > 0 {
			secs := float64(d) / float64(time.Second)
			sb.WriteString(strconv.FormatFloat(secs, 'f', -1, 64) + "S")
		}
	} else if days == 0 {
		sb.WriteString("T0S")
	}
	return sb.String()
}

// ToAbsoluteUrl converts a relative path into an absolute URL of the baseUrl.
func ToAbsoluteUrl(baseUrl, pathOrUrl string) string {
	if strings.HasPrefix(pathOrUrl, "http://") || strings.HasPrefix(pathOrUrl, "https://") {
		return pathOrUrl
	}
	return CombineWith(baseUrl, pathOrUrl)
}

func equalsIgnoreCase(a, b string) bool {
	return strings.EqualFold(a, b)
}
