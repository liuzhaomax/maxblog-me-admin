package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math"
	"maxblog-me-admin/internal/core"
	"maxblog-me-admin/src/pb"
	"maxblog-me-admin/src/schema"
	"maxblog-me-admin/src/utils"
	"time"
)

var UserSet = wire.NewSet(wire.Struct(new(BUser), "*"))

type BUser struct{}

func (bUser *BUser) ValidateLoginInfo(c *gin.Context, loginInfo *schema.LoginInfo) error {
	addr := core.GetDownstreamMaxblogBEUserAddr()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Info(core.FormatError(300, err).Error())
		return err
	}
	mobile, err := core.RSADecrypt(core.GetPrivateKey(), loginInfo.Mobile)
	if err != nil {
		return core.FormatError(202, err)
	}
	password, err := core.RSADecrypt(core.GetPrivateKey(), loginInfo.Password)
	if err != nil {
		return core.FormatError(202, err)
	}
	client := pb.NewUserServiceClient(conn)
	pbRes, err := client.ValidateLoginInfo(context.Background(), &pb.LoginRequest{
		Mobile:   mobile,
		Password: password,
	})
	if err != nil {
		return err
	}
	check := core.VerifyEncodedPwd(password, pbRes.Salt, pbRes.EncodedPwd)
	if check == false {
		return core.FormatError(200, nil)
	}
	return nil
}

func (bUser *BUser) SetLoginCookie(c *gin.Context, loginInfo *schema.LoginInfo) (string, string, error) {
	duration := time.Hour * 24 * 30 // 30天
	cipherToken, mobile, err := genToken(loginInfo.Mobile, duration)
	if err != nil {
		return EmptyStr, EmptyStr, err
	}
	b64Token := core.BASE64EncodeStr(cipherToken)
	targetDomain := core.GetUpstreamDomain()
	secure := core.GetUpstreamSecure()
	durationInt := int(duration) / int(math.Pow10(9))
	c.SetCookie(
		"TOKEN",
		b64Token,
		durationInt,
		"/",
		targetDomain,
		secure,
		true)
	return b64Token, mobile, nil
}

func (bUser *BUser) GetUserById(c *gin.Context, id uint32) (*schema.UserRes, error) {
	addr := core.GetDownstreamMaxblogBEUserAddr()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Info(core.FormatError(300, err).Error())
		return nil, err
	}
	client := pb.NewUserServiceClient(conn)
	pbRes, err := client.GetUserById(context.Background(), &pb.IdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	dataRes := schema.Pb2UserRes(pbRes)
	return &dataRes, nil
}

func (bUser *BUser) CreateUser(c *gin.Context, userReq *schema.UserReq) (bool, string, error) {
	mobile, err := core.RSADecrypt(core.GetPrivateKey(), userReq.Mobile)
	if err != nil {
		return false, EmptyStr, core.FormatError(202, err)
	}
	password, err := core.RSADecrypt(core.GetPrivateKey(), userReq.Mobile)
	if err != nil {
		return false, EmptyStr, core.FormatError(202, err)
	}
	salt, encodedPwd := core.GetEncodedPwd(password)
	if userReq.NickName == "" {
		userReq.NickName = utils.GenCNName()
	}

	addr := core.GetDownstreamMaxblogBEUserAddr()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Info(core.FormatError(300, err).Error())
		return false, EmptyStr, err
	}
	client := pb.NewUserServiceClient(conn)
	pbRes, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{
		Mobile:   mobile,
		Password: encodedPwd,
		NickName: userReq.NickName,
		Salt:     salt,
	})
	if err != nil {
		return false, EmptyStr, err
	}
	if pbRes == nil {
		return false, EmptyStr, core.FormatError(997, nil)
	}
	return pbRes.Result, mobile, nil
}

func (bUser *BUser) ClearLoginCookie(c *gin.Context) (string, error) {
	j := core.NewJWT()
	b64token, err := c.Cookie("TOKEN")
	if err != nil {
		return EmptyStr, core.FormatError(208, err)
	}
	token, err := core.BASE64DecodeStr(b64token)
	if err != nil {
		return EmptyStr, core.FormatError(202, err)
	}
	decryptedToken, err := core.RSADecrypt(core.GetPrivateKey(), token)
	if err != nil {
		return EmptyStr, core.FormatError(202, err)
	}
	parsedToken, err := j.ParseToken(decryptedToken)
	if err != nil {
		return EmptyStr, core.FormatError(202, err)
	}
	targetDomain := core.GetUpstreamDomain()
	secure := core.GetUpstreamSecure()
	c.SetCookie(
		"TOKEN",
		"",
		1,
		"/",
		targetDomain,
		secure,
		true)
	return parsedToken, nil
}
