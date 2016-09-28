package main

import (
	"net"

	db "github.com/crbaker/libre/libre_server/database"

	pb "github.com/crbaker/libre/libre"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	_ "github.com/mattn/go-sqlite3"
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

func main() {
	lis, err := net.Listen("tcp", port)
	checkErr(err)

	db.InitDatabase()

	s := grpc.NewServer()
	pb.RegisterLibreServer(s, &server{})
	s.Serve(lis)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
