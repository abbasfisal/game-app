package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/game-app/config"
	"github.com/abbasfisal/game-app/delivery/httpserver"
	"github.com/abbasfisal/game-app/repository/mysql"
	"github.com/abbasfisal/game-app/service/authservice"
	"github.com/abbasfisal/game-app/service/userservice"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "at"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
	GracefulShutDownTimeOut    = 5 * time.Second
)

func main() {

	cfg := config.Config{
		HttpServer: config.HttpServer{Port: 8080},
		Mysql: mysql.Config{
			Username: "root",
			Password: "password",
			Port:     3307,
			Host:     "localhost",
			DbName:   "gameapp_db",
		},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Application: config.Application{GracefulTimeOutShutDown: GracefulShutDownTimeOut},
	}
	//redis
	//adp := redis.New(cfg)
	//redismatching.New(adp)
	//matchingservice.New()
	//
	authSvc, UserSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, UserSvc)

	var httpServer *echo.Echo
	go func() {
		httpServer = server.Serve()
	}()

	//graceful shut down

	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), cfg.Application.GracefulTimeOutShutDown)
	defer cancel()

	gracefullyShutdown := make(chan os.Signal)
	signal.Notify(gracefullyShutdown, os.Interrupt)
	<-gracefullyShutdown

	err := httpServer.Shutdown(ctxWithTimeOut)
	if err != nil {
		fmt.Println("http server shutdown error: ", err)
		return
	}

	fmt.Println("gracefully shutdown ... ")

	<-ctxWithTimeOut.Done()
	//time.Sleep(cfg.Application.GracefulTimeOutShutDown)
}
func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	mysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, mysqlRepo)

	return authSvc, userSvc
}

//func profileHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("content-type", "application/json")
//
//	if r.Method != http.MethodGet {
//		fmt.Errorf(`{"error":"invalid method "}`)
//		return
//	}
//
//	authSrv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	authToken := r.Header.Get("Authorization")
//
//	claims, err := authSrv.VerifyToken(authToken)
//	if err != nil {
//		fmt.Errorf(`{"error":"token isnot valid "}`)
//	}
//	mysqlRepo := mysql.New()
//	userSvc := userservice.New(authSrv, mysqlRepo)
//	resp, err := userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//		return
//	}
//
//	data, err := json.Marshal(resp)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//		return
//	}
//
//	w.Write(data)
//}

//func loginHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("content-type", "application/json")
//	if r.Method != http.MethodPost {
//		fmt.Fprintf(w, `{"error":"invalid method "}`)
//		return
//	}
//
//	data, err := io.ReadAll(r.Body)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//		return
//	}
//
//	var req userservice.LoginRequest
//	err = json.Unmarshal(data, &req)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//		return
//	}
//	authSrv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	mysqlRepo := mysql.New()
//	userSvc := userservice.New(authSrv, mysqlRepo)
//	res, err := userSvc.Login(req)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//		return
//	}
//
//	data, err = json.Marshal(res)
//	if err != nil {
//		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//		return
//	}
//
//	w.Write(data)
//}
