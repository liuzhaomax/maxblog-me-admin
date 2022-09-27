package schema

import "maxblog-me-admin/src/pb"

type UserRes struct {
	Mobile string `json:"mobile"`
}

func Pb2Res(dataRes *pb.UserRes) UserRes {
	return UserRes{
		Mobile: dataRes.Mobile,
	}
}
