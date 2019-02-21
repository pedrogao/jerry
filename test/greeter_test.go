package test

import (
	"context"
	"github.com/PedroGao/jerry/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

const (
	address = "localhost:3000"
)

var (
	conn *grpc.ClientConn
	err  error
)

func init() {
	// Set up a connection to the server.
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	//defer conn.Close()
}

func TestGreeter(t *testing.T) {
	c := pb.NewBookClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r, err := c.SearchBook(ctx, &pb.SearchBookRequest{Keyword: "c"})
	assert.Equal(t, nil, err)
	assert.Equal(t, "C程序设计语言", r.Title)
}
