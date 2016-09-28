package main

import (
	"log"

	pb "github.com/crbaker/libre/libre"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewLibreClient(conn)

	_, err = c.SaveBook(context.Background(), &pb.SaveBookRequest{Book: &pb.Book{Title: "Some Funky Book"}})

	fetchReply, err := c.FetchBooks(context.Background(), &pb.Empty{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Print(len(fetchReply.Books))
	// log.Printf("Books: %s", r.Message)
}
