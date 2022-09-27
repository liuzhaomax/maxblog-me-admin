package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"maxblog-me-admin/internal/core"
	"maxblog-me-admin/src/pb"
	"maxblog-me-admin/src/schema"
)

var UserSet = wire.NewSet(wire.Struct(new(BUser), "*"))

type BUser struct{}

func (bUser *BUser) GetUserById(c *gin.Context, id uint32) (*schema.UserRes, error) {
	addr := core.GetDownstreamBEUserAddr()
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
