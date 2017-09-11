package trigger

import (
	"time"
)

//Trigger is a trigger entity
type Trigger struct {
	// start time
	StartTime time.Time
	// the interval of job
	Interval int
	// the unit of job ... ex: day/month/week...
	Unit string
}

//NewTrigger provides a new func of Trigger
func NewTrigger(st time.Time, interval int, unit string) *Trigger {
	return &Trigger{StartTime: st, Interval: interval, Unit: unit}
}

//NextTime get the next run time
func (t *Trigger) NextTime() time.Time {
	//The unit is below
	//year month week day hour minute second
	var nextTime time.Time
	switch t.Unit {
	case "year":
		nextTime = time.Now().AddDate(1*t.Interval, 0, 0)
	case "month":
		nextTime = time.Now().AddDate(0, 1*t.Interval, 0)
	case "week":
		nextTime = time.Now().AddDate(0, 0, 7*t.Interval)
	case "day":
		nextTime = time.Now().AddDate(0, 0, 1*t.Interval)
	case "hour":
		nextTime = time.Now().Add(time.Hour * time.Duration(t.Interval))
	case "minute":
		nextTime = time.Now().Add(time.Minute * time.Duration(t.Interval))
	case "second":
		nextTime = time.Now().Add(time.Second * time.Duration(t.Interval))
	}
	return nextTime
}
