package job

//Jober is a job interface
type Jober interface {
	//Start jobs
	Start()
}

//JobFunc common func
type JobFunc func()

//Job is a common struct
type Job struct {
	jobFunc JobFunc
}

//NewJob is produce a job
func NewJob(job JobFunc) *Job {
	return &Job{jobFunc: job}
}

//Start of job
func (j *Job) Start() {
	go j.jobFunc()
}
