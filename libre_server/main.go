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

	dummyBook := pb.Book{Title: "Some Book", Description: "Really long read about the moon"}
	otherBook := pb.Book{Title: "Some Book", Description: "Really long read about the moon"}

	books := []*pb.Book{&dummyBook, &otherBook}

	return &pb.FetchBooksReply{Books: books}, nil
}

func (s *server) SaveBook(ctx context.Context, in *pb.SaveBookRequest) (*pb.SaveBookReply, error) {

	db.PersistBook(in.Book)

	return &pb.SaveBookReply{ErrorCode: pb.SaveBookReply_OK, Message: in.Book.Title}, nil
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
