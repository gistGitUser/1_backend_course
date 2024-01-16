package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	db "github.com/gistGitUser/course/project/sqlc"
	"github.com/gistGitUser/course/project/token"
	"github.com/gistGitUser/course/project/util"
	"github.com/go-playground/validator/v10"
	"log"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker *token.PasetoMaker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenAuth, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenAuth,
		router:     nil,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			log.Fatal(err)
		}
	}

	server.SetupRouter()

	return server, nil
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))
	{
		authRoutes.POST("/accounts", server.createAccount)
		authRoutes.GET("/accounts/:id", server.getAccount)
		authRoutes.GET("/accounts", server.listAccount)
		authRoutes.POST("/transfers", server.createTransfer)
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
