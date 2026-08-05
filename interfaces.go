package servicestack

// IReturn is implemented by Request DTOs that return a typed TResponse.
//
// Generated DTOs implement it with a CreateResponse() method, which lets the
// generic client functions infer the Response Type from the Request DTO, e.g:
//
//	func (Hello) CreateResponse() (r HelloResponse) { return }
//
//	res, err := servicestack.Send(client, Hello{Name: "World"}) // res is HelloResponse
type IReturn[TResponse any] interface {
	CreateResponse() TResponse
}

// IReturnVoid is implemented by Request DTOs that don't return a Response Body.
type IReturnVoid interface {
	CreateResponseVoid()
}

// IVerb is implemented by Request DTOs annotated with the HTTP Method they
// should be sent with, e.g:
//
//	func (Hello) HttpMethod() string { return "GET" }
type IVerb interface {
	HttpMethod() string
}
