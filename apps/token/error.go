package token

import "github.com/infraboard/mcube/exception"

var (
	ErrUnauthorized = exception.NewUnauthorized("Auth Header Required, Format: %s: Bearer ${access_token}, or Format: %s: Bearer${access_token}", ACCESS_TOKEN_HEADER_KEY, ACCESS_TOKEN_HEADER_KEY_SOCKET)
)
