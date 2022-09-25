package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"maxblog-me-template/internal/core"
	"maxblog-me-template/src/pb"
	"maxblog-me-template/src/schema"
)

var DataSet = wire.NewSet(wire.Struct(new(BData), "*"))

type BData struct{}

func (b *BData) GetDataById(c *gin.Context, id uint32) (*schema.DataRes, error) {
	addr := core.GetDownstreamMaxblogBETemplateAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Fatal(core.FormatError(300, err).Error())
		return nil, err
	}
	client := pb.NewDataServiceClient(conn)
	pbRes, err := client.GetDataById(context.Background(), &pb.IdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	dataRes := schema.Pb2Res(pbRes)
	return &dataRes, nil
}
