package scheduler

import (
	"fmt"
	"time"
)

type Scheduler struct {
}

func New() Scheduler {
	return Scheduler{}
}

func (s Scheduler) Start(done <-chan bool) {
	for {
		select {
		case d := <-done:
			fmt.Println("exiting scheduler :", d)
			return
		default:
			fmt.Println("scheduler now ", time.Now())
			return
		}
	}
}
