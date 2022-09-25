package interceptor

import "github.com/google/wire"

var InterceptorSet = wire.NewSet(
	AuthSet,
)
