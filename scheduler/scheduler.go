package scheduler

import (
	"fmt"
	"github.com/abbasfisal/game-app/deliver/dto"
	"github.com/abbasfisal/game-app/service/matchingservice"
	"github.com/go-co-op/gocron/v2"
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

func (s Scheduler) Start(done <-chan bool) {

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
	s.matchSvc.MatchWaitedUsers(dto.MatchWaitedRequest{})
}
