package interceptor

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-template/internal/core"
	"net/http"
	"strings"
	"sync"
)

var AuthSet = wire.NewSet(wire.Struct(new(Interceptor), "*"))

var interceptor *Interceptor
var once sync.Once

func init() {
	once.Do(func() {
		interceptor = &Interceptor{}
	})
}

func GetInstanceOfContext() *Interceptor {
	return interceptor
}

type Interceptor struct{}

// emails in tokens need to be equal
func (inter *Interceptor) CheckTwoTokens() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token in req header
		headerToken := ctx.Request.Header.Get("Authorization")
		headerToken, _ = core.RSADecrypt(core.GetPrivateKey(), headerToken)
		headerTokenEmail, _ := core.ParseToken(headerToken)
		headerTokenEmail = strings.Split(headerTokenEmail, "|")[0]
		// token in req cookie
		cookieToken, _ := ctx.Cookie("TOKEN")
		cookieToken, _ = core.RSADecrypt(core.GetPrivateKey(), cookieToken)
		cookieTokenEmail, _ := core.ParseToken(cookieToken)
		// checking tokens info
		if headerTokenEmail != cookieTokenEmail || headerTokenEmail == "" || cookieTokenEmail == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(200, nil))
		}
		ctx.Next()
	}
}
