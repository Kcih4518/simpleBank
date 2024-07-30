package gapi

import (
	"log"

	db "github.com/Kcih4518/simpleBank/db/sqlc"
	"github.com/Kcih4518/simpleBank/pb"
	"github.com/Kcih4518/simpleBank/token"
	"github.com/Kcih4518/simpleBank/util"
)

// Server serves gRPC request for out banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %v", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker}

	return server, nil
}
