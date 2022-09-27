package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-admin/internal/core"
	"maxblog-me-admin/src/service"
	"maxblog-me-admin/src/utils"
	"net/http"
)

var UserSet = wire.NewSet(wire.Struct(new(HUser), "*"))

type HUser struct {
	BUser *service.BUser
	IRes  core.IResponse
}

func (hUser *HUser) GetUserById(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := utils.Str2Uint32(idRaw)
	if err != nil {
		hUser.IRes.ResFailure(c, core.GetFuncName(), http.StatusBadRequest, core.FormatError(299, err))
		return
	}
	dataRes, err := hUser.BUser.GetUserById(c, id)
	if err != nil {
		hUser.IRes.ResFailure(c, core.GetFuncName(), http.StatusInternalServerError, core.FormatError(399, err))
		return
	}
	hUser.IRes.ResSuccess(c, core.GetFuncName(), dataRes)
}
