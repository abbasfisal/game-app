package config

import (
	"github.com/abbasfisal/game-app/repository/mysql"
	"github.com/abbasfisal/game-app/service/authservice"
	"github.com/abbasfisal/game-app/service/matchingservice"
	"time"
)

type Config struct {
	Application     Application
	HttpServer      HttpServer         //port
	Mysql           mysql.Config       //username,pass,dbname,port,host
	Auth            authservice.Config //accesstoken,signkey,refreshtoken ,etc
	MatchingService matchingservice.Config
}

type HttpServer struct {
	Port int
}

type Application struct {
	GracefulTimeOutShutDown time.Duration `json:"graceful_time_out_shut_down"`
}
