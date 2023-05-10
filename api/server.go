package api

import (
	"fmt"

	"github.com/Kcih4518/simpleBank/token"
	"github.com/Kcih4518/simpleBank/util"

	db "github.com/Kcih4518/simpleBank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	// register various handler functions
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

// setupRouter setups all the routes for the HTTP server.
func (server *Server) setupRouter() {
	router := gin.Default()

	// If put multi func last will be handler and other will be middleware
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRotues := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRotues.POST("/accounts", server.createAccount)
	authRotues.GET("/accounts/:id", server.getAccount)
	authRotues.GET("/accounts", server.listAccounts)

	authRotues.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
