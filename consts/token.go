package consts

import "time"

const (
	AccessTokenHeader    = "Authorization"
	HeaderForwardedProto = "X-Forwarded-Proto"
	MaxAge               = 3600 * 24
)

const (
	AccessTokenExpireDuration = time.Hour * 24
)
