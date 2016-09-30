package main

import (
	"net"

	db "github.com/crbaker/libre/libre_server/database"
	se "github.com/crbaker/libre/libre_server/search"

	pb "github.com/crbaker/libre/libre"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) FetchBooks(ctx context.Context, in *pb.Empty) (*pb.FetchBooksReply, error) {
	return &pb.FetchBooksReply{Books: db.FetchBooks()}, nil
}

func (s *server) SaveBook(ctx context.Context, in *pb.SaveBookRequest) (*pb.SaveBookReply, error) {
	code := db.PersistBook(in.Book)
	return &pb.SaveBookReply{ErrorCode: code}, nil
}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	books := se.Search(in.Keyword)

	return &pb.SearchReply{
		Books: books,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	checkErr(err)

	s := grpc.NewServer()
	pb.RegisterLibreServer(s, &server{})
	s.Serve(lis)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
