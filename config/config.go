package config

import (
	"github.com/abbasfisal/game-app/repository/mysql"
	"github.com/abbasfisal/game-app/service/authservice"
)

type Config struct {
	HttpServer HttpServer
	Mysql      mysql.Config
	Auth       authservice.Config
}

type HttpServer struct {
	Port int
}
