package token

import (
	"errors"

	"github.com/emicklei/go-restful/v3"
)

var ErrTokenExpoired = errors.New("refresh_token expoired")

// var ErrOtherPlaceLoggedIn = errors.New("异地登陆")
var ErrOtherPlaceLoggedIn = errors.New("ErrOtherPlaceLoggedIn")

func GetTokenFromRequest(r *restful.Request) *Token {
	tk := r.Attribute(TOKEN_ATTRIBUTE_NAME)
	if tk == nil {
		return nil
	}
	return tk.(*Token)
}
