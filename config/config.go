package config

import (
	"github.com/abbasfisal/game-app/repository/mysql"
	"github.com/abbasfisal/game-app/service/authservice"
)

type Config struct {
	HttpServer HttpServer         //port
	Mysql      mysql.Config       //username,pass,dbname,port,host
	Auth       authservice.Config //accesstoken,signkey,refreshtoken ,etc
}

type HttpServer struct {
	Port int
}
