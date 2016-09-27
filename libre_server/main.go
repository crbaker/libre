package main

import (
	"log"
	"net"

	pb "github.com/libre/libre"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) FetchBooks(ctx context.Context, in *pb.Empty) (*pb.FetchBooksReply, error) {
	return &pb.FetchBooksReply{}, nil
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
