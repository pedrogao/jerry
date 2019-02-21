package main

import (
	"github.com/PedroGao/jerry/config"
	"github.com/PedroGao/jerry/controller/rpc"
	"github.com/PedroGao/jerry/model"
	"github.com/PedroGao/jerry/pb"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	rcfg = pflag.StringP("config", "c", "", "config file path")
)

const (
	port = ":3000"
)

func main() {

	// parse the flags
	pflag.Parse()

	// init config from file
	if err := config.Init(*rcfg); err != nil {
		panic(err)
	}

	// init db
	model.Init()
	defer model.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// apply Interceptor middleware for authorization
	//s := grpc.NewServer(grpc.UnaryInterceptor(middleware.Interceptor))
	// remove authority just for test, use it in production
	s := grpc.NewServer()

	pb.RegisterUserServer(s, &rpc.User{})
	pb.RegisterBookServer(s, &rpc.Book{})

	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("listening at: %s", port)

	}
}
