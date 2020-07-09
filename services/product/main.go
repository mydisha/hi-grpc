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

func (*server) InsertProduct(ctx context.Context, req *proto.Product) (*proto.Product, error) {
	productName := req.GetName()
	productPrice := req.GetPrice()
	productImage := req.GetImage()

	fmt.Printf("Product name %s, price %d, and image %s", productName, productPrice, productImage)

	response := proto.Product{
		Name:  productName,
		Price: productPrice,
		Image: "lemonilo.com",
	}

	return &response, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0"+config.PRODUCT_SERVICE_PORT)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterProductServiceServer(s, &server{})

	fmt.Printf("Starting product server")

	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
