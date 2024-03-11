package scheduler

import (
	"context"
	"fmt"
	"github.com/abbasfisal/game-app/delivery/dto"
	"github.com/abbasfisal/game-app/service/matchingservice"
	"github.com/go-co-op/gocron/v2"
	"sync"
	"time"
)

type Scheduler struct {
	sch      gocron.Scheduler
	matchSvc matchingservice.Service
}

func New(matchingSvc matchingservice.Service) Scheduler {
	sch, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println("scheduler failed: ", err)
		return Scheduler{}
	}
	return Scheduler{
		sch:      sch,
		matchSvc: matchingSvc,
	}
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {

	defer wg.Done()

	s.sch.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(s.MatchWaitedUser),
	)

	s.sch.Start()

	<-done
	fmt.Println("exiting scheduler ")
	s.sch.StopJobs()

}
func (s Scheduler) MatchWaitedUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	_, err := s.matchSvc.MatchWaitedUsers(ctx, dto.MatchWaitedRequest{})
	if err != nil {
		//todo:log error
		//todo:update metrics
		fmt.Println("matchSvc.MatchWaitedUsers : ", err)
	}
}
