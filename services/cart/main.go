package main

import (
	"context"
	"fmt"
	"github.com/mydisha/hi-grpc/config"
	proto "github.com/mydisha/hi-grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) InsertCart(ctx context.Context, req *proto.Cart) (*proto.Cart, error) {
	productName := req.GetName()
	status := req.GetStatus()
	project := req.GetProjects()

	fmt.Printf("Berhasil di terima -> Product name %s and cart status %s", productName, status)

	response := proto.Cart{
		Name:   productName,
		Status: project["test"],
	}

	return &response, nil
}

func main() {
	listen, err := net.Listen("tcp", config.CART_SERVICE_PORT)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterCartServiceServer(s, &server{})

	fmt.Printf("Starting cart server")

	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
