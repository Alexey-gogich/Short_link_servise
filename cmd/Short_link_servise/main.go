package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"short_link_servise/internal/handlers"
	"short_link_servise/internal/repository"
	inmemoryRepo "short_link_servise/internal/repository/inmemory"
	postgreRepo "short_link_servise/internal/repository/postgre"
	"short_link_servise/internal/servise"
	"time"

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

	DatabaseTypeFlag := flag.String("database", "nil", "Type of using database")
	flag.Parse()
	//Переменная выбора базы данных
	envDatabaseType := os.Getenv("DATABASETYPE")
	var repositoryes *repository.Repos
	if envDatabaseType == "postgre" || *DatabaseTypeFlag == "postgre" {
		//postgre database
		db, _ := sqlx.Open("postgres", "postgres://postgres:postgres@0.0.0.0:8888/short_links_servise?sslmode=disable")
		if err := db.Ping(); err != nil {
			panic(fmt.Errorf("Database connetion error: " + err.Error()))
		}
		repositoryes = postgreRepo.NewRepos(db)
	} else if envDatabaseType == "inmemory" || *DatabaseTypeFlag == "inmemory" {
		// inmemory database
		db := inmemoryRepo.NewCache(time.Duration(30)*time.Minute, time.Duration(1)*time.Minute)
		db.StartCleaner()
		repositoryes = inmemoryRepo.NewRepos(db)
	}

	servises := servise.NewServises(repositoryes)
	handlers := handlers.NewHandlers(servises)

	pb.RegisterShortLiksServer(grpcServer, handlers.Links)
	grpcServer.Serve(listener)
}
