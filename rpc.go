package main

import (
	"context"
	"fmt"
	"github.com/mydisha/hi-grpc/config"
	proto "github.com/mydisha/hi-grpc/pb"
	"google.golang.org/grpc"
	"log"
)

func serviceCartClient() proto.CartServiceClient {
	conn, err := grpc.Dial(config.CART_SERVICE_PORT, grpc.WithInsecure())

	if err != nil {
		log.Fatal("failed connect to cart service")
	}

	return proto.NewCartServiceClient(conn)
}

func serviceProductClient() proto.ProductServiceClient {
	conn, err := grpc.Dial(config.PRODUCT_SERVICE_PORT, grpc.WithInsecure())

	if err != nil {
		log.Fatal("failed to connect product service")
	}

	return proto.NewProductServiceClient(conn)
}

func main() {
	ctx := context.Background()
	cart := serviceCartClient()
	cartResp, err := cart.InsertCart(ctx, &proto.Cart{
		Name:   "mie lemonilo",
		Status: "deleted",
		Projects: map[string]string{
			"test": "map test",
		},
	})

	if err != nil {
		log.Fatal("failed to insert cart")
	}

	fmt.Printf("cart response : product name %s, status %s \n", cartResp.Name, cartResp.Status)

	product := serviceProductClient()
	productResp, err := product.InsertProduct(ctx, &proto.Product{
		Name:  "Produk mie lemonilo",
		Price: 7000,
		Image: "google.com",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("product response : product name %s, image %s \n", productResp.Name, productResp.GetImage())
}
