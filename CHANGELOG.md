# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.1]

### Added

- `npm run bump` / `npm run release` scripts that tag a version and create the
  GitHub Release, which the `release` workflow publishes to the Go module proxy

## [0.1.0]

### Added

- `Client` for consuming ServiceStack APIs with generated typed DTOs
- Response Types inferred from Request DTOs via `IReturn[T]`/`CreateResponse()`
- HTTP Verb resolved from the DTO's `HttpMethod()`, falling back to its name
- Structured `ResponseStatus` errors in `WebServiceException`, incl. field errors
- `Api()` results that return errors instead of a separate error
- Auth with Basic Auth, API Keys, Bearer Tokens, Refresh Tokens and Cookies
- Batched, one-way and multipart file upload Requests
- `context.Context` variants for every API
- Built-in ServiceStack DTOs referenced by generated DTOs

[0.1.1]: https://github.com/ServiceStack/servicestack-go/releases/tag/v0.1.1
[0.1.0]: https://github.com/ServiceStack/servicestack-go/releases/tag/v0.1.0
