package schema

import "maxblog-me-admin/src/pb"

type UserRes struct {
	Id       uint32 `json:"id"`
	Mobile   string `json:"mobile"`
	NickName string `json:"nickName"`
	Role     uint32 `json:"role"`
}

func Pb2UserRes(userRes *pb.UserRes) UserRes {
	return UserRes{
		Id:       userRes.Id,
		Mobile:   userRes.Mobile,
		NickName: userRes.NickName,
		Role:     userRes.Role,
	}
}
