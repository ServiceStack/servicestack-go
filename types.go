package servicestack

import "time"

// ResponseStatus is ServiceStack's structured error response, returned in the
// `responseStatus` property of failed API Responses.
type ResponseStatus struct {
	ErrorCode  string            `json:"errorCode,omitempty"`
	Message    string            `json:"message,omitempty"`
	StackTrace string            `json:"stackTrace,omitempty"`
	Errors     []ResponseError   `json:"errors,omitempty"`
	Meta       map[string]string `json:"meta,omitempty"`
}

// FieldError returns the error message for the specified field, or "" if the
// field has no error. Field names are matched case-insensitively.
func (s *ResponseStatus) FieldError(fieldName string) string {
	if err := s.GetFieldError(fieldName); err != nil {
		return err.Message
	}
	return ""
}

// GetFieldError returns the ResponseError for the specified field, or nil if
// the field has no error. Field names are matched case-insensitively.
func (s *ResponseStatus) GetFieldError(fieldName string) *ResponseError {
	if s == nil {
		return nil
	}
	for i := range s.Errors {
		if equalsIgnoreCase(s.Errors[i].FieldName, fieldName) {
			return &s.Errors[i]
		}
	}
	return nil
}

// Error implements the error interface so a ResponseStatus can be returned as an error.
func (s *ResponseStatus) Error() string {
	if s == nil {
		return ""
	}
	if s.Message != "" {
		return s.Message
	}
	return s.ErrorCode
}

// ResponseError is a field validation error within a ResponseStatus.
type ResponseError struct {
	ErrorCode string            `json:"errorCode,omitempty"`
	FieldName string            `json:"fieldName,omitempty"`
	Message   string            `json:"message,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
}

// EmptyResponse is returned by APIs with no Response Body.
type EmptyResponse struct {
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
}

// IdResponse is returned by APIs that return the Id of the created or updated entity.
type IdResponse struct {
	Id             string          `json:"id,omitempty"`
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
}

// StringResponse is returned by APIs that return a single string result.
type StringResponse struct {
	Result         string            `json:"result,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// StringsResponse is returned by APIs that return a list of string results.
type StringsResponse struct {
	Results        []string          `json:"results,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// AuditBase contains the audit fields of AutoQuery CRUD data models.
type AuditBase struct {
	CreatedDate  time.Time  `json:"createdDate,omitempty"`
	CreatedBy    string     `json:"createdBy,omitempty"`
	ModifiedDate time.Time  `json:"modifiedDate,omitempty"`
	ModifiedBy   string     `json:"modifiedBy,omitempty"`
	DeletedDate  *time.Time `json:"deletedDate,omitempty"`
	DeletedBy    *string    `json:"deletedBy,omitempty"`
}

// QueryBase contains the query params supported by all AutoQuery Requests.
type QueryBase struct {
	Skip        *int              `json:"skip,omitempty"`
	Take        *int              `json:"take,omitempty"`
	OrderBy     string            `json:"orderBy,omitempty"`
	OrderByDesc string            `json:"orderByDesc,omitempty"`
	Include     string            `json:"include,omitempty"`
	Fields      string            `json:"fields,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
}

// QueryDb is the base type of AutoQuery RDBMS Requests, embedded in generated
// AutoQuery Request DTOs.
type QueryDb struct {
	QueryBase
}

// QueryData is the base type of AutoQuery Data Requests, embedded in generated
// AutoQuery Data Request DTOs.
type QueryData struct {
	QueryBase
}

// QueryResponse is the typed response of AutoQuery Requests.
type QueryResponse[T any] struct {
	Offset         int               `json:"offset,omitempty"`
	Total          int               `json:"total,omitempty"`
	Results        []T               `json:"results,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// Authenticate is the Request DTO for authenticating with a ServiceStack Service.
type Authenticate struct {
	Provider          string            `json:"provider,omitempty"`
	UserName          string            `json:"userName,omitempty"`
	Password          string            `json:"password,omitempty"`
	RememberMe        *bool             `json:"rememberMe,omitempty"`
	AccessToken       string            `json:"accessToken,omitempty"`
	AccessTokenSecret string            `json:"accessTokenSecret,omitempty"`
	ReturnUrl         string            `json:"returnUrl,omitempty"`
	ErrorView         string            `json:"errorView,omitempty"`
	Meta              map[string]string `json:"meta,omitempty"`
}

func (Authenticate) CreateResponse() (r AuthenticateResponse) { return }
func (Authenticate) HttpMethod() string                       { return HttpPost }

// AuthenticateResponse is the Response DTO of a successful Authenticate Request.
type AuthenticateResponse struct {
	UserId             string            `json:"userId,omitempty"`
	SessionId          string            `json:"sessionId,omitempty"`
	UserName           string            `json:"userName,omitempty"`
	DisplayName        string            `json:"displayName,omitempty"`
	ReferrerUrl        string            `json:"referrerUrl,omitempty"`
	BearerToken        string            `json:"bearerToken,omitempty"`
	RefreshToken       string            `json:"refreshToken,omitempty"`
	RefreshTokenExpiry *time.Time        `json:"refreshTokenExpiry,omitempty"`
	ProfileUrl         string            `json:"profileUrl,omitempty"`
	Roles              []string          `json:"roles,omitempty"`
	Permissions        []string          `json:"permissions,omitempty"`
	AuthProvider       string            `json:"authProvider,omitempty"`
	ResponseStatus     *ResponseStatus   `json:"responseStatus,omitempty"`
	Meta               map[string]string `json:"meta,omitempty"`
}

// Register is the Request DTO for registering a new User.
type Register struct {
	UserName        string            `json:"userName,omitempty"`
	FirstName       string            `json:"firstName,omitempty"`
	LastName        string            `json:"lastName,omitempty"`
	DisplayName     string            `json:"displayName,omitempty"`
	Email           string            `json:"email,omitempty"`
	Password        string            `json:"password,omitempty"`
	ConfirmPassword string            `json:"confirmPassword,omitempty"`
	AutoLogin       *bool             `json:"autoLogin,omitempty"`
	ErrorView       string            `json:"errorView,omitempty"`
	Meta            map[string]string `json:"meta,omitempty"`
}

func (Register) CreateResponse() (r RegisterResponse) { return }
func (Register) HttpMethod() string                   { return HttpPost }

// RegisterResponse is the Response DTO of a successful Register Request.
type RegisterResponse struct {
	UserId             string            `json:"userId,omitempty"`
	SessionId          string            `json:"sessionId,omitempty"`
	UserName           string            `json:"userName,omitempty"`
	ReferrerUrl        string            `json:"referrerUrl,omitempty"`
	BearerToken        string            `json:"bearerToken,omitempty"`
	RefreshToken       string            `json:"refreshToken,omitempty"`
	RefreshTokenExpiry *time.Time        `json:"refreshTokenExpiry,omitempty"`
	Roles              []string          `json:"roles,omitempty"`
	Permissions        []string          `json:"permissions,omitempty"`
	RedirectUrl        string            `json:"redirectUrl,omitempty"`
	ResponseStatus     *ResponseStatus   `json:"responseStatus,omitempty"`
	Meta               map[string]string `json:"meta,omitempty"`
}

// AssignRoles is the Request DTO for assigning Roles and Permissions to a User.
type AssignRoles struct {
	UserName    string            `json:"userName,omitempty"`
	Permissions []string          `json:"permissions,omitempty"`
	Roles       []string          `json:"roles,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
}

func (AssignRoles) CreateResponse() (r AssignRolesResponse) { return }
func (AssignRoles) HttpMethod() string                      { return HttpPost }

// AssignRolesResponse is the Response DTO of a successful AssignRoles Request.
type AssignRolesResponse struct {
	AllRoles       []string          `json:"allRoles,omitempty"`
	AllPermissions []string          `json:"allPermissions,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// UnAssignRoles is the Request DTO for removing Roles and Permissions from a User.
type UnAssignRoles struct {
	UserName    string            `json:"userName,omitempty"`
	Permissions []string          `json:"permissions,omitempty"`
	Roles       []string          `json:"roles,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
}

func (UnAssignRoles) CreateResponse() (r UnAssignRolesResponse) { return }
func (UnAssignRoles) HttpMethod() string                        { return HttpPost }

// UnAssignRolesResponse is the Response DTO of a successful UnAssignRoles Request.
type UnAssignRolesResponse struct {
	AllRoles       []string          `json:"allRoles,omitempty"`
	AllPermissions []string          `json:"allPermissions,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// ConvertSessionToToken converts an authenticated Session into a JWT Bearer Token.
type ConvertSessionToToken struct {
	PreserveSession bool              `json:"preserveSession,omitempty"`
	Meta            map[string]string `json:"meta,omitempty"`
}

func (ConvertSessionToToken) CreateResponse() (r ConvertSessionToTokenResponse) { return }
func (ConvertSessionToToken) HttpMethod() string                                { return HttpPost }

// ConvertSessionToTokenResponse is the Response DTO of ConvertSessionToToken.
type ConvertSessionToTokenResponse struct {
	Meta           map[string]string `json:"meta,omitempty"`
	AccessToken    string            `json:"accessToken,omitempty"`
	RefreshToken   string            `json:"refreshToken,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// GetAccessToken exchanges a Refresh Token for a new JWT Bearer Token.
type GetAccessToken struct {
	RefreshToken string            `json:"refreshToken,omitempty"`
	Meta         map[string]string `json:"meta,omitempty"`
}

func (GetAccessToken) CreateResponse() (r GetAccessTokenResponse) { return }
func (GetAccessToken) HttpMethod() string                         { return HttpPost }

// GetAccessTokenResponse is the Response DTO of GetAccessToken.
type GetAccessTokenResponse struct {
	AccessToken    string            `json:"accessToken,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// UserApiKey is an API Key assigned to a User.
type UserApiKey struct {
	Key        string            `json:"key,omitempty"`
	KeyType    string            `json:"keyType,omitempty"`
	ExpiryDate *time.Time        `json:"expiryDate,omitempty"`
	Meta       map[string]string `json:"meta,omitempty"`
}

// GetApiKeys returns the API Keys assigned to the authenticated User.
type GetApiKeys struct {
	Environment string            `json:"environment,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
}

func (GetApiKeys) CreateResponse() (r GetApiKeysResponse) { return }
func (GetApiKeys) HttpMethod() string                     { return HttpGet }

// GetApiKeysResponse is the Response DTO of GetApiKeys.
type GetApiKeysResponse struct {
	Results        []UserApiKey      `json:"results,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// RegenerateApiKeys regenerates the API Keys of the authenticated User.
type RegenerateApiKeys struct {
	Environment string            `json:"environment,omitempty"`
	Meta        map[string]string `json:"meta,omitempty"`
}

func (RegenerateApiKeys) CreateResponse() (r RegenerateApiKeysResponse) { return }
func (RegenerateApiKeys) HttpMethod() string                            { return HttpPost }

// RegenerateApiKeysResponse is the Response DTO of RegenerateApiKeys.
type RegenerateApiKeysResponse struct {
	Results        []UserApiKey      `json:"results,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// NavItem is a navigation item returned by GetNavItems.
type NavItem struct {
	Label     string            `json:"label,omitempty"`
	Href      string            `json:"href,omitempty"`
	Exact     *bool             `json:"exact,omitempty"`
	Id        string            `json:"id,omitempty"`
	ClassName string            `json:"className,omitempty"`
	IconClass string            `json:"iconClass,omitempty"`
	IconSrc   string            `json:"iconSrc,omitempty"`
	Show      string            `json:"show,omitempty"`
	Hide      string            `json:"hide,omitempty"`
	Children  []NavItem         `json:"children,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
}

// GetNavItems returns the Site's navigation items.
type GetNavItems struct {
	Name string `json:"name,omitempty"`
}

func (GetNavItems) CreateResponse() (r GetNavItemsResponse) { return }
func (GetNavItems) HttpMethod() string                      { return HttpGet }

// GetNavItemsResponse is the Response DTO of GetNavItems.
type GetNavItemsResponse struct {
	BaseUrl        string               `json:"baseUrl,omitempty"`
	Results        []NavItem            `json:"results,omitempty"`
	NavItemsMap    map[string][]NavItem `json:"navItemsMap,omitempty"`
	Meta           map[string]string    `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus      `json:"responseStatus,omitempty"`
}
