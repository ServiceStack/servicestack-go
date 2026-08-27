/* Options:
Date: 2026-08-07 02:41:34
Version: 10.09
Tip: To override a DTO option, remove "//" prefix before updating
BaseUrl: https://test.servicestack.net

//GlobalNamespace:
//MakePropertiesOptional: False
//AddServiceStackTypes: True
//AddResponseStatus: False
//AddImplicitVersion:
//AddDescriptionAsComments: True
//IncludeTypes:
//ExcludeTypes:
//DefaultImports:
*/

package dtos

import (
	ss "github.com/ServiceStack/servicestack-go"
	"time"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Poco struct {
	Name string `json:"name"`
}

type CustomType struct {
	Id   int     `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SetterType struct {
	Id   int     `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type DeclarativeChildValidation struct {
	Name string `json:"name"`
	// @Validate(Validator="MaximumLength(20)")
	Value string `json:"value"`
}

type FluentChildValidation struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DeclarativeSingleValidation struct {
	Name string `json:"name"`
	// @Validate(Validator="MaximumLength(20)")
	Value string `json:"value"`
}

type FluentSingleValidation struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// @DataContract
type CancelRequest struct {
	// @DataMember(Order=1)
	Tag *string `json:"tag,omitempty"`
	// @DataMember(Order=2)
	Meta map[string]string `json:"meta,omitempty"`
}

// @DataContract
type CancelRequestResponse struct {
	// @DataMember(Order=1)
	Tag *string `json:"tag,omitempty"`
	// @DataMember(Order=2)
	Elapsed time.Duration `json:"elapsed,omitempty"`
	// @DataMember(Order=3)
	Meta map[string]string `json:"meta,omitempty"`
	// @DataMember(Order=4)
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

// @DataContract
type UpdateEventSubscriber struct {
	// @DataMember(Order=1)
	Id *string `json:"id,omitempty"`
	// @DataMember(Order=2)
	SubscribeChannels []string `json:"subscribeChannels,omitempty"`
	// @DataMember(Order=3)
	UnsubscribeChannels []string `json:"unsubscribeChannels,omitempty"`
}

// @DataContract
type UpdateEventSubscriberResponse struct {
	// @DataMember(Order=1)
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type IAuthTokens struct {
	Provider           *string           `json:"provider,omitempty"`
	UserId             *string           `json:"userId,omitempty"`
	AccessToken        *string           `json:"accessToken,omitempty"`
	AccessTokenSecret  *string           `json:"accessTokenSecret,omitempty"`
	RefreshToken       *string           `json:"refreshToken,omitempty"`
	RefreshTokenExpiry *time.Time        `json:"refreshTokenExpiry,omitempty"`
	RequestToken       *string           `json:"requestToken,omitempty"`
	RequestTokenSecret *string           `json:"requestTokenSecret,omitempty"`
	Items              map[string]string `json:"items,omitempty"`
}

// @DataContract
type AuthUserSession struct {
	// @DataMember(Order=1)
	ReferrerUrl *string `json:"referrerUrl,omitempty"`
	// @DataMember(Order=2)
	Id *string `json:"id,omitempty"`
	// @DataMember(Order=3)
	UserAuthId *string `json:"userAuthId,omitempty"`
	// @DataMember(Order=4)
	UserAuthName *string `json:"userAuthName,omitempty"`
	// @DataMember(Order=5)
	UserName *string `json:"userName,omitempty"`
	// @DataMember(Order=6)
	TwitterUserId *string `json:"twitterUserId,omitempty"`
	// @DataMember(Order=7)
	TwitterScreenName *string `json:"twitterScreenName,omitempty"`
	// @DataMember(Order=8)
	FacebookUserId *string `json:"facebookUserId,omitempty"`
	// @DataMember(Order=9)
	FacebookUserName *string `json:"facebookUserName,omitempty"`
	// @DataMember(Order=10)
	FirstName *string `json:"firstName,omitempty"`
	// @DataMember(Order=11)
	LastName *string `json:"lastName,omitempty"`
	// @DataMember(Order=12)
	DisplayName *string `json:"displayName,omitempty"`
	// @DataMember(Order=13)
	Company *string `json:"company,omitempty"`
	// @DataMember(Order=14)
	Email *string `json:"email,omitempty"`
	// @DataMember(Order=15)
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	// @DataMember(Order=16)
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// @DataMember(Order=17)
	BirthDate *time.Time `json:"birthDate,omitempty"`
	// @DataMember(Order=18)
	BirthDateRaw *string `json:"birthDateRaw,omitempty"`
	// @DataMember(Order=19)
	Address *string `json:"address,omitempty"`
	// @DataMember(Order=20)
	Address2 *string `json:"address2,omitempty"`
	// @DataMember(Order=21)
	City *string `json:"city,omitempty"`
	// @DataMember(Order=22)
	State *string `json:"state,omitempty"`
	// @DataMember(Order=23)
	Country *string `json:"country,omitempty"`
	// @DataMember(Order=24)
	Culture *string `json:"culture,omitempty"`
	// @DataMember(Order=25)
	FullName *string `json:"fullName,omitempty"`
	// @DataMember(Order=26)
	Gender *string `json:"gender,omitempty"`
	// @DataMember(Order=27)
	Language *string `json:"language,omitempty"`
	// @DataMember(Order=28)
	MailAddress *string `json:"mailAddress,omitempty"`
	// @DataMember(Order=29)
	Nickname *string `json:"nickname,omitempty"`
	// @DataMember(Order=30)
	PostalCode *string `json:"postalCode,omitempty"`
	// @DataMember(Order=31)
	TimeZone *string `json:"timeZone,omitempty"`
	// @DataMember(Order=32)
	RequestTokenSecret *string `json:"requestTokenSecret,omitempty"`
	// @DataMember(Order=33)
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// @DataMember(Order=34)
	LastModified time.Time `json:"lastModified,omitempty"`
	// @DataMember(Order=35)
	Roles []string `json:"roles,omitempty"`
	// @DataMember(Order=36)
	Permissions []string `json:"permissions,omitempty"`
	// @DataMember(Order=37)
	IsAuthenticated bool `json:"isAuthenticated,omitempty"`
	// @DataMember(Order=38)
	FromToken bool `json:"fromToken,omitempty"`
	// @DataMember(Order=39)
	ProfileUrl *string `json:"profileUrl,omitempty"`
	// @DataMember(Order=40)
	Sequence *string `json:"sequence,omitempty"`
	// @DataMember(Order=41)
	Tag int64 `json:"tag,omitempty"`
	// @DataMember(Order=42)
	AuthProvider *string `json:"authProvider,omitempty"`
	// @DataMember(Order=43)
	ProviderOAuthAccess []IAuthTokens `json:"providerOAuthAccess,omitempty"`
	// @DataMember(Order=44)
	Meta map[string]string `json:"meta,omitempty"`
	// @DataMember(Order=45)
	Audiences []string `json:"audiences,omitempty"`
	// @DataMember(Order=46)
	Scopes []string `json:"scopes,omitempty"`
	// @DataMember(Order=47)
	Dns *string `json:"dns,omitempty"`
	// @DataMember(Order=48)
	Rsa *string `json:"rsa,omitempty"`
	// @DataMember(Order=49)
	Sid *string `json:"sid,omitempty"`
	// @DataMember(Order=50)
	Hash *string `json:"hash,omitempty"`
	// @DataMember(Order=51)
	HomePhone *string `json:"homePhone,omitempty"`
	// @DataMember(Order=52)
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// @DataMember(Order=53)
	Webpage *string `json:"webpage,omitempty"`
	// @DataMember(Order=54)
	EmailConfirmed *bool `json:"emailConfirmed,omitempty"`
	// @DataMember(Order=55)
	PhoneNumberConfirmed *bool `json:"phoneNumberConfirmed,omitempty"`
	// @DataMember(Order=56)
	TwoFactorEnabled *bool `json:"twoFactorEnabled,omitempty"`
	// @DataMember(Order=57)
	SecurityStamp *string `json:"securityStamp,omitempty"`
	// @DataMember(Order=58)
	Type *string `json:"type,omitempty"`
	// @DataMember(Order=59)
	RecoveryToken *string `json:"recoveryToken,omitempty"`
	// @DataMember(Order=60)
	RefId *int `json:"refId,omitempty"`
	// @DataMember(Order=61)
	RefIdStr *string `json:"refIdStr,omitempty"`
}

type NestedClass struct {
	Value string `json:"value"`
}

type EnumType string

const (
	EnumTypeValue1 EnumType = "Value1"
	EnumTypeValue2          = "Value2"
	EnumTypeValue3          = "Value3"
)

// @Flags()
type EnumTypeFlags int

const (
	EnumTypeFlagsValue1 EnumTypeFlags = 0
	EnumTypeFlagsValue2 EnumTypeFlags = 1
	EnumTypeFlagsValue3 EnumTypeFlags = 2
)

type EnumWithValues string

const (
	EnumWithValuesNone   EnumWithValues = "None"
	EnumWithValuesValue1                = "Member 1"
	EnumWithValuesValue2                = "Value2"
)

// @Flags()
type EnumFlags int

const (
	EnumFlagsValue0   EnumFlags = 0
	EnumFlagsValue1   EnumFlags = 1
	EnumFlagsValue2   EnumFlags = 2
	EnumFlagsValue3   EnumFlags = 4
	EnumFlagsValue123 EnumFlags = 7
)

type EnumAsInt string

const (
	EnumAsIntValue1 EnumAsInt = "Value1"
	EnumAsIntValue2           = "Value2"
	EnumAsIntValue3           = "Value3"
)

type EnumStyle string

const (
	EnumStylelower       EnumStyle = "lower"
	EnumStyleUPPER                 = "UPPER"
	EnumStylePascalCase            = "PascalCase"
	EnumStylecamelCase             = "camelCase"
	EnumStylecamelUPPER            = "camelUPPER"
	EnumStylePascalUPPER           = "PascalUPPER"
)

type EnumStyleMembers string

const (
	EnumStyleMembersLower       EnumStyleMembers = "lower"
	EnumStyleMembersUpper                        = "UPPER"
	EnumStyleMembersPascalCase                   = "PascalCase"
	EnumStyleMembersCamelCase                    = "camelCase"
	EnumStyleMembersCamelUpper                   = "camelUPPER"
	EnumStyleMembersPascalUpper                  = "PascalUPPER"
)

type KeyValuePair[TKey any, TValue any] struct {
	Key   TKey   `json:"key"`
	Value TValue `json:"value"`
}

type SubType struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type AllTypesBase struct {
	Id               int                          `json:"id,omitempty"`
	NullableId       *int                         `json:"nullableId,omitempty"`
	Byte             byte                         `json:"byte,omitempty"`
	Short            int16                        `json:"short,omitempty"`
	Int              int                          `json:"int,omitempty"`
	Long             int64                        `json:"long,omitempty"`
	UShort           uint16                       `json:"uShort,omitempty"`
	UInt             uint32                       `json:"uInt,omitempty"`
	ULong            uint64                       `json:"uLong,omitempty"`
	Float            float32                      `json:"float,omitempty"`
	Double           float64                      `json:"double,omitempty"`
	Decimal          float64                      `json:"decimal,omitempty"`
	String           string                       `json:"string"`
	DateTime         time.Time                    `json:"dateTime,omitempty"`
	TimeSpan         time.Duration                `json:"timeSpan,omitempty"`
	DateTimeOffset   time.Time                    `json:"dateTimeOffset,omitempty"`
	Guid             string                       `json:"guid,omitempty"`
	Char             string                       `json:"char,omitempty"`
	KeyValuePair     KeyValuePair[string, string] `json:"keyValuePair,omitempty"`
	NullableDateTime *time.Time                   `json:"nullableDateTime,omitempty"`
	NullableTimeSpan *time.Duration               `json:"nullableTimeSpan,omitempty"`
	StringList       []string                     `json:"stringList"`
	StringArray      []string                     `json:"stringArray"`
	StringMap        map[string]string            `json:"stringMap"`
	IntStringMap     map[int]string               `json:"intStringMap"`
	SubType          SubType                      `json:"subType"`
}

type HelloBase struct {
	Id int `json:"id,omitempty"`
}

type HelloBase_1[T any] struct {
	Items  []T   `json:"items"`
	Counts []int `json:"counts"`
}

type IPoco struct {
	Name string `json:"name"`
}

type IEmptyInterface struct {
}

type EmptyClass struct {
}

type DayOfWeek string

const (
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekMonday              = "Monday"
	DayOfWeekTuesday             = "Tuesday"
	DayOfWeekWednesday           = "Wednesday"
	DayOfWeekThursday            = "Thursday"
	DayOfWeekFriday              = "Friday"
	DayOfWeekSaturday            = "Saturday"
)

// @DataContract
type ScopeType string

const (
	ScopeTypeGlobal ScopeType = "Global"
	ScopeTypeSale             = "Sale"
)

type Channel struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Device struct {
	Id        int64     `json:"id,omitempty"`
	Type      string    `json:"type"`
	TimeStamp int64     `json:"timeStamp,omitempty"`
	Channels  []Channel `json:"channels"`
}

type Logger struct {
	Id      int64    `json:"id,omitempty"`
	Devices []Device `json:"devices"`
}

type Rockstar struct {
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       *int   `json:"age,omitempty"`
}

// @DataContract
type AiContent struct {
	/** @description The type of the content part. */
	// @DataMember(Name="type")
	Type string `json:"type"`
}

/** @description The function that the model called. */
// @DataContract
type ToolFunction struct {
	/** @description The name of the function to call. */
	// @DataMember(Name="name")
	Name string `json:"name"`
	/** @description The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function. */
	// @DataMember(Name="arguments")
	Arguments string `json:"arguments"`
}

/** @description The tool calls generated by the model, such as function calls. */
// @DataContract
type ToolCall struct {
	/** @description The ID of the tool call. */
	// @DataMember(Name="id")
	Id string `json:"id"`
	/** @description The type of the tool. Currently, only `function` is supported. */
	// @DataMember(Name="type")
	Type string `json:"type"`
	/** @description The function that the model called. */
	// @DataMember(Name="function")
	Function ToolFunction `json:"function"`
}

/** @description A list of messages comprising the conversation so far. */
// @DataContract
type AiMessage struct {
	/** @description The contents of the message. */
	// @DataMember(Name="content")
	Content []interface{} `json:"content,omitempty"`
	/** @description The role of the author of this message. Valid values are `system`, `user`, `assistant` and `tool`. */
	// @DataMember(Name="role")
	Role string `json:"role"`
	/** @description An optional name for the participant. Provides the model information to differentiate between participants of the same role. */
	// @DataMember(Name="name")
	Name *string `json:"name,omitempty"`
	/** @description The tool calls generated by the model, such as function calls. */
	// @DataMember(Name="tool_calls")
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	/** @description Tool call that this message is responding to. */
	// @DataMember(Name="tool_call_id")
	ToolCallId *string `json:"tool_call_id,omitempty"`
	/** @description The reasoning an assistant message was generated with, normalized per provider when replayed as history. */
	// @DataMember(Name="reasoning")
	Reasoning *string `json:"reasoning,omitempty"`
	/** @description The reasoning an assistant message was generated with, as emitted by Gemini and most OpenAI-compatible providers. */
	// @DataMember(Name="reasoning_content")
	ReasoningContent *string `json:"reasoning_content,omitempty"`
	/** @description Unix timestamp (in milliseconds) the message was generated. */
	// @DataMember(Name="timestamp")
	Timestamp *int64 `json:"timestamp,omitempty"`
	/** @description Images attached to the message. Folded into `content` parts before sending to a provider. */
	// @DataMember(Name="images")
	Images []interface{} `json:"images,omitempty"`
}

/** @description Parameters for audio output. Required when audio output is requested with modalities: [audio] */
// @DataContract
type AiChatAudio struct {
	/** @description Specifies the output audio format. Must be one of wav, mp3, flac, opus, or pcm16. */
	// @DataMember(Name="format")
	Format string `json:"format"`
	/** @description The voice the model uses to respond. Supported voices are alloy, ash, ballad, coral, echo, fable, nova, onyx, sage, and shimmer. */
	// @DataMember(Name="voice")
	Voice string `json:"voice"`
}

type ResponseFormat string

const (
	ResponseFormatText       ResponseFormat = "text"
	ResponseFormatJsonObject                = "json_object"
)

// @DataContract
type AiResponseFormat struct {
	/** @description An object specifying the format that the model must output. Compatible with GPT-4 Turbo and all GPT-3.5 Turbo models newer than gpt-3.5-turbo-1106. */
	// @DataMember(Name="type")
	Type ResponseFormat `json:"type,omitempty"`
}

type ToolType string

const (
	ToolTypeFunction ToolType = "function"
)

// @DataContract
type AiToolFunction struct {
	/** @description The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64. */
	// @DataMember(Name="name")
	Name *string `json:"name,omitempty"`
	/** @description A description of what the function does, used by the model to choose when and how to call the function. */
	// @DataMember(Name="description")
	Description *string `json:"description,omitempty"`
	/** @description The parameters the functions accepts, described as a JSON Schema object. See the guide for examples, and the JSON Schema reference for documentation about the format. */
	// @DataMember(Name="parameters")
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// @DataContract
type Tool struct {
	/** @description The type of the tool. Currently, only function is supported. */
	// @DataMember(Name="type")
	Type ToolType `json:"type,omitempty"`
	/** @description The function definition the model may call. */
	// @DataMember(Name="function")
	Function *AiToolFunction `json:"function,omitempty"`
}

type RoomType string

const (
	RoomTypeSingle RoomType = "Single"
	RoomTypeDouble          = "Double"
	RoomTypeQueen           = "Queen"
	RoomTypeTwin            = "Twin"
	RoomTypeSuite           = "Suite"
)

/** @description Discount Coupons */
type Coupon struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Discount    int       `json:"discount,omitempty"`
	ExpiryDate  time.Time `json:"expiryDate,omitempty"`
}

type Address struct {
	Id          int64   `json:"id,omitempty"`
	AddressText *string `json:"addressText,omitempty"`
}

/** @description Booking Details */
type Booking struct {
	ss.AuditBase
	Id               int        `json:"id,omitempty"`
	Name             string     `json:"name"`
	RoomType         RoomType   `json:"roomType,omitempty"`
	RoomNumber       int        `json:"roomNumber,omitempty"`
	BookingStartDate time.Time  `json:"bookingStartDate,omitempty"`
	BookingEndDate   *time.Time `json:"bookingEndDate,omitempty"`
	Cost             float64    `json:"cost,omitempty"`
	// @References("typeof(MyApp.ServiceModel.Coupon)")
	CouponId  *string `json:"couponId,omitempty"`
	Discount  Coupon  `json:"discount"`
	Notes     *string `json:"notes,omitempty"`
	Cancelled *bool   `json:"cancelled,omitempty"`
	// @References("typeof(MyApp.ServiceModel.Address)")
	PermanentAddressId *int64   `json:"permanentAddressId,omitempty"`
	PermanentAddress   *Address `json:"permanentAddress,omitempty"`
	// @References("typeof(MyApp.ServiceModel.Address)")
	PostalAddressId *int64   `json:"postalAddressId,omitempty"`
	PostalAddress   *Address `json:"postalAddress,omitempty"`
}

type QueryDbTenant[From any, Into any] struct {
	ss.QueryDb
}

type LivingStatus string

const (
	LivingStatusAlive LivingStatus = "Alive"
	LivingStatusDead               = "Dead"
)

type RockstarAuditTenant struct {
	ss.AuditBase
	TenantId     int          `json:"tenantId,omitempty"`
	Id           int          `json:"id,omitempty"`
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Age          *int         `json:"age,omitempty"`
	DateOfBirth  time.Time    `json:"dateOfBirth,omitempty"`
	DateDied     *time.Time   `json:"dateDied,omitempty"`
	LivingStatus LivingStatus `json:"livingStatus,omitempty"`
}

type RockstarBase struct {
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Age          *int         `json:"age,omitempty"`
	DateOfBirth  time.Time    `json:"dateOfBirth,omitempty"`
	DateDied     *time.Time   `json:"dateDied,omitempty"`
	LivingStatus LivingStatus `json:"livingStatus,omitempty"`
}

type RockstarAuto struct {
	RockstarBase
	Id int `json:"id,omitempty"`
}

type OnlyDefinedInGenericType struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type OnlyDefinedInGenericTypeFrom struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type OnlyDefinedInGenericTypeInto struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type RockstarAudit struct {
	RockstarBase
	Id           int       `json:"id,omitempty"`
	CreatedDate  time.Time `json:"createdDate,omitempty"`
	CreatedBy    string    `json:"createdBy"`
	CreatedInfo  string    `json:"createdInfo"`
	ModifiedDate time.Time `json:"modifiedDate,omitempty"`
	ModifiedBy   string    `json:"modifiedBy"`
	ModifiedInfo string    `json:"modifiedInfo"`
}

type CreateAuditBase[Table any, TResponse any] struct {
}

type CreateAuditTenantBase[Table any, TResponse any] struct {
	CreateAuditBase[Table, TResponse]
}

type UpdateAuditBase[Table any, TResponse any] struct {
}

type UpdateAuditTenantBase[Table any, TResponse any] struct {
	UpdateAuditBase[Table, TResponse]
}

type PatchAuditBase[Table any, TResponse any] struct {
}

type PatchAuditTenantBase[Table any, TResponse any] struct {
	PatchAuditBase[Table, TResponse]
}

type SoftDeleteAuditBase[Table any, TResponse any] struct {
}

type SoftDeleteAuditTenantBase[Table any, TResponse any] struct {
	SoftDeleteAuditBase[Table, TResponse]
}

type RockstarVersion struct {
	RockstarBase
	Id         int    `json:"id,omitempty"`
	RowVersion uint64 `json:"rowVersion,omitempty"`
}

// @Route("/messages/crud/{Id}", "PUT")
type MessageCrud struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func (MessageCrud) CreateResponseVoid() {}
func (MessageCrud) HttpMethod() string  { return "PUT" }

type QueryResponseAlt[T any] struct {
	Offset         int               `json:"offset,omitempty"`
	Total          int               `json:"total,omitempty"`
	Results        []T               `json:"results"`
	Meta           map[string]string `json:"meta"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

/** @description Output object for generated text */
type TextOutput struct {
	/** @description The generated text */
	// @ApiMember(Description="The generated text")
	Text *string `json:"text,omitempty"`
}

type UploadInfo struct {
	Name          string `json:"name"`
	FileName      string `json:"fileName"`
	ContentLength int64  `json:"contentLength,omitempty"`
	ContentType   string `json:"contentType"`
}

type MetadataTestNestedChild struct {
	Name string `json:"name"`
}

type MetadataTestChild struct {
	Name    string                    `json:"name"`
	Results []MetadataTestNestedChild `json:"results"`
}

type MenuItemExampleItem struct {
	// @DataMember(Order=1)
	// @ApiMember()
	Name1 string `json:"name1"`
}

type MenuItemExample struct {
	// @DataMember(Order=1)
	// @ApiMember()
	Name1               string              `json:"name1"`
	MenuItemExampleItem MenuItemExampleItem `json:"menuItemExampleItem"`
}

// @DataContract
type MenuExample struct {
	// @DataMember(Order=1)
	// @ApiMember()
	MenuItemExample1 MenuItemExample `json:"menuItemExample1"`
}

type ListResult struct {
	Result string `json:"result"`
}

type ArrayResult struct {
	Result string `json:"result"`
}

type HelloResponseBase struct {
	RefId int `json:"refId,omitempty"`
}

type HelloWithReturnResponse struct {
	Result string `json:"result"`
}

type HelloType struct {
	Result string `json:"result"`
}

type InnerType struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}

type InnerEnum string

const (
	InnerEnumFoo InnerEnum = "Foo"
	InnerEnumBar           = "Bar"
	InnerEnumBaz           = "Baz"
)

type ReturnedDto struct {
	Id int `json:"id,omitempty"`
}

type CustomUserSession struct {
	AuthUserSession
	// @DataMember
	CustomName *string `json:"customName,omitempty"`
	// @DataMember
	CustomInfo *string `json:"customInfo,omitempty"`
}

type UnAuthInfo struct {
	CustomInfo *string `json:"customInfo,omitempty"`
}

/** @description Annotations for the message, when applicable, as when using the web search tool. */
// @DataContract
type UrlCitation struct {
	/** @description The index of the last character of the URL citation in the message. */
	// @DataMember(Name="end_index")
	EndIndex int `json:"end_index,omitempty"`
	/** @description The index of the first character of the URL citation in the message. */
	// @DataMember(Name="start_index")
	StartIndex int `json:"start_index,omitempty"`
	/** @description The title of the web resource. */
	// @DataMember(Name="title")
	Title string `json:"title"`
	/** @description The URL of the web resource. */
	// @DataMember(Name="url")
	Url string `json:"url"`
}

/** @description Annotations for the message, when applicable, as when using the web search tool. */
// @DataContract
type ChoiceAnnotation struct {
	/** @description The type of the URL citation. Always url_citation. */
	// @DataMember(Name="type")
	Type string `json:"type"`
	/** @description A URL citation when using web search. */
	// @DataMember(Name="url_citation")
	UrlCitation UrlCitation `json:"url_citation"`
}

/** @description If the audio output modality is requested, this object contains data about the audio response from the model. */
// @DataContract
type ChoiceAudio struct {
	/** @description Base64 encoded audio bytes generated by the model, in the format specified in the request. */
	// @DataMember(Name="data")
	Data string `json:"data"`
	/** @description The Unix timestamp (in seconds) for when this audio response will no longer be accessible on the server for use in multi-turn conversations. */
	// @DataMember(Name="expires_at")
	ExpiresAt int64 `json:"expires_at,omitempty"`
	/** @description Unique identifier for this audio response. */
	// @DataMember(Name="id")
	Id string `json:"id"`
	/** @description Transcript of the audio generated by the model. */
	// @DataMember(Name="transcript")
	Transcript string `json:"transcript"`
}

// @DataContract
type ChoiceMessage struct {
	/** @description The contents of the message. */
	// @DataMember(Name="content")
	Content string `json:"content"`
	/** @description The refusal message generated by the model. */
	// @DataMember(Name="refusal")
	Refusal *string `json:"refusal,omitempty"`
	/** @description The reasoning process used by the model. */
	// @DataMember(Name="reasoning")
	Reasoning *string `json:"reasoning,omitempty"`
	/** @description The reasoning process used by the model, as emitted by Gemini and most OpenAI-compatible providers. */
	// @DataMember(Name="reasoning_content")
	ReasoningContent *string `json:"reasoning_content,omitempty"`
	/** @description The reasoning process used by the model, as emitted by Anthropic. */
	// @DataMember(Name="thinking")
	Thinking *string `json:"thinking,omitempty"`
	/** @description The role of the author of this message. */
	// @DataMember(Name="role")
	Role string `json:"role"`
	/** @description Unix timestamp (in milliseconds) the message was generated. */
	// @DataMember(Name="timestamp")
	Timestamp *int64 `json:"timestamp,omitempty"`
	/** @description The tool call this message is responding to, set on `tool` role messages in tool_history. */
	// @DataMember(Name="tool_call_id")
	ToolCallId *string `json:"tool_call_id,omitempty"`
	/** @description Images generated by the model or produced by a tool call. */
	// @DataMember(Name="images")
	Images []interface{} `json:"images,omitempty"`
	/** @description Audio generated by the model or produced by a tool call. */
	// @DataMember(Name="audios")
	Audios []interface{} `json:"audios,omitempty"`
	/** @description Files produced by a tool call. */
	// @DataMember(Name="files")
	Files []interface{} `json:"files,omitempty"`
	/** @description Annotations for the message, when applicable, as when using the web search tool. */
	// @DataMember(Name="annotations")
	Annotations []ChoiceAnnotation `json:"annotations,omitempty"`
	/** @description If the audio output modality is requested, this object contains data about the audio response from the model. */
	// @DataMember(Name="audio")
	Audio *ChoiceAudio `json:"audio,omitempty"`
	/** @description The tool calls generated by the model, such as function calls. */
	// @DataMember(Name="tool_calls")
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

/** @description A list of message content tokens with log probability information. */
// @DataContract
type LogprobItem struct {
	/** @description The token. */
	// @DataMember(Name="token")
	Token string `json:"token"`
	/** @description The log probability of this token, if it is within the top 20 most likely tokens. Otherwise, the value `-9999`.0 is used to signify that the token is very unlikely. */
	// @DataMember(Name="logprob")
	Logprob float64 `json:"logprob,omitempty"`
	/** @description A list of integers representing the UTF-8 bytes representation of the token. Useful in instances where characters are represented by multiple tokens and their byte representations must be combined to generate the correct text representation. Can be `null` if there is no bytes representation for the token. */
	// @DataMember(Name="bytes")
	Bytes []byte `json:"bytes"`
	/** @description List of the most likely tokens and their log probability, at this token position. In rare cases, there may be fewer than the number of requested `top_logprobs` returned. */
	// @DataMember(Name="top_logprobs")
	TopLogprobs []LogprobItem `json:"top_logprobs"`
}

/** @description Log probability information for the choice. */
// @DataContract
type Logprobs struct {
	/** @description A list of message content tokens with log probability information. */
	// @DataMember(Name="content")
	Content []LogprobItem `json:"content"`
}

// @DataContract
type Choice struct {
	/** @description The reason the model stopped generating tokens. This will be stop if the model hit a natural stop point or a provided stop sequence, length if the maximum number of tokens specified in the request was reached, content_filter if content was omitted due to a flag from our content filters, tool_calls if the model called a tool */
	// @DataMember(Name="finish_reason")
	FinishReason string `json:"finish_reason"`
	/** @description The index of the choice in the list of choices. */
	// @DataMember(Name="index")
	Index int `json:"index,omitempty"`
	/** @description A chat completion message generated by the model. */
	// @DataMember(Name="message")
	Message ChoiceMessage `json:"message"`
	/** @description Log probability information for the choice. */
	// @DataMember(Name="logprobs")
	Logprobs *Logprobs `json:"logprobs,omitempty"`
}

/** @description Usage statistics for the completion request. */
// @DataContract
type AiCompletionUsage struct {
	/** @description When using Predicted Outputs, the number of tokens in the prediction that appeared in the completion. */
	// @DataMember(Name="accepted_prediction_tokens")
	AcceptedPredictionTokens int64 `json:"accepted_prediction_tokens,omitempty"`
	/** @description Audio input tokens generated by the model. */
	// @DataMember(Name="audio_tokens")
	AudioTokens int64 `json:"audio_tokens,omitempty"`
	/** @description Tokens generated by the model for reasoning. */
	// @DataMember(Name="reasoning_tokens")
	ReasoningTokens int64 `json:"reasoning_tokens,omitempty"`
	/** @description When using Predicted Outputs, the number of tokens in the prediction that did not appear in the completion. */
	// @DataMember(Name="rejected_prediction_tokens")
	RejectedPredictionTokens int64 `json:"rejected_prediction_tokens,omitempty"`
}

/** @description Breakdown of tokens used in the prompt. */
// @DataContract
type AiPromptUsage struct {
	/** @description When using Predicted Outputs, the number of tokens in the prediction that appeared in the completion. */
	// @DataMember(Name="accepted_prediction_tokens")
	AcceptedPredictionTokens int64 `json:"accepted_prediction_tokens,omitempty"`
	/** @description Audio input tokens present in the prompt. */
	// @DataMember(Name="audio_tokens")
	AudioTokens int64 `json:"audio_tokens,omitempty"`
	/** @description Cached tokens present in the prompt. */
	// @DataMember(Name="cached_tokens")
	CachedTokens int64 `json:"cached_tokens,omitempty"`
}

/** @description Usage statistics for the completion request. */
// @DataContract
type AiUsage struct {
	/** @description Number of tokens in the generated completion. */
	// @DataMember(Name="completion_tokens")
	CompletionTokens int64 `json:"completion_tokens,omitempty"`
	/** @description Number of tokens in the prompt. */
	// @DataMember(Name="prompt_tokens")
	PromptTokens int64 `json:"prompt_tokens,omitempty"`
	/** @description Total number of tokens used in the request (prompt + completion). */
	// @DataMember(Name="total_tokens")
	TotalTokens int64 `json:"total_tokens,omitempty"`
	/** @description Breakdown of tokens used in a completion. */
	// @DataMember(Name="completion_tokens_details")
	CompletionTokensDetails *AiCompletionUsage `json:"completion_tokens_details,omitempty"`
	/** @description Breakdown of tokens used in the prompt. */
	// @DataMember(Name="prompt_tokens_details")
	PromptTokensDetails *AiPromptUsage `json:"prompt_tokens_details,omitempty"`
	/** @description Seconds spent servicing the completion, including every request in the tool loop. */
	// @DataMember(Name="duration")
	Duration *int64 `json:"duration,omitempty"`
}

type TypesGroup struct {
}

/** @description Text content part */
// @DataContract
type AiTextContent struct {
	AiContent
	/** @description The text content. */
	// @DataMember(Name="text")
	Text string `json:"text"`
}

// @DataContract
type AiImageUrl struct {
	/** @description Either a URL of the image or the base64 encoded image data. */
	// @DataMember(Name="url")
	Url string `json:"url"`
}

/** @description Image content part */
// @DataContract
type AiImageContent struct {
	AiContent
	/** @description The image for this content. */
	// @DataMember(Name="image_url")
	ImageUrl AiImageUrl `json:"image_url"`
}

/** @description Audio content part */
// @DataContract
type AiInputAudio struct {
	/** @description URL or Base64 encoded audio data. */
	// @DataMember(Name="data")
	Data string `json:"data"`
	/** @description The format of the encoded audio data. Currently supports 'wav' and 'mp3'. */
	// @DataMember(Name="format")
	Format string `json:"format"`
}

/** @description Audio content part */
// @DataContract
type AiAudioContent struct {
	AiContent
	/** @description The audio input for this content. */
	// @DataMember(Name="input_audio")
	InputAudio AiInputAudio `json:"input_audio"`
}

/** @description File content part */
// @DataContract
type AiFile struct {
	/** @description The URL or base64 encoded file data, used when passing the file to the model as a string. */
	// @DataMember(Name="file_data")
	FileData string `json:"file_data"`
	/** @description The name of the file, used when passing the file to the model as a string. */
	// @DataMember(Name="filename")
	Filename string `json:"filename"`
	/** @description The ID of an uploaded file to use as input. */
	// @DataMember(Name="file_id")
	FileId *string `json:"file_id,omitempty"`
}

/** @description File content part */
// @DataContract
type AiFileContent struct {
	AiContent
	/** @description The file input for this content. */
	// @DataMember(Name="file")
	File AiFile `json:"file"`
}

// @DataContract
type AiAudioUrl struct {
	/** @description Either a URL of the audio or the base64 encoded audio data. */
	// @DataMember(Name="url")
	Url string `json:"url"`
}

/** @description Generated audio content part, referenced by URL (emitted by tool calls and audio models) */
// @DataContract
type AiAudioUrlContent struct {
	AiContent
	/** @description The audio for this content. */
	// @DataMember(Name="audio_url")
	AudioUrl AiAudioUrl `json:"audio_url"`
}

type ChatMessage struct {
	Id          int64   `json:"id,omitempty"`
	Channel     *string `json:"channel,omitempty"`
	FromUserId  *string `json:"fromUserId,omitempty"`
	FromName    *string `json:"fromName,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Message     *string `json:"message,omitempty"`
	UserAuthId  *string `json:"userAuthId,omitempty"`
	Private     bool    `json:"private,omitempty"`
}

type GetChatHistoryResponse struct {
	Results        []ChatMessage      `json:"results,omitempty"`
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type GetUserDetailsResponse struct {
	Provider     *string    `json:"provider,omitempty"`
	UserId       *string    `json:"userId,omitempty"`
	UserName     *string    `json:"userName,omitempty"`
	FullName     *string    `json:"fullName,omitempty"`
	DisplayName  *string    `json:"displayName,omitempty"`
	FirstName    *string    `json:"firstName,omitempty"`
	LastName     *string    `json:"lastName,omitempty"`
	Company      *string    `json:"company,omitempty"`
	Email        *string    `json:"email,omitempty"`
	PhoneNumber  *string    `json:"phoneNumber,omitempty"`
	BirthDate    *time.Time `json:"birthDate,omitempty"`
	BirthDateRaw *string    `json:"birthDateRaw,omitempty"`
	Address      *string    `json:"address,omitempty"`
	Address2     *string    `json:"address2,omitempty"`
	City         *string    `json:"city,omitempty"`
	State        *string    `json:"state,omitempty"`
	Country      *string    `json:"country,omitempty"`
	Culture      *string    `json:"culture,omitempty"`
	Gender       *string    `json:"gender,omitempty"`
	Language     *string    `json:"language,omitempty"`
	MailAddress  *string    `json:"mailAddress,omitempty"`
	Nickname     *string    `json:"nickname,omitempty"`
	PostalCode   *string    `json:"postalCode,omitempty"`
	TimeZone     *string    `json:"timeZone,omitempty"`
}

type CustomHttpErrorResponse struct {
	Custom         string            `json:"custom"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type Items struct {
	Results []Item `json:"results"`
}

type ReturnCustom400Response struct {
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type ThrowTypeResponse struct {
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type ThrowValidationResponse struct {
	Age            int               `json:"age,omitempty"`
	Required       string            `json:"required"`
	Email          string            `json:"email"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type ThrowBusinessErrorResponse struct {
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

/** @description Response object for text generation requests */
// @Api(Description="Response object for text generation requests")
type TextGenerationResponse struct {
	/** @description List of generated text outputs */
	// @ApiMember(Description="List of generated text outputs")
	Results []TextOutput `json:"results,omitempty"`
	/** @description Detailed response status information */
	// @ApiMember(Description="Detailed response status information")
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type TestFileUploadsResponse struct {
	Id             *int               `json:"id,omitempty"`
	RefId          *string            `json:"refId,omitempty"`
	Files          []UploadInfo       `json:"files"`
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type TestUploadWithDto struct {
	Int                   int                          `json:"int,omitempty"`
	NullableId            *int                         `json:"nullableId,omitempty"`
	Long                  int64                        `json:"long,omitempty"`
	Double                float64                      `json:"double,omitempty"`
	String                string                       `json:"string"`
	DateTime              time.Time                    `json:"dateTime,omitempty"`
	IntArray              []int                        `json:"intArray,omitempty"`
	IntList               []int                        `json:"intList,omitempty"`
	StringArray           []string                     `json:"stringArray,omitempty"`
	StringList            []string                     `json:"stringList,omitempty"`
	PocoArray             []Poco                       `json:"pocoArray,omitempty"`
	PocoList              []Poco                       `json:"pocoList,omitempty"`
	NullableByteArray     []*byte                      `json:"nullableByteArray,omitempty"`
	NullableByteList      []byte                       `json:"nullableByteList,omitempty"`
	NullableDateTimeArray []*time.Time                 `json:"nullableDateTimeArray,omitempty"`
	NullableDateTimeList  []time.Time                  `json:"nullableDateTimeList,omitempty"`
	PocoLookup            map[string][]Poco            `json:"pocoLookup,omitempty"`
	PocoLookupMap         map[string][]map[string]Poco `json:"pocoLookupMap,omitempty"`
	MapList               map[string][]string          `json:"mapList,omitempty"`
}

func (TestUploadWithDto) HttpMethod() string { return "POST" }

type Account struct {
	Name *string `json:"name,omitempty"`
}

type Project struct {
	Account *string `json:"account,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type SecuredResponse struct {
	Result         *string            `json:"result,omitempty"`
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type CreateJwtResponse struct {
	Token          *string            `json:"token,omitempty"`
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type CreateRefreshJwtResponse struct {
	Token          *string            `json:"token,omitempty"`
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type MetadataTestResponse struct {
	Id      int                 `json:"id,omitempty"`
	Results []MetadataTestChild `json:"results"`
}

// @DataContract
type GetExampleResponse struct {
	// @DataMember(Order=1)
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
	// @DataMember(Order=2)
	// @ApiMember()
	MenuExample1 MenuExample `json:"menuExample1"`
}

// @Route("/messages/{Id}", "PUT")
type Message struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func (Message) HttpMethod() string { return "PUT" }

type GetRandomIdsResponse struct {
	Results []string `json:"results"`
}

type HelloResponse struct {
	Result string `json:"result"`
}

type AllTypes struct {
	Id               int                          `json:"id,omitempty"`
	NullableId       *int                         `json:"nullableId,omitempty"`
	Byte             byte                         `json:"byte,omitempty"`
	Short            int16                        `json:"short,omitempty"`
	Int              int                          `json:"int,omitempty"`
	Long             int64                        `json:"long,omitempty"`
	UShort           uint16                       `json:"uShort,omitempty"`
	UInt             uint32                       `json:"uInt,omitempty"`
	ULong            uint64                       `json:"uLong,omitempty"`
	Float            float32                      `json:"float,omitempty"`
	Double           float64                      `json:"double,omitempty"`
	Decimal          float64                      `json:"decimal,omitempty"`
	String           string                       `json:"string"`
	DateTime         time.Time                    `json:"dateTime,omitempty"`
	TimeSpan         time.Duration                `json:"timeSpan,omitempty"`
	DateTimeOffset   time.Time                    `json:"dateTimeOffset,omitempty"`
	Guid             string                       `json:"guid,omitempty"`
	Char             string                       `json:"char,omitempty"`
	KeyValuePair     KeyValuePair[string, string] `json:"keyValuePair,omitempty"`
	NullableDateTime *time.Time                   `json:"nullableDateTime,omitempty"`
	NullableTimeSpan *time.Duration               `json:"nullableTimeSpan,omitempty"`
	StringList       []string                     `json:"stringList"`
	StringArray      []string                     `json:"stringArray"`
	StringMap        map[string]string            `json:"stringMap"`
	IntStringMap     map[int]string               `json:"intStringMap"`
	SubType          SubType                      `json:"subType"`
}

func (AllTypes) HttpMethod() string { return "POST" }

type AllCollectionTypes struct {
	IntArray      []int                        `json:"intArray"`
	IntList       []int                        `json:"intList"`
	StringArray   []string                     `json:"stringArray"`
	StringList    []string                     `json:"stringList"`
	FloatArray    []float32                    `json:"floatArray"`
	DoubleList    []float64                    `json:"doubleList"`
	ByteArray     []byte                       `json:"byteArray"`
	CharArray     []string                     `json:"charArray"`
	DecimalList   []float64                    `json:"decimalList"`
	PocoArray     []Poco                       `json:"pocoArray"`
	PocoList      []Poco                       `json:"pocoList"`
	PocoLookup    map[string][]Poco            `json:"pocoLookup"`
	PocoLookupMap map[string][]map[string]Poco `json:"pocoLookupMap"`
}

func (AllCollectionTypes) HttpMethod() string { return "POST" }

type HelloAllTypesResponse struct {
	Result             string             `json:"result"`
	AllTypes           AllTypes           `json:"allTypes"`
	AllCollectionTypes AllCollectionTypes `json:"allCollectionTypes"`
}

type SubAllTypes struct {
	AllTypesBase
	Hierarchy int `json:"hierarchy,omitempty"`
}

type HelloDateTime struct {
	DateTime time.Time `json:"dateTime,omitempty"`
}

func (HelloDateTime) HttpMethod() string { return "POST" }

// @DataContract
type HelloWithDataContractResponse struct {
	// @DataMember(Name="result", Order=1, IsRequired=true, EmitDefaultValue=false)
	Result string `json:"result"`
}

/** @description Description on HelloWithDescriptionResponse type */
type HelloWithDescriptionResponse struct {
	Result string `json:"result"`
}

type HelloWithInheritanceResponse struct {
	HelloResponseBase
	Result string `json:"result"`
}

type HelloWithAlternateReturnResponse struct {
	HelloWithReturnResponse
	AltResult string `json:"altResult"`
}

type HelloWithRouteResponse struct {
	Result string `json:"result"`
}

type HelloWithTypeResponse struct {
	Result HelloType `json:"result"`
}

type HelloInnerTypesResponse struct {
	InnerType InnerType `json:"innerType"`
	InnerEnum InnerEnum `json:"innerEnum,omitempty"`
}

type HelloVerbResponse struct {
	Result string `json:"result"`
}

type EnumResponse struct {
	Operator ScopeType `json:"operator,omitempty"`
}

// @Route("/hellotypes/{Name}")
type HelloTypes struct {
	String string `json:"string"`
	Bool   bool   `json:"bool,omitempty"`
	Int    int    `json:"int,omitempty"`
}

func (HelloTypes) HttpMethod() string { return "POST" }

// @DataContract
type HelloZipResponse struct {
	// @DataMember
	Result string `json:"result"`
}

type PingResponse struct {
	Responses      map[string]ss.ResponseStatus `json:"responses,omitempty"`
	ResponseStatus *ss.ResponseStatus           `json:"responseStatus,omitempty"`
}

type RequiresRoleResponse struct {
	Result         string            `json:"result"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type SendVerbResponse struct {
	Id            int     `json:"id,omitempty"`
	PathInfo      *string `json:"pathInfo,omitempty"`
	RequestMethod *string `json:"requestMethod,omitempty"`
}

type GetSessionResponse struct {
	Result         *CustomUserSession `json:"result,omitempty"`
	UnAuthInfo     *UnAuthInfo        `json:"unAuthInfo,omitempty"`
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

// @DataContract(Namespace="http://schemas.servicestack.net/types")
type GetStuffResponse struct {
	// @DataMember
	SummaryDate *time.Time `json:"summaryDate,omitempty"`
	// @DataMember
	SummaryEndDate *time.Time `json:"summaryEndDate,omitempty"`
	// @DataMember
	Symbol *string `json:"symbol,omitempty"`
	// @DataMember
	Email *string `json:"email,omitempty"`
	// @DataMember
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

type StoreLogsResponse struct {
	ExistingLogs   []Logger          `json:"existingLogs"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type TestAuthResponse struct {
	UserId         *string            `json:"userId,omitempty"`
	SessionId      *string            `json:"sessionId,omitempty"`
	UserName       *string            `json:"userName,omitempty"`
	DisplayName    *string            `json:"displayName,omitempty"`
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type RequiresAdmin struct {
	Id int `json:"id,omitempty"`
}

func (RequiresAdmin) HttpMethod() string { return "POST" }

// @Route("/custom")
// @Route("/custom/{Data}")
type CustomRoute struct {
	Data *string `json:"data,omitempty"`
}

func (CustomRoute) HttpMethod() string { return "POST" }

// @Route("/wait/{ForMs}")
type Wait struct {
	ForMs int `json:"forMs,omitempty"`
}

func (Wait) HttpMethod() string { return "POST" }

// @Route("/echo/types")
type EchoTypes struct {
	Byte           byte          `json:"byte,omitempty"`
	Short          int16         `json:"short,omitempty"`
	Int            int           `json:"int,omitempty"`
	Long           int64         `json:"long,omitempty"`
	UShort         uint16        `json:"uShort,omitempty"`
	UInt           uint32        `json:"uInt,omitempty"`
	ULong          uint64        `json:"uLong,omitempty"`
	Float          float32       `json:"float,omitempty"`
	Double         float64       `json:"double,omitempty"`
	Decimal        float64       `json:"decimal,omitempty"`
	String         *string       `json:"string,omitempty"`
	DateTime       time.Time     `json:"dateTime,omitempty"`
	TimeSpan       time.Duration `json:"timeSpan,omitempty"`
	DateTimeOffset time.Time     `json:"dateTimeOffset,omitempty"`
	Guid           string        `json:"guid,omitempty"`
	Char           string        `json:"char,omitempty"`
}

func (EchoTypes) HttpMethod() string { return "POST" }

// @Route("/echo/collections")
type EchoCollections struct {
	StringList   []string          `json:"stringList,omitempty"`
	StringArray  []string          `json:"stringArray,omitempty"`
	StringMap    map[string]string `json:"stringMap,omitempty"`
	IntStringMap map[int]string    `json:"intStringMap,omitempty"`
}

func (EchoCollections) HttpMethod() string { return "POST" }

// @Route("/echo/complex")
type EchoComplexTypes struct {
	SubType      *SubType           `json:"subType,omitempty"`
	SubTypes     []SubType          `json:"subTypes,omitempty"`
	SubTypeMap   map[string]SubType `json:"subTypeMap,omitempty"`
	StringMap    map[string]string  `json:"stringMap,omitempty"`
	IntStringMap map[int]string     `json:"intStringMap,omitempty"`
}

func (EchoComplexTypes) HttpMethod() string { return "POST" }

// @Route("/rockstars", "POST")
type StoreRockstars []Rockstar

func (StoreRockstars) HttpMethod() string { return "POST" }

// @DataContract
type ChatResponse struct {
	/** @description A unique identifier for the chat completion. */
	// @DataMember(Name="id")
	Id string `json:"id"`
	/** @description A list of chat completion choices. Can be more than one if n is greater than 1. */
	// @DataMember(Name="choices")
	Choices []Choice `json:"choices"`
	/** @description The Unix timestamp (in seconds) of when the chat completion was created. */
	// @DataMember(Name="created")
	Created int64 `json:"created,omitempty"`
	/** @description The model used for the chat completion. */
	// @DataMember(Name="model")
	Model string `json:"model"`
	/** @description This fingerprint represents the backend configuration that the model runs with. */
	// @DataMember(Name="system_fingerprint")
	SystemFingerprint *string `json:"system_fingerprint,omitempty"`
	/** @description The object type, which is always chat.completion. */
	// @DataMember(Name="object")
	Object string `json:"object"`
	/** @description Specifies the processing type used for serving the request. */
	// @DataMember(Name="service_tier")
	ServiceTier *string `json:"service_tier,omitempty"`
	/** @description Usage statistics for the completion request. */
	// @DataMember(Name="usage")
	Usage AiUsage `json:"usage"`
	/** @description The provider used for the chat completion. */
	// @DataMember(Name="provider")
	Provider *string `json:"provider,omitempty"`
	/** @description Total cost of the completion in USD, accumulated across every request in the tool loop. */
	// @DataMember(Name="cost")
	Cost *float64 `json:"cost,omitempty"`
	/** @description The assistant and tool messages exchanged during the tool-execution loop, in order. */
	// @DataMember(Name="tool_history")
	ToolHistory []ChoiceMessage `json:"tool_history,omitempty"`
	/** @description Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. */
	// @DataMember(Name="metadata")
	Metadata map[string]string `json:"metadata,omitempty"`
	// @DataMember(Name="responseStatus")
	ResponseStatus *ss.ResponseStatus `json:"responseStatus,omitempty"`
}

type RockstarWithIdResponse struct {
	Id             int               `json:"id,omitempty"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type RockstarWithIdAndResultResponse struct {
	Id             int               `json:"id,omitempty"`
	Result         RockstarAuto      `json:"result"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type RockstarWithIdAndCountResponse struct {
	Id             int               `json:"id,omitempty"`
	Count          int               `json:"count,omitempty"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type RockstarWithIdAndRowVersionResponse struct {
	Id             int               `json:"id,omitempty"`
	RowVersion     uint32            `json:"rowVersion,omitempty"`
	ResponseStatus ss.ResponseStatus `json:"responseStatus"`
}

type QueryItems struct {
	ss.QueryDb
}

func (QueryItems) CreateResponse() (r ss.QueryResponse[Poco]) { return }
func (QueryItems) HttpMethod() string                         { return "GET" }

// @Route("/channels/{Channel}/raw")
type PostRawToChannel struct {
	From     *string `json:"from,omitempty"`
	ToUserId *string `json:"toUserId,omitempty"`
	Channel  *string `json:"channel,omitempty"`
	Message  *string `json:"message,omitempty"`
	Selector *string `json:"selector,omitempty"`
}

func (PostRawToChannel) CreateResponseVoid() {}
func (PostRawToChannel) HttpMethod() string  { return "POST" }

// @Route("/channels/{Channel}/chat")
type PostChatToChannel struct {
	From     *string `json:"from,omitempty"`
	ToUserId *string `json:"toUserId,omitempty"`
	Channel  *string `json:"channel,omitempty"`
	Message  *string `json:"message,omitempty"`
	Selector *string `json:"selector,omitempty"`
}

func (PostChatToChannel) CreateResponse() (r ChatMessage) { return }
func (PostChatToChannel) HttpMethod() string              { return "POST" }

// @Route("/chathistory")
type GetChatHistory struct {
	Channels []string `json:"channels,omitempty"`
	AfterId  *int64   `json:"afterId,omitempty"`
	Take     *int     `json:"take,omitempty"`
}

func (GetChatHistory) CreateResponse() (r GetChatHistoryResponse) { return }
func (GetChatHistory) HttpMethod() string                         { return "POST" }

// @Route("/reset")
type ClearChatHistory struct {
}

func (ClearChatHistory) CreateResponseVoid() {}
func (ClearChatHistory) HttpMethod() string  { return "POST" }

// @Route("/reset-serverevents")
type ResetServerEvents struct {
}

func (ResetServerEvents) CreateResponseVoid() {}
func (ResetServerEvents) HttpMethod() string  { return "POST" }

// @Route("/channels/{Channel}/object")
type PostObjectToChannel struct {
	ToUserId   *string     `json:"toUserId,omitempty"`
	Channel    *string     `json:"channel,omitempty"`
	Selector   *string     `json:"selector,omitempty"`
	CustomType *CustomType `json:"customType,omitempty"`
	SetterType *SetterType `json:"setterType,omitempty"`
}

func (PostObjectToChannel) CreateResponseVoid() {}
func (PostObjectToChannel) HttpMethod() string  { return "POST" }

// @Route("/account")
type GetUserDetails struct {
}

func (GetUserDetails) CreateResponse() (r GetUserDetailsResponse) { return }
func (GetUserDetails) HttpMethod() string                         { return "GET" }

type CustomHttpError struct {
	StatusCode        int    `json:"statusCode,omitempty"`
	StatusDescription string `json:"statusDescription"`
}

func (CustomHttpError) CreateResponse() (r CustomHttpErrorResponse) { return }
func (CustomHttpError) HttpMethod() string                          { return "POST" }

type AltQueryItems struct {
	Name string `json:"name"`
}

func (AltQueryItems) CreateResponse() (r QueryResponseAlt[Item]) { return }
func (AltQueryItems) HttpMethod() string                         { return "POST" }

type GetItems struct {
}

func (GetItems) CreateResponse() (r Items) { return }
func (GetItems) HttpMethod() string        { return "GET" }

type GetNakedItems struct {
}

func (GetNakedItems) CreateResponse() (r []Item) { return }
func (GetNakedItems) HttpMethod() string         { return "GET" }

// @ValidateRequest(Validator="IsAuthenticated")
type DeclarativeValidationAuth struct {
	Name string `json:"name"`
}

func (DeclarativeValidationAuth) CreateResponseVoid() {}
func (DeclarativeValidationAuth) HttpMethod() string  { return "POST" }

type DeclarativeCollectiveValidationTest struct {
	// @Validate(Validator="NotEmpty")
	// @Validate(Validator="MaximumLength(20)")
	Site                   string                       `json:"site"`
	DeclarativeValidations []DeclarativeChildValidation `json:"declarativeValidations"`
	FluentValidations      []FluentChildValidation      `json:"fluentValidations"`
}

func (DeclarativeCollectiveValidationTest) CreateResponse() (r ss.EmptyResponse) { return }
func (DeclarativeCollectiveValidationTest) HttpMethod() string                   { return "POST" }

type DeclarativeSingleValidationTest struct {
	// @Validate(Validator="NotEmpty")
	// @Validate(Validator="MaximumLength(20)")
	Site                        string                      `json:"site"`
	DeclarativeSingleValidation DeclarativeSingleValidation `json:"declarativeSingleValidation"`
	FluentSingleValidation      FluentSingleValidation      `json:"fluentSingleValidation"`
}

func (DeclarativeSingleValidationTest) CreateResponse() (r ss.EmptyResponse) { return }
func (DeclarativeSingleValidationTest) HttpMethod() string                   { return "POST" }

type DummyTypes struct {
	HelloResponses                []HelloResponse                   `json:"helloResponses,omitempty"`
	ListResult                    []ListResult                      `json:"listResult,omitempty"`
	ArrayResult                   []ArrayResult                     `json:"arrayResult,omitempty"`
	CancelRequest                 *CancelRequest                    `json:"cancelRequest,omitempty"`
	CancelRequestResponse         *CancelRequestResponse            `json:"cancelRequestResponse,omitempty"`
	UpdateEventSubscriber         *UpdateEventSubscriber            `json:"updateEventSubscriber,omitempty"`
	UpdateEventSubscriberResponse *UpdateEventSubscriberResponse    `json:"updateEventSubscriberResponse,omitempty"`
	GetApiKeys                    *ss.GetApiKeys                    `json:"getApiKeys,omitempty"`
	GetApiKeysResponse            *ss.GetApiKeysResponse            `json:"getApiKeysResponse,omitempty"`
	RegenerateApiKeys             *ss.RegenerateApiKeys             `json:"regenerateApiKeys,omitempty"`
	RegenerateApiKeysResponse     *ss.RegenerateApiKeysResponse     `json:"regenerateApiKeysResponse,omitempty"`
	UserApiKey                    *ss.UserApiKey                    `json:"userApiKey,omitempty"`
	ConvertSessionToToken         *ss.ConvertSessionToToken         `json:"convertSessionToToken,omitempty"`
	ConvertSessionToTokenResponse *ss.ConvertSessionToTokenResponse `json:"convertSessionToTokenResponse,omitempty"`
	GetAccessToken                *ss.GetAccessToken                `json:"getAccessToken,omitempty"`
	GetAccessTokenResponse        *ss.GetAccessTokenResponse        `json:"getAccessTokenResponse,omitempty"`
	NavItem                       *ss.NavItem                       `json:"navItem,omitempty"`
	GetNavItems                   *ss.GetNavItems                   `json:"getNavItems,omitempty"`
	GetNavItemsResponse           *ss.GetNavItemsResponse           `json:"getNavItemsResponse,omitempty"`
	EmptyResponse                 *ss.EmptyResponse                 `json:"emptyResponse,omitempty"`
	IdResponse                    *ss.IdResponse                    `json:"idResponse,omitempty"`
	StringResponse                *ss.StringResponse                `json:"stringResponse,omitempty"`
	StringsResponse               *ss.StringsResponse               `json:"stringsResponse,omitempty"`
	AuditBase                     *ss.AuditBase                     `json:"auditBase,omitempty"`
}

func (DummyTypes) CreateResponseVoid() {}
func (DummyTypes) HttpMethod() string  { return "POST" }

// @Route("/throwhttperror/{Status}")
type ThrowHttpError struct {
	Status  *int   `json:"status,omitempty"`
	Message string `json:"message"`
}

func (ThrowHttpError) CreateResponseVoid() {}
func (ThrowHttpError) HttpMethod() string  { return "POST" }

// @Route("/throw404")
// @Route("/throw404/{Message}")
type Throw404 struct {
	Message string `json:"message"`
}

func (Throw404) CreateResponseVoid() {}
func (Throw404) HttpMethod() string  { return "POST" }

// @Route("/throwcustom400")
// @Route("/throwcustom400/{Message}")
type ThrowCustom400 struct {
	Message string `json:"message"`
}

func (ThrowCustom400) CreateResponseVoid() {}
func (ThrowCustom400) HttpMethod() string  { return "POST" }

// @Route("/returncustom400")
// @Route("/returncustom400/{Message}")
type ReturnCustom400 struct {
	Message string `json:"message"`
}

func (ReturnCustom400) CreateResponse() (r ReturnCustom400Response) { return }
func (ReturnCustom400) HttpMethod() string                          { return "POST" }

// @Route("/throw/{Type}")
type ThrowType struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (ThrowType) CreateResponse() (r ThrowTypeResponse) { return }
func (ThrowType) HttpMethod() string                    { return "POST" }

// @Route("/throwvalidation")
type ThrowValidation struct {
	Age      int    `json:"age,omitempty"`
	Required string `json:"required"`
	Email    string `json:"email"`
}

func (ThrowValidation) CreateResponse() (r ThrowValidationResponse) { return }
func (ThrowValidation) HttpMethod() string                          { return "POST" }

// @Route("/throwbusinesserror")
type ThrowBusinessError struct {
}

func (ThrowBusinessError) CreateResponse() (r ThrowBusinessErrorResponse) { return }
func (ThrowBusinessError) HttpMethod() string                             { return "POST" }

/** @description Convert speech to text */
// @Api(Description="Convert speech to text")
type SpeechToText struct {
	/** @description The audio stream containing the speech to be transcribed */
	// @ApiMember(Description="The audio stream containing the speech to be transcribed")
	// @Required()
	Audio string `json:"audio"`
	/** @description Optional client-provided identifier for the request */
	// @ApiMember(Description="Optional client-provided identifier for the request")
	RefId *string `json:"refId,omitempty"`
	/** @description Tag to identify the request */
	// @ApiMember(Description="Tag to identify the request")
	Tag *string `json:"tag,omitempty"`
}

func (SpeechToText) CreateResponse() (r TextGenerationResponse) { return }
func (SpeechToText) HttpMethod() string                         { return "POST" }

type TestFileUploads struct {
	Id    *int    `json:"id,omitempty"`
	RefId *string `json:"refId,omitempty"`
}

func (TestFileUploads) CreateResponse() (r TestFileUploadsResponse) { return }
func (TestFileUploads) HttpMethod() string                          { return "POST" }

type RootPathRoutes struct {
	Path *string `json:"path,omitempty"`
}

func (RootPathRoutes) CreateResponseVoid() {}
func (RootPathRoutes) HttpMethod() string  { return "POST" }

type GetAccount struct {
	Account *string `json:"account,omitempty"`
}

func (GetAccount) CreateResponse() (r Account) { return }
func (GetAccount) HttpMethod() string          { return "POST" }

type GetProject struct {
	Account *string `json:"account,omitempty"`
	Project *string `json:"project,omitempty"`
}

func (GetProject) CreateResponse() (r Project) { return }
func (GetProject) HttpMethod() string          { return "POST" }

// @Route("/image-stream")
type ImageAsStream struct {
	Format *string `json:"format,omitempty"`
}

func (ImageAsStream) CreateResponse() (r []byte) { return }
func (ImageAsStream) HttpMethod() string         { return "POST" }

// @Route("/image-bytes")
type ImageAsBytes struct {
	Format *string `json:"format,omitempty"`
}

func (ImageAsBytes) CreateResponse() (r []byte) { return }
func (ImageAsBytes) HttpMethod() string         { return "POST" }

// @Route("/image-custom")
type ImageAsCustomResult struct {
	Format *string `json:"format,omitempty"`
}

func (ImageAsCustomResult) CreateResponse() (r []byte) { return }
func (ImageAsCustomResult) HttpMethod() string         { return "POST" }

// @Route("/image-response")
type ImageWriteToResponse struct {
	Format *string `json:"format,omitempty"`
}

func (ImageWriteToResponse) CreateResponse() (r []byte) { return }
func (ImageWriteToResponse) HttpMethod() string         { return "POST" }

// @Route("/image-file")
type ImageAsFile struct {
	Format *string `json:"format,omitempty"`
}

func (ImageAsFile) CreateResponse() (r []byte) { return }
func (ImageAsFile) HttpMethod() string         { return "POST" }

// @Route("/image-redirect")
type ImageAsRedirect struct {
	Format *string `json:"format,omitempty"`
}

func (ImageAsRedirect) CreateResponseVoid() {}
func (ImageAsRedirect) HttpMethod() string  { return "POST" }

// @Route("/hello-image/{Name}")
type HelloImage struct {
	Name       *string `json:"name,omitempty"`
	Format     *string `json:"format,omitempty"`
	Width      *int    `json:"width,omitempty"`
	Height     *int    `json:"height,omitempty"`
	FontSize   *int    `json:"fontSize,omitempty"`
	FontFamily *string `json:"fontFamily,omitempty"`
	Foreground *string `json:"foreground,omitempty"`
	Background *string `json:"background,omitempty"`
}

func (HelloImage) CreateResponse() (r []byte) { return }
func (HelloImage) HttpMethod() string         { return "GET" }

// @Route("/secured")
// @ValidateRequest(Validator="IsAuthenticated")
type Secured struct {
	Name *string `json:"name,omitempty"`
}

func (Secured) CreateResponse() (r SecuredResponse) { return }
func (Secured) HttpMethod() string                  { return "POST" }

// @Route("/jwt")
type CreateJwt struct {
	AuthUserSession
	JwtExpiry *time.Time `json:"jwtExpiry,omitempty"`
}

func (CreateJwt) CreateResponse() (r CreateJwtResponse) { return }
func (CreateJwt) HttpMethod() string                    { return "POST" }

// @Route("/jwt-refresh")
type CreateRefreshJwt struct {
	UserAuthId *string    `json:"userAuthId,omitempty"`
	JwtExpiry  *time.Time `json:"jwtExpiry,omitempty"`
}

func (CreateRefreshJwt) CreateResponse() (r CreateRefreshJwtResponse) { return }
func (CreateRefreshJwt) HttpMethod() string                           { return "POST" }

// @Route("/jwt-invalidate")
type InvalidateLastAccessToken struct {
}

func (InvalidateLastAccessToken) CreateResponse() (r ss.EmptyResponse) { return }
func (InvalidateLastAccessToken) HttpMethod() string                   { return "POST" }

// @Route("/logs")
type ViewLogs struct {
	Clear bool `json:"clear,omitempty"`
}

func (ViewLogs) CreateResponse() (r string) { return }
func (ViewLogs) HttpMethod() string         { return "POST" }

// @Route("/metadatatest")
type MetadataTest struct {
	Id int `json:"id,omitempty"`
}

func (MetadataTest) CreateResponse() (r MetadataTestResponse) { return }
func (MetadataTest) HttpMethod() string                       { return "POST" }

// @Route("/metadatatest-array")
type MetadataTestArray struct {
	Id int `json:"id,omitempty"`
}

func (MetadataTestArray) CreateResponse() (r []MetadataTestChild) { return }
func (MetadataTestArray) HttpMethod() string                      { return "POST" }

// @Route("/example", "GET")
// @DataContract
type GetExample struct {
}

func (GetExample) CreateResponse() (r GetExampleResponse) { return }
func (GetExample) HttpMethod() string                     { return "GET" }

// @Route("/messages/{Id}", "GET")
type RequestMessage struct {
	Id int `json:"id,omitempty"`
}

func (RequestMessage) CreateResponse() (r Message) { return }
func (RequestMessage) HttpMethod() string          { return "GET" }

// @Route("/randomids")
type GetRandomIds struct {
	Take *int `json:"take,omitempty"`
}

func (GetRandomIds) CreateResponse() (r GetRandomIdsResponse) { return }
func (GetRandomIds) HttpMethod() string                       { return "POST" }

// @Route("/textfile-test")
type TextFileTest struct {
	AsAttachment bool `json:"asAttachment,omitempty"`
}

func (TextFileTest) CreateResponseVoid() {}
func (TextFileTest) HttpMethod() string  { return "POST" }

// @Route("/return/text")
type ReturnText struct {
	Text *string `json:"text,omitempty"`
}

func (ReturnText) CreateResponseVoid() {}
func (ReturnText) HttpMethod() string  { return "POST" }

// @Route("/return/html")
type ReturnHtml struct {
	Text *string `json:"text,omitempty"`
}

func (ReturnHtml) CreateResponseVoid() {}
func (ReturnHtml) HttpMethod() string  { return "POST" }

// @Route("/hello")
// @Route("/hello/{Name}")
type Hello struct {
	// @Required()
	Name  string `json:"name"`
	Title string `json:"title"`
}

func (Hello) CreateResponse() (r HelloResponse) { return }
func (Hello) HttpMethod() string                { return "POST" }

// @Route("/hello-secure/{Name}")
// @ValidateRequest(Validator="IsAuthenticated")
type HelloSecure struct {
	Name string `json:"name"`
}

func (HelloSecure) CreateResponse() (r HelloResponse) { return }
func (HelloSecure) HttpMethod() string                { return "POST" }

type HelloWithNestedClass struct {
	Name            string      `json:"name"`
	NestedClassProp NestedClass `json:"nestedClassProp"`
}

func (HelloWithNestedClass) CreateResponse() (r HelloResponse) { return }
func (HelloWithNestedClass) HttpMethod() string                { return "GET" }

type HelloList struct {
	Names []string `json:"names"`
}

func (HelloList) CreateResponse() (r []ListResult) { return }
func (HelloList) HttpMethod() string               { return "POST" }

type HelloArray struct {
	Names []string `json:"names"`
}

func (HelloArray) CreateResponse() (r []ArrayResult) { return }
func (HelloArray) HttpMethod() string                { return "POST" }

type HelloMap struct {
	Names []string `json:"names"`
}

func (HelloMap) CreateResponse() (r map[string]ArrayResult) { return }
func (HelloMap) HttpMethod() string                         { return "POST" }

type HelloQueryResponse struct {
	Names []string `json:"names"`
}

func (HelloQueryResponse) CreateResponse() (r ss.QueryResponse[string]) { return }
func (HelloQueryResponse) HttpMethod() string                           { return "POST" }

type HelloWithEnum struct {
	EnumProp         EnumType         `json:"enumProp,omitempty"`
	EnumTypeFlags    EnumTypeFlags    `json:"enumTypeFlags,omitempty"`
	EnumWithValues   EnumWithValues   `json:"enumWithValues,omitempty"`
	NullableEnumProp *EnumType        `json:"nullableEnumProp,omitempty"`
	EnumFlags        EnumFlags        `json:"enumFlags,omitempty"`
	EnumAsInt        EnumAsInt        `json:"enumAsInt,omitempty"`
	EnumStyle        EnumStyle        `json:"enumStyle,omitempty"`
	EnumStyleMembers EnumStyleMembers `json:"enumStyleMembers,omitempty"`
}

func (HelloWithEnum) CreateResponseVoid() {}
func (HelloWithEnum) HttpMethod() string  { return "POST" }

type HelloWithEnumList struct {
	EnumProp         []EnumType       `json:"enumProp"`
	EnumWithValues   []EnumWithValues `json:"enumWithValues"`
	NullableEnumProp []EnumType       `json:"nullableEnumProp"`
	EnumFlags        []EnumFlags      `json:"enumFlags"`
	EnumStyle        []EnumStyle      `json:"enumStyle"`
}

func (HelloWithEnumList) CreateResponseVoid() {}
func (HelloWithEnumList) HttpMethod() string  { return "POST" }

type HelloWithEnumMap struct {
	EnumProp         map[EnumType]EnumType             `json:"enumProp"`
	EnumWithValues   map[EnumWithValues]EnumWithValues `json:"enumWithValues"`
	NullableEnumProp map[EnumType]EnumType             `json:"nullableEnumProp"`
	EnumFlags        map[EnumFlags]EnumFlags           `json:"enumFlags"`
	EnumStyle        map[EnumStyle]EnumStyle           `json:"enumStyle"`
}

func (HelloWithEnumMap) CreateResponseVoid() {}
func (HelloWithEnumMap) HttpMethod() string  { return "POST" }

type HelloExternal struct {
	Name string `json:"name"`
}

func (HelloExternal) CreateResponseVoid() {}
func (HelloExternal) HttpMethod() string  { return "POST" }

/** @description AllowedAttributes Description */
// @Route("/allowed-attributes", "GET")
// @Api(Description="AllowedAttributes Description")
// @ApiResponse(Description="Your request was not understood", StatusCode=400)
// @DataContract
type AllowedAttributes struct {
	/** @description Range Description */
	// @DataMember(Name="Aliased")
	// @ApiMember(DataType="double", Description="Range Description", IsRequired=true, ParameterType="path")
	Aliased float64 `json:"Aliased"`
}

func (AllowedAttributes) CreateResponseVoid() {}
func (AllowedAttributes) HttpMethod() string  { return "GET" }

// @Route("/all-types")
type HelloAllTypes struct {
	Name               string             `json:"name"`
	AllTypes           AllTypes           `json:"allTypes"`
	AllCollectionTypes AllCollectionTypes `json:"allCollectionTypes"`
}

func (HelloAllTypes) CreateResponse() (r HelloAllTypesResponse) { return }
func (HelloAllTypes) HttpMethod() string                        { return "POST" }

type HelloSubAllTypes struct {
	AllTypesBase
	Hierarchy int `json:"hierarchy,omitempty"`
}

func (HelloSubAllTypes) CreateResponse() (r SubAllTypes) { return }
func (HelloSubAllTypes) HttpMethod() string              { return "POST" }

type HelloString struct {
	Name string `json:"name"`
}

func (HelloString) CreateResponse() (r string) { return }
func (HelloString) HttpMethod() string         { return "POST" }

type HelloVoid struct {
	Name string `json:"name"`
}

func (HelloVoid) CreateResponseVoid() {}
func (HelloVoid) HttpMethod() string  { return "POST" }

// @DataContract
type HelloWithDataContract struct {
	// @DataMember(Name="name", Order=1, IsRequired=true, EmitDefaultValue=false)
	Name string `json:"name"`
	// @DataMember(Name="id", Order=2, EmitDefaultValue=false)
	Id int `json:"id,omitempty"`
}

func (HelloWithDataContract) CreateResponse() (r HelloWithDataContractResponse) { return }
func (HelloWithDataContract) HttpMethod() string                                { return "POST" }

/** @description Description on HelloWithDescription type */
type HelloWithDescription struct {
	Name string `json:"name"`
}

func (HelloWithDescription) CreateResponse() (r HelloWithDescriptionResponse) { return }
func (HelloWithDescription) HttpMethod() string                               { return "POST" }

type HelloWithInheritance struct {
	HelloBase
	Name string `json:"name"`
}

func (HelloWithInheritance) CreateResponse() (r HelloWithInheritanceResponse) { return }
func (HelloWithInheritance) HttpMethod() string                               { return "POST" }

type HelloWithGenericInheritance struct {
	HelloBase_1[Poco]
	Result string `json:"result"`
}

func (HelloWithGenericInheritance) CreateResponseVoid() {}
func (HelloWithGenericInheritance) HttpMethod() string  { return "POST" }

type HelloWithGenericInheritance2 struct {
	HelloBase_1[Hello]
	Result string `json:"result"`
}

func (HelloWithGenericInheritance2) CreateResponseVoid() {}
func (HelloWithGenericInheritance2) HttpMethod() string  { return "POST" }

type HelloWithReturn struct {
	Name string `json:"name"`
}

func (HelloWithReturn) CreateResponse() (r HelloWithAlternateReturnResponse) { return }
func (HelloWithReturn) HttpMethod() string                                   { return "POST" }

// @Route("/helloroute")
type HelloWithRoute struct {
	Name string `json:"name"`
}

func (HelloWithRoute) CreateResponse() (r HelloWithRouteResponse) { return }
func (HelloWithRoute) HttpMethod() string                         { return "POST" }

type HelloWithType struct {
	Name string `json:"name"`
}

func (HelloWithType) CreateResponse() (r HelloWithTypeResponse) { return }
func (HelloWithType) HttpMethod() string                        { return "POST" }

type HelloInterface struct {
	Poco           IPoco           `json:"poco"`
	EmptyInterface IEmptyInterface `json:"emptyInterface"`
	EmptyClass     EmptyClass      `json:"emptyClass"`
}

func (HelloInterface) CreateResponseVoid() {}
func (HelloInterface) HttpMethod() string  { return "POST" }

type HelloInnerTypes struct {
}

func (HelloInnerTypes) CreateResponse() (r HelloInnerTypesResponse) { return }
func (HelloInnerTypes) HttpMethod() string                          { return "POST" }

type HelloBuiltin struct {
	DayOfWeek DayOfWeek `json:"dayOfWeek,omitempty"`
}

func (HelloBuiltin) CreateResponseVoid() {}
func (HelloBuiltin) HttpMethod() string  { return "POST" }

type HelloGet struct {
	Id int `json:"id,omitempty"`
}

func (HelloGet) CreateResponse() (r HelloVerbResponse) { return }
func (HelloGet) HttpMethod() string                    { return "GET" }

type HelloPost struct {
	HelloBase
}

func (HelloPost) CreateResponse() (r HelloVerbResponse) { return }
func (HelloPost) HttpMethod() string                    { return "POST" }

type HelloPut struct {
	Id int `json:"id,omitempty"`
}

func (HelloPut) CreateResponse() (r HelloVerbResponse) { return }
func (HelloPut) HttpMethod() string                    { return "PUT" }

type HelloDelete struct {
	Id int `json:"id,omitempty"`
}

func (HelloDelete) CreateResponse() (r HelloVerbResponse) { return }
func (HelloDelete) HttpMethod() string                    { return "DELETE" }

type HelloPatch struct {
	Id int `json:"id,omitempty"`
}

func (HelloPatch) CreateResponse() (r HelloVerbResponse) { return }
func (HelloPatch) HttpMethod() string                    { return "PATCH" }

type HelloReturnVoid struct {
	Id int `json:"id,omitempty"`
}

func (HelloReturnVoid) CreateResponseVoid() {}
func (HelloReturnVoid) HttpMethod() string  { return "POST" }

type EnumRequest struct {
	Operator ScopeType `json:"operator,omitempty"`
}

func (EnumRequest) CreateResponse() (r EnumResponse) { return }
func (EnumRequest) HttpMethod() string               { return "PUT" }

// @Route("/hellozip")
// @DataContract
type HelloZip struct {
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Test []string `json:"test"`
}

func (HelloZip) CreateResponse() (r HelloZipResponse) { return }
func (HelloZip) HttpMethod() string                   { return "POST" }

// @Route("/ping")
type Ping struct {
}

func (Ping) CreateResponse() (r PingResponse) { return }
func (Ping) HttpMethod() string               { return "POST" }

// @Route("/reset-connections")
type ResetConnections struct {
}

func (ResetConnections) CreateResponseVoid() {}
func (ResetConnections) HttpMethod() string  { return "POST" }

// @Route("/requires-role")
type RequiresRole struct {
}

func (RequiresRole) CreateResponse() (r RequiresRoleResponse) { return }
func (RequiresRole) HttpMethod() string                       { return "POST" }

// @Route("/return/string")
type ReturnString struct {
	Data string `json:"data"`
}

func (ReturnString) CreateResponse() (r string) { return }
func (ReturnString) HttpMethod() string         { return "POST" }

// @Route("/return/bytes")
type ReturnBytes struct {
	Data []byte `json:"data"`
}

func (ReturnBytes) CreateResponse() (r []byte) { return }
func (ReturnBytes) HttpMethod() string         { return "POST" }

// @Route("/return/stream")
type ReturnStream struct {
	Data []byte `json:"data"`
}

func (ReturnStream) CreateResponse() (r []byte) { return }
func (ReturnStream) HttpMethod() string         { return "POST" }

// @Route("/return/json")
type ReturnJson struct {
}

func (ReturnJson) CreateResponseVoid() {}
func (ReturnJson) HttpMethod() string  { return "POST" }

// @Route("/return/json/header")
type ReturnJsonHeader struct {
}

func (ReturnJsonHeader) CreateResponseVoid() {}
func (ReturnJsonHeader) HttpMethod() string  { return "POST" }

// @Route("/write/json")
type WriteJson struct {
}

func (WriteJson) CreateResponseVoid() {}
func (WriteJson) HttpMethod() string  { return "POST" }

// @Route("/Request1", "GET")
type GetRequest1 struct {
}

func (GetRequest1) CreateResponse() (r []ReturnedDto) { return }
func (GetRequest1) HttpMethod() string                { return "GET" }

// @Route("/Request2", "GET")
type GetRequest2 struct {
}

func (GetRequest2) CreateResponse() (r []ReturnedDto) { return }
func (GetRequest2) HttpMethod() string                { return "GET" }

// @Route("/sendjson")
type SendJson struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name"`
	RequestStream []byte `json:"requestStream"`
}

func (SendJson) CreateResponse() (r string) { return }
func (SendJson) HttpMethod() string         { return "POST" }

// @Route("/sendtext")
type SendText struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name"`
	ContentType   string `json:"contentType"`
	RequestStream []byte `json:"requestStream"`
}

func (SendText) CreateResponse() (r string) { return }
func (SendText) HttpMethod() string         { return "POST" }

// @Route("/sendraw")
type SendRaw struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name"`
	ContentType   string `json:"contentType"`
	RequestStream []byte `json:"requestStream"`
}

func (SendRaw) CreateResponse() (r []byte) { return }
func (SendRaw) HttpMethod() string         { return "POST" }

type SendDefault struct {
	Id int `json:"id,omitempty"`
}

func (SendDefault) CreateResponse() (r SendVerbResponse) { return }
func (SendDefault) HttpMethod() string                   { return "POST" }

// @Route("/sendrestget/{Id}", "GET")
type SendRestGet struct {
	Id int `json:"id,omitempty"`
}

func (SendRestGet) CreateResponse() (r SendVerbResponse) { return }
func (SendRestGet) HttpMethod() string                   { return "GET" }

type SendGet struct {
	Id int `json:"id,omitempty"`
}

func (SendGet) CreateResponse() (r SendVerbResponse) { return }
func (SendGet) HttpMethod() string                   { return "GET" }

type SendPost struct {
	Id int `json:"id,omitempty"`
}

func (SendPost) CreateResponse() (r SendVerbResponse) { return }
func (SendPost) HttpMethod() string                   { return "POST" }

type SendPut struct {
	Id int `json:"id,omitempty"`
}

func (SendPut) CreateResponse() (r SendVerbResponse) { return }
func (SendPut) HttpMethod() string                   { return "PUT" }

type SendReturnVoid struct {
	Id int `json:"id,omitempty"`
}

func (SendReturnVoid) CreateResponseVoid() {}
func (SendReturnVoid) HttpMethod() string  { return "POST" }

// @Route("/session")
type GetSession struct {
}

func (GetSession) CreateResponse() (r GetSessionResponse) { return }
func (GetSession) HttpMethod() string                     { return "POST" }

// @Route("/session/edit/{CustomName}")
type UpdateSession struct {
	CustomName *string `json:"customName,omitempty"`
}

func (UpdateSession) CreateResponse() (r GetSessionResponse) { return }
func (UpdateSession) HttpMethod() string                     { return "POST" }

// @Route("/Stuff")
// @DataContract(Namespace="http://schemas.servicestack.net/types")
type GetStuff struct {
	// @DataMember
	// @ApiMember(DataType="DateTime", Name="Summary Date")
	SummaryDate *time.Time `json:"summaryDate,omitempty"`
	// @DataMember
	// @ApiMember(DataType="DateTime", Name="Summary End Date")
	SummaryEndDate *time.Time `json:"summaryEndDate,omitempty"`
	// @DataMember
	// @ApiMember(DataType="string", Name="Symbol")
	Symbol *string `json:"symbol,omitempty"`
	// @DataMember
	// @ApiMember(DataType="string", Name="Email")
	Email *string `json:"email,omitempty"`
	// @DataMember
	// @ApiMember(DataType="bool", Name="Is Enabled")
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

func (GetStuff) CreateResponse() (r GetStuffResponse) { return }
func (GetStuff) HttpMethod() string                   { return "POST" }

type StoreLogs struct {
	Loggers []Logger `json:"loggers"`
}

func (StoreLogs) CreateResponse() (r StoreLogsResponse) { return }
func (StoreLogs) HttpMethod() string                    { return "POST" }

type HelloAuth struct {
	Name *string `json:"name,omitempty"`
}

func (HelloAuth) CreateResponse() (r HelloResponse) { return }
func (HelloAuth) HttpMethod() string                { return "POST" }

// @Route("/testauth")
type TestAuth struct {
}

func (TestAuth) CreateResponse() (r TestAuthResponse) { return }
func (TestAuth) HttpMethod() string                   { return "POST" }

// @Route("/testdata/AllTypes")
type TestDataAllTypes struct {
}

func (TestDataAllTypes) CreateResponse() (r AllTypes) { return }
func (TestDataAllTypes) HttpMethod() string           { return "POST" }

// @Route("/testdata/AllCollectionTypes")
type TestDataAllCollectionTypes struct {
}

func (TestDataAllCollectionTypes) CreateResponse() (r AllCollectionTypes) { return }
func (TestDataAllCollectionTypes) HttpMethod() string                     { return "POST" }

// @Route("/void-response")
type TestVoidResponse struct {
}

func (TestVoidResponse) CreateResponseVoid() {}
func (TestVoidResponse) HttpMethod() string  { return "POST" }

// @Route("/null-response")
type TestNullResponse struct {
}

func (TestNullResponse) CreateResponseVoid() {}
func (TestNullResponse) HttpMethod() string  { return "POST" }

/** @description Chat Completions API (OpenAI-Compatible) */
// @Route("/v1/chat/completions", "POST")
// @DataContract
type ChatCompletion struct {
	/** @description The messages to generate chat completions for. */
	// @DataMember(Name="messages")
	Messages []AiMessage `json:"messages"`
	/** @description ID of the model to use. See the model endpoint compatibility table for details on which models work with the Chat API */
	// @DataMember(Name="model")
	Model string `json:"model"`
	/** @description Parameters for audio output. Required when audio output is requested with modalities: [audio] */
	// @DataMember(Name="audio")
	Audio *AiChatAudio `json:"audio,omitempty"`
	/** @description Modify the likelihood of specified tokens appearing in the completion. */
	// @DataMember(Name="logit_bias")
	LogitBias map[int]int `json:"logit_bias,omitempty"`
	/** @description Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. */
	// @DataMember(Name="metadata")
	Metadata map[string]string `json:"metadata,omitempty"`
	/** @description Constrains effort on reasoning for reasoning models. Currently supported values are minimal, low, medium, and high (none, default). Reducing reasoning effort can result in faster responses and fewer tokens used on reasoning in a response. */
	// @DataMember(Name="reasoning_effort")
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	/** @description An object specifying the format that the model must output. Compatible with GPT-4 Turbo and all GPT-3.5 Turbo models newer than `gpt-3.5-turbo-1106`. Setting Type to ResponseFormat.JsonObject enables JSON mode, which guarantees the message the model generates is valid JSON. */
	// @DataMember(Name="response_format")
	ResponseFormat *AiResponseFormat `json:"response_format,omitempty"`
	/** @description Specifies the processing type used for serving the request. */
	// @DataMember(Name="service_tier")
	ServiceTier *string `json:"service_tier,omitempty"`
	/** @description A stable identifier used to help detect users of your application that may be violating OpenAI's usage policies. The IDs should be a string that uniquely identifies each user. */
	// @DataMember(Name="safety_identifier")
	SafetyIdentifier *string `json:"safety_identifier,omitempty"`
	/** @description Up to 4 sequences where the API will stop generating further tokens. */
	// @DataMember(Name="stop")
	Stop []string `json:"stop,omitempty"`
	/** @description Output types that you would like the model to generate. Most models are capable of generating text, which is the default: */
	// @DataMember(Name="modalities")
	Modalities []string `json:"modalities,omitempty"`
	/** @description Used by OpenAI to cache responses for similar requests to optimize your cache hit rates. */
	// @DataMember(Name="prompt_cache_key")
	PromptCacheKey *string `json:"prompt_cache_key,omitempty"`
	/** @description A list of tools the model may call. Currently, only functions are supported as a tool. Use this to provide a list of functions the model may generate JSON inputs for. A max of 128 functions are supported. */
	// @DataMember(Name="tools")
	Tools []Tool `json:"tools,omitempty"`
	/** @description Constrains the verbosity of the model's response. Lower values will result in more concise responses, while higher values will result in more verbose responses. Currently supported values are low, medium, and high. */
	// @DataMember(Name="verbosity")
	Verbosity *string `json:"verbosity,omitempty"`
	/** @description What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. */
	// @DataMember(Name="temperature")
	Temperature *float64 `json:"temperature,omitempty"`
	/** @description An upper bound for the number of tokens that can be generated for a completion, including visible output tokens and reasoning tokens. */
	// @DataMember(Name="max_completion_tokens")
	MaxCompletionTokens *int `json:"max_completion_tokens,omitempty"`
	/** @description An integer between 0 and 20 specifying the number of most likely tokens to return at each token position, each with an associated log probability. logprobs must be set to true if this parameter is used. */
	// @DataMember(Name="top_logprobs")
	TopLogprobs *int `json:"top_logprobs,omitempty"`
	/** @description An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered. */
	// @DataMember(Name="top_p")
	TopP *float64 `json:"top_p,omitempty"`
	/** @description Number between `-2.0` and `2.0`. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim. */
	// @DataMember(Name="frequency_penalty")
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
	/** @description Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics. */
	// @DataMember(Name="presence_penalty")
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`
	/** @description This feature is in Beta. If specified, our system will make a best effort to sample deterministically, such that repeated requests with the same seed and parameters should return the same result. Determinism is not guaranteed, and you should refer to the system_fingerprint response parameter to monitor changes in the backend. */
	// @DataMember(Name="seed")
	Seed *int `json:"seed,omitempty"`
	/** @description How many chat completion choices to generate for each input message. Note that you will be charged based on the number of generated tokens across all of the choices. Keep `n` as `1` to minimize costs. */
	// @DataMember(Name="n")
	N *int `json:"n,omitempty"`
	/** @description Whether or not to store the output of this chat completion request for use in our model distillation or evals products. */
	// @DataMember(Name="store")
	Store *bool `json:"store,omitempty"`
	/** @description Whether to return log probabilities of the output tokens or not. If true, returns the log probabilities of each output token returned in the content of message. */
	// @DataMember(Name="logprobs")
	Logprobs *bool `json:"logprobs,omitempty"`
	/** @description Whether to enable parallel function calling during tool use. */
	// @DataMember(Name="parallel_tool_calls")
	ParallelToolCalls *bool `json:"parallel_tool_calls,omitempty"`
	/** @description Whether to enable thinking mode for some Qwen models and providers. */
	// @DataMember(Name="enable_thinking")
	EnableThinking *bool `json:"enable_thinking,omitempty"`
	/** @description If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only server-sent events as they become available, with the stream terminated by a `data: [DONE]` message. */
	// @DataMember(Name="stream")
	Stream *bool `json:"stream,omitempty"`
}

func (ChatCompletion) CreateResponse() (r ChatResponse) { return }
func (ChatCompletion) HttpMethod() string               { return "POST" }

/** @description Find Bookings */
// @Route("/bookings", "GET")
// @Route("/bookings/{Id}", "GET")
type QueryBookings struct {
	ss.QueryDb
	Id *int `json:"id,omitempty"`
}

func (QueryBookings) CreateResponse() (r ss.QueryResponse[Booking]) { return }
func (QueryBookings) HttpMethod() string                            { return "GET" }

/** @description Find Coupons */
// @Route("/coupons", "GET")
type QueryCoupons struct {
	ss.QueryDb
	Id string `json:"id"`
}

func (QueryCoupons) CreateResponse() (r ss.QueryResponse[Coupon]) { return }
func (QueryCoupons) HttpMethod() string                           { return "GET" }

type QueryAddresses struct {
	ss.QueryDb
	Ids []int64 `json:"ids"`
}

func (QueryAddresses) CreateResponse() (r ss.QueryResponse[Address]) { return }
func (QueryAddresses) HttpMethod() string                            { return "GET" }

type QueryRockstarAudit struct {
	QueryDbTenant[RockstarAuditTenant, RockstarAuto]
	Id *int `json:"id,omitempty"`
}

func (QueryRockstarAudit) CreateResponse() (r ss.QueryResponse[RockstarAuto]) { return }
func (QueryRockstarAudit) HttpMethod() string                                 { return "GET" }

type QueryRockstarAuditSubOr struct {
	ss.QueryDb
	FirstNameStartsWith string `json:"firstNameStartsWith"`
	AgeOlderThan        *int   `json:"ageOlderThan,omitempty"`
}

func (QueryRockstarAuditSubOr) CreateResponse() (r ss.QueryResponse[RockstarAuto]) { return }
func (QueryRockstarAuditSubOr) HttpMethod() string                                 { return "GET" }

type QueryPocoBase struct {
	ss.QueryDb
	Id int `json:"id,omitempty"`
}

func (QueryPocoBase) CreateResponse() (r ss.QueryResponse[OnlyDefinedInGenericType]) { return }
func (QueryPocoBase) HttpMethod() string                                             { return "GET" }

type QueryPocoIntoBase struct {
	ss.QueryDb
	Id int `json:"id,omitempty"`
}

func (QueryPocoIntoBase) CreateResponse() (r ss.QueryResponse[OnlyDefinedInGenericTypeInto]) { return }
func (QueryPocoIntoBase) HttpMethod() string                                                 { return "GET" }

// @Route("/message/query/{Id}", "GET")
type MessageQuery struct {
	ss.QueryDb
	Id int `json:"id,omitempty"`
}

func (MessageQuery) CreateResponse() (r ss.QueryResponse[MessageQuery]) { return }
func (MessageQuery) HttpMethod() string                                 { return "GET" }

// @Route("/rockstars", "GET")
type QueryRockstars struct {
	ss.QueryDb
}

func (QueryRockstars) CreateResponse() (r ss.QueryResponse[Rockstar]) { return }
func (QueryRockstars) HttpMethod() string                             { return "GET" }

/** @description Create a new Booking */
// @Route("/bookings", "POST")
// @ValidateRequest(Validator="HasRole(`Employee`)")
type CreateBooking struct {
	/** @description Name this Booking is for */
	// @Validate(Validator="NotEmpty")
	Name     string   `json:"name"`
	RoomType RoomType `json:"roomType,omitempty"`
	// @Validate(Validator="GreaterThan(0)")
	RoomNumber int `json:"roomNumber,omitempty"`
	// @Validate(Validator="GreaterThan(0)")
	Cost float64 `json:"cost,omitempty"`
	// @Required()
	BookingStartDate   time.Time  `json:"bookingStartDate"`
	BookingEndDate     *time.Time `json:"bookingEndDate,omitempty"`
	Notes              *string    `json:"notes,omitempty"`
	CouponId           *string    `json:"couponId,omitempty"`
	PermanentAddressId *int64     `json:"permanentAddressId,omitempty"`
	PostalAddressId    *int64     `json:"postalAddressId,omitempty"`
}

func (CreateBooking) CreateResponse() (r ss.IdResponse) { return }
func (CreateBooking) HttpMethod() string                { return "POST" }

/** @description Update an existing Booking */
// @Route("/booking/{Id}", "PATCH")
// @ValidateRequest(Validator="HasRole(`Employee`)")
// @ValidateRequest(Validator="HasRole(`Manager`)")
type UpdateBooking struct {
	Id       int       `json:"id,omitempty"`
	Name     *string   `json:"name,omitempty"`
	RoomType *RoomType `json:"roomType,omitempty"`
	// @Validate(Validator="GreaterThan(0)")
	RoomNumber *int `json:"roomNumber,omitempty"`
	// @Validate(Validator="GreaterThan(0)")
	Cost               *float64   `json:"cost,omitempty"`
	BookingStartDate   *time.Time `json:"bookingStartDate,omitempty"`
	BookingEndDate     *time.Time `json:"bookingEndDate,omitempty"`
	Notes              *string    `json:"notes,omitempty"`
	CouponId           *string    `json:"couponId,omitempty"`
	Cancelled          *bool      `json:"cancelled,omitempty"`
	PermanentAddressId *int64     `json:"permanentAddressId,omitempty"`
	PostalAddressId    *int64     `json:"postalAddressId,omitempty"`
}

func (UpdateBooking) CreateResponse() (r ss.IdResponse) { return }
func (UpdateBooking) HttpMethod() string                { return "PATCH" }

/** @description Delete a Booking */
// @Route("/booking/{Id}", "DELETE")
type DeleteBooking struct {
	Id int `json:"id,omitempty"`
}

func (DeleteBooking) CreateResponseVoid() {}
func (DeleteBooking) HttpMethod() string  { return "DELETE" }

// @Route("/coupons", "POST")
// @ValidateRequest(Validator="HasRole(`Employee`)")
type CreateCoupon struct {
	// @Validate(Validator="NotEmpty")
	Id string `json:"id"`
	// @Validate(Validator="NotEmpty")
	Description string `json:"description"`
	// @Validate(Validator="GreaterThan(0)")
	Discount int `json:"discount,omitempty"`
	// @Validate(Validator="NotNull")
	ExpiryDate time.Time `json:"expiryDate"`
}

func (CreateCoupon) CreateResponse() (r ss.IdResponse) { return }
func (CreateCoupon) HttpMethod() string                { return "POST" }

// @Route("/coupons/{Id}", "PATCH")
// @ValidateRequest(Validator="HasRole(`Employee`)")
type UpdateCoupon struct {
	Id string `json:"id"`
	// @Validate(Validator="NotEmpty")
	Description string `json:"description"`
	// @Validate(Validator="NotNull")
	// @Validate(Validator="GreaterThan(0)")
	Discount *int `json:"discount"`
	// @Validate(Validator="NotNull")
	ExpiryDate *time.Time `json:"expiryDate"`
}

func (UpdateCoupon) CreateResponse() (r ss.IdResponse) { return }
func (UpdateCoupon) HttpMethod() string                { return "PATCH" }

/** @description Delete a Coupon */
// @Route("/coupons/{Id}", "DELETE")
// @ValidateRequest(Validator="HasRole(`Manager`)")
type DeleteCoupon struct {
	Id string `json:"id"`
}

func (DeleteCoupon) CreateResponseVoid() {}
func (DeleteCoupon) HttpMethod() string  { return "DELETE" }

type CreateAddress struct {
	AddressText *string `json:"addressText,omitempty"`
}

func (CreateAddress) CreateResponse() (r ss.IdResponse) { return }
func (CreateAddress) HttpMethod() string                { return "POST" }

type UpdateAddress struct {
	Id          int     `json:"id,omitempty"`
	AddressText *string `json:"addressText,omitempty"`
}

func (UpdateAddress) CreateResponse() (r ss.IdResponse) { return }
func (UpdateAddress) HttpMethod() string                { return "PATCH" }

type CreateRockstarAudit struct {
	RockstarBase
}

func (CreateRockstarAudit) CreateResponse() (r RockstarWithIdResponse) { return }
func (CreateRockstarAudit) HttpMethod() string                         { return "POST" }

type CreateRockstarAuditTenant struct {
	CreateAuditTenantBase[RockstarAuditTenant, RockstarWithIdAndResultResponse]
	SessionId    string       `json:"sessionId"`
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Age          *int         `json:"age,omitempty"`
	DateOfBirth  time.Time    `json:"dateOfBirth,omitempty"`
	DateDied     *time.Time   `json:"dateDied,omitempty"`
	LivingStatus LivingStatus `json:"livingStatus,omitempty"`
}

func (CreateRockstarAuditTenant) CreateResponse() (r RockstarWithIdAndResultResponse) { return }
func (CreateRockstarAuditTenant) HttpMethod() string                                  { return "POST" }

type UpdateRockstarAuditTenant struct {
	UpdateAuditTenantBase[RockstarAuditTenant, RockstarWithIdAndResultResponse]
	SessionId    string        `json:"sessionId"`
	Id           int           `json:"id,omitempty"`
	FirstName    string        `json:"firstName"`
	LivingStatus *LivingStatus `json:"livingStatus,omitempty"`
}

func (UpdateRockstarAuditTenant) CreateResponse() (r RockstarWithIdAndResultResponse) { return }
func (UpdateRockstarAuditTenant) HttpMethod() string                                  { return "PUT" }

type PatchRockstarAuditTenant struct {
	PatchAuditTenantBase[RockstarAuditTenant, RockstarWithIdAndResultResponse]
	SessionId    string        `json:"sessionId"`
	Id           int           `json:"id,omitempty"`
	FirstName    string        `json:"firstName"`
	LivingStatus *LivingStatus `json:"livingStatus,omitempty"`
}

func (PatchRockstarAuditTenant) CreateResponse() (r RockstarWithIdAndResultResponse) { return }
func (PatchRockstarAuditTenant) HttpMethod() string                                  { return "PATCH" }

type SoftDeleteAuditTenant struct {
	SoftDeleteAuditTenantBase[RockstarAuditTenant, RockstarWithIdAndResultResponse]
	Id int `json:"id,omitempty"`
}

func (SoftDeleteAuditTenant) CreateResponse() (r RockstarWithIdAndResultResponse) { return }
func (SoftDeleteAuditTenant) HttpMethod() string                                  { return "PUT" }

type CreateRockstarAuditMqToken struct {
	RockstarBase
	BearerToken string `json:"bearerToken"`
}

func (CreateRockstarAuditMqToken) CreateResponse() (r RockstarWithIdResponse) { return }
func (CreateRockstarAuditMqToken) HttpMethod() string                         { return "POST" }

type RealDeleteAuditTenant struct {
	SessionId string `json:"sessionId"`
	Id        int    `json:"id,omitempty"`
	Age       *int   `json:"age,omitempty"`
}

func (RealDeleteAuditTenant) CreateResponse() (r RockstarWithIdAndCountResponse) { return }
func (RealDeleteAuditTenant) HttpMethod() string                                 { return "DELETE" }

type CreateRockstarVersion struct {
	RockstarBase
}

func (CreateRockstarVersion) CreateResponse() (r RockstarWithIdAndRowVersionResponse) { return }
func (CreateRockstarVersion) HttpMethod() string                                      { return "POST" }
