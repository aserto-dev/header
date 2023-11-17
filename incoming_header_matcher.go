package header

import (
	"github.com/go-http-utils/headers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// IncomingHeaderMatcher is a matcher that makes it so that HTTP clients do not have to prefix
// the header key with Grpc-Metadata-.
// see https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/#mapping-from-http-request-headers-to-grpc-client-metadata
func IncomingHeaderMatcher(key string) (string, bool) {
	switch key {
	case string(HeaderAsertoTenantID),
		string(HeaderAsertoAccountID),
		string(HeaderAsertoRequestID),
		string(HeaderAsertoTenantKey),
		headers.IfMatch,
		headers.IfNoneMatch:
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
