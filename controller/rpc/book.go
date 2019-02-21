package rpc

import (
	"github.com/PedroGao/jerry/model"
	"github.com/PedroGao/jerry/pb"
	"golang.org/x/net/context"
)

type Book struct {
}

func (b *Book) GetBook(ctx context.Context, in *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	book, err := model.GetBookById(int(in.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetBookResponse{
		Id:      int32(book.Id),
		Title:   book.Title,
		Author:  book.Author,
		Image:   book.Image,
		Summary: book.Summary,
	}, nil
}

func (b *Book) SearchBook(ctx context.Context, in *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	var (
		book *model.Book
		err  error
	)
	book, err = model.SearchBookByKeyword(in.Keyword)
	if err != nil {
		return nil, err
	}
	return &pb.SearchBookResponse{
		Id:      int32(book.Id),
		Title:   book.Title,
		Author:  book.Author,
		Image:   book.Image,
		Summary: book.Summary,
	}, nil
}
