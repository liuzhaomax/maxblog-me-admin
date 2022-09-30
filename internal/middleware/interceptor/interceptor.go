package interceptor

import (
	"github.com/google/wire"
)

var InterceptorSet = wire.NewSet(wire.Struct(new(Interceptor), "*"))

type Interceptor struct {
	InterceptorAuth *Auth
}
