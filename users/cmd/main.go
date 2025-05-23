package main

import (
	"log"
	"net"

	"github.com/Binit-Dhakal/e-fit/users/gen"
	"github.com/Binit-Dhakal/e-fit/users/internal/controller"
	grpchandler "github.com/Binit-Dhakal/e-fit/users/internal/handler/grpc"
	"github.com/Binit-Dhakal/e-fit/users/internal/repository/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo, err := postgres.New()
	if err != nil {
		panic(err)
	}

	ctrl := controller.NewUserService(repo)
	h := grpchandler.NewHandler(ctrl)

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterUserServiceServer(srv, h)

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
