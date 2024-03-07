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
	Router  *echo.Echo
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service) Server {
	return Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
		Router:  echo.New(),
	}
}

func (s Server) Serve() {

	//middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	//routes
	s.Router.GET("/health-check", s.healthcheck)

	userGroup := s.Router.Group("/users")

	userGroup.POST("/register", s.registerHandler)
	//userGroup.POST("/login", s.loginHandler)
	//userGroup.POST("/profileHandler", s.loginHandler)
	//	http.HandleFunc("/users/profile", profileHandler)

	//run server
	s.Router.Logger.Fatal(s.Router.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)))

}
