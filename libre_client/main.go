package main

import (
	"fmt"
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

	c.SaveBook(context.Background(), &pb.SaveBookRequest{Book: &pb.Book{Title: "Some Funky Book"}})

	fetchReply, err := c.FetchBooks(context.Background(), &pb.Empty{})

	fmt.Println(len(fetchReply.Books))

	// c.DeleteBook(context.Background(), &pb.DeleteBookRequest{Book: saveReply.Book})

	//fetchReply, err = c.FetchBooks(context.Background(), &pb.Empty{})
	//fmt.Println(len(fetchReply.Books))

	// sr, err := c.Search(context.Background(), &pb.SearchRequest{Keyword: "978-0393338102"})

	// log.Println(sr.Books)
	// log.Print(len(fetchReply.Books))
	// log.Printf("Books: %s", r.Message)
}
