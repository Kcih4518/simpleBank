// main.go
package main

import (
	"database/sql"
	"log"
	"net"

	db "github.com/Kcih4518/simpleBank/db/sqlc"
	"github.com/Kcih4518/simpleBank/gapi"
	"github.com/Kcih4518/simpleBank/pb"
	"github.com/Kcih4518/simpleBank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to db
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	store := db.NewStore(conn)
	runGrpcServer(config, store)

}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterSimpleBankServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		log.Fatal("cannot create listener:")
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:")
	}

}

// func runGinServer(config util.Config, store db.Store) {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("cannot create server:", err)
// 	}
// 	err = server.Start(config.HTTPServerAddress)
// 	if err != nil {
// 		log.Fatal("cannot start server:", err)
// 	}
// }
