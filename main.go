package main

import (
	"TaskScheduler/scheduler"
	"TaskScheduler/trigger"
	"log"
	"time"
)

type job struct{}

func (j job) Start() {
	log.Println("Hello World")
}

func main() {
	trigger := trigger.NewTrigger(time.Now().Add(30*time.Second), 5, "second")
	job1 := job{}
	job2 := job{}
	scheduler := scheduler.NewScheduler(trigger, job1, job2)
	scheduler.Start()
}
