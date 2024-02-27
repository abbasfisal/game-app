package httpserver

import (
	"fmt"
	"github.com/abbasfisal/game-app/config"
	"github.com/abbasfisal/game-app/service/authservice"
	"github.com/abbasfisal/game-app/service/userservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config  config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service) Server {
	return Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}

func (s Server) Serve() {
	e := echo.New()

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routes
	e.GET("/health-check", s.healthcheck)

	userGroup := e.Group("/users")

	userGroup.POST("/register", s.registerHandler)
	userGroup.POST("/login", s.loginHandler)
	userGroup.POST("/profileHandler", s.loginHandler)

	//	http.HandleFunc("/users/profile", profileHandler)

	//run server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)))
}
