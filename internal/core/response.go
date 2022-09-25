package core

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var ResponseSet = wire.NewSet(wire.Struct(new(Response), "*"), wire.Bind(new(IResponse), new(*Response)))

type Response struct {
	ILogger ILogger
}

type IResponse interface {
	ResSuccess(ctx *gin.Context, funcName string, sth interface{})
	ResFailure(ctx *gin.Context, funcName string, code int, err error)
}

func (res *Response) ResSuccess(ctx *gin.Context, funcName string, sth interface{}) {
	res.ILogger.LogSuccess(funcName)
	res.ResJson(ctx, http.StatusOK, gin.H{
		"status": gin.H{
			"code": 0,
			"desc": "success",
		},
		"data": sth,
	})
}

func (res *Response) ResJson(ctx *gin.Context, status int, sth interface{}) {
	ctx.JSON(status, sth)
}

func (res *Response) ResFailure(ctx *gin.Context, funcName string, status int, err error) {
	res.ILogger.LogFailure(funcName, err)
	res.ResError(ctx, status, err)
}

func (res *Response) ResError(ctx *gin.Context, status int, err error) {
	res.ResJson(ctx, status, err)
}
