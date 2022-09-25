package schema

import "maxblog-me-template/src/pb"

type DataRes struct {
	Mobile string `json:"mobile"`
}

func Pb2Res(dataRes *pb.DataRes) DataRes {
	return DataRes{
		Mobile: dataRes.Mobile,
	}
}
