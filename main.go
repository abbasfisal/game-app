package main

import (
	"encoding/json"
	"fmt"
	"github.com/abbasfisal/game-app/repository/mysql"
	"github.com/abbasfisal/game-app/service/authservice"
	"github.com/abbasfisal/game-app/service/userservice"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "at"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/users/register", registerHandler)
	http.HandleFunc("/users/login", loginHandler)
	http.HandleFunc("/users/profile", profileHandler)

	println("localhost:8080 is running")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Method != http.MethodGet {
		fmt.Errorf(`{"error":"invalid method "}`)
		return
	}

	authSrv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	authToken := r.Header.Get("Authorization")

	claims, err := authSrv.VerifyToken(authToken)
	if err != nil {
		fmt.Errorf(`{"error":"token isnot valid "}`)
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(authSrv, mysqlRepo)
	resp, err := userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}

	w.Write(data)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, `{"error":"invalid method "}`)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}

	var req userservice.LoginRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}
	authSrv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	mysqlRepo := mysql.New()
	userSvc := userservice.New(authSrv, mysqlRepo)
	res, err := userSvc.Login(req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}

	data, err = json.Marshal(res)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}

	w.Write(data)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"message":"every thing is ok "}`)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, `{"er	ror":"invalid method "}`)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}

	var req userservice.RegisterRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}
	authSrv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	mysqlRepo := mysql.New()
	userSvc := userservice.New(authSrv, mysqlRepo)
	_, err = userSvc.Register(req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		return
	}
	w.Write([]byte(`{"message":"user created"}`))
}
