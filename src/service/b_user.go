package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"math"
	"maxblog-me-admin/internal/core"
	"maxblog-me-admin/src/pb"
	"maxblog-me-admin/src/schema"
	"time"
)

var UserSet = wire.NewSet(wire.Struct(new(BUser), "*"))

type BUser struct{}

func (bUser *BUser) ValidateLoginInfo(c *gin.Context, loginInfo *schema.LoginInfo) error {
	return nil
}

func (bUser *BUser) SetLoginCookie(c *gin.Context, loginInfo *schema.LoginInfo) error {
	duration := time.Hour * 24 * 7 // 一周
	cipherToken, _, err := genToken(loginInfo.Mobile, duration)
	if err != nil {
		return err
	}
	targetDomain := core.GetUpstreamDomain()
	secure := core.GetUpstreamSecure()
	durationInt := int(duration) / int(math.Pow10(9))
	c.SetCookie(
		"TOKEN",
		cipherToken,
		durationInt,
		"/",
		targetDomain,
		secure,
		true)
	return nil
}

func (bUser *BUser) SetLoginJWT(c *gin.Context, loginInfo *schema.LoginInfo) (string, string, error) {
	duration := time.Hour * 24 * 7 // 一周
	cipherToken, mobile, err := genToken(loginInfo.Mobile, duration)
	if err != nil {
		return EmptyStr, EmptyStr, err
	}
	return cipherToken, mobile, nil
}

func (bUser *BUser) GetUserById(c *gin.Context, id uint32) (*schema.UserRes, error) {
	addr := core.GetDownstreamMaxblogBEUserAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Fatal(core.FormatError(300, err).Error())
		return nil, err
	}
	client := pb.NewUserServiceClient(conn)
	pbRes, err := client.GetUserById(context.Background(), &pb.IdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	dataRes := schema.Pb2Res(pbRes)
	return &dataRes, nil
}
