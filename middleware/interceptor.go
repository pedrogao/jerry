package middleware

import (
	"context"
	"github.com/PedroGao/jerry/libs/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	err = auth(ctx)
	if err != nil {
		return
	}
	// 继续处理请求
	return handler(ctx, req)
}

func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		tokenStr string
	)
	if val, ok := md["token"]; ok {
		tokenStr = val[0]
	}
	indentify, err := token.JwtInstance.VerifySingleToken(tokenStr)
	log.Println(indentify)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Token认证信息无效: ", err.Error())
	}
	return nil
}
