package interceptor

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-admin/internal/core"
	"net/http"
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

type Interceptor struct {
	ILogger core.ILogger
}

func (itcpt *Interceptor) CheckTokens() gin.HandlerFunc {
	return func(c *gin.Context) {
		j := core.NewJWT()
		// token in req header
		headerToken := c.Request.Header.Get("Authorization")
		if headerToken == "" || len(headerToken) == 0 {
			itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, nil))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, nil))
			return
		}
		headerDecryptedToken, err := core.RSADecrypt(core.GetPrivateKey(), headerToken)
		if err != nil {
			itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		headerParsedToken, err := j.ParseToken(headerDecryptedToken)
		if err != nil {
			if err.Error() == core.TokenExpired {
				itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
				c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
				return
			}
			itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		// token in req cookie
		cookieToken, err := c.Cookie("TOKEN")
		if cookieToken == "" || len(cookieToken) == 0 {
			itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, nil))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		cookieDecryptedToken, err := core.RSADecrypt(core.GetPrivateKey(), cookieToken)
		if err != nil {
			itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		cookieParsedToken, err := j.ParseToken(cookieDecryptedToken)
		if err != nil {
			if err.Error() == core.TokenExpired {
				itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
				c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
				return
			}
			itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, err))
			return
		}
		// checking tokens info
		if headerParsedToken != cookieParsedToken {
			itcpt.ILogger.LogFailure(core.GetFuncName(), core.FormatError(206, nil))
			c.AbortWithStatusJSON(http.StatusUnauthorized, core.FormatError(206, nil))
			return
		}
		c.Next()
	}
}
