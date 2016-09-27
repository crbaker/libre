package main

import (
	"log"
	"net"

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

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLibreServer(s, &server{})
	s.Serve(lis)
}
