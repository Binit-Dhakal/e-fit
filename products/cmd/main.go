package main

import (
	"log"
	"net"

	"github.com/Binit-Dhakal/e-fit/products/gen"
	"github.com/Binit-Dhakal/e-fit/products/internal/controller"
	grpchandler "github.com/Binit-Dhakal/e-fit/products/internal/handler/grpc"
	"github.com/Binit-Dhakal/e-fit/products/internal/repository/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo, err := postgres.New()
	if err != nil {
		panic(err)
	}

	ctrl := controller.NewProductService(repo)
	h := grpchandler.NewHandler(ctrl)

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterProductServiceServer(srv, h)

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
