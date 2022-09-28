package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-admin/internal/core"
)

var HandlerSet = wire.NewSet(
	UserSet,
)

func GetMobile(ctx *gin.Context) string {
	cookieToken, _ := ctx.Cookie("TOKEN")
	cookieToken, _ = core.RSADecrypt(core.GetPrivateKey(), cookieToken)
	cookieTokenEmail, _ := core.ParseToken(cookieToken)
	return cookieTokenEmail
}
