package token

const (
	// token默认过期时长
	DEFAULT_ACCESS_TOKEN_EXPIRE_SECOND = 600
	// 刷新token默认过期时间
	DEFAULT_REFRESH_TOKEN_EXPIRE_SECOND = DEFAULT_ACCESS_TOKEN_EXPIRE_SECOND * 6
)

const (
	ACCESS_TOKEN_HEADER_KEY        = "Authorization"
	ACCESS_TOKEN_COOKIE_KEY        = "mcenter.access_token"
	VALIDATE_TOKEN_HEADER_KEY      = "X-VALIDATE-TOKEN"
	ACCESS_TOKEN_HEADER_KEY_SOCKET = "sec-WebSocket-Protocol"
)

const (
	TOKEN_ATTRIBUTE_NAME = "token"
)
