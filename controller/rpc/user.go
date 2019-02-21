package rpc

import (
	"github.com/PedroGao/jerry/model"
	"github.com/PedroGao/jerry/pb"
	"golang.org/x/net/context"
)

type User struct {
}

func (u *User) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := &model.UserModel{
		Id: uint64(in.Id),
	}
	ok, err := model.DB.Get(user)
	if !ok {
		return nil, err
	}
	return &pb.GetUserResponse{
		Id:       int32(user.Id),
		Nickname: user.Username,
		Super:    1,
		Active:   1,
		Email:    user.SayHello,
		GroupId:  1,
	}, nil
}
