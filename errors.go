package servicestack

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// WebServiceException is the error returned for failed API Requests, containing
// the HTTP Status Code and the structured ResponseStatus error (when available).
//
// Use errors.As to inspect it:
//
//	if _, err := servicestack.Send(client, Hello{}); err != nil {
//	    var webEx *servicestack.WebServiceException
//	    if errors.As(err, &webEx) {
//	        fmt.Println(webEx.StatusCode, webEx.ErrorCode(), webEx.ErrorMessage())
//	    }
//	}
type WebServiceException struct {
	StatusCode        int
	StatusDescription string
	ResponseStatus    *ResponseStatus
	// ResponseBody is the raw error response body, useful when the Service
	// returned a non ResponseStatus error (e.g. a plain text or HTML error).
	ResponseBody string
	// Inner is the underlying transport or serialization error, if any.
	Inner error
}

func (e *WebServiceException) Error() string {
	msg := e.ErrorMessage()
	if msg == "" {
		msg = e.StatusDescription
	}
	if msg == "" {
		msg = http.StatusText(e.StatusCode)
	}
	errorCode := e.ErrorCode()
	if errorCode != "" && errorCode != msg {
		return fmt.Sprintf("%d %s: %s", e.StatusCode, errorCode, msg)
	}
	return fmt.Sprintf("%d %s", e.StatusCode, msg)
}

// Unwrap returns the underlying transport error, if any.
func (e *WebServiceException) Unwrap() error { return e.Inner }

// ErrorCode returns the ResponseStatus ErrorCode, e.g. "NotFound", "ArgumentNullException".
func (e *WebServiceException) ErrorCode() string {
	if e.ResponseStatus != nil {
		return e.ResponseStatus.ErrorCode
	}
	return ""
}

// ErrorMessage returns the ResponseStatus error message.
func (e *WebServiceException) ErrorMessage() string {
	if e.ResponseStatus != nil {
		return e.ResponseStatus.Message
	}
	return ""
}

// StackTrace returns the Server StackTrace, only populated in DebugMode.
func (e *WebServiceException) StackTrace() string {
	if e.ResponseStatus != nil {
		return e.ResponseStatus.StackTrace
	}
	return ""
}

// FieldErrors returns any field validation errors.
func (e *WebServiceException) FieldErrors() []ResponseError {
	if e.ResponseStatus != nil {
		return e.ResponseStatus.Errors
	}
	return nil
}

// FieldError returns the error message for the specified field, or "" if the
// field has no error.
func (e *WebServiceException) FieldError(fieldName string) string {
	if e.ResponseStatus == nil {
		return ""
	}
	return e.ResponseStatus.FieldError(fieldName)
}

// IsAny reports whether the Response returned any of the HTTP Status Codes.
func (e *WebServiceException) IsAny(statusCodes ...int) bool {
	for _, statusCode := range statusCodes {
		if e.StatusCode == statusCode {
			return true
		}
	}
	return false
}

// IsUnauthorized reports whether the Request requires Authentication (401).
func (e *WebServiceException) IsUnauthorized() bool { return e.StatusCode == http.StatusUnauthorized }

// IsForbidden reports whether the User was denied access to the Request (403).
func (e *WebServiceException) IsForbidden() bool { return e.StatusCode == http.StatusForbidden }

// IsNotFound reports whether the Request returned 404 NotFound.
func (e *WebServiceException) IsNotFound() bool { return e.StatusCode == http.StatusNotFound }

// IsValidationError reports whether the Response contains field validation errors.
func (e *WebServiceException) IsValidationError() bool { return len(e.FieldErrors()) > 0 }

// AsWebServiceException returns the *WebServiceException in err's chain, if any.
func AsWebServiceException(err error) (*WebServiceException, bool) {
	var webEx *WebServiceException
	if errors.As(err, &webEx) {
		return webEx, true
	}
	return nil, false
}

// GetStatusCode returns the HTTP Status Code of a failed API Request, or 0 if
// err wasn't a *WebServiceException.
func GetStatusCode(err error) int {
	if webEx, ok := AsWebServiceException(err); ok {
		return webEx.StatusCode
	}
	return 0
}

// GetResponseStatus returns the structured error of a failed API Request, or
// nil if err wasn't a *WebServiceException with a ResponseStatus.
func GetResponseStatus(err error) *ResponseStatus {
	if webEx, ok := AsWebServiceException(err); ok {
		return webEx.ResponseStatus
	}
	return nil
}

func toErrorStatus(err error) *ResponseStatus {
	if err == nil {
		return nil
	}
	if webEx, ok := AsWebServiceException(err); ok {
		if webEx.ResponseStatus != nil {
			return webEx.ResponseStatus
		}
		return &ResponseStatus{
			ErrorCode: strings.TrimSpace(http.StatusText(webEx.StatusCode)),
			Message:   webEx.Error(),
		}
	}
	return &ResponseStatus{
		ErrorCode: "Exception",
		Message:   err.Error(),
	}
}
