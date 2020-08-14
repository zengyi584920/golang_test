package main

type Worker struct {
	JobQueue chan Job
}

func NewWorker() Worker {
	return Worker{JobQueue: make(chan Job)}
}
func (w Worker) Run(wq chan chan Job) {
	go func() {
		for {
			//timeout := time.NewTimer(time.Microsecond * 500)
			wq <- w.JobQueue//放入工作池
			select {
			case job := <-w.JobQueue:
				job.Do()
			//case <-timeout.C:
			//	fmt.Println("worker time out")
			}
		}
	}()
}