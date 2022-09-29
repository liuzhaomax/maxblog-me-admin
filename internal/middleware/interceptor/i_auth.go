package interceptor

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-admin/internal/core"
	"net/http"
)

var AuthSet = wire.NewSet(wire.Struct(new(Auth), "*"))

type Auth struct {
	ILogger core.ILogger
}

func (auth *Auth) CheckTokens() gin.HandlerFunc {
	return func(c *gin.Context) {
		j := core.NewJWT()
		// token in req header
		headerToken := c.Request.Header.Get("Authorization")
		if headerToken == "" || len(headerToken) == 0 {
			auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, nil))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, nil))
			return
		}
		headerDecryptedToken, err := core.RSADecrypt(core.GetPrivateKey(), headerToken)
		if err != nil {
			auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		headerParsedToken, err := j.ParseToken(headerDecryptedToken)
		if err != nil {
			if err.Error() == core.TokenExpired {
				auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
				c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
				return
			}
			auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		// token in req cookie
		cookieToken, err := c.Cookie("TOKEN")
		if cookieToken == "" || len(cookieToken) == 0 {
			auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, nil))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		cookieDecryptedToken, err := core.RSADecrypt(core.GetPrivateKey(), cookieToken)
		if err != nil {
			auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		cookieParsedToken, err := j.ParseToken(cookieDecryptedToken)
		if err != nil {
			if err.Error() == core.TokenExpired {
				auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
				c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
				return
			}
			auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		// checking tokens info
		if headerParsedToken != cookieParsedToken {
			auth.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, nil))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, nil))
			return
		}
		c.Next()
	}
}
