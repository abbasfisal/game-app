package main

import (
	"fmt"
	"github.com/abbasfisal/game-app/scheduler"
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

	done := make(chan bool)
	sch := scheduler.New()
	go func() {
		sch.Start(done)
	}()
	done <- true

	fmt.Println("received interrupt , shutting down gracefully")
	
	time.Sleep(GracefulShutDownTimeOut)
}
