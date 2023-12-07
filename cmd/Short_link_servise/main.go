package main

import (
	"fmt"
	"net"
	"short_link_servise/internal/handlers"
	postgreRepo "short_link_servise/internal/repository/postgre"
	"short_link_servise/internal/servise"

	pb "short_link_servise/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	listener, err := net.Listen("tcp", ":5248")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	//postgre database
	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@0.0.0.0:8888/short_links_servise?sslmode=disable")
	if err != nil {
		panic(fmt.Errorf("Database connetion error: " + err.Error()))
	}
	repositoryes := postgreRepo.NewRepos(db)
	//inmemory database
	// db := inmemoryRepo.NewCache(time.Duration(30)*time.Minute, time.Duration(1)*time.Minute)
	// db.StartCleaner()
	// repositoryes := inmemoryRepo.NewRepos(db)
	servises := servise.NewServises(repositoryes)
	handlers := handlers.NewHandlers(servises)

	pb.RegisterShortLiksServer(grpcServer, handlers.Links)
	grpcServer.Serve(listener)
}
