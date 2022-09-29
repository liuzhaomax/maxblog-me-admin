package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"maxblog-me-admin/internal/core"
	"maxblog-me-admin/src/schema"
	"maxblog-me-admin/src/service"
	"maxblog-me-admin/src/utils"
	"net/http"
)

var UserSet = wire.NewSet(wire.Struct(new(HUser), "*"))

type HUser struct {
	BUser *service.BUser
	IRes  core.IResponse
}

func (hUser *HUser) GetIndex(c *gin.Context) {
	hUser.IRes.ResSuccess(c, core.GetFuncName(), http.StatusOK, "Hello MaxBlog")
}

func (hUser *HUser) GetPuk(c *gin.Context) {
	puk := core.GetPublicKeyStr()
	hUser.IRes.ResSuccess(c, core.GetFuncName(), http.StatusOK, puk)
}

func (hUser *HUser) PostLogin(c *gin.Context) {
	var loginInfo schema.LoginInfo
	err := c.ShouldBind(&loginInfo)
	if err != nil {
		hUser.IRes.ResFailure(c, core.GetFuncName(), http.StatusBadRequest, core.FormatError(201, err))
		return
	}
	err = hUser.BUser.ValidateLoginInfo(c, &loginInfo)
	if err != nil {
		hUser.IRes.ResFailure(c, core.GetFuncName(), http.StatusUnauthorized, core.FormatError(200, err))
		return
	}
	cipherToken, mobile, err := hUser.BUser.SetLoginCookie(c, &loginInfo)
	if err != nil {
		hUser.IRes.ResFailure(c, core.GetFuncName(), http.StatusInternalServerError, core.FormatError(200, err))
		return
	}
	logger.WithFields(logger.Fields{"用户": mobile}).Info(core.FormatInfo(108))
	hUser.IRes.ResSuccess(c, core.GetFuncName(), http.StatusOK, cipherToken)
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
	hUser.IRes.ResSuccess(c, core.GetFuncName(), http.StatusOK, dataRes)
}

func (hUser *HUser) PostUser(c *gin.Context) {
	var userReq schema.UserReq
	err := c.ShouldBind(&userReq)
	if err != nil {
		hUser.IRes.ResFailure(c, core.GetFuncName(), http.StatusBadRequest, core.FormatError(201, err))
		return
	}
	success, mobile, err := hUser.BUser.CreateUser(c, &userReq)
	if err != nil || success != true {
		hUser.IRes.ResFailure(c, core.GetFuncName(), http.StatusInternalServerError, core.FormatError(205, err))
		return
	}
	logger.WithFields(logger.Fields{"用户": mobile}).Info(core.FormatInfo(108))
	hUser.IRes.ResSuccess(c, core.GetFuncName(), http.StatusOK, gin.H{"msg": success})
}

func (hUser *HUser) DeleteLogout(c *gin.Context) {
	mobile, err := hUser.BUser.ClearLoginCookie(c)
	if err != nil {
		hUser.IRes.ResFailure(c, core.GetFuncName(), http.StatusInternalServerError, core.FormatError(207, err))
		return
	}
	logger.WithFields(logger.Fields{"用户": mobile}).Info(core.FormatInfo(109))
	hUser.IRes.ResSuccess(c, core.GetFuncName(), http.StatusOK, gin.H{"msg": "ok"})
}
