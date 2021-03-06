package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/config"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/context"
	"github.com/shubhamjagdhane/myBooksShelf/server/users/token"
)

// Server serves HTTP requests for our signup service
type Server struct {
	store      *context.Database
	tokenMaker token.Maker
	router     *gin.Engine
}

// ENV stores our all environmental variables
var ENV config.Config

func loadConfig(path string) {
	cnf, err := config.InitConfig(path)
	if err != nil {
		log.Fatal("Failed to load configuration")
	}
	ENV = cnf
}

// NewServer creates new HTTP server and creates routing
func NewServer() (*Server, error) {
	loadConfig(".")
	tokenMaker, err := token.NewPasetoMaker(ENV.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      context.InitDB(ENV.DBUser, ENV.DBPassword, ENV.DBName, ENV.Host, ENV.Port),
		tokenMaker: tokenMaker,
	}

	server.store.Migrate()
	server.setupRouter()

	return server, nil

}

// setupRouter set the router for the users
func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/api/signup", server.createAccount)
	router.POST("/api/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/api/address/", server.getAddress)
	authRoutes.POST("/api/address/", server.setAddress)
	authRoutes.PUT("/api/address/", server.setDefaultAddress)

	server.router = router
}

// Start starts the server by specifying the address
func (server *Server) Start() error {
	return server.router.Run(ENV.ServerAddress)
}

// errorResponse gives key value pair as error with mentioned error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
