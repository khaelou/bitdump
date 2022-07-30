package worker

import (
	"bitdump/macro"
	"log"
)

type Job struct {
	ID       int
	Focus    string
	ExecFunc macro.ExecMacro
}

type Worker struct {
	ID            int
	Factory       string
	Role          string
	WorkerChannel chan chan Job
	JobChannel    chan Job
	EndShift      chan bool
}

func (w *Worker) StartWorker() {
	go func() {
		for {
			w.WorkerChannel <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				macro.ExecuteMacro(w.ID, w.Factory, w.Role, job.Focus, job.ExecFunc)
			case <-w.EndShift:
				return
			}
		}
	}()
}

func (w *Worker) StopWorker() {
	log.Printf("Worker [%d @ %s] has halted!", w.ID, w.Factory)
	w.EndShift <- true
}
