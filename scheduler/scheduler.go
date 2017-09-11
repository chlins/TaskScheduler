package scheduler

import (
	"TaskScheduler/job"
	"TaskScheduler/trigger"
	"log"
	"time"
)

//Scheduler is a implments of schedule
type Scheduler struct {
	Trigger *trigger.Trigger
	Jobs    []job.Jober
	Ender   chan string
}

//NewScheduler new a scheduler
func NewScheduler(trigger *trigger.Trigger, jobs ...job.Jober) Scheduler {
	return Scheduler{Trigger: trigger, Jobs: jobs}
}

//Start a scheduler
func (s *Scheduler) Start() {
	log.Println("Scheduler service start...")
	<-time.After(s.Trigger.StartTime.Sub(time.Now()))
	var nextTime time.Time
	for {
		select {
		case <-s.Ender:
			return
		default:
		}
		for _, job := range s.Jobs {
			go job.Start()
		}
		nextTime = s.Trigger.NextTime()
		<-time.After(nextTime.Sub(time.Now()))
	}
}

//Stop a scheduler
func (s *Scheduler) Stop() {
	s.Ender <- "exit"
}
