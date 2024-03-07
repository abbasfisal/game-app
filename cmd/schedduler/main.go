package main

import (
	"fmt"
	"github.com/abbasfisal/game-app/scheduler"
	"github.com/abbasfisal/game-app/service/matchingservice"
	"sync"
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

	matchSvc := matchingservice.New()
	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan bool)
	sch := scheduler.New(matchSvc)
	go func() {
		sch.Start(done, &wg)
	}()
	done <- true

	wg.Wait()
	fmt.Println("received interrupt , shutting down gracefully")

	time.Sleep(GracefulShutDownTimeOut)
}
