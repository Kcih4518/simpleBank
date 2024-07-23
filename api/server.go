package api

import (
	"log"

	db "github.com/Kcih4518/simpleBank/db/sqlc"
	"github.com/Kcih4518/simpleBank/token"
	"github.com/Kcih4518/simpleBank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
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

	// validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("currency", validCurrency); err != nil {
			log.Fatalf("Failed to register 'currency' custom validation: %v", err)

		}
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// TODO: add routers
	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.DELETE("/accounts/:id", server.delAccount)
	authRoutes.PATCH("/accounts/:id", server.updateAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
