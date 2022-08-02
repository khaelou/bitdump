package worker

import (
	"context"
	"fmt"
	"log"

	"bitdump/macro"
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

func checkProductQuality(ctx context.Context, worker *Worker, job *Job) {
	if ctx.Err() != nil {
		log.Fatalln("product error:", ctx.Err().Error())
	}

	fmt.Println("[✓]", worker.Role, "@", job.Focus)
}

func (w *Worker) StartWorker(ctx context.Context) {
	go func() {
		for {
			w.WorkerChannel <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				macro.ExecuteMacro(w.ID, w.Factory, w.Role, job.Focus, job.ExecFunc)

				ctx = context.WithValue(ctx, w, job)
				checkProductQuality(ctx, w, &job)
			case <-w.EndShift:
				<-ctx.Done()
				return
			}
		}
	}()
}

func (w *Worker) StopWorker(ctx context.Context) {
	log.Printf("Worker [%d @ %s] has halted!", w.ID, w.Factory)
	w.EndShift <- true
}
